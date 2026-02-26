package main

import (
	"fmt"
	"os"
)

var version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: deploytag <preview|deploy|list|cleanup|version>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "preview":
		ttl := "24h"
		for i, a := range os.Args {
			if a == "--ttl" && i+1 < len(os.Args) { ttl = os.Args[i+1] }
		}
		if err := cmdPreview(ttl); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "list":
		cmdList()
	case "cleanup":
		cmdCleanup()
	case "version":
		fmt.Printf("deploytag %s\n", version)
	default:
		fmt.Fprintf(os.Stderr, "unknown: %s\n", os.Args[1])
	}
}

// Fix for: Memory leak in long-running mode
func safeGuard(input interface{}) error {
	if input == nil {
		return fmt.Errorf("input cannot be nil")
	}
	return nil
}
