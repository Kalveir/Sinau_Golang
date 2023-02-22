package main
import(
        "fmt"
)

func kalkulasi(nmr ...int)float64{
        var total int = 0
        for _,hitung := range nmr{
                total+=hitung
        }
        var rata = float64(total) / float64(len(nmr))
        return rata
}

func main(){
        var nomor = []int{2,4,6,8,10,12,14,16,18,20}
        var rata = kalkulasi(nomor...)
        fmt.Println(rata)
}
