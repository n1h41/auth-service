package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")
	if err != nil || config == nil {
		t.Fatalf("Error loading config: %s", err)
	}
}
