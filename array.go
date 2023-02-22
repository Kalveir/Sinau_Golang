package main
import "fmt"

func main(){
        var names [4]string
        names[0]="Yoga"
        names[1]="Alif"
        names[2]="Khikam"
        names[3]="Aries"
        fmt.Println("menampilkan array = ",names[0],names[2], names[3])

        //inisilisasi array 
        var buah = [5]string{"apel","pepaya","nangka","durian","mangga"}
        fmt.Println("Jumlah buah = ",len(buah))
        fmt.Println("Isi semua element buah = ",buah)

        //inisialisasi array tanpa jumlah element
        var angka = [...]int{
                1,
                2,
                3,
                4,
        }
        fmt.Println("Jumlah angka=\t", len(angka))
        fmt.Println("Isi elemen array=\t",angka)

        //array multidimensi
        var angka1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
        var angka2= [2][3]int{{3, 2, 3}, {3, 4, 5}}
        fmt.Println(angka1)
        fmt.Println(angka2)

        //perulangan array menggunakan for
        for a:=0;a<len(buah);a++{
                fmt.Println("nama buah ke",a+1,"= ",buah[a])
        }
        fmt.Println(" ")
        //perulangan array menggunakan for dan range
        for _,fruit:=range buah{
                fmt.Println("nama buah = ", fruit)
        }
        fmt.Println(" ")
        //perulangan array menggunakan for dan range dengan 1 variabel
        for i:= range buah{  
                fmt.Println(buah[i])
        }
        //Alokasi Elemen Array Menggunakan Keyword make
        fmt.Println(" ")
        var kendaraan = make([]string,2)
        kendaraan[0] = "mobil"
        kendaraan[1] = "motor"
        fmt.Println(kendaraan)


}
