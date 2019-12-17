// Code generated by monitor-code-gen. DO NOT EDIT.

package marathon

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/marathon"

var groupSet = map[string]bool{}

const (
	gaugeMarathonAppCPUAllocated               = "gauge.marathon.app.cpu.allocated"
	gaugeMarathonAppCPUAllocatedPerInstance    = "gauge.marathon.app.cpu.allocated.per.instance"
	gaugeMarathonAppDelayed                    = "gauge.marathon.app.delayed"
	gaugeMarathonAppDeploymentsTotal           = "gauge.marathon.app.deployments.total"
	gaugeMarathonAppDiskAllocated              = "gauge.marathon.app.disk.allocated"
	gaugeMarathonAppDiskAllocatedPerInstance   = "gauge.marathon.app.disk.allocated.per.instance"
	gaugeMarathonAppGpuAllocated               = "gauge.marathon.app.gpu.allocated"
	gaugeMarathonAppGpuAllocatedPerInstance    = "gauge.marathon.app.gpu.allocated.per.instance"
	gaugeMarathonAppInstancesTotal             = "gauge.marathon.app.instances.total"
	gaugeMarathonAppMemoryAllocated            = "gauge.marathon.app.memory.allocated"
	gaugeMarathonAppMemoryAllocatedPerInstance = "gauge.marathon.app.memory.allocated.per.instance"
	gaugeMarathonAppTasksRunning               = "gauge.marathon.app.tasks.running"
	gaugeMarathonAppTasksStaged                = "gauge.marathon.app.tasks.staged"
	gaugeMarathonAppTasksUnhealthy             = "gauge.marathon.app.tasks.unhealthy"
	gaugeMarathonTaskHealthchecksFailingTotal  = "gauge.marathon.task.healthchecks.failing.total"
	gaugeMarathonTaskHealthchecksPassingTotal  = "gauge.marathon.task.healthchecks.passing.total"
	gaugeMarathonTaskStagedTimeElapsed         = "gauge.marathon.task.staged.time.elapsed"
	gaugeMarathonTaskStartTimeElapsed          = "gauge.marathon.task.start.time.elapsed"
)

var metricSet = map[string]monitors.MetricInfo{
	gaugeMarathonAppCPUAllocated:               {Type: datapoint.Gauge},
	gaugeMarathonAppCPUAllocatedPerInstance:    {Type: datapoint.Gauge},
	gaugeMarathonAppDelayed:                    {Type: datapoint.Gauge},
	gaugeMarathonAppDeploymentsTotal:           {Type: datapoint.Gauge},
	gaugeMarathonAppDiskAllocated:              {Type: datapoint.Gauge},
	gaugeMarathonAppDiskAllocatedPerInstance:   {Type: datapoint.Gauge},
	gaugeMarathonAppGpuAllocated:               {Type: datapoint.Gauge},
	gaugeMarathonAppGpuAllocatedPerInstance:    {Type: datapoint.Gauge},
	gaugeMarathonAppInstancesTotal:             {Type: datapoint.Gauge},
	gaugeMarathonAppMemoryAllocated:            {Type: datapoint.Gauge},
	gaugeMarathonAppMemoryAllocatedPerInstance: {Type: datapoint.Gauge},
	gaugeMarathonAppTasksRunning:               {Type: datapoint.Gauge},
	gaugeMarathonAppTasksStaged:                {Type: datapoint.Gauge},
	gaugeMarathonAppTasksUnhealthy:             {Type: datapoint.Gauge},
	gaugeMarathonTaskHealthchecksFailingTotal:  {Type: datapoint.Gauge},
	gaugeMarathonTaskHealthchecksPassingTotal:  {Type: datapoint.Gauge},
	gaugeMarathonTaskStagedTimeElapsed:         {Type: datapoint.Gauge},
	gaugeMarathonTaskStartTimeElapsed:          {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	gaugeMarathonAppCPUAllocated:               true,
	gaugeMarathonAppCPUAllocatedPerInstance:    true,
	gaugeMarathonAppDiskAllocated:              true,
	gaugeMarathonAppDiskAllocatedPerInstance:   true,
	gaugeMarathonAppInstancesTotal:             true,
	gaugeMarathonAppMemoryAllocated:            true,
	gaugeMarathonAppMemoryAllocatedPerInstance: true,
	gaugeMarathonAppTasksRunning:               true,
	gaugeMarathonAppTasksStaged:                true,
	gaugeMarathonAppTasksUnhealthy:             true,
	gaugeMarathonTaskHealthchecksFailingTotal:  true,
	gaugeMarathonTaskHealthchecksPassingTotal:  true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "collectd/marathon",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
