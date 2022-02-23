/**
	Nama: Panji Ahmad Nurhusni
	Tugas: Assignment 1
**/

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func createDaftarBiodata(siswa ...Biodata) []Biodata {
	var daftarBiodata = []Biodata{}

	for _, biodataSiswa := range siswa {
		daftarBiodata = append(daftarBiodata, biodataSiswa)
	}

	return daftarBiodata
}

func showBiodata(biodata Biodata) {
	fmt.Println("Nama: ", biodata.nama)
	fmt.Println("Alamat: ", biodata.alamat)
	fmt.Println("Pekerjaan: ", biodata.pekerjaan)
	fmt.Println("Alasan memilih kelas Golang: ", biodata.alasan)
}

func main() {
	siswa1 := Biodata{
		nama:      "Irvandy Hartono",
		alamat:    "Indonesia",
		pekerjaan: "Back-End Engineer",
		alasan:    "Ingin mempelajari Golang untuk back end",
	}

	siswa2 := Biodata{
		nama:      "Panji Ahmad Nurhusni",
		alamat:    "Tangerang",
		pekerjaan: "Back-End Engineer",
		alasan:    "Ingin mempelajari Golang untuk back end",
	}

	siswa3 := Biodata{
		nama:      "Arif Kurniawan",
		alamat:    "Indonesia",
		pekerjaan: "Back-End Engineer",
		alasan:    "Ingin mempelajari Golang untuk back end",
	}

	siswa4 := Biodata{
		nama:      "Fachrul Mustofa",
		alamat:    "Indonesia",
		pekerjaan: "Back-End Engineer",
		alasan:    "Ingin mempelajari Golang untuk back end",
	}

	siswa5 := Biodata{
		nama:      "James Thomas Widjaja",
		alamat:    "Indonesia",
		pekerjaan: "Back-End Engineer",
		alasan:    "Ingin mempelajari Golang untuk back end",
	}

	siswa6 := Biodata{
		nama:      "Harka Dienillah",
		alamat:    "Indonesia",
		pekerjaan: "Back-End Engineer",
		alasan:    "Ingin mempelajari Golang untuk back end",
	}

	daftarBiodata := createDaftarBiodata(siswa1, siswa2, siswa3, siswa4, siswa5, siswa6)

	getNomorAbsen := os.Args[1]
	nomorAbsen, _ := strconv.ParseInt(getNomorAbsen, 0, 64)

	showBiodata(daftarBiodata[nomorAbsen])
}
