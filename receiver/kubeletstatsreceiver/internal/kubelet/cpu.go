// Copyright 2020, OpenTelemetry Authors
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

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

func addCPUMetrics(dest pmetric.MetricSlice, cpuMetrics metadata.CPUMetrics, s *stats.CPUStats, startTime pcommon.Timestamp, currentTime pcommon.Timestamp) {
	if s == nil {
		return
	}
	addCPUUsageMetric(dest, cpuMetrics.Utilization, s, currentTime)
	addCPUTimeMetric(dest, cpuMetrics.Time, s, startTime, currentTime)
}

func addCPUUsageMetric(dest pmetric.MetricSlice, metricInt metadata.MetricIntf, s *stats.CPUStats, currentTime pcommon.Timestamp) {
	if s.UsageNanoCores == nil {
		return
	}
	value := float64(*s.UsageNanoCores) / 1_000_000_000
	fillDoubleGauge(dest.AppendEmpty(), metricInt, value, currentTime)
}

func addCPUTimeMetric(dest pmetric.MetricSlice, metricInt metadata.MetricIntf, s *stats.CPUStats, startTime pcommon.Timestamp, currentTime pcommon.Timestamp) {
	if s.UsageCoreNanoSeconds == nil {
		return
	}
	value := float64(*s.UsageCoreNanoSeconds) / 1_000_000_000
	fillDoubleSum(dest.AppendEmpty(), metricInt, value, startTime, currentTime)
}
