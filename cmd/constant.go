package cmd

import (
	"time"
)

const (
	intervalCpuWorkerCheckContextDone = 10000
	durationMemoryWorkerDoRefresh     = 5 * time.Minute
	durationEachSignCheck             = 100 * time.Millisecond
)
