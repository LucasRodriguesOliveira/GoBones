package server

import (
	"context"
	"fmt"
	"net"
	"time"
)

func isPortAvailable(port string) bool {
	var d net.Dialer
  addres := fmt.Sprintf("localhost:%s", port)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
  
	conn, err := d.DialContext(ctx, "tcp", addres)

  if conn != nil {
    defer conn.Close()
  }

	return err != nil
}
