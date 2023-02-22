package lib2
import f "fmt" //alias import package

var Siswas = struct{
        Name string
        Kelas int
}{}
func init(){
        Siswas.Name = "Mayangka"
        Siswas.Kelas = 6
        f.Println("-->lib/hai.go")
}
func Helo(){
        f.Println("haihai")
}
