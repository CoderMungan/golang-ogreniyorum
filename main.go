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

	// Syntax Kolaylıkları Vol1
	var a, b = 6, "Coder"
	c,d := 7, "Mungan"
	fmt.Println(a,b,c,d)
	// Syntax Kolaylıkları Vol2
	var x,k,j string = "Siktir", "Yeni Dil", "Öğreniyorum"
	fmt.Println(x,k,j)
	// Syntax Kolaylıkları Vol3
	var(
		l int // Eğer bir atama yapmaz isen başlangıcı 0'dır
		i string = "L kaldı yukarıda"
		t bool = false
	)
	fmt.Println(l, i, t)

	// Printf komutu %v Value değeri verir %T type verir
	fmt.Printf("i'nin değeri : `%v` ve type: `%T` şeklindedir. \n",i,i)

	// Unsigned integers
	var z1 uint = 500
	fmt.Printf("%T and %v şeklindedir.",z1,z1)

	// Arrayler

	var dizi = [5]string{"Naber","Nasılsın","GoLang","Ogreniyorum","Mola ver"}
	fmt.Println("\n",dizi)

	numaralar := [3]int{1,2,3}
	numaralar[0] = 5
	fmt.Println(numaralar)


	
}