package main

import (
    "encoding/json"
    "fmt"

    // "github.com/fatih/structs"
)

type Test01 interface{}

type test1 struct {
    User        string `json:"user"`
    Success     bool   `json:"success"`
    Continent   string `json:"continent"`
    Country     string `json:"country"`
    CountryCode string `json:"country_code"`
}

type test2 struct {
    User        string `json:"user"`
    CountryCode string `json:"country_code"`
    CountryName string `json:"country_name"`
    Status      string `json:"status"`
}

type test3 struct {
    User        string `json:"user"`
    CountryCode string `json:"country_code"`
    CountryName string `json:"country_name"`
    ZipCode     string `json:"zip_code"`
}

func test() string {

    str1 := `{
   "user" : "vasilii",
   "Success" : "true",
   "continent" : "Europe",
   "country" : "Russia",
   "countryCode" : "RU",
}`

    str2 := `{
   "user" : "vasilii",
   "country_name" : "Russia",
   "countryCode" : "RU",
   "status" : "success"
}`
    str3 := `{
   "user" : "vasilii",
   "country_name" : "Russia",
   "countryCode" : "RU",
   "zip_code" : "4564646"
}`

    allInfo := ""
    rs := []string{str1, str2, str3}

    structsArr := []Test01{test1{}, test2{}, test3{}}
    i := 1

    for _, info := range structsArr {
        jsons := rs[i]

        bytes := []byte(jsons)

        json.Unmarshal(bytes, &info)

        s := structs.New(info)
        for _, f := range s.Fields() {
            if f.IsExported() {
                allInfo += fmt.Sprintf(" %+v : %+v \n", f.Name(), f.Value())
            }

        }
        i++
    }
    return allInfo
}

func main() {

    fmt.Println(test())
}