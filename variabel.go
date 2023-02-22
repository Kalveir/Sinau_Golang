package main
import "fmt"
func main(){
        var nama = "Yoga"
        var akhir string
        akhir = "krisna"
        var umur int32 = 20
        fmt.Println(nama,akhir,"umur =",umur)
        fmt.Printf("Selamat datang kembali tuan \n")
        
        /*
        %s digunakan untuk mengganti value string di akhir
        */
        fmt.Printf("hallo %s %s!\n",nama,akhir)
        fmt.Println("halo",nama,akhir+"!")

        /*
        multi variabel
        */
        var first,second,third int32
        first,second,third = 1,2,3
        fmt.Println(first,second,third)

        var pertama,kedua,ketiga string = "satu","dua","tiga"
        fmt.Println(pertama,kedua,ketiga)

        anjay,anjir,anjrit := 1,2,3
        fmt.Println( anjay,anjir,anjrit )

        /*
        variabel undescore
        */
        name, _:="yoga","krisna"
        fmt.Println(name)

}
