package main

import "fmt"
import "time"

func recoverWithMessage() {
	if r := recover(); r != nil {
		fmt.Println("Tiklanish", r)
	}
}

func fullName(firstName *string, lastName *string) string {
	defer recoverWithMessage()
	if firstName == nil {
		panic("ism bo'lishi mumkin emas")
	}
	if lastName == nil {
		panic("familiya bo'lishi mumkin emas")
	}
	return fmt.Sprintf("%s %s\n", *firstName, *lastName)
}

func main() {
	firstName := "G'olibbek"
	lastName := "Ashirov"
	fmt.Println(fullName(&firstName, &lastName))
	fmt.Println(fullName(nil, nil))

	t := time.Date(2022, 5, 13, 3, 23, 43, 02, time.UTC)

	loc := time.FixedZone("UTC", 6*54*44)

	res := t.In(loc)

	fmt.Printf("%v\n", res)
}
