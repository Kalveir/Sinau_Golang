package main
import "fmt"

func main(){
        var poin int64 = 83

        if poin <= 30{
                fmt.Println("Nilai anda = C")
        }else if poin <= 50{
                fmt.Println("Nilai anda = C++")
        }else if poin <= 79{
                fmt.Println("Nilai anda = B")
        }else if poin <= 85{
                fmt.Println("Nilai anda = B++")
        }else if poin <= 100 {
                fmt.Println("Nilai anda = A")
        }else {
                fmt.Println("Anda gagal")
        }
}
