package main

import"fmt"

type identitas struct{
	nim,umur int
	nama_panggilan,asal,jurusan string
}
type mahasiswa struct{
	id int
	nama string
	ipk float32
}

func main() {

	//MAP
	jajan:=map[string]string{
		"Jenis" : "Wafer",
		"Nama" : "Tango",
		"Rasa" : "Manis",
		"Harga" : "Mulai dari 2000",
		"Tekstur" : "Renyah",
	}
	fmt.Println(jajan)
	fmt.Println("Review jajan : ",jajan)
	fmt.Println("Makanan ringan ",jajan["Nama"], "Memiliki rasa ", jajan["Rasa"], "dan bertekstur", jajan["Tekstur"])

	//menghitung data
	fmt.Println("Jumlah data dari jajan ada sebanyak ",len(jajan))

	//Mengubah data dengan key
	jajan["Nama"]="Wafello"
	fmt.Println("Nama jajannya sekarang menjadi",jajan["Nama"])

	//menghapus data
	delete(jajan, "Harga")
	fmt.Println("Jumlah data sekarang adalah :",len(jajan))


	//struct
	user:= identitas{
		nim : 9972900,
		nama_panggilan : "Jena",
		umur : 19,
		asal : "Jakarta",
		jurusan : "Matematika",
	}

	fmt.Println(user)
	fmt.Println("Dia bernama ",user.nama_panggilan, "berasal dari", user.asal, "berumur ",user.umur)

	//struct dengan pointer
	data1 := mahasiswa {1, "Jena azzahra", 3.75}
	fmt.Println(data1)
	wisuda(&data1)
	fmt.Println(data1.nama)
}

func wisuda(Mahasiswa *mahasiswa){
	Mahasiswa.nama = Mahasiswa.nama + "S.mat"
	fmt.Println(Mahasiswa.nama)
}