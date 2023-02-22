package main
import "fmt"

func main(){
        //inisialisasi slice
        var elektronik = []string{ "Smartphone","Laptop","Ac","Tv" }
        fmt.Println(elektronik[2])

        //perbedaan slice dan array
        fmt.Println(" ")
        var app1 = [2]string{"youtube","google"}//array
        var app2 = []string{"Gojek","Grab"}//slice
        var app3 = [...]string{"Chrome","Firefox"}//array
        fmt.Println(app1,app2,app3)

        //Hubungan array dan slice
        fmt.Println(" ")
        var nomor = []int{1,2,3,4} //array
        var newNomor = nomor[1:3] //slice
        var angka = nomor[:] //semua elemen akan ditampilkan
        fmt.Println(newNomor)
        fmt.Println(angka)
        
        //mengubah data elemen slice baru, yang terbentuk dari slice lama
        fmt.Println(" ")
        var pasukan = [5]string{"barbarian","archer","wizard","PEKKA"}
        /* menampilkan index 1 dan urutan array ke-4 */
        fmt.Println(pasukan[1:4])
        var newpasukan = pasukan[0:4]
        newpasukan[1]="Giant"
        fmt.Println(pasukan)
        fmt.Println(newpasukan[1])

        //penggunaan len dan cap
        fmt.Println(" ")
        fmt.Println("penggunaan fungsi len() = ",len(pasukan))
        fmt.Println("penggunaan fungsi cap() = ",cap(pasukan))

        //append()
        fmt.Println(" ")
        var nama string
        var mahasiswa = []string{"yoga","alif","khikam","aries"}
        fmt.Print("masukan nama : ")
        fmt.Scanf("%s",&nama)
        var newMahasiswa = mahasiswa[0:2]
        fmt.Println("slice asli = ",mahasiswa)
        fmt.Println("slice [0:2] = ",newMahasiswa)
        var mkeMahasiswa = append(newMahasiswa,nama)
        fmt.Println("hasil penambahan data baru : ")
        fmt.Println(mahasiswa)
        fmt.Println(newMahasiswa)
        fmt.Println(mkeMahasiswa)

        //copy()
        dst := make([]string, 3)
        src := []string{"watermelon", "pinnaple", "apple", "orange"}
        n := copy(dst, src)
        fmt.Println(dst) // watermelon pinnaple apple
        fmt.Println(src) // watermelon pinnaple apple orange
        fmt.Println(n)// 3
        //seleksi 3 index
        fmt.Println(dst[0:1:1])



}
