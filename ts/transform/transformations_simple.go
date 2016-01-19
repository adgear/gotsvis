// Copyright (c) 2014 Datacratic. All rights reserved.

package transform

import (
	"fmt"
	"gotsvis/ts"
	"math"
)

type CumulativeSum struct {
	sum float64
}

func (cs *CumulativeSum) Name() string {
	return "CumulativeSum"
}

func (cs *CumulativeSum) Transform(val float64) float64 {
	if math.IsNaN(val) {
		return cs.sum
	}
	cs.sum += val
	return cs.sum
}

type DivideBy struct {
	By float64
}

func (div *DivideBy) Name() string {
	return fmt.Sprintf("DivideBy(%f)", div.By)
}

func (div *DivideBy) Transform(val float64) float64 {
	if math.IsNaN(val) || math.IsInf(val, 0) {
		return val
	}
	return val / div.By
}

type MultiplyBy struct {
	By float64
}

func (mult *MultiplyBy) Name() string {
	return fmt.Sprintf("MultiplyBy(%f)", mult.By)
}

func (mult *MultiplyBy) Transform(val float64) float64 {
	if math.IsNaN(val) || math.IsInf(val, 0) {
		return val
	}
	return val * mult.By
}

type MarkRaise struct {
	count int64
	last  float64
}

func (raise *MarkRaise) Name() string {
	return "MarkRaise"
}

func (raise *MarkRaise) Transform(val float64) float64 {
	if math.IsNaN(val) {
		return val
	}
	if raise.count == 0 {
		raise.count++
		raise.last = val
		return 0.0
	}
	raise.count++
	if raise.last == val {
		raise.last = val
		return 0.0
	}
	if raise.last-val < 0.0 {
		raise.last = val
		return 1.0
	}
	raise.last = val
	return 0.0
}

type MarkLarger struct {
	Than float64
}

func (mark *MarkLarger) Name() string {
	return fmt.Sprintf("MarkLarger(%f)", mark.Than)
}

func (mark *MarkLarger) Transform(val float64) float64 {
	if math.IsNaN(val) {
		return val
	}
	if val > mark.Than {
		return 1.0
	}
	return 0.0
}

// Mark Equal works except with NaN, NaN is always false,
// even when compared to itself.
type MarkEqual struct {
	To float64
}

func (mark *MarkEqual) Name() string {
	return fmt.Sprintf("MarkEqual(%f)", mark.To)
}
func (mark *MarkEqual) Transform(val float64) float64 {
	if math.IsNaN(val) {
		return val
	}
	if val == mark.To {
		return 1.0
	}
	return 0.0
}

type DiffPrevious struct {
	count int64
	last  float64
}

func (diff *DiffPrevious) Name() string {
	return "MarkRaise"
}

func (diff *DiffPrevious) Transform(val float64) float64 {
	panic("not finished")

	if math.IsNaN(val) {
		return val
	}

	if diff.count == 0 {
		diff.count++
		diff.last = val
		return 0.0
	}
	diff.count++
	if diff.last == val {
		diff.last = val
		return 0.0
	}
	if diff.last-val < 0.0 {
		diff.last = val
		return 1.0
	}
	diff.last = val
	return 0.0
}

type Transforms struct {
	Transforms []ts.Transform
}

func (ts *Transforms) Name() string {
	s := "Transforms["
	for _, t := range ts.Transforms {
		s += "(" + t.Name() + ")"
	}
	s += "]"
	return s
}

func (ts *Transforms) Transform(val float64) float64 {
	for _, t := range ts.Transforms {
		val = t.Transform(val)
	}
	return val
}

type IfTrueSet struct {
	Condition ts.Transform
	Value     float64
}

func (ies *IfTrueSet) Name() string {
	return fmt.Sprintf("IfEqual(%s)Set(%f)Over", ies.Condition.Name(), ies.Value)
}

func (ies *IfTrueSet) Transform(val float64) float64 {
	cond := ies.Condition.Transform(val)
	if math.IsNaN(cond) || cond == 0 {
		return val
	}
	return ies.Value
}
