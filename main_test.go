// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/nbio/st"
	"os"
	"testing"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/goplg", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.dist.yml")

	err := config.LoadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}
}

// TestMain test cases
func TestMain(t *testing.T) {
	st.Expect(t, 8000, config.Get("app", "port").Int(8080))
}
