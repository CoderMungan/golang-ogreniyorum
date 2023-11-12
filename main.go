package main
import "fmt"

func main(){
	fmt.Println("Hello World!")
	// Tekli yorum satırı
	/* 
		Satırları kapsayan yorum satırı!
		@CoderMungan
	*/
	// Types
	var name string
	var car string = "Mercedes"
	var price float64 = 15000.99
	var isMarried bool = true
	y := 2 // Bu değişkende ":=" operatörü sonradan değer alabileceğini gösterir. JS'de let var gibi.
	// y string = "Naber" // Ama type ilk olarak ilk olduğundan hata döndürür aynı type'de değer almalıdır.
	// ":=" Fonksiyon dışında kullanılamaz hata döndürür
	fmt.Println(car,price,isMarried,y) // Mercedes 15000.99 true 2
	name = "CoderMungan"
	fmt.Println(name) // CoderMungan Çıktısı alınıyor
}