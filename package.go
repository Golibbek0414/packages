package main
import "fmt"
func main() {
    var n,a int
    fmt.Scan(&n)
    fmt.Scan(&a)
    var b[]int
    b=append(n)
    var c[]int
    for i:=0;i<b;i++{
        if b[i]==a{
            continue
        }
        c=append(b[i])
    }
    fmt.Println(c)
}

   



// package main
// import( "fmt"
//     "os"
//     "encoding/json"
// )

// type myStruct struct {
	
// 	Name string `json:"name"`
	
	
// 	Age int `json:",omitempty"`
	
// 	Status bool `json:"-"`
// }

// func main() {
//     m := myStruct{Name: "John Connor", Age: 0, Status: true}

//     data, err := json.Marshal(m)
//     if err != nil {
// 	    fmt.Println(err)
// 	    return
//     }
//     fmt.Printf("%s", data) 
// }


