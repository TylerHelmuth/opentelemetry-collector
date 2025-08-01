// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package internal

import (
	otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

type EntityRefSlice struct {
	orig  *[]*otlpcommon.EntityRef
	state *State
}

func GetOrigEntityRefSlice(ms EntityRefSlice) *[]*otlpcommon.EntityRef {
	return ms.orig
}

func GetEntityRefSliceState(ms EntityRefSlice) *State {
	return ms.state
}

func NewEntityRefSlice(orig *[]*otlpcommon.EntityRef, state *State) EntityRefSlice {
	return EntityRefSlice{orig: orig, state: state}
}

func CopyOrigEntityRefSlice(dest, src []*otlpcommon.EntityRef) []*otlpcommon.EntityRef {
	var newDest []*otlpcommon.EntityRef
	if cap(dest) < len(src) {
		newDest = make([]*otlpcommon.EntityRef, len(src))
		// Copy old pointers to re-use.
		copy(newDest, dest)
		// Add new pointers for missing elements from len(dest) to len(srt).
		for i := len(dest); i < len(src); i++ {
			newDest[i] = &otlpcommon.EntityRef{}
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
			newDest[i] = &otlpcommon.EntityRef{}
		}
	}
	for i := range src {
		CopyOrigEntityRef(newDest[i], src[i])
	}
	return newDest
}

func GenerateTestEntityRefSlice() EntityRefSlice {
	orig := []*otlpcommon.EntityRef(nil)
	state := StateMutable
	es := NewEntityRefSlice(&orig, &state)
	FillTestEntityRefSlice(es)
	return es
}

func FillTestEntityRefSlice(es EntityRefSlice) {
	*es.orig = make([]*otlpcommon.EntityRef, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpcommon.EntityRef{}
		FillTestEntityRef(NewEntityRef((*es.orig)[i], es.state))
	}
}

// MarshalJSONStreamEntityRefSlice marshals all properties from the current struct to the destination stream.
func MarshalJSONStreamEntityRefSlice(ms EntityRefSlice, dest *json.Stream) {
	dest.WriteArrayStart()
	if len(*ms.orig) > 0 {
		MarshalJSONStreamEntityRef(NewEntityRef((*ms.orig)[0], ms.state), dest)
	}
	for i := 1; i < len((*ms.orig)); i++ {
		dest.WriteMore()
		MarshalJSONStreamEntityRef(NewEntityRef((*ms.orig)[i], ms.state), dest)
	}
	dest.WriteArrayEnd()
}

// UnmarshalJSONIterEntityRefSlice unmarshals all properties from the current struct from the source iterator.
func UnmarshalJSONIterEntityRefSlice(ms EntityRefSlice, iter *json.Iterator) {
	iter.ReadArrayCB(func(iter *json.Iterator) bool {
		*ms.orig = append(*ms.orig, &otlpcommon.EntityRef{})
		UnmarshalJSONIterEntityRef(NewEntityRef((*ms.orig)[len(*ms.orig)-1], ms.state), iter)
		return true
	})
}
