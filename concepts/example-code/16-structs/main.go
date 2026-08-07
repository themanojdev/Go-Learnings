package main

import "fmt"

// Problem Statement:
// Complete the newUser function. It takes a name and a membershipType
// ("premium" or "standard") and returns a User with an embedded Membership.
//   - "premium" membership gets a MessageCharLimit of 1000
//   - "standard" membership gets a MessageCharLimit of 100
//   - any other membershipType should return an empty (zero-value) User

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	const premium = 1000
	const standard = 100

	switch membershipType {
	case "premium":
		return User{
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: premium,
			},
			Name: name,
		}
	case "standard":
		return User{
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: standard,
			},
			Name: name,
		}
	default:
		return User{}
	}
}

func main() {
	fmt.Printf("%+v\n", newUser("Alice", "premium"))
	fmt.Printf("%+v\n", newUser("Bob", "standard"))
	fmt.Printf("%+v\n", newUser("Charlie", "gold"))
	fmt.Printf("%+v\n", newUser("Dana", ""))
}
