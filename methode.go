package main
import(
        "fmt"
        "strings"
)

type student struct{
        nama string
        umur int
}

func (s student) sayHello(){
        fmt.Println("halo", s.nama)
}
func (s student) getName(i int) []string{
        return strings.Split(s.nama," ")
}

func main(){
        var siswa = student{"Reza",7}
        siswa.sayHello()
        var siswi = student{"Ririn",6}
        fmt.Printf("Nama anda : ")
        fmt.Scan(siswi)
}
