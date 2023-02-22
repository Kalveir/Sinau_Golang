package lib
import "fmt"

type Student struct{
        Name string
        Age int
}

func Sayhello(){
        fmt.Println("Halo user")
}
func Intro(nama string){
        fmt.Println("selamat datang",nama)
}
