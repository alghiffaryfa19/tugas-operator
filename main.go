package main

import (
	"fmt"
)

func rupiahToDollar(money float64) string {
	konversi := (money/14251.55)
	return fmt.Sprintf("USD %.2f",konversi)
}

func rupiahToEuro(money float64) string {
	konversi := (money/16846.72)
	return fmt.Sprintf("%.2f EURO",konversi)
}

func bonus(gbp int) string {
	var knut int
	var sickle int
	var si int
	var galleon int
	var galleons int
	knut = gbp*100
	sickle = (knut%29)
	si = knut/29
	galleon = (si/17)
	galleons = (si%17)
	
	return fmt.Sprintf(`Jumlah Knut yang kamu dapat : %d
	Hasil penukaran yang didapat : %d galleon(s)
	Sisa ditukar menjadi : %d sickle(s)
	Keping knut yang tersisa %d`, knut,galleon,galleons,sickle)
	//return fmt.Sprintf("%d",galleon)
}

func main() {
	var i int
	var nominal float64

	var gbp int

	fmt.Printf("Pilih:\n")
	fmt.Printf("1. Rupiah to Dollar:\n")
	fmt.Printf("2. Rupiah to Euro :\n")
	fmt.Printf("3. Bonus :\n")
	fmt.Scanf("%d",&i)

	switch i{
		case 1:
			fmt.Print("Tulis nominal : ")
			fmt.Scan(&nominal)		
			fmt.Println(rupiahToDollar(nominal))
		case 2:
			fmt.Print("Tulis nominal : ")
			fmt.Scan(&nominal)
			fmt.Println(rupiahToEuro(nominal))
		case 3:
			fmt.Print("GBP : ")
			fmt.Scan(&gbp)
			fmt.Println(bonus(gbp))
		default:fmt.Println("Tidak ada")
	}
}

