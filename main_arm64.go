// go:build arm64

package main

import (
	"fmt"

	_ "github.com/google/go-github/v45/github"
)

func init() { fmt.Println("arm64") }
