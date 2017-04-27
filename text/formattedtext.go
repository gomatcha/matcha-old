package text

// import "sort"

// type _range struct {
// 	min, max int
// }

// func (r _range) intersect(r2 _range) _range {
// 	var max, min int
// 	if r.min > r2.min {
// 		min = r.min
// 	}
// 	if r.max > r2.max {
// 		max = r.max
// 	}
// 	return _range{min: min, max: max}
// }

// func (r _range) isValid() bool {
// 	return r.max >= r.min
// }

// type formatRange struct {
// 	min    *Position
// 	max    *Position
// 	style *Style
// }

// type StyletedText struct {
// 	Text         Text
// 	formatRanges []*formatRange
// }

// func (ft *StyletedText) At(byteIdx int) *Style {
// 	return nil
// }

// func (ft *StyletedText) Update(style *Style, min, max int) {
// 	for _, i := range ft.ranges(min, max) {
// 		i.style.Update(style)
// 	}
// }

// func (ft *StyletedText) Set(style *Style, min, max int) {
// 	for _, i := range ft.ranges(min, max) {
// 		i.style = style
// 	}
// }

// func (ft *StyletedText) ranges(min, max int) []*formatRange {
// 	r := _range{min, max}

// 	// Get formatRanges that intersect.
// 	ranges := []*formatRange{}
// 	for _, i := range ft.formatRanges {
// 		r2 := _range{i.min.Index(), i.max.Index()}
// 		if r.intersect(r2).isValid() {
// 			ranges = append(ranges, i)
// 		}
// 	}

// 	toAdd := []*formatRange{}
// 	if len(ranges) > 0 {
// 		// Trim the first formatRange to min.
// 		first := ranges[0]
// 		if first.min.Index() < min {
// 			toAdd = append(toAdd, &formatRange{
// 				style: first.style.Copy(),
// 				min:    ft.Text.Position(first.min.Index()),
// 				max:    ft.Text.Position(min - 1),
// 			})
// 			first.min = ft.Text.Position(min)
// 		}

// 		// Trim the last formatRange to max.
// 		last := ranges[len(ranges)-1]
// 		if last.max.Index() > max {
// 			toAdd = append(toAdd, &formatRange{
// 				style: last.style.Copy(),
// 				min:    ft.Text.Position(max + 1),
// 				max:    ft.Text.Position(last.max.Index()),
// 			})
// 			last.max = ft.Text.Position(max)
// 		}
// 	}

// 	// Fill in any gaps.
// 	gaps := []*formatRange{}
// 	for _, i := range ranges {
// 		if i.min.Index() < r.min {
// 			fr := &formatRange{
// 				style: &Style{},
// 				min:    ft.Text.Position(r.min),
// 				max:    ft.Text.Position(i.min.Index()),
// 			}
// 			gaps = append(gaps, fr)
// 			r.min = i.max.Index() + 1
// 		}
// 	}
// 	if r.min != r.max {
// 		fr := &formatRange{
// 			style: &Style{},
// 			min:    ft.Text.Position(r.min),
// 			max:    ft.Text.Position(r.max),
// 		}
// 		gaps = append(gaps, fr)
// 	}
// 	ranges = append(ranges, gaps...)

// 	// Update formatRanges.
// 	ft.formatRanges = append(ft.formatRanges, toAdd...)
// 	ft.formatRanges = append(ft.formatRanges, gaps...)
// 	sort.Slice(ft.formatRanges, func(i, j int) bool {
// 		return ft.formatRanges[i].min.Index() < ft.formatRanges[j].min.Index()
// 	})

// 	return ranges
// }
