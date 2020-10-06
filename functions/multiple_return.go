package functions

import "fmt"

func vals() (int, int) {
    return 3, 7
}

//Multiplereturnfunc helps in understanding how go returns multiple values
func Multiplereturnfunc() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}