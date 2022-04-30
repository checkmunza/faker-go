package faker

import "testing"

func TestRandomLogin(t *testing.T) {
	_, err := RandomLogin()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}

func TestMultiRandomLogin(t *testing.T) {
	_, err := MultiRandomLogin(5)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
}
