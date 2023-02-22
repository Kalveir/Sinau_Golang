package main
import "fmt"

func main(){
        var numberA = 4
        var numberB *int = &numberA

        fmt.Println("Number A (value) : ", numberA)
        fmt.Println("Number A (address) : ", &numberA)
        fmt.Println("Number B (value) : ", *numberB)
        fmt.Println("Number B (address) : ",numberB)

        fmt.Println(" ")
        numberA = 5
        fmt.Println("Number A (value) : ", numberA)
        fmt.Println("Number A (address) : ", &numberA)
        fmt.Println("Number B (value) : ", *numberB)
        fmt.Println("Number B (address) : ",numberB)


}

