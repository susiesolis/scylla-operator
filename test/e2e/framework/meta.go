// Copyright (C) 2021 ScyllaDB

package framework

import (
	g "github.com/onsi/ginkgo/v2"
)

const (
	SerialLabelName            = "Serial"
	RequiresClusterIPLabelName = "RequiresClusterIP"
)

var (
	Serial = []interface{}{
		g.Serial,
		g.Label(SerialLabelName),
	}
	RequiresClusterIP = g.Label(RequiresClusterIPLabelName)
)
