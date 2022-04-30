package faker

import "testing"

func TestRandomLocation(t *testing.T) {
	_, err := RandomLocation()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}

func TestMultiRandomLocation(t *testing.T) {
	_, err := MultiRandomLocation(5)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}
