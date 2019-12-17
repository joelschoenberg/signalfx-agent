// Code generated by monitor-code-gen. DO NOT EDIT.

package docker

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "docker-container-stats"

const (
	groupBlkio   = "blkio"
	groupCPU     = "cpu"
	groupMemory  = "memory"
	groupNetwork = "network"
)

var groupSet = map[string]bool{
	groupBlkio:   true,
	groupCPU:     true,
	groupMemory:  true,
	groupNetwork: true,
}

const (
	blkioIoMergedRecursiveAsync        = "blkio.io_merged_recursive.async"
	blkioIoMergedRecursiveRead         = "blkio.io_merged_recursive.read"
	blkioIoMergedRecursiveSync         = "blkio.io_merged_recursive.sync"
	blkioIoMergedRecursiveTotal        = "blkio.io_merged_recursive.total"
	blkioIoMergedRecursiveWrite        = "blkio.io_merged_recursive.write"
	blkioIoQueueRecursiveAsync         = "blkio.io_queue_recursive.async"
	blkioIoQueueRecursiveRead          = "blkio.io_queue_recursive.read"
	blkioIoQueueRecursiveSync          = "blkio.io_queue_recursive.sync"
	blkioIoQueueRecursiveTotal         = "blkio.io_queue_recursive.total"
	blkioIoQueueRecursiveWrite         = "blkio.io_queue_recursive.write"
	blkioIoServiceBytesRecursiveAsync  = "blkio.io_service_bytes_recursive.async"
	blkioIoServiceBytesRecursiveRead   = "blkio.io_service_bytes_recursive.read"
	blkioIoServiceBytesRecursiveSync   = "blkio.io_service_bytes_recursive.sync"
	blkioIoServiceBytesRecursiveTotal  = "blkio.io_service_bytes_recursive.total"
	blkioIoServiceBytesRecursiveWrite  = "blkio.io_service_bytes_recursive.write"
	blkioIoServiceTimeRecursiveAsync   = "blkio.io_service_time_recursive.async"
	blkioIoServiceTimeRecursiveRead    = "blkio.io_service_time_recursive.read"
	blkioIoServiceTimeRecursiveSync    = "blkio.io_service_time_recursive.sync"
	blkioIoServiceTimeRecursiveTotal   = "blkio.io_service_time_recursive.total"
	blkioIoServiceTimeRecursiveWrite   = "blkio.io_service_time_recursive.write"
	blkioIoServicedRecursiveAsync      = "blkio.io_serviced_recursive.async"
	blkioIoServicedRecursiveRead       = "blkio.io_serviced_recursive.read"
	blkioIoServicedRecursiveSync       = "blkio.io_serviced_recursive.sync"
	blkioIoServicedRecursiveTotal      = "blkio.io_serviced_recursive.total"
	blkioIoServicedRecursiveWrite      = "blkio.io_serviced_recursive.write"
	blkioIoTimeRecursiveAsync          = "blkio.io_time_recursive.async"
	blkioIoTimeRecursiveRead           = "blkio.io_time_recursive.read"
	blkioIoTimeRecursiveSync           = "blkio.io_time_recursive.sync"
	blkioIoTimeRecursiveTotal          = "blkio.io_time_recursive.total"
	blkioIoTimeRecursiveWrite          = "blkio.io_time_recursive.write"
	blkioIoWaitTimeRecursiveAsync      = "blkio.io_wait_time_recursive.async"
	blkioIoWaitTimeRecursiveRead       = "blkio.io_wait_time_recursive.read"
	blkioIoWaitTimeRecursiveSync       = "blkio.io_wait_time_recursive.sync"
	blkioIoWaitTimeRecursiveTotal      = "blkio.io_wait_time_recursive.total"
	blkioIoWaitTimeRecursiveWrite      = "blkio.io_wait_time_recursive.write"
	cpuPercent                         = "cpu.percent"
	cpuPercpuUsage                     = "cpu.percpu.usage"
	cpuThrottlingDataPeriods           = "cpu.throttling_data.periods"
	cpuThrottlingDataThrottledPeriods  = "cpu.throttling_data.throttled_periods"
	cpuThrottlingDataThrottledTime     = "cpu.throttling_data.throttled_time"
	cpuUsageKernelmode                 = "cpu.usage.kernelmode"
	cpuUsageSystem                     = "cpu.usage.system"
	cpuUsageTotal                      = "cpu.usage.total"
	cpuUsageUsermode                   = "cpu.usage.usermode"
	memoryPercent                      = "memory.percent"
	memoryStatsActiveAnon              = "memory.stats.active_anon"
	memoryStatsActiveFile              = "memory.stats.active_file"
	memoryStatsCache                   = "memory.stats.cache"
	memoryStatsDirty                   = "memory.stats.dirty"
	memoryStatsHierarchicalMemoryLimit = "memory.stats.hierarchical_memory_limit"
	memoryStatsHierarchicalMemswLimit  = "memory.stats.hierarchical_memsw_limit"
	memoryStatsInactiveAnon            = "memory.stats.inactive_anon"
	memoryStatsInactiveFile            = "memory.stats.inactive_file"
	memoryStatsMappedFile              = "memory.stats.mapped_file"
	memoryStatsPgfault                 = "memory.stats.pgfault"
	memoryStatsPgmajfault              = "memory.stats.pgmajfault"
	memoryStatsPgpgin                  = "memory.stats.pgpgin"
	memoryStatsPgpgout                 = "memory.stats.pgpgout"
	memoryStatsRss                     = "memory.stats.rss"
	memoryStatsRssHuge                 = "memory.stats.rss_huge"
	memoryStatsShmem                   = "memory.stats.shmem"
	memoryStatsSwap                    = "memory.stats.swap"
	memoryStatsTotalActiveAnon         = "memory.stats.total_active_anon"
	memoryStatsTotalActiveFile         = "memory.stats.total_active_file"
	memoryStatsTotalCache              = "memory.stats.total_cache"
	memoryStatsTotalDirty              = "memory.stats.total_dirty"
	memoryStatsTotalInactiveAnon       = "memory.stats.total_inactive_anon"
	memoryStatsTotalInactiveFile       = "memory.stats.total_inactive_file"
	memoryStatsTotalMappedFile         = "memory.stats.total_mapped_file"
	memoryStatsTotalPgfault            = "memory.stats.total_pgfault"
	memoryStatsTotalPgmajfault         = "memory.stats.total_pgmajfault"
	memoryStatsTotalPgpgin             = "memory.stats.total_pgpgin"
	memoryStatsTotalPgpgout            = "memory.stats.total_pgpgout"
	memoryStatsTotalRss                = "memory.stats.total_rss"
	memoryStatsTotalRssHuge            = "memory.stats.total_rss_huge"
	memoryStatsTotalShmem              = "memory.stats.total_shmem"
	memoryStatsTotalSwap               = "memory.stats.total_swap"
	memoryStatsTotalUnevictable        = "memory.stats.total_unevictable"
	memoryStatsTotalWriteback          = "memory.stats.total_writeback"
	memoryStatsUnevictable             = "memory.stats.unevictable"
	memoryStatsWriteback               = "memory.stats.writeback"
	memoryUsageLimit                   = "memory.usage.limit"
	memoryUsageMax                     = "memory.usage.max"
	memoryUsageTotal                   = "memory.usage.total"
	networkUsageRxBytes                = "network.usage.rx_bytes"
	networkUsageRxDropped              = "network.usage.rx_dropped"
	networkUsageRxErrors               = "network.usage.rx_errors"
	networkUsageRxPackets              = "network.usage.rx_packets"
	networkUsageTxBytes                = "network.usage.tx_bytes"
	networkUsageTxDropped              = "network.usage.tx_dropped"
	networkUsageTxErrors               = "network.usage.tx_errors"
	networkUsageTxPackets              = "network.usage.tx_packets"
)

