package faker

import "testing"

func TestRandomData(t *testing.T) {
	_, err := RandomData()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}

func TestMultiRandomData(t *testing.T) {
	_, err := MultiRandomData(5)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}
