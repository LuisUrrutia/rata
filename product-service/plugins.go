package main

import (
	_ "github.com/micro/go-micro/registry/mdns"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/kubernetes"
)
