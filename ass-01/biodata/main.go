package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// )

// type Participant struct {
// 	Nama      string `json:"Nama"`
// 	Alamat    string `json:"Alamat"`
// 	Pekerjaan string `json:"Pekerjaan"`
// 	Alasan    string `json:"Alasan"`
// }

// func main() {
// 	fmt.Println("hello")
// 	var argsRaw = os.Args
// 	var argsSlice = argsRaw[1:]
// 	var args = argsSlice[0]
// 	no_absen, err := strconv.Atoi(args)
// 	fmt.Println(no_absen)

// 	content, err := ioutil.ReadFile("participant.json")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	var participants []Participant

// 	err2 := json.Unmarshal(content, &participants)
// 	if err2 != nil {
// 		fmt.Println("Error JSON Unmarshalling")
// 		fmt.Println(err2.Error())
// 		return
// 	}
// 	// getParticipant(participants, no_absen)
// }

// func getParticipant(no_absen int) {

// }
