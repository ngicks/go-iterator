package base

// Range is numeric range iterator
// which iterates over contiguous number in range of [start, end).
//
// Each Next call advances iterator 1 ahead to its tail.
// If you need to skip some number, use with Exclude or Map.
type Range struct {
	start int
	end   int
}

func NewRange(start, end int) *Range {
	return &Range{
		start: start,
		end:   end,
	}
}

func (r *Range) Next() (next int, ok bool) {
	if r.start >= r.end {
		return 0, false
	}
	next = r.start
	r.start++
	return next, true
}

func (r *Range) NextBack() (next int, ok bool) {
	if r.start >= r.end {
		return 0, false
	}
	next = r.end - 1
	r.end--
	return next, true
}

func (r *Range) SizeHint() int {
	return r.end - r.start
}
