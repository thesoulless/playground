// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

var (
	addr = flag.String("addr", "http://localhost:9090/noise", "address of server")

	count = flag.Int("count", math.MaxInt, "Number of requests to send")
)

// generateLoad sends count requests to the server.
func generateLoad(count int) error {
	if *addr == "" {
		return fmt.Errorf("-addr must be set to the address of the server (e.g., http://localhost:9090)")
	}

	url := *addr

	for i := 0; i < count; i++ {
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error writing request: %v", err)
		}
		if _, err := io.Copy(io.Discard, resp.Body); err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}
		resp.Body.Close()
	}

	return nil
}

func main() {
	flag.Parse()

	if err := generateLoad(*count); err != nil {
		log.Fatal(err)
	}
}
