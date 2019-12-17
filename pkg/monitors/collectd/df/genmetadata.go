// Code generated by monitor-code-gen. DO NOT EDIT.

package df

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/df"

const (
	groupInodes     = "inodes"
	groupPercentage = "percentage"
)

var groupSet = map[string]bool{
	groupInodes:     true,
	groupPercentage: true,
}

const (
	dfComplexFree         = "df_complex.free"
	dfComplexReserved     = "df_complex.reserved"
	dfComplexUsed         = "df_complex.used"
	dfInodesFree          = "df_inodes.free"
	dfInodesReserved      = "df_inodes.reserved"
	dfInodesUsed          = "df_inodes.used"
	percentBytesFree      = "percent_bytes.free"
	percentBytesReserved  = "percent_bytes.reserved"
	percentBytesUsed      = "percent_bytes.used"
	percentInodesFree     = "percent_inodes.free"
	percentInodesReserved = "percent_inodes.reserved"
	percentInodesUsed     = "percent_inodes.used"
)

var metricSet = map[string]monitors.MetricInfo{
	dfComplexFree:         {Type: datapoint.Gauge},
	dfComplexReserved:     {Type: datapoint.Gauge},
	dfComplexUsed:         {Type: datapoint.Gauge},
	dfInodesFree:          {Type: datapoint.Gauge, Group: groupInodes},
	dfInodesReserved:      {Type: datapoint.Gauge, Group: groupInodes},
	dfInodesUsed:          {Type: datapoint.Gauge, Group: groupInodes},
	percentBytesFree:      {Type: datapoint.Gauge, Group: groupPercentage},
	percentBytesReserved:  {Type: datapoint.Gauge, Group: groupPercentage},
	percentBytesUsed:      {Type: datapoint.Gauge, Group: groupPercentage},
	percentInodesFree:     {Type: datapoint.Gauge},
	percentInodesReserved: {Type: datapoint.Gauge},
	percentInodesUsed:     {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	dfComplexFree: true,
	dfComplexUsed: true,
}

var groupMetricsMap = map[string][]string{
	groupInodes: []string{
		dfInodesFree,
		dfInodesReserved,
		dfInodesUsed,
	},
	groupPercentage: []string{
		percentBytesFree,
		percentBytesReserved,
		percentBytesUsed,
	},
}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "collectd/df",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: true,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
