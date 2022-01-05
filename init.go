package main

import (
	// "crypto/sha256"
	"encoding/json"
	"fmt"
	// "time"
)

type User struct {
	Nama string `json:"Nama"`
	Umur string `json:"Umur"`
}

func main() {
	// Object
	data := `{"Nama":"Eko","Umur":"20"}`
	jsonData := []byte(data)
	var dataUser User

	err := json.Unmarshal(jsonData, &dataUser)

	if err != nil {
		panic("Error")
	}

	fmt.Println("Nama: ", dataUser.Nama)
	fmt.Println("Umur: ", dataUser.Umur) 

	// Object dalam array
	data2 := `[
		{"Nama":"Eko","Umur":"20"},
		{"Nama":"Kevin","Umur":"22"}
		]`
	jsonData2 := []byte(data2)
	var dataUser2 []User

	err2 := json.Unmarshal(jsonData2, &dataUser2)

	if err2 != nil {
		panic("Error")
	}
	for _, i := range dataUser2 {
		fmt.Println("Nama: ", i.Nama)
		fmt.Println("Umur: ", i.Umur)
	} 

}
