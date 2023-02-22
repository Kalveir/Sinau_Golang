package main
import(
        "fmt"
)
type human struct {
        nama string
        umur int
} 
func main(){
        //anonymous struct tanpa pengisian property
        var orang = struct{
                human
                pekerjaan string
        }{}
        orang.human = human{"Zaky",21}
        orang.pekerjaan = "Mahasiswa"
        var personn = orang.pekerjaan
        fmt.Println(orang.human.nama," ",orang.human.umur)
        fmt.Println(personn)
        //anonymous struct dengan pengisian property
        var manusia = struct{
                human
                pekerjaan string
        }{
        human : human{"Luluk",20},
        pekerjaan : "pedagang",
        }
        fmt.Println(manusia.human.nama,"",manusia.human.umur)
        fmt.Println(manusia.pekerjaan)
        //inisialisasi anonymous struct dengan slice
        var dospem = []struct{
                human
                pekerjaan string
        }{
                {human:human{"Dian",28},pekerjaan:"Dosen"},
                {human:human{"Nanda",27},pekerjaan:"Dosen"},
                {human:human{"Anang",32},pekerjaan:"Dosen"},
                {human:human{"Zoqi",28},pekerjaan:"Dosen"},
        }
        for _,i:=range dospem{
                fmt.Println(i)
        }
}
