// go:build ppc64le

package main

import (
	"fmt"

	_ "github.com/google/go-cmp/cmp"
)

func init() { fmt.Println("arm64") }
