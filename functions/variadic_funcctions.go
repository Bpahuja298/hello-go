package functions

import "fmt"

func sum(nums ...int) {              //syntax
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

//Variablefunc helps to understand how variable functions works in golang
func Variablefunc() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}