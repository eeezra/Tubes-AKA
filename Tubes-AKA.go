package main

import "fmt"

const MAXX int = 500       //jml array maksimum buat diisi data
const MAX_YEARS int = 1000 //array maksimum untuk menghitung jml kegiatan pertahunnya

type Kegiatan struct {
	ketua         string
	anggota       [4]string
	prodiFakultas string
	judul         string
	sumberDana    string
	luaran        string
	tahun         int
}

type tabKegiatan [MAXX]Kegiatan

var dataPenelitian tabKegiatan
var dataAbdimas tabKegiatan
var countPenelitian int
var countAbdimas int

func main() {
	var pilihan int
	for {
		menu()
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			addData()
		} else if pilihan == 2 {
			editData()
		} else if pilihan == 3 {
			deleteData()
		} else if pilihan == 4 {
			printData()
		} else if pilihan == 5 {
			sortingData()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		}
	}
}

func menu() {
	fmt.Println("<== APLIKASI TRI-DARMA PERGURUAN TINGGI 2 ==>")
	fmt.Println("| 1. Tambah Data                            |")
	fmt.Println("| 2. Edit Data                              |")
	fmt.Println("| 3. Hapus Data                             |")
	fmt.Println("| 4. Tampilkan Data                         |")
	fmt.Println("| 5. Urutkan Data                           |")
	fmt.Println("| 0. Keluar                                 |")
	fmt.Println("<===========================================>")
	fmt.Print("Pilih menu: ")
}

func addData() {
	var kegiatan Kegiatan
	fmt.Print("Masukkan jenis data (1: Penelitian, 2: Abdimas): ")
	var tipe int
	var jmlAnggota int
	fmt.Scanln(&tipe)
	if tipe != 1 && tipe != 2 {
		fmt.Print("Pilihan tidak valid, silahkan pilih kembali (1/2): ")
		fmt.Scanln(&tipe)
	}

	fmt.Print("Masukkan Ketua: ")
	fmt.Scanln(&kegiatan.ketua)
	fmt.Print("Ada berapa anggota? (maksimal 4): ")
	fmt.Scanln(&jmlAnggota)
	for i := 0; i < jmlAnggota; i++ {
		fmt.Printf("Masukkan Anggota %d: ", i+1)
		fmt.Scanln(&kegiatan.anggota[i])
	}
	fmt.Print("Masukkan Prodi/Fakultas: ")
	fmt.Scanln(&kegiatan.prodiFakultas)
	fmt.Print("Masukkan Judul: ")
	fmt.Scanln(&kegiatan.judul)
	fmt.Print("Masukkan Sumber Dana: ")
	fmt.Scanln(&kegiatan.sumberDana)
	fmt.Print("Masukkan Tahun Kegiatan: ")
	fmt.Scanln(&kegiatan.tahun)
	if tipe == 1 {
		fmt.Print("Masukkan Luaran (Publikasi/Produk): ")
	} else if tipe == 2 {
		fmt.Print("Masukkan Luaran (Publikasi/Produk/Seminar/Pelatihan): ")
	}
	fmt.Scanln(&kegiatan.luaran)

	if tipe == 1 {
		if countPenelitian < MAXX {
			dataPenelitian[countPenelitian] = kegiatan
			countPenelitian++
			fmt.Println("Data Penelitian berhasil ditambahkan.")
		} else {
			fmt.Println("Data Penelitian penuh!")
		}
	} else {
		if countAbdimas < MAXX {
			dataAbdimas[countAbdimas] = kegiatan
			countAbdimas++
			fmt.Println("Data Abdimas berhasil ditambahkan.")
		} else {
			fmt.Println("Data Abdimas penuh!")
		}
	}
}

