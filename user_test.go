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

func TestRandomUserByNationality(t *testing.T) {
	got, err := RandomUserByNationality([]Nationality{US})
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	if got.Nationality != "US" {
		t.Errorf("Gender: got %q, wanted US", got.Nationality)
	}
}

func TestMultiRandomUserByNationality(t *testing.T) {
	got, err := MultiRandomUserByNationality([]Nationality{US}, 5)
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	for _, value := range got {
		if value.Nationality != "US" {
			t.Errorf("Gender: got %q, wanted US", value.Nationality)
		}
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

func TestRandomMaleUserByNationality(t *testing.T) {
	got, err := RandomMaleUserByNationality([]Nationality{US})
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	if got.Nationality != "US" {
		t.Errorf("Gender: got %q, wanted US", got.Nationality)
	}
	if got.Gender != "male" {
		t.Errorf("Gender: got %q, wanted male", got.Gender)
	}
}

func TestMultiRandomMaleUserByNationality(t *testing.T) {
	got, err := MultiRandomMaleUserByNationality([]Nationality{US}, 5)
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	for _, value := range got {
		if value.Nationality != "US" {
			t.Errorf("Gender: got %q, wanted US", value.Nationality)
		}
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

func TestRandomFemaleUserByNationality(t *testing.T) {
	got, err := RandomFemaleUserByNationality([]Nationality{US})
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	if got.Nationality != "US" {
		t.Errorf("Gender: got %q, wanted US", got.Nationality)
	}
	if got.Gender != "female" {
		t.Errorf("Gender: got %q, wanted female", got.Gender)
	}
}

func TestMultiRandomFemaleUserByNationality(t *testing.T) {
	got, err := MultiRandomFemaleUserByNationality([]Nationality{US}, 5)
	if err != nil {
		t.Errorf("Error raised: %q", err)
	}
	for _, value := range got {
		if value.Nationality != "US" {
			t.Errorf("Gender: got %q, wanted US", value.Nationality)
		}
		if value.Gender != "female" {
			t.Errorf("Gender: got %q, wanted female", value.Gender)
		}
	}
}
