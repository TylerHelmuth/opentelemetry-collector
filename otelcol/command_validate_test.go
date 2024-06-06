// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/featuregate"
)

func TestValidateSubCommandNoConfig(t *testing.T) {
	cmd := newValidateSubCommand(CollectorSettings{Factories: nopFactories}, flags(featuregate.GlobalRegistry()))
	err := cmd.Execute()
	require.Error(t, err)
	require.Contains(t, err.Error(), "at least one config flag must be provided")
}

func TestValidateSubCommandInvalidComponents(t *testing.T) {
	fileLocation := filepath.Join("testdata", "otelcol-invalid-components.yaml")
	fileProvider := newFakeProvider("file", func(_ context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
		return confmap.NewRetrieved(newConfFromFile(t, fileLocation))
	})
	cmd := newValidateSubCommand(CollectorSettings{Factories: nopFactories, ConfigProviderSettings: ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:              []string{fileLocation},
			ProviderFactories: []confmap.ProviderFactory{fileProvider},
		},
	}}, flags(featuregate.GlobalRegistry()))
	err := cmd.Execute()
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown type: \"nosuchprocessor\"")
}
