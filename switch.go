package main
import "fmt"

func main(){
        var nilai float32
        var poling int64
        var nama string
        fmt.Print("Masukan nama anda : ")
        fmt.Scanf("%s",&nama)

        fmt.Print("masukan nilai UTS anda : ")
        fmt.Scanf("%f",&nilai)

        fmt.Print("masukan nilai Ujian : ")
        fmt.Scanf("%d",&poling)

        fmt.Println("Nama : ",nama)

        switch{
                //switch case dengan kondisi
                case (nilai <= 40) || (nilai <= 50):
                        fmt.Println("ikut Remidi")
                case (nilai >= 50) && (nilai <=100):
                        fmt.Println("UTS lulus")
                default:
                        fmt.Println("UTS gagal")
        }
        switch{
                //fallthrough
                case (poling <= 35):
                        //jika kondisi benar, Skiip
                        fallthrough
                case (poling <= 45):
                        fmt.Println("Nilai : C")
                case (poling <=60 ):
                        fmt.Println("Nilai : C++")
                case (poling <=79):
                        fmt.Println("Nilai : B")
                case (poling <=85):
                        fmt.Println("B++")
                case (poling <=100):
                        fmt.Println("A")
                default:
                        fmt.Println("Anda tidak lulus ujian")
        }
}
