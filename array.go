package main
import "fmt"



func main(){

	// tipleri
	/*
		var değişken adı = operatörü + [length]type{elemanları}
		eğer ki shorthand ile declarations yapacaksan da
		değişkent adı := operatörü [...]type{elemanları}
	*/

	var num = [...]int{1,2,3,4,5,6,7}
	num2 := [3]string{"Mehmet","Halil","MUNGAN"} //Shorthand declaration method


	fmt.Println(num,num2)

	// Array işlemler

	arr := [5]int{10,20,30,40,50}
	arr[3] = 42
	fmt.Println(arr)

	// Bir kaç örnek

	stringArr := [...]string{"GoLang Öğreniyorum","Öğrendiklerimi Paylaşıyorum","Ali Ata Bak"}
	fmt.Println("Öncesi; ", stringArr)

	cliFor := stringArr

	cliFor[2] = "Emel eve gel"

	fmt.Println("stringArr: ", stringArr)
	fmt.Println("cliFor: ", cliFor)
}