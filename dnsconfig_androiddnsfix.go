// +build androiddnsfix

package androiddnsfix

import (
	_ "net"
	"sync"
	"time"
	_ "unsafe"
)

//go:linkname defaultNs net.defaultNs
var defaultNs = []string{"8.8.8.8:53", "8.8.4.4:53"}
