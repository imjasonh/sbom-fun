// go:build windows

package main

import (
	"fmt"

	_ "gopkg.in/yaml.v3"
)

func init() { fmt.Println("windows") }
