package faker

import "testing"

func TestRandomUser(t *testing.T) {
	_, err := RandomUser()
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
}

func TestMultiRandomUser(t *testing.T) {
	_, err := MultiRandomUser(5)
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
}

func TestRandomMaleUser(t *testing.T) {
	got, err := RandomMaleUser()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
	if got.Gender != "male" {
		t.Errorf("Gender: got %q, wanted male", got.Gender)
	}
}

func TestMultiRandomMaleUser(t *testing.T) {
	got, err := MultiRandomMaleUser(5)
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	for _, value := range got {
		if value.Gender != "male" {
			t.Errorf("Gender: got %q, wanted male", value.Gender)
		}
	}
}

func TestRandomFemaleUser(t *testing.T) {
	got, err := RandomFemaleUser()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
	if got.Gender != "female" {
		t.Errorf("Gender: got %q, wanted female", got.Gender)
	}
}

func TestMultiRandomFemaleUser(t *testing.T) {
	got, err := MultiRandomFemaleUser(5)
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	for _, value := range got {
		if value.Gender != "female" {
			t.Errorf("Gender: got %q, wanted male", value.Gender)
		}
	}
}