func editData() {
	fmt.Println("Edit Data Penelitian / Abdimas")
	fmt.Print("Masukkan jenis data yang ingin diedit (1: Penelitian, 2: Abdimas): ")
	var tipe int
	fmt.Scanln(&tipe)

	if tipe != 1 && tipe != 2 {
		fmt.Print("Pilihan tidak valid, silahkan pilih kembali (1/2): ")
		fmt.Scanln(&tipe)
	}

	fmt.Print("Masukkan judul kegiatan yang ingin diedit: ")
	var judul string
	fmt.Scanln(&judul)

	kegiatan, idx, found := SequentialSearch(tipe, judul)
	if !found {
		fmt.Println("Kegiatan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan Ketua baru: ")
	fmt.Scanln(&kegiatan.ketua)
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan Anggota %d baru (kosongkan jika tidak ada): ", i+1)
		fmt.Scanln(&kegiatan.anggota[i])
	}
	fmt.Print("Masukkan Prodi/Fakultas baru: ")
	fmt.Scanln(&kegiatan.prodiFakultas)
	fmt.Print("Masukkan Judul baru: ")
	fmt.Scanln(&kegiatan.judul)
	fmt.Print("Masukkan Sumber Dana baru: ")
	fmt.Scanln(&kegiatan.sumberDana)
	fmt.Print("Masukkan Tahun Kegiatan baru: ")
	fmt.Scanln(&kegiatan.tahun)
	if tipe == 1 {
		fmt.Print("Masukkan Luaran baru (Publikasi/Produk): ")
	} else {
		fmt.Print("Masukkan Luaran baru (Publikasi/Produk/Seminar/Pelatihan): ")
	}
	fmt.Scanln(&kegiatan.luaran)

	if tipe == 1 {
		dataPenelitian[idx] = *kegiatan
	} else {
		dataAbdimas[idx] = *kegiatan
	}
	fmt.Println("Data berhasil diedit.")
}

