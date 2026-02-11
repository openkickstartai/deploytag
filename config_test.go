package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	tmp := t.TempDir()
	old, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(old)

	cfg, err := LoadDeployConfig()
	if err != nil { t.Fatal(err) }
	if cfg.Provider != "docker" { t.Errorf("want docker, got %s", cfg.Provider) }
	if cfg.Port != 8080 { t.Errorf("want 8080, got %d", cfg.Port) }
	if cfg.TTL != "24h" { t.Errorf("want 24h, got %s", cfg.TTL) }
}

func TestLoadCustomConfig(t *testing.T) {
	tmp := t.TempDir()
	old, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(old)

	os.WriteFile(filepath.Join(tmp, ".deploytag.yml"), []byte("provider: static\nport: 3000\nttl: 48h\n"), 0644)
	cfg, err := LoadDeployConfig()
	if err != nil { t.Fatal(err) }
	if cfg.Provider != "static" { t.Errorf("want static, got %s", cfg.Provider) }
	if cfg.Port != 3000 { t.Errorf("want 3000, got %d", cfg.Port) }
}

func TestRandomID(t *testing.T) {
	a := randomID(8)
	b := randomID(8)
	if a == b { t.Error("IDs should be unique") }
	if len(a) != 8 { t.Errorf("want len 8, got %d", len(a)) }
}
