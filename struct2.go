package main

import "fmt"

func main() {
	user1 := User{
		ID:        1,
		FirstName: "Fahrettin",
		LastName:  "Emre",
		UserName:  "fahrettinemre",
		Age:       29,
		Pay:       &Payment{Salary: 3000, Bonus: 750},
	}
	fmt.Println(user1)
	fmt.Println(user1.Pay.Bonus)
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetPayment())

	user2 := new(User)
	user2.FirstName = "Emre"
	fmt.Println(user2.FirstName)

	user3 := NewUser()
	user3.Pay.Bonus = 780
	fmt.Println(user3.Pay.Bonus)
	user3.Pay = &Payment{Salary: 900, Bonus: 500}
	fmt.Println(user3.Pay.Bonus)

}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

type Payment struct {
	Salary float64
	Bonus  float64
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
func (u User) GetUserName() string {
	return u.UserName
}
func (u User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
