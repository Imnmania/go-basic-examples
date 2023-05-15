package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	/// --------------------
	/// In golang there are:
	/// * no inheritence
	/// * no super or parent
	/// * no child
	/// --------------------

	niloy := User{"Niloy", "n@b.com", true, 27}
	fmt.Println(niloy)
	fmt.Printf("Niloy details are: %+v \n", niloy)
	fmt.Printf("Name is %v and email is %v \n", niloy.Name, niloy.Email)

	niloy.GetStatus()
	niloy.NewMail()
	fmt.Printf("Name is %v and email is %v \n", niloy.Name, niloy.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active => ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is => ", u.Email)
}
