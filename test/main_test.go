package main

import (
	"code.byted.org/gopkg/metrics"
	"testing"
	"time"
)

func TestEmitToSpecificHost(t *testing.T) {
	cli := metrics.NewDefaultMetricsClientV2("p.s.m", true)
	for i := 0; i < 10; i++ {
		_ = cli.EmitTimer("timer.test", 1)
		time.Sleep(2 * time.Second)
		println("执行了一次")
	}
	cli.Flush()
}
