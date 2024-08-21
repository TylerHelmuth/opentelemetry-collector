// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package connector

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/testdata"
)

type mutatingTracesSink struct {
	*consumertest.TracesSink
}

func (mts *mutatingTracesSink) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

func TestTracesRouterMultiplexing(t *testing.T) {
	var max = 20
	for numIDs := 1; numIDs < max; numIDs++ {
		for numCons := 1; numCons < max; numCons++ {
			for numTraces := 1; numTraces < max; numTraces++ {
				t.Run(
					fmt.Sprintf("%d-ids/%d-cons/%d-logs", numIDs, numCons, numTraces),
					fuzzTraces(numIDs, numCons, numTraces),
				)
			}
		}
	}
}

func fuzzTraces(numIDs, numCons, numTraces int) func(*testing.T) {
	return func(t *testing.T) {
		allIDs := make([]component.PipelineID, 0, numCons)
		allCons := make([]consumer.Traces, 0, numCons)
		allConsMap := make(map[component.PipelineID]consumer.Traces)

		// If any consumer is mutating, the router must report mutating
		for i := 0; i < numCons; i++ {
			allIDs = append(allIDs, component.NewPipelineIDWithName("sink", strconv.Itoa(numCons)))
			// Random chance for each consumer to be mutating
			if (numCons+numTraces+i)%4 == 0 {
				allCons = append(allCons, &mutatingTracesSink{TracesSink: new(consumertest.TracesSink)})
			} else {
				allCons = append(allCons, new(consumertest.TracesSink))
			}
			allConsMap[allIDs[i]] = allCons[i]
		}

		r := NewTracesRouter(allConsMap)
		td := testdata.GenerateTraces(1)

		// Keep track of how many logs each consumer should receive.
		// This will be validated after every call to RouteTraces.
		expected := make(map[component.PipelineID]int, numCons)

		for i := 0; i < numTraces; i++ {
			// Build a random set of ids (no duplicates)
			randCons := make(map[component.PipelineID]bool, numIDs)
			for j := 0; j < numIDs; j++ {
				// This number should be pretty random and less than numCons
				conNum := (numCons + numIDs + i + j) % numCons
				randCons[allIDs[conNum]] = true
			}

			// Convert to slice, update expectations
			conIDs := make([]component.PipelineID, 0, len(randCons))
			for id := range randCons {
				conIDs = append(conIDs, id)
				expected[id]++
			}

			// Route to list of consumers
			fanout, err := r.Consumer(conIDs...)
			assert.NoError(t, err)
			assert.NoError(t, fanout.ConsumeTraces(context.Background(), td))

			// Validate expectations for all consumers
			for id := range expected {
				traces := []ptrace.Traces{}
				switch con := allConsMap[id].(type) {
				case *consumertest.TracesSink:
					traces = con.AllTraces()
				case *mutatingTracesSink:
					traces = con.AllTraces()
				}
				assert.Len(t, traces, expected[id])
				for n := 0; n < len(traces); n++ {
					assert.EqualValues(t, td, traces[n])
				}
			}
		}
	}
}

func TestTracesRouterConsumer(t *testing.T) {
	ctx := context.Background()
	td := testdata.GenerateTraces(1)

	fooID := component.NewPipelineID("foo")
	barID := component.NewPipelineID("bar")

	foo := new(consumertest.TracesSink)
	bar := new(consumertest.TracesSink)
	r := NewTracesRouter(map[component.PipelineID]consumer.Traces{fooID: foo, barID: bar})

	rcs := r.PipelineIDs()
	assert.Len(t, rcs, 2)
	assert.ElementsMatch(t, []component.PipelineID{fooID, barID}, rcs)

	assert.Len(t, foo.AllTraces(), 0)
	assert.Len(t, bar.AllTraces(), 0)

	both, err := r.Consumer(fooID, barID)
	assert.NotNil(t, both)
	assert.NoError(t, err)

	assert.NoError(t, both.ConsumeTraces(ctx, td))
	assert.Len(t, foo.AllTraces(), 1)
	assert.Len(t, bar.AllTraces(), 1)

	fooOnly, err := r.Consumer(fooID)
	assert.NotNil(t, fooOnly)
	assert.NoError(t, err)

	assert.NoError(t, fooOnly.ConsumeTraces(ctx, td))
	assert.Len(t, foo.AllTraces(), 2)
	assert.Len(t, bar.AllTraces(), 1)

	barOnly, err := r.Consumer(barID)
	assert.NotNil(t, barOnly)
	assert.NoError(t, err)

	assert.NoError(t, barOnly.ConsumeTraces(ctx, td))
	assert.Len(t, foo.AllTraces(), 2)
	assert.Len(t, bar.AllTraces(), 2)

	none, err := r.Consumer()
	assert.Nil(t, none)
	assert.Error(t, err)

	fake, err := r.Consumer(component.NewPipelineID("fake"))
	assert.Nil(t, fake)
	assert.Error(t, err)
}
