// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-config"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var configFile string

	flag.StringVar(&configFile, "config", "config.dist.yml", "config")
	flag.Parse()

	err := config.LoadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("PORT", strconv.Itoa(config.Get("app", "port").Int(8080)))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
