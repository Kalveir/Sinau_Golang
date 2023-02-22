package main
import(
        "fmt"
)

//nilai value yang bisa didefinisikan lebih awal
func persegi(sisi,lebar int) (luas int,keliling int){
        keliling = sisi * 4
        luas = sisi * lebar
        return
}
//return variabel
func penjumlahan(angka,nomer int) int{
        var hitung = angka + nomer
        return hitung
}
//return basic
func sayHello(nama string)string{
        return nama
}
//return variabel different data type & return double variable
func dataDiri(nama string,umur int) (string, int){
        return nama,umur
}
//basic function
func pembagian(){
        var ang = 34
        var nmr = 8
        var hasil = ang / nmr
        fmt.Println(hasil)
}

//variadic function
func Genap(nbr ...int) int{
        var j int
        for _,i := range nbr{
                j = i
        }
        return j

}
func main(){
        var jumlah int
        jumlah = penjumlahan(4,5)
        fmt.Println(jumlah)
        fmt.Println(penjumlahan(9,8))

        var orang string
        var nama = "Agnes"
        orang = sayHello("yoga")
        fmt.Println("hallo",orang)
        fmt.Println("hallo",sayHello(nama))

        var name string = "Yoga"
        var umur int = 20
        var data_nama,data_umur = dataDiri(name,umur)
        fmt.Println("Nama : ",data_nama)
        fmt.Println("Umur : ",data_umur)

        pembagian()
        
        var luas,keliling int
        luas,keliling = persegi(12,16)
        fmt.Println("keliling = ",keliling)
        fmt.Println("luas = ",luas)

        var number  = Genap(1,2,3,4,5,6,7,8,9)
        fmt.Println(number)
        

}
