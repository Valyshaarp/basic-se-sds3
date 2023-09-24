package main

import (
	"fmt"
	"math"
)
//Valysha Alhamdaniya R.P.
func main() {

	//luas persegi panjang
	var panjang, lebar, luas_pp int
	fmt.Println("Masukkan panjang persegi : ")
	fmt.Scanln(&panjang)
	fmt.Println("Masukkan lebar persegi : ")
	fmt.Scanln(&lebar)
	luas_pp = panjang * lebar
	fmt.Println("Luas persegi panjang : ", luas_pp)

	//luas segitiga
	var alas, tinggi, luas_segitiga float32
	fmt.Println("Masukkan alas segitiga : ")
	fmt.Scanln(&alas)
	fmt.Println("Masukkan tinggi segitiga : ")
	fmt.Scanln(&tinggi)
	luas_segitiga = alas * tinggi / 2
	fmt.Println("Luas Segitiga : ", luas_segitiga)

	//luas lingkaran
	var r, luas_lingk, phi float32
	phi = math.Phi
	fmt.Print("Masukkan panjang jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas_lingk = phi * r * r
	fmt.Println("Luas lingkaran adalah ", luas_lingk)

	//Energi Potensial
	//Suatu benda yang bermassa 0,075 kg yang terlempar ke atas dengan ketinggian maksimum 1,4 m. Jika diketahui percepatan gravitasi adalah 10 m/s2 maka besarnya energi potensial tersebut adalah … Joule.
	var m, g, h, EP float32
	m = 0.075
	fmt.Println("massa benda : ", m, "kg")
	g = 10
	fmt.Println("percepatan gravitasi : ", g, "m/s^2")
	h = 1.4
	fmt.Println("Masukkan ketinggian benda : ", h, "m")
	EP = m * g * h
	fmt.Println("Energi Potensial benda adalah : ", EP, "Joule")

	//Energi Kinetik
	//Sebuah benda bermassa 60 kg bergerak dengan kecepatan 5 m/s. Besar energi kinetik benda tersebut adalah …. joule.
	var m2, v, EK float32
	m2 = 60
	fmt.Println("massa benda : ", m2, "kg")
	v = 5
	fmt.Println("kecepatan gravitasi : ", v, "m/s")
	EK = 0.5 * m * v * v
	fmt.Println("Energi Kinetik benda adalah : ", EK, "Joule")

	/*Frekuensi getaran
	Terdapat sebuah Bandul yang mempunyai Getaran sebanyak 60 Kali dengan waktu selama 15 detik. Maka dari itu hitunglah jumlah Frekuensi Bandul tersebut ?.
	Frekuensi = Jumlah Getaran / Waktu (satuan hertz) */
	var jumlah_getaran, waktu, Frekuensi float32
	jumlah_getaran = 60
	waktu = 15
	Frekuensi = jumlah_getaran / waktu
	fmt.Println("Jumlah getaran : ", jumlah_getaran, "kali")
	fmt.Println("Waktu : ", waktu, "detik")
	fmt.Println("Frekuensi getaran benda adalah : ", Frekuensi, "Hertz")

	//Konversi satuan suhu dari celcius ke fahrenheit, kelvin, reamur
	var Celcius, Fahrenheit, Kelvin, Reamur float32
	fmt.Println("Masukkan derajat Celcius : ")
	fmt.Scanln(&Celcius)
	fmt.Println("Suhu awal : ", Celcius, "C")
	Fahrenheit = (9/5)*Celcius + 32
	Kelvin = Celcius + 273.15
	Reamur = (4 / 5) * Celcius
	fmt.Println("Derajat Celcius ", Celcius, "C ke derajat Fahrenheit ", Fahrenheit, "F")
	fmt.Println("Derajat Celcius ", Celcius, "C ke derajat Reamur ", Reamur, "R")
	fmt.Println("Derajat Celcius ", Celcius, "C ke derajat Kelvin ", Kelvin, "F")
}
