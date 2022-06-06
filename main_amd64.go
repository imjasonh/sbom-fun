// go:build amd64

package main

import (
	"fmt"

	_ "github.com/google/go-containerregistry/pkg/name"
)

func init() { fmt.Println("amd64") }
