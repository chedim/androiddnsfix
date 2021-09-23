// +build androiddnsfix

package androiddnsfix

import (
	"net"
  "context"
  "time"
)

func init() {
  net.DefaultResolver = &net.Resolver{
      PreferGo: true,
      Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
          d := net.Dialer{
              Timeout: time.Millisecond * time.Duration(3000),
          }
          return d.DialContext(ctx, "udp", "8.8.8.8:53")
      },
  }
}
