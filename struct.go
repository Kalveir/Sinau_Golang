package main
import "fmt"

//tag property pada struct
type riwayat struct{
        alamat string "json: alamat"
        lahir int "json: lahir"
}

type person struct{
        age int
        name string
}

type data struct{
        nama string
        npm int
        kelas string
        semester int
}
//embeded struct
type info struct{
        alamat string
        umur int
        data
}

func main(){
        //deklarasi struct
        var mahasiswa data
        mahasiswa.nama = "Yoga"
        var maba = data{nama:"Dani",npm:2055201001044,kelas:"pagi"}
        //output
        fmt.Println(mahasiswa.nama)
        fmt.Println(maba.nama," ",maba.npm)
        //pointer
        var anggota *data = &maba       
        anggota.nama = "William" 
        fmt.Println(&maba.nama)
        fmt.Println(anggota.nama)       
        //emmbedded struct
        var senat = info{}
        senat.nama = "Dimas"
        senat.npm = 2055201001044
        senat.alamat = "Lekok"
        senat.umur = 24
        fmt.Println(senat.data.nama,senat.data.npm,senat.alamat)
        //kombinasi slice dan struct
        var mhs = []info{
                {alamat : "Sumberanyar",umur: 20},
                {alamat : "Sedarum ",umur:24},
                {alamat : "Grati", umur:23},
        }
        for _,j := range mhs{
                fmt.Println(j.alamat,j.umur)
        }
         //type alias di struct
        type orang = person
        var manusia = orang{20,"Krisna"}
        fmt.Println(manusia.name, manusia.age)
}
