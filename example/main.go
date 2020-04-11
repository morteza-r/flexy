package main

import (
	"fmt"
	"github.com/morteza-r/flexy"
	"time"
)

func main() {
	client := flexy.NewClient("http://127.0.0.1:9888")
	for i := float64(1); i <= 5; i++ {
		addExample(client, i)
	}
	getExample(client, 2)
	multiGetExample(client)
}

func addExample(client *flexy.Client, i float64) {
	start := time.Now()
	user := User{
		Id:     i,
		Name:   "morteza",
		Family: RandStringBytes(10),
		Active: false,
		Account: Account{
			Name: RandStringBytes(5),
			Credit: 1200,
		},
	}

	err := client.Query().
		Table("users").
		Model(&user).
		Add()
	elapsed := time.Since(start)

	fmt.Println("addExample took", elapsed)
	if err != nil {
		fmt.Println(err)
		return
	}
	//jsonPrint(user)
}

func getExample(client *flexy.Client, id float64) {
	user := User{
		Id: id,
	}
	err := client.Query().
		Table("users").
		Model(&user).
		Get()

	if err != nil {
		fmt.Println(err)
		return
	}
	jsonPrint(user)
}

func multiGetExample(client *flexy.Client) {
	var users []User
	err := client.Query().
		Table("users").
		Model(&users).
		OrderString("family").
		Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonPrint(users)
}

func updateExample(client *flexy.Client) {
	tempUser := User{
		Id: 1,
	}
	start := time.Now()
	err := client.Query().
		Table("users").
		Model(&tempUser).
		Update()
	elapsed := time.Since(start)
	fmt.Println("updateExample took", elapsed)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonPrint(tempUser)
}
