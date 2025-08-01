// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestEntityRef_MoveTo(t *testing.T) {
	ms := generateTestEntityRef()
	dest := NewEntityRef()
	ms.MoveTo(dest)
	assert.Equal(t, NewEntityRef(), ms)
	assert.Equal(t, generateTestEntityRef(), dest)
	dest.MoveTo(dest)
	assert.Equal(t, generateTestEntityRef(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newEntityRef(&otlpcommon.EntityRef{}, &sharedState)) })
	assert.Panics(t, func() { newEntityRef(&otlpcommon.EntityRef{}, &sharedState).MoveTo(dest) })
}

func TestEntityRef_CopyTo(t *testing.T) {
	ms := NewEntityRef()
	orig := NewEntityRef()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestEntityRef()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newEntityRef(&otlpcommon.EntityRef{}, &sharedState)) })
}

func TestEntityRef_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestEntityRef()
	internal.MarshalJSONStreamEntityRef(internal.EntityRef(src), stream)
	require.NoError(t, stream.Error())

	// Append an unknown field at the start to ensure unknown fields are skipped
	// and the unmarshal logic continues.
	buf := stream.Buffer()
	assert.EqualValues(t, '{', buf[0])
	iter := json.BorrowIterator(append([]byte(`{"unknown": "string",`), buf[1:]...))
	defer json.ReturnIterator(iter)
	dest := NewEntityRef()
	internal.UnmarshalJSONIterEntityRef(internal.EntityRef(dest), iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestEntityRef_SchemaUrl(t *testing.T) {
	ms := NewEntityRef()
	assert.Empty(t, ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		newEntityRef(&otlpcommon.EntityRef{}, &sharedState).SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	})
}

func TestEntityRef_Type(t *testing.T) {
	ms := NewEntityRef()
	assert.Empty(t, ms.Type())
	ms.SetType("host")
	assert.Equal(t, "host", ms.Type())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newEntityRef(&otlpcommon.EntityRef{}, &sharedState).SetType("host") })
}

func TestEntityRef_IdKeys(t *testing.T) {
	ms := NewEntityRef()
	assert.Equal(t, pcommon.NewStringSlice(), ms.IdKeys())
	internal.FillTestStringSlice(internal.StringSlice(ms.IdKeys()))
	assert.Equal(t, pcommon.StringSlice(internal.GenerateTestStringSlice()), ms.IdKeys())
}

func TestEntityRef_DescriptionKeys(t *testing.T) {
	ms := NewEntityRef()
	assert.Equal(t, pcommon.NewStringSlice(), ms.DescriptionKeys())
	internal.FillTestStringSlice(internal.StringSlice(ms.DescriptionKeys()))
	assert.Equal(t, pcommon.StringSlice(internal.GenerateTestStringSlice()), ms.DescriptionKeys())
}

func generateTestEntityRef() EntityRef {
	return EntityRef(internal.GenerateTestEntityRef())
}
