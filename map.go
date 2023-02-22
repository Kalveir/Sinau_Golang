package main
import(
        "fmt"
)

func main(){
        //pengenalah
        var mahasiswa  map[string]int
        mahasiswa = map[string]int{}

        mahasiswa["yoga"] = 20
        mahasiswa["yunita"] = 19

        fmt.Println("Nama : ",mahasiswa["yoga"])        //benar
        fmt.Println("Nama : ",mahasiswa["yunita"])      //benar
        fmt.Println("Nama  : ",mahasiswa["alif"])       //salah

        //inisialiasi tunggal
        var data map[float32]int32      // <--- Wajib
        data = map[float32]int32{}      // <--- Wajib
        data[3.14] = 22
        fmt.Println(data[3.14])

        //map style JSON
        var database = map[string]int64{
                "Tepung": 12000,
                "Beras" : 15000,
                "Minyak goreng": 20000,
        }
        fmt.Println("jumlah data tadi : ",len(database))
        //itteration
        for key,val := range database{
                fmt.Println(key," : ",val)
        }
        //dellet item
        delete(database,"Minyak goreng")
        fmt.Println("jumlah data sekarang : ",len(database))
        fmt.Println(database)

        //search item 
        var google = map[string]int32{
                "Sepeda":70000,
                "Scooter":40000,
                "Skateboard":20000,
                "Sepatu Roda":100000,
        }
        var text string
        fmt.Print("Carilah : ")
        fmt.Scanf("%s",&text)
        var search,find = google[text]
        if find{
                fmt.Println(search)
        }else{
                fmt.Println("tidak ditemukan")
        }
}
