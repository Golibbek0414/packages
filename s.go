package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

type User struct {
    Abbos string
    Age     string
}

func createUserList(data [][]string) []User {
    var UserList []User
    for i, line := range data {
        if i > 0 { 
            var rec User
            for j, field := range line {
                if j == 0 {
                    rec.Abbos = field
                } else if j == 1 {
                    rec.Age = field
                }
            }
            UserList = append(UserList, rec)
        }
    }
    return UserList
}

func main() {
    
    f, err := os.Open("data.csv")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

   
    UserList := createUserList(data)

    fmt.Printf("%+v\n", UserList)
	data := UserList{}
 
	_ = json.Unmarshal([]byte(file), &UserList)
 
	for i := 0; i < len(data.UserList); i++ {
		fmt.Println("Name: ", data.UserList[i].Name)
		fmt.Println("Age: ", data.UserList[i].Age)
	}
    
}