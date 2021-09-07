package metric

// Bucket 桶
type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

// Append 桶末尾添加值
func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++
}

// Add 桶内某个位置添加值
func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] += val
	b.Count++
}

// Reset 清空桶
func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0
}

// Next 下一个桶
func (b *Bucket) Next() *Bucket {
	return b.next
}
