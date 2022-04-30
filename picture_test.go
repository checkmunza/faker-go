package faker

import "testing"

func TestRandomPicture(t *testing.T) {
	_, err := RandomPicture()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}

func TestMultiRandomPicture(t *testing.T) {
	_, err := MultiRandomPicture(5)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}
