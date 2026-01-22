package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type profile struct {
	URL string `json:"url"`
}

type employee struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Position string `json:"position"`
	IsActive bool   `json:"is_active"`
	Role     string `json:"-"`
	Profile  profile
}

var jsonPayload = `{
 "name": "Tomi",
 "email": "exmaple@mail.com",
 "position": "Software Developer",
 "is_active": true,
 "profile": {
	"url": "https://dummy.com/profile/1/img.png"
 }
}`

func main() {
	var e employee

	err := json.Unmarshal([]byte(jsonPayload), &e)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", e)

	bs, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
