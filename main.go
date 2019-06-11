package main

import (
	"flag"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()
	flag.Parse()

	klog.Info("nice to meet you")
}
