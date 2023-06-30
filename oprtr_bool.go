package main

import "fmt"

func main() {
	nilaiUjian := 88
	absensi := 95

	var lulusUjian bool = nilaiUjian > 80
	var lulusAbsen bool = absensi > 90

	var lulus = lulusAbsen && lulusUjian

	fmt.Println("lulus ujian = ", lulusUjian)
	fmt.Println("lulus Absen = ", lulusAbsen)
	fmt.Println("lulus = ", lulus)
}
