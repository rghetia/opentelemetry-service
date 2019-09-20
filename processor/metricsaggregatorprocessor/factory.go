// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metricsaggregatorprocessor

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/config/configerror"
	"github.com/open-telemetry/opentelemetry-service/config/configmodels"
	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/processor"
)

const (
	// The value of "type" key in configuration.
	typeStr = "aggregator"
)

// Factory is the factory for batch processor.
type Factory struct {
}

// Type gets the type of the config created by this factory.
func (f *Factory) Type() string {
	return typeStr
}

// CreateDefaultConfig creates the default configuration for processor.
func (f *Factory) CreateDefaultConfig() configmodels.Processor {
	reportingInterval := defaultReportingInterval

	return &Config{
		ProcessorSettings: configmodels.ProcessorSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		Reportinginterval: &reportingInterval,
	}
}

// CreateTraceProcessor creates a trace processor based on this config.
func (f *Factory) CreateTraceProcessor(
	logger *zap.Logger,
	nextConsumer consumer.TraceConsumer,
	c configmodels.Processor,
) (processor.TraceProcessor, error) {
	return nil, configerror.ErrDataTypeIsNotSupported
}

// CreateMetricsProcessor creates a metrics processor based on this config.
func (f *Factory) CreateMetricsProcessor(
	logger *zap.Logger,
	nextConsumer consumer.MetricsConsumer,
	c configmodels.Processor,
) (processor.MetricsProcessor, error) {
	cfg := c.(*Config)

	var aggregatorOptions []Option
	if cfg.Reportinginterval != nil {
		aggregatorOptions = append(aggregatorOptions, WithReportingInterval(*cfg.Reportinginterval))
	}

	if len(cfg.DropResourceKeys) > 0 {
		aggregatorOptions = append(aggregatorOptions, WithDropResourceKeys(cfg.DropResourceKeys))
	}

	if len(cfg.DropLabelKeys) > 0 {
		aggregatorOptions = append(aggregatorOptions, WithDropLabelKeys(cfg.DropLabelKeys))
	}

	return NewAggregator(cfg.NameVal, logger, nextConsumer, aggregatorOptions...), nil
}
