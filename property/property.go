package main
import "property/lib"
import ."property/lib2"//mempersingkat pemanggilan package
import "fmt"

func main(){
        lib.Sayhello()
        lib.Intro("Yoga")
        var s1 = lib.Student{"wick",24}
        fmt.Println(s1.Name)
        fmt.Println(s1.Age)
        Helo() // <-fungsi dari package lib2
        fmt.Println(Siswas.Name)
        fmt.Println(Siswas.Kelas)
}
