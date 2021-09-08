package metrics

import (
	"math"
	"sort"
	"sync"
	"time"
)

// Timing 维护每个时间桶的时间长短, 持续时间保存在一个数组中，以允许从源数据计算各种统计数据。
type Timing struct {
	Buckets map[int64]*timingBucket
	Mutex   *sync.RWMutex

	CachedSortedDurations []time.Duration
	LastCachedTime        int64
}

// 时间桶
type timingBucket struct {
	Durations []time.Duration
}

// NewTiming 新建滑动时间结构
func NewTiming() *Timing {
	r := &Timing{
		Buckets: make(map[int64]*timingBucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

// 创建一个可排序的数组,并实现排序相关的接口
type byDuration []time.Duration

func (c byDuration) Len() int {
	return len(c)
}
func (c byDuration) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byDuration) Less(i, j int) bool {
	return c[i] < c[j]
}

// SortedDurations 返回一个 time.Duration 数组，从最近 60 秒内发生的最短到最长排序。
func (r *Timing) SortedDurations() []time.Duration {
	r.Mutex.RLock()
	t := r.LastCachedTime
	r.Mutex.RUnlock()

	if t+time.Duration(1*time.Second).Nanoseconds() > time.Now().UnixNano() {
		return r.CachedSortedDurations
	}

	var durations byDuration
	now := time.Now()

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp > now.Unix()-60 {
			for _, duration := range bucket.Durations {
				durations = append(durations, duration)
			}
		}
	}

	sort.Sort(durations)
	r.CachedSortedDurations = durations
	r.LastCachedTime = time.Now().UnixNano()

	return r.CachedSortedDurations
}

// 获取当前时间桶
func (r *Timing) getCurrentBucket() *timingBucket {
	r.Mutex.RLock()
	now := time.Now()
	bucket, exists := r.Buckets[now.Unix()]
	r.Mutex.RUnlock()

	if !exists {
		r.Mutex.Lock()
		defer r.Mutex.Unlock()

		r.Buckets[now.Unix()] = &timingBucket{}
		bucket = r.Buckets[now.Unix()]
	}
	return bucket
}

// 移除所有桶
func (r *Timing) removeOldBuckets() {
	now := time.Now()

	for timestamp := range r.Buckets {
		if timestamp <= now.Unix()-60 {
			delete(r.Buckets, timestamp)
		}
	}
}

// Add 增加时间段
func (r *Timing) Add(duration time.Duration) {
	b := r.getCurrentBucket()

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b.Durations = append(b.Durations, duration)
	r.removeOldBuckets()
}

// Percentile 计算百分位
func (r *Timing) Percentile(p float64) uint32 {
	sortedDurations := r.SortedDurations()
	length := len(sortedDurations)
	if length <= 0 {
		return 0
	}
	pos := r.ordinal(len(sortedDurations), p) - 1
	return uint32(sortedDurations[pos].Nanoseconds() / 1000000)
}

// 计算排序
func (r *Timing) ordinal(length int, percentile float64) int64 {
	if percentile == 0 && length > 0 {
		return 1
	}
	return int64(math.Ceil((percentile / float64(100)) * float64(length)))
}

// Mean 含义
func (r *Timing) Mean() uint32 {
	sortedDurations := r.SortedDurations()
	var sum time.Duration
	for _, d := range sortedDurations {
		sum += d
	}

	length := int64(len(sortedDurations))
	if length == 0 {
		return 0
	}

	return uint32(sum.Nanoseconds()/length) / 1000000
}
