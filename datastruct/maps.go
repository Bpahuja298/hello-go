package datastruct

import "fmt"

//Mapsfunc to understand maps
func Mapsfunc() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

	_, prs := m["k2"]        //The optional second return value when getting a value 
							 //from a map indicates if the key was present in the map.
	fmt.Println("prs:", prs) // This can be used to disambiguate between missing keys
							 // and keys with zero values like 0 or "". 
							 //Here we didn’t need the value itself, so we ignored it
							 // with the blank identifier _.
    

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}