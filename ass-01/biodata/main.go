package main

import (
	"fmt"
	"os"
	"strconv"
)

type Participant struct {
	No_Absen  int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	//participants data
	var participants = []Participant{
		{No_Absen: 1, Nama: "Ariel", Alamat: "Under the sea", Pekerjaan: "Escape", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 2, Nama: "Aurora", Alamat: "Wood Castle", Pekerjaan: "Sleeping", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 3, Nama: "Cinderella", Alamat: "Village house", Pekerjaan: "Dancing", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 4, Nama: "Rapunzel", Alamat: "Abandoned Tower", Pekerjaan: "Painter", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 5, Nama: "Tatiana", Alamat: "Homie Restaurant", Pekerjaan: "Restaurant Owner", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 6, Nama: "Mulan", Alamat: "China", Pekerjaan: "Warior", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 7, Nama: "Red Riding Hood", Alamat: "Grandma House in the Wood", Pekerjaan: "Food Delivery", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 8, Nama: "Anna", Alamat: "Arendalle", Pekerjaan: "Saving Her Sister", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 9, Nama: "Belle", Alamat: "Cursed Castle", Pekerjaan: "Beast Prisoner", Alasan: "Lorem ipsum dolor sit amet"},
		{No_Absen: 10, Nama: "Vanellope Von Schweetz", Alamat: "Sugar Rush Data", Pekerjaan: "Racer", Alasan: "Lorem ipsum dolor sit amet"},
	}

	//getting argument
	var argsRaw = os.Args
	var args = argsRaw[1]
	no_absen, err := strconv.Atoi(args)
	_ = err

	getParticipant(participants, no_absen)
}

func getParticipant(p []Participant, no_absen int) {

	if no_absen <= 10 {
		fmt.Println("Data Siswa:")
		fmt.Println("Nama:", p[no_absen-1].No_Absen)
		fmt.Println("Nama:", p[no_absen-1].Nama)
		fmt.Println("Alamat:", p[no_absen-1].Alamat)
		fmt.Println("Pekerjaan:", p[no_absen-1].Pekerjaan)
		fmt.Println("Alasan:", p[no_absen-1].Alasan)
	} else {
		fmt.Println("Student not exist!")
	}
}