var metricSet = map[string]monitors.MetricInfo{
	blkioIoMergedRecursiveAsync:        {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoMergedRecursiveRead:         {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoMergedRecursiveSync:         {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoMergedRecursiveTotal:        {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoMergedRecursiveWrite:        {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoQueueRecursiveAsync:         {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoQueueRecursiveRead:          {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoQueueRecursiveSync:          {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoQueueRecursiveTotal:         {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoQueueRecursiveWrite:         {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceBytesRecursiveAsync:  {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceBytesRecursiveRead:   {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceBytesRecursiveSync:   {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceBytesRecursiveTotal:  {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceBytesRecursiveWrite:  {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceTimeRecursiveAsync:   {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceTimeRecursiveRead:    {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceTimeRecursiveSync:    {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceTimeRecursiveTotal:   {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServiceTimeRecursiveWrite:   {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServicedRecursiveAsync:      {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServicedRecursiveRead:       {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServicedRecursiveSync:       {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServicedRecursiveTotal:      {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoServicedRecursiveWrite:      {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoTimeRecursiveAsync:          {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoTimeRecursiveRead:           {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoTimeRecursiveSync:           {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoTimeRecursiveTotal:          {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoTimeRecursiveWrite:          {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoWaitTimeRecursiveAsync:      {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoWaitTimeRecursiveRead:       {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoWaitTimeRecursiveSync:       {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoWaitTimeRecursiveTotal:      {Type: datapoint.Counter, Group: groupBlkio},
	blkioIoWaitTimeRecursiveWrite:      {Type: datapoint.Counter, Group: groupBlkio},
	cpuPercent:                         {Type: datapoint.Gauge, Group: groupCPU},
	cpuPercpuUsage:                     {Type: datapoint.Counter, Group: groupCPU},
	cpuThrottlingDataPeriods:           {Type: datapoint.Counter, Group: groupCPU},
	cpuThrottlingDataThrottledPeriods:  {Type: datapoint.Counter, Group: groupCPU},
	cpuThrottlingDataThrottledTime:     {Type: datapoint.Counter, Group: groupCPU},
	cpuUsageKernelmode:                 {Type: datapoint.Counter, Group: groupCPU},
	cpuUsageSystem:                     {Type: datapoint.Counter, Group: groupCPU},
	cpuUsageTotal:                      {Type: datapoint.Counter, Group: groupCPU},
	cpuUsageUsermode:                   {Type: datapoint.Counter, Group: groupCPU},
	memoryPercent:                      {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsActiveAnon:              {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsActiveFile:              {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsCache:                   {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsDirty:                   {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsHierarchicalMemoryLimit: {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsHierarchicalMemswLimit:  {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsInactiveAnon:            {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsInactiveFile:            {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsMappedFile:              {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsPgfault:                 {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsPgmajfault:              {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsPgpgin:                  {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsPgpgout:                 {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsRss:                     {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsRssHuge:                 {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsShmem:                   {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsSwap:                    {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalActiveAnon:         {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalActiveFile:         {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalCache:              {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalDirty:              {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalInactiveAnon:       {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalInactiveFile:       {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalMappedFile:         {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalPgfault:            {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsTotalPgmajfault:         {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsTotalPgpgin:             {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsTotalPgpgout:            {Type: datapoint.Counter, Group: groupMemory},
	memoryStatsTotalRss:                {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalRssHuge:            {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalShmem:              {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalSwap:               {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalUnevictable:        {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsTotalWriteback:          {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsUnevictable:             {Type: datapoint.Gauge, Group: groupMemory},
	memoryStatsWriteback:               {Type: datapoint.Gauge, Group: groupMemory},
	memoryUsageLimit:                   {Type: datapoint.Gauge, Group: groupMemory},
	memoryUsageMax:                     {Type: datapoint.Gauge, Group: groupMemory},
	memoryUsageTotal:                   {Type: datapoint.Gauge, Group: groupMemory},
	networkUsageRxBytes:                {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageRxDropped:              {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageRxErrors:               {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageRxPackets:              {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageTxBytes:                {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageTxDropped:              {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageTxErrors:               {Type: datapoint.Counter, Group: groupNetwork},
	networkUsageTxPackets:              {Type: datapoint.Counter, Group: groupNetwork},
}

var defaultMetrics = map[string]bool{
	blkioIoServiceBytesRecursiveRead:  true,
	blkioIoServiceBytesRecursiveWrite: true,
	cpuUsageSystem:                    true,
	cpuUsageTotal:                     true,
	memoryUsageLimit:                  true,
	memoryUsageTotal:                  true,
	networkUsageRxBytes:               true,
	networkUsageTxBytes:               true,
}

var groupMetricsMap = map[string][]string{
	groupBlkio: []string{
		blkioIoMergedRecursiveAsync,
		blkioIoMergedRecursiveRead,
		blkioIoMergedRecursiveSync,
		blkioIoMergedRecursiveTotal,
		blkioIoMergedRecursiveWrite,
		blkioIoQueueRecursiveAsync,
		blkioIoQueueRecursiveRead,
		blkioIoQueueRecursiveSync,
		blkioIoQueueRecursiveTotal,
		blkioIoQueueRecursiveWrite,
		blkioIoServiceBytesRecursiveAsync,
		blkioIoServiceBytesRecursiveRead,
		blkioIoServiceBytesRecursiveSync,
		blkioIoServiceBytesRecursiveTotal,
		blkioIoServiceBytesRecursiveWrite,
		blkioIoServiceTimeRecursiveAsync,
		blkioIoServiceTimeRecursiveRead,
		blkioIoServiceTimeRecursiveSync,
		blkioIoServiceTimeRecursiveTotal,
		blkioIoServiceTimeRecursiveWrite,
		blkioIoServicedRecursiveAsync,
		blkioIoServicedRecursiveRead,
		blkioIoServicedRecursiveSync,
		blkioIoServicedRecursiveTotal,
		blkioIoServicedRecursiveWrite,
		blkioIoTimeRecursiveAsync,
		blkioIoTimeRecursiveRead,
		blkioIoTimeRecursiveSync,
		blkioIoTimeRecursiveTotal,
		blkioIoTimeRecursiveWrite,
		blkioIoWaitTimeRecursiveAsync,
		blkioIoWaitTimeRecursiveRead,
		blkioIoWaitTimeRecursiveSync,
		blkioIoWaitTimeRecursiveTotal,
		blkioIoWaitTimeRecursiveWrite,
	},
	groupCPU: []string{
		cpuPercent,
		cpuPercpuUsage,
		cpuThrottlingDataPeriods,
		cpuThrottlingDataThrottledPeriods,
		cpuThrottlingDataThrottledTime,
		cpuUsageKernelmode,
		cpuUsageSystem,
		cpuUsageTotal,
		cpuUsageUsermode,
	},
	groupMemory: []string{
		memoryPercent,
		memoryStatsActiveAnon,
		memoryStatsActiveFile,
		memoryStatsCache,
		memoryStatsDirty,
		memoryStatsHierarchicalMemoryLimit,
		memoryStatsHierarchicalMemswLimit,
		memoryStatsInactiveAnon,
		memoryStatsInactiveFile,
		memoryStatsMappedFile,
		memoryStatsPgfault,
		memoryStatsPgmajfault,
		memoryStatsPgpgin,
		memoryStatsPgpgout,
		memoryStatsRss,
		memoryStatsRssHuge,
		memoryStatsShmem,
		memoryStatsSwap,
		memoryStatsTotalActiveAnon,
		memoryStatsTotalActiveFile,
		memoryStatsTotalCache,
		memoryStatsTotalDirty,
		memoryStatsTotalInactiveAnon,
		memoryStatsTotalInactiveFile,
		memoryStatsTotalMappedFile,
		memoryStatsTotalPgfault,
		memoryStatsTotalPgmajfault,
		memoryStatsTotalPgpgin,
		memoryStatsTotalPgpgout,
		memoryStatsTotalRss,
		memoryStatsTotalRssHuge,
		memoryStatsTotalShmem,
		memoryStatsTotalSwap,
		memoryStatsTotalUnevictable,
		memoryStatsTotalWriteback,
		memoryStatsUnevictable,
		memoryStatsWriteback,
		memoryUsageLimit,
		memoryUsageMax,
		memoryUsageTotal,
	},
	groupNetwork: []string{
		networkUsageRxBytes,
		networkUsageRxDropped,
		networkUsageRxErrors,
		networkUsageRxPackets,
		networkUsageTxBytes,
		networkUsageTxDropped,
		networkUsageTxErrors,
		networkUsageTxPackets,
	},
}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "docker-container-stats",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
