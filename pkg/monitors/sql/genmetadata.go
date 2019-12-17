// Code generated by monitor-code-gen. DO NOT EDIT.

package sql

import (
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "sql"

var groupSet = map[string]bool{}

var metricSet = map[string]monitors.MetricInfo{}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "sql",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           true,
}
