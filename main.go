package main //Ana Paket başlangıç noktası
import "fmt"



func main(){
	fmt.Println("Hello World!")


	// Değişkenler

	// var name string
	// var email string
	// var age int

	// Farklı esnek kullanımlar
	var (
		isim, email string = "Mehmet Halil", "halilmungan@gmail.com"
		yas int = 25
	)

	fmt.Println("İsmi :",isim,"Email Adresi :",email,"Yaşı :",yas)
}

