package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(inputName string) Profile {
	var NewProfile Profile
	NewProfile.Name = inputName

	if inputName == "Sasuke" {
		NewProfile.Health = 200
		NewProfile.Power = 100
		NewProfile.Exp = 0
	}

	if inputName == "Goku" {
		NewProfile.Health = 400
		NewProfile.Power = 300
		NewProfile.Exp = 100
	}

	if inputName == "Naruto" {
		NewProfile.Health = 150
		NewProfile.Power = 200
		NewProfile.Exp = 50
	}
	return NewProfile
}

func PowerUp(profile Profile, multiplier int) Profile {
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)

	return profile

}

func main() {
	Newprofile := MakeProfile("Naruto")
	fmt.Println("Name: ", Newprofile.Name)
	fmt.Println("Health: ", Newprofile.Health)
	fmt.Println("Power: ", Newprofile.Power)
	fmt.Println("Exp: ", Newprofile.Exp)

	fmt.Println("==============Heroes Power Up =============")

	powerUp := PowerUp(Newprofile, 4)
	fmt.Println("Name: ", powerUp.Name)
	fmt.Println("Health: ", powerUp.Health)
	fmt.Println("Power: ", powerUp.Power)
	fmt.Println("Exp: ", powerUp.Exp)

}
