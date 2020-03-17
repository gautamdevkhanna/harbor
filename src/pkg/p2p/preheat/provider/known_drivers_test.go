package provider

import "testing"

func TestListProviders(t *testing.T) {
	metadata, err := ListProviders()
	if err != nil {
		t.Fatal(err)
	}
	if len(metadata) != len(knownDrivers) {
		t.Errorf("expect %d provider but got %d", len(knownDrivers), len(metadata))
	}
	m := metadata[0]
	if m.ID != "dragonfly" {
		t.Errorf("expect dragonfly provider but got %s", m.ID)
	}
}

func TestGetProvider(t *testing.T) {
	f, ok := GetProvider("dragonfly")
	if !ok {
		t.Fatal("expect dragonfly provider existing but not")
	}

	_, err := f(nil)
	if err != nil {
		t.Error(err)
	}
}
