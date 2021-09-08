package metrics

import (
	"sync"
	"time"
)

// Number 在有限的时间段内跟踪桶的个数, 当前是桶时长是 1 秒,且只保存最后 10 秒
type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

// 个数桶
type numberBucket struct {
	Value float64
}

// NewNumber 创建个数桶
func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

// 获取当前个数桶
func (u *Number) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool
	if bucket, ok = u.Buckets[now]; !ok {
		bucket = &numberBucket{}
		u.Buckets[now] = bucket
	}
	return bucket
}

// RemoveOldBuckets 移除旧桶
func (u *Number) RemoveOldBuckets() {
	now := time.Now().Unix() - 10
	for timestamp := range u.Buckets {
		if timestamp <= now {
			delete(u.Buckets, timestamp)
		}
	}
}

// Increment 桶数量增加
func (u *Number) Increment(i float64) {
	if i == 0 {
		return
	}

	u.Mutex.Lock()
	defer u.Mutex.Unlock()

	b := u.getCurrentBucket()
	b.Value += i
	u.RemoveOldBuckets()
}

// UpdateMax 更新桶最大值
func (u *Number) UpdateMax(n float64) {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()

	b := u.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	u.RemoveOldBuckets()
}

// Sum 返回最后 10 秒的和
func (u *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	u.Mutex.RLocker()
	defer u.Mutex.RUnlock()

	for timestamp, bucket := range u.Buckets {
		if timestamp >= now.Unix()-10 {
			sum += bucket.Value
		}
	}

	return sum
}

// Max 返回最后 10 秒内的最大值
func (u *Number) Max(now time.Time) float64 {
	var max float64

	u.Mutex.RLock()
	defer u.Mutex.RUnlock()

	for timestamp, bucket := range u.Buckets {
		if timestamp >= now.Unix()-10 {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}
	return max
}

// Avg 某段时间求平均值
func (u *Number) Avg(now time.Time) float64 {
	return u.Sum(now) / 10
}
