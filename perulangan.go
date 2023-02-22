package main
import "fmt"

func main(){
        var angka,number int32 = 5,7

        //perulangan for pada umumnya
        for nomor := 1;nomor<10;nomor++{
                fmt.Print(" ",nomor)
        }
        fmt.Println(" ")
        //perulangan for hanya dengan kondisi
        for angka < 20{
                fmt.Print(" ",angka)
                angka++
        }
        fmt.Println(" ")
        //perulangan for tanpa argumen
        for {
                fmt.Print(" ",number)
                number++
                if number > 30{
                        break
                }
        }
        fmt.Println(" ")
}

