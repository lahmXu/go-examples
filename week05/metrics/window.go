package metrics

// Window 窗口
type Window struct {
	window []Bucket
	size   int
}

// WindowOpts 创建窗口参数
type WindowOpts struct {
	Size int
}

// NewWindow 创建新窗口
func NewWindow(opts WindowOpts) *Window {
	buckets := make([]Bucket, opts.Size)
	for offset := range buckets {
		buckets[offset] = Bucket{Points: make([]float64, 0)}
		nextOffset := offset + 1
		if nextOffset == opts.Size {
			nextOffset = 0
		}
		buckets[offset].next = &buckets[nextOffset]
	}
	return &Window{window: buckets, size: opts.Size}
}

// ResetWindow 重置窗口
func (w *Window) ResetWindow() {
	for offset := range w.window {
		w.ResetBucket(offset)
	}
}

// ResetBucket 重置某个桶
func (w *Window) ResetBucket(offset int) {
	w.window[offset].Reset()
}

// ResetBuckets 重置多个桶
func (w *Window) ResetBuckets(offsets []int) {
	for _, offset := range offsets {
		w.ResetBucket(offset)
	}
}

// Append 窗口新增值
func (w *Window) Append(offset int, val float64) {
	w.window[offset].Append(val)
}

// Add 窗口新增值
func (w *Window) Add(offset int, val float64) {
	if w.window[offset].Count == 0 {
		w.window[offset].Append(val)
		return
	}
	w.window[offset].Add(0, val)
}

// Bucket 返回桶
func (w *Window) Bucket(offset int) Bucket {
	return w.window[offset]
}

// Size 返回窗口大小
func (w *Window) Size() int {
	return w.size
}
