package main

import (
	"fmt"
	"os"
)

var version = "0.1.0"

func printUsage() {
	fmt.Println(`deploytag - Zero-config deploy previews from git tags

Usage:
  deploytag <command> [flags]

Commands:
  preview   Create a deploy preview from the current branch
  deploy    Deploy a specific tag (e.g. v1.2.3-preview)
  list      List active preview deployments
  cleanup   Remove expired preview deployments
  version   Print the deploytag version

Flags:
  --ttl <duration>   Lifetime of the preview (default: 24h)
  --help, -h         Show this help message`)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "preview":
		ttl := "24h"
		for i, a := range os.Args {
			if a == "--ttl" && i+1 < len(os.Args) {
				ttl = os.Args[i+1]
			}
		}
		if err := cmdPreview(ttl); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "deploy":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: deploytag deploy <tag>")
			os.Exit(1)
		}
		if err := cmdDeploy(os.Args[2]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "list":
		cmdList()
	case "cleanup":
		cmdCleanup()
	case "version":
		fmt.Printf("deploytag %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}
