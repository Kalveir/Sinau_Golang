package main
import "fmt"

func main(){
        //deklarasi
        var hasil,angka,nomer float32
        //input
        angka  = 26
        nomer  = 56

        //penjumlahan
        hasil = angka + nomer
        fmt.Println("hasil penjumlahan = ",hasil)
        //pengurangan
        hasil = nomer - angka
        fmt.Println("hasil pengurangan = ",hasil)
        //perkalian
        hasil = angka * nomer
        fmt.Println("hasil pembagian = ",hasil)
        //pembagian
        hasil  = nomer/angka
        fmt.Println("hasil pembagian = ",hasil)
}
