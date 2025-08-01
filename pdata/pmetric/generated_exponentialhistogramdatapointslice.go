// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"iter"
	"sort"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

// ExponentialHistogramDataPointSlice logically represents a slice of ExponentialHistogramDataPoint.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewExponentialHistogramDataPointSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ExponentialHistogramDataPointSlice struct {
	orig  *[]*otlpmetrics.ExponentialHistogramDataPoint
	state *internal.State
}

func newExponentialHistogramDataPointSlice(orig *[]*otlpmetrics.ExponentialHistogramDataPoint, state *internal.State) ExponentialHistogramDataPointSlice {
	return ExponentialHistogramDataPointSlice{orig: orig, state: state}
}

// NewExponentialHistogramDataPointSlice creates a ExponentialHistogramDataPointSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewExponentialHistogramDataPointSlice() ExponentialHistogramDataPointSlice {
	orig := []*otlpmetrics.ExponentialHistogramDataPoint(nil)
	state := internal.StateMutable
	return newExponentialHistogramDataPointSlice(&orig, &state)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewExponentialHistogramDataPointSlice()".
func (es ExponentialHistogramDataPointSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ExponentialHistogramDataPointSlice) At(i int) ExponentialHistogramDataPoint {
	return newExponentialHistogramDataPoint((*es.orig)[i], es.state)
}

// All returns an iterator over index-value pairs in the slice.
//
//	for i, v := range es.All() {
//	    ... // Do something with index-value pair
//	}
func (es ExponentialHistogramDataPointSlice) All() iter.Seq2[int, ExponentialHistogramDataPoint] {
	return func(yield func(int, ExponentialHistogramDataPoint) bool) {
		for i := 0; i < es.Len(); i++ {
			if !yield(i, es.At(i)) {
				return
			}
		}
	}
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ExponentialHistogramDataPointSlice can be initialized:
//
//	es := NewExponentialHistogramDataPointSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ExponentialHistogramDataPointSlice) EnsureCapacity(newCap int) {
	es.state.AssertMutable()
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlpmetrics.ExponentialHistogramDataPoint, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ExponentialHistogramDataPoint.
// It returns the newly added ExponentialHistogramDataPoint.
func (es ExponentialHistogramDataPointSlice) AppendEmpty() ExponentialHistogramDataPoint {
	es.state.AssertMutable()
	*es.orig = append(*es.orig, &otlpmetrics.ExponentialHistogramDataPoint{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ExponentialHistogramDataPointSlice) MoveAndAppendTo(dest ExponentialHistogramDataPointSlice) {
	es.state.AssertMutable()
	dest.state.AssertMutable()
	// If they point to the same data, they are the same, nothing to do.
	if es.orig == dest.orig {
		return
	}
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ExponentialHistogramDataPointSlice) RemoveIf(f func(ExponentialHistogramDataPoint) bool) {
	es.state.AssertMutable()
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			(*es.orig)[i] = nil
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		(*es.orig)[i] = nil
		newLen++
	}
	*es.orig = (*es.orig)[:newLen]
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ExponentialHistogramDataPointSlice) CopyTo(dest ExponentialHistogramDataPointSlice) {
	dest.state.AssertMutable()
	*dest.orig = copyOrigExponentialHistogramDataPointSlice(*dest.orig, *es.orig)
}

// Sort sorts the ExponentialHistogramDataPoint elements within ExponentialHistogramDataPointSlice given the
// provided less function so that two instances of ExponentialHistogramDataPointSlice
// can be compared.
func (es ExponentialHistogramDataPointSlice) Sort(less func(a, b ExponentialHistogramDataPoint) bool) {
	es.state.AssertMutable()
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}

// marshalJSONStream marshals all properties from the current struct to the destination stream.
func (ms ExponentialHistogramDataPointSlice) marshalJSONStream(dest *json.Stream) {
	dest.WriteArrayStart()
	if len(*ms.orig) > 0 {
		ms.At(0).marshalJSONStream(dest)
	}
	for i := 1; i < len(*ms.orig); i++ {
		dest.WriteMore()
		ms.At(i).marshalJSONStream(dest)
	}
	dest.WriteArrayEnd()
}

// unmarshalJSONIter unmarshals all properties from the current struct from the source iterator.
func (ms ExponentialHistogramDataPointSlice) unmarshalJSONIter(iter *json.Iterator) {
	iter.ReadArrayCB(func(iter *json.Iterator) bool {
		*ms.orig = append(*ms.orig, &otlpmetrics.ExponentialHistogramDataPoint{})
		ms.At(ms.Len() - 1).unmarshalJSONIter(iter)
		return true
	})
}

func copyOrigExponentialHistogramDataPointSlice(dest, src []*otlpmetrics.ExponentialHistogramDataPoint) []*otlpmetrics.ExponentialHistogramDataPoint {
	var newDest []*otlpmetrics.ExponentialHistogramDataPoint
	if cap(dest) < len(src) {
		newDest = make([]*otlpmetrics.ExponentialHistogramDataPoint, len(src))
		// Copy old pointers to re-use.
		copy(newDest, dest)
		// Add new pointers for missing elements from len(dest) to len(srt).
		for i := len(dest); i < len(src); i++ {
			newDest[i] = &otlpmetrics.ExponentialHistogramDataPoint{}
		}
	} else {
		newDest = dest[:len(src)]
		// Cleanup the rest of the elements so GC can free the memory.
		// This can happen when len(src) < len(dest) < cap(dest).
		for i := len(src); i < len(dest); i++ {
			dest[i] = nil
		}
		// Add new pointers for missing elements.
		// This can happen when len(dest) < len(src) < cap(dest).
		for i := len(dest); i < len(src); i++ {
			newDest[i] = &otlpmetrics.ExponentialHistogramDataPoint{}
		}
	}
	for i := range src {
		copyOrigExponentialHistogramDataPoint(newDest[i], src[i])
	}
	return newDest
}