func deleteData() {
	fmt.Println("Hapus Data Penelitian / Abdimas")
	fmt.Print("Masukkan jenis data yang ingin dihapus (1: Penelitian, 2: Abdimas): ")
	var tipe int
	fmt.Scanln(&tipe)
	if tipe != 1 && tipe != 2 {
		fmt.Print("Pilihan tidak valid, silahkan pilih kembali (1/2): ")
		fmt.Scanln(&tipe)
	}
	fmt.Print("Masukkan judul kegiatan yang ingin dihapus: ")
	var judul string
	fmt.Scanln(&judul)

	if tipe == 1 {
		for i := 0; i < countPenelitian; i++ {
			if dataPenelitian[i].judul == judul {
				dataPenelitian[i] = dataPenelitian[countPenelitian-1]
				dataPenelitian[countPenelitian-1] = Kegiatan{}
				countPenelitian--
				fmt.Println("Data berhasil dihapus.")
				return
			}
		}
	} else {
		for i := 0; i < countAbdimas; i++ {
			if dataAbdimas[i].judul == judul {
				dataAbdimas[i] = dataAbdimas[countAbdimas-1]
				dataAbdimas[countAbdimas-1] = Kegiatan{}
				countAbdimas--
				fmt.Println("Data berhasil dihapus.")
				return
			}
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func printData() {
	fmt.Println("Tampilkan Data Penelitian / Abdimas")
	fmt.Println("1. Berdasarkan Tahun")
	fmt.Println("2. Berdasarkan Fakultas/Prodi")
	fmt.Println("3. Tampilkan semua data")
	fmt.Print("Pilih menu: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan Tahun: ")
		var tahun int
		fmt.Scanln(&tahun)

		idxPenelitian := binarySearchByTahunRecursive(dataPenelitian, 0, countPenelitian-1, tahun)
		idxAbdimas := binarySearchByTahunRecursive(dataAbdimas, 0, countAbdimas-1, tahun)

		fmt.Println("Data Penelitian:")
		if idxPenelitian != -1 {
			for i := idxPenelitian; i >= 0 && dataPenelitian[i].tahun == tahun; i-- {
				printKegiatan(dataPenelitian[i])
			}
			for i := idxPenelitian + 1; i < countPenelitian && dataPenelitian[i].tahun == tahun; i++ {
				printKegiatan(dataPenelitian[i])
			}
		} else {
			fmt.Println("Tidak ada data penelitian untuk tahun tersebut.")
		}

		fmt.Println("Data Abdimas:")
		if idxAbdimas != -1 {
			for i := idxAbdimas; i >= 0 && dataAbdimas[i].tahun == tahun; i-- {
				printKegiatan(dataAbdimas[i])
			}
			for i := idxAbdimas + 1; i < countAbdimas && dataAbdimas[i].tahun == tahun; i++ {
				printKegiatan(dataAbdimas[i])
			}
		} else {
			fmt.Println("Tidak ada data abdimas untuk tahun tersebut.")
		}
	} else if pilihan == 2 {
		fmt.Print("Masukkan Prodi/Fakultas: ")
		var prodiFakultas string
		fmt.Scanln(&prodiFakultas)

		fmt.Println("Data Penelitian:")
		found := false
		for i := 0; i < countPenelitian; i++ {
			if dataPenelitian[i].prodiFakultas == prodiFakultas {
				printKegiatan(dataPenelitian[i])
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada data penelitian untuk Prodi/Fakultas tersebut.")
		}

		fmt.Println("Data Abdimas:")
		found = false
		for i := 0; i < countAbdimas; i++ {
			if dataAbdimas[i].prodiFakultas == prodiFakultas {
				printKegiatan(dataAbdimas[i])
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada data abdimas untuk Prodi/Fakultas tersebut.")
		}
	} else if pilihan == 3 {
		fmt.Println("Data Penelitian:")
		for i := 0; i < countPenelitian; i++ {
			printKegiatan(dataPenelitian[i])
		}

		fmt.Println("Data Abdimas:")
		for i := 0; i < countAbdimas; i++ {
			printKegiatan(dataAbdimas[i])
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func sortingData() {
	fmt.Println("Urutkan Data Penelitian / Abdimas")
	fmt.Println("1. Berdasarkan Tahun")
	fmt.Println("2. Berdasarkan Jumlah Kegiatan")
	fmt.Print("Pilih menu: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilih urutan: ")
		var urutan int
		fmt.Scanln(&urutan)
		if urutan == 1 {
			insertionSortByTahun(true)
		} else if urutan == 2 {
			insertionSortByTahun(false)
		}
	} else if pilihan == 2 {
		sortByJumlahKegiatan()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func SequentialSearch(tipe int, judul string) (*Kegiatan, int, bool) {
	var kegiatan *Kegiatan
	var idx int
	var found bool

	if tipe == 1 {
		for i := 0; i < countPenelitian && !found; i++ {
			if dataPenelitian[i].judul == judul {
				kegiatan = &dataPenelitian[i]
				idx = i
				found = true
			}
		}
	} else {
		for i := 0; i < countAbdimas && !found; i++ {
			if dataAbdimas[i].judul == judul {
				kegiatan = &dataAbdimas[i]
				idx = i
				found = true
			}
		}
	}

	return kegiatan, idx, found
}

// fungsi untuk mencari data di tahun yang sesuai dengan tahun input
func binarySearchByTahunRecursive(data tabKegiatan, low, high, tahun int) int {
	// basis: Jika low melewati high, maka elemen tidak ditemukan
	if low > high {
		return -1
	}

	// hitung indeks tengah
	mid := (low + high) / 2

	// jika elemen ditemukan di indeks tengah
	if data[mid].tahun == tahun {
		return mid
	} else if data[mid].tahun < tahun {
		// pencarian di sisi kanan
		return binarySearchByTahunRecursive(data, mid+1, high, tahun)
	} else {
		// pencarian di sisi kiri
		return binarySearchByTahunRecursive(data, low, mid-1, tahun)
	}
}

func recursiveInsertionSort(arr *tabKegiatan, n int, ascending bool) {
	if n <= 1 {
		return
	}

	recursiveInsertionSort(arr, n-1, ascending)

	last := (*arr)[n-1]
	j := n - 2

	for j >= 0 && ((ascending && last.tahun < (*arr)[j].tahun) || (!ascending && last.tahun > (*arr)[j].tahun)) {
		(*arr)[j+1] = (*arr)[j]
		j--
	}
	(*arr)[j+1] = last
}

func insertionSortByTahun(ascending bool) {
	if countPenelitian > 0 {
		recursiveInsertionSort(&dataPenelitian, countPenelitian, ascending)
		fmt.Println("Data Penelitian setelah pengurutan:")
		for i := 0; i < countPenelitian; i++ {
			printKegiatan(dataPenelitian[i])
		}
	} else {
		fmt.Println("Data Penelitian kosong, tidak ada yang diurutkan.")
	}

	if countAbdimas > 0 {
		// Memanggil fungsi rekursif untuk dataAbdimas
		recursiveInsertionSort(&dataAbdimas, countAbdimas, ascending)
		fmt.Println("Data Abdimas setelah pengurutan:")
		for i := 0; i < countAbdimas; i++ {
			printKegiatan(dataAbdimas[i])
		}
	} else {
		fmt.Println("Data Abdimas kosong, tidak ada yang diurutkan.")
	}
}

// fungsi untuk menghitung jumlah kegiatan per tahunnya
func countKegiatanPerYear(data tabKegiatan, count int) ([MAX_YEARS]int, [MAX_YEARS]int) {
	var years [MAX_YEARS]int
	var counts [MAX_YEARS]int
	var yearIndex int

	for i := 0; i < count; i++ {
		year := data[i].tahun
		found := false

		for j := 0; j < yearIndex && !found; j++ {
			if years[j] == year {
				counts[j]++
				found = true
			}
		}

		if !found {
			years[yearIndex] = year
			counts[yearIndex] = 1
			yearIndex++
		}
	}

	return years, counts
}

//selection sort
func sortByJumlahKegiatan() {
	yearsPenelitian, countsPenelitian := countKegiatanPerYear(dataPenelitian, countPenelitian)
	yearsAbdimas, countsAbdimas := countKegiatanPerYear(dataAbdimas, countAbdimas)

	if countPenelitian > 0 {
		for i := 0; i < countPenelitian-1; i++ {
			maxIdx := i
			for j := i + 1; j < countPenelitian; j++ {
				if countsPenelitian[j] > countsPenelitian[maxIdx] {
					maxIdx = j
				}
			}
			countsPenelitian[i], countsPenelitian[maxIdx] = countsPenelitian[maxIdx], countsPenelitian[i]
			yearsPenelitian[i], yearsPenelitian[maxIdx] = yearsPenelitian[maxIdx], yearsPenelitian[i]
		}

		var sortedData tabKegiatan
		sortedCount := 0

		for i := 0; i < countPenelitian; i++ {
			for j := 0; j < countPenelitian; j++ {
				if dataPenelitian[j].tahun == yearsPenelitian[i] {
					sortedData[sortedCount] = dataPenelitian[j]
					sortedCount++
				}
			}
		}

		for i := 0; i < countPenelitian; i++ {
			dataPenelitian[i] = sortedData[i]
		}

		fmt.Println("Data penelitian setelah diurutkan berdasarkan jumlah kegiatan per tahun:")
		for i := 0; i < countPenelitian; i++ {
			printKegiatan(dataPenelitian[i])
		}
	}

	if countAbdimas > 0 {
		for i := 0; i < countAbdimas-1; i++ {
			maxIdx := i
			for j := i + 1; j < countAbdimas; j++ {
				if countsAbdimas[j] > countsAbdimas[maxIdx] {
					maxIdx = j
				}
			}
			countsAbdimas[i], countsAbdimas[maxIdx] = countsAbdimas[maxIdx], countsAbdimas[i]
			yearsAbdimas[i], yearsAbdimas[maxIdx] = yearsAbdimas[maxIdx], yearsAbdimas[i]
		}

		var sortedData tabKegiatan
		sortedCount := 0

		for i := 0; i < countAbdimas; i++ {
			for j := 0; j < countAbdimas; j++ {
				if dataAbdimas[j].tahun == yearsAbdimas[i] {
					sortedData[sortedCount] = dataAbdimas[j]
					sortedCount++
				}
			}
		}

		for i := 0; i < countAbdimas; i++ {
			dataAbdimas[i] = sortedData[i]
		}

		fmt.Println("Data abdimas setelah diurutkan berdasarkan jumlah kegiatan per tahun:")
		for i := 0; i < countAbdimas; i++ {
			printKegiatan(dataAbdimas[i])
		}
	}
}

func printKegiatan(kegiatan Kegiatan) {
	fmt.Println("=================================")
	fmt.Println("Ketua         :", kegiatan.ketua)
	fmt.Println("Anggota")
	for j := 0; j < 4; j++ {
		if kegiatan.anggota[j] != "" {
			fmt.Printf("Anggota %d     : %s\n", j+1, kegiatan.anggota[j])
		}
	}
	fmt.Println("Prodi/Fakultas:", kegiatan.prodiFakultas)
	fmt.Println("Judul         :", kegiatan.judul)
	fmt.Println("Sumber Dana   :", kegiatan.sumberDana)
	fmt.Println("Tahun         :", kegiatan.tahun)
	fmt.Println("Luaran        :", kegiatan.luaran)
	fmt.Println("=================================")
}
