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
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/consumer/consumerdata"
)

const (
	defaultReportingInterval           = 60 * time.Second
)

// aggregator is a component that accepts metrics, and aggregates them over ReportingInterval.
//
// aggregator implements consumer.MetricsConsumer
type aggregator struct {
	sender  consumer.MetricsConsumer
	name    string
	logger  *zap.Logger

	reportingInterval time.Duration
	dropResourceKeys []string
	dropLabelKeys []string
}

var _ consumer.MetricsConsumer = (*aggregator)(nil)

// NewAggregator creates a new aggregator that batches spans by node and resource
func NewAggregator(name string, logger *zap.Logger, sender consumer.MetricsConsumer, opts ...Option) consumer.MetricsConsumer {
	// Init with defaults
	b := &aggregator{
		name:   name,
		sender: sender,
		logger: logger,

		reportingInterval: defaultReportingInterval,
	}

	// Override with options
	for _, opt := range opts {
		opt(b)
	}

	//Start timer to export metrics
	return b
}

// ConsumeTraceData implements aggregator as a SpanProcessor and takes the provided spans and adds them to
// batches
func (b *aggregator) ConsumeMetricsData(ctx context.Context, td consumerdata.MetricsData) error {
	return nil
}
