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
}