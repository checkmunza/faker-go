package faker

import (
	"time"
)

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type DOB struct {
	Date time.Time `json:"date"`
	Age  uint      `json:"age"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type User struct {
	Gender      string  `json:"gender"`
	Name        Name    `json:"name"`
	Email       string  `json:"email"`
	DOB         DOB     `json:"dob"`
	Phone       string  `json:"phone"`
	Cell        string  `json:"cell"`
	Picture     Picture `json:"picture"`
	Nationality string  `json:"nat"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Location struct {
	Street      string      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Postcode    string      `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

type Data struct {
	Gender      string   `json:"gender"`
	Name        Name     `json:"name"`
	Location    Location `json:"location"`
	Email       string   `json:"email"`
	Login       Login    `json:"login"`
	DOB         DOB      `json:"dob"`
	Picture     Picture  `json:"picture"`
	Nationality string   `json:"nat"`
}
