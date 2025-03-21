package config

import (
	"testing"
)

func TestGetVersionFromAtlasAddress(t *testing.T) {
	t.Parallel()

	v10 := "1.0"
	v11 := "1.1"
	v12 := "1.2"
	v13 := "1.3"
	v15 := "1.5"

	if ok, err := IsVersionAtLeast(&v10, &v10); err != nil || !ok {
		t.Fatalf("failed to get version from atlas address: %v", err)
	}

	if ok, err := IsVersionAtLeast(&v13, &v15); err != nil || ok {
		t.Fatalf("failed to get version from atlas address: %v", err)
	}

	if ok, err := IsVersionAtLeast(&v12, &v15); err != nil || ok {
		t.Fatalf("failed to get version from atlas address: %v", err)
	}

	if ok, err := IsVersionAtLeast(&v11, &v15); err != nil || ok {
		t.Fatalf("failed to get version from atlas address: %v", err)
	}
}
