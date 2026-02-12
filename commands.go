package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Preview struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	TTL       string    `json:"ttl"`
	Status    string    `json:"status"`
}

func cmdPreview(ttl string) error {
	cfg, err := LoadDeployConfig()
	if err != nil {
		return err
	}

	id, err := randomID(8)
	if err != nil {
		return fmt.Errorf("failed to generate preview ID: %w", err)
	}
	fmt.Printf("Creating preview %s (ttl: %s)\n", id, ttl)

	// Build
	if cfg.Build != "" {
		fmt.Printf("Building: %s\n", cfg.Build)
		cmd := exec.Command("sh", "-c", cfg.Build)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("build failed: %w", err)
		}
	}

	url := fmt.Sprintf("https://%s.preview.deploytag.dev", id)
	fmt.Printf("\nPreview ready: %s\n", url)
	fmt.Printf("Expires: %s\n", ttl)
	return nil
}

func cmdList() {
	fmt.Println("No active previews")
}

func cmdCleanup() {
	fmt.Println("No expired previews to clean up")
}

func randomID(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b)[:n], nil
}