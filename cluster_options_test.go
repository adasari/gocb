package gocb

import "testing"

func TestClusterFromOptionsTCPNoDelay(t *testing.T) {
	cluster := clusterFromOptions(ClusterOptions{
		IoConfig: IoConfig{
			EnableTCPNoDelay: true,
		},
		Tracer: &NoopTracer{},
		Meter:  &NoopMeter{},
	})

	if !cluster.enableTCPNoDelay {
		t.Fatal("TCP_NODELAY is disabled")
	}
}
