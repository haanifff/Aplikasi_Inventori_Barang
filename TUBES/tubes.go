package main

import (
	"fmt"
)

type Barang struct {
	ID       int
	Nama     string
	Jumlah   int
	Kategori string
	tanggal  int
	bulan    int
	tahun    int
}

const MaxBarang int = 1000

var inventori tabBarang

type tabBarang [MaxBarang]Barang

var jumlahBarang int

func main() {
	fungsi_menu()
}

func fungsi_menu() {
	//menampilkan menu utama
	i := 11
	for i != 0 {
		fmt.Println()
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|              ~~ APLIKASI INVENTORI BARANG ~~             |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  Menu :                                                  |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  1. Transaksi & Input Barang                             |")
		fmt.Println("|  2. Daftar & List Barang                                 |")
		fmt.Println("|  3. Ubah Data Barang                                     |")
		fmt.Println("|  0. Quit                                                 |")
		fmt.Println("------------------------------------------------------------")
		fmt.Print("Masukkan Pilihan (1/2/3/0) : ")
		fmt.Scan(&i)
		fmt.Println()

		for i != 1 && i != 2 && i != 3 && i != 0 {
			fmt.Println("Pastikan Pilihan Anda Benar")
			fmt.Print("Masukkan Pilihan (1/2/3/0) : ")
			fmt.Scan(&i)
			fmt.Println()
		}
		if i == 1 {
			menu_transaksi()
		} else if i == 2 {
			menu_daftar_barang()
		} else if i == 3 {
			menu_ubah_barang()
		}
	}
	fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Inventori Barang")
}

func menu_transaksi() {
	//menampilkan menu transaksi
	i := 11
	for i != 0 {
		fmt.Println()
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|              ~~ APLIKASI INVENTORI BARANG ~~             |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  Transaksi & Input Barang :                              |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  1. Input Barang Baru Gudang                             |")
		fmt.Println("|  2. Tambah Stok Barang                                   |")
		fmt.Println("|  3. Barang Keluar Gudang                                 |")
		fmt.Println("|  0. Kembali ke Menu                                      |")
		fmt.Println("------------------------------------------------------------")
		fmt.Print("Masukkan Pilihan (1/2/3/0) : ")
		fmt.Scan(&i)
		fmt.Println()

		for i != 1 && i != 2 && i != 3 && i != 0 {
			fmt.Println("Pastikan Pilihan Anda Benar")
			fmt.Print("Masukkan Pilihan (1/2/3/0) : ")
			fmt.Scan(&i)
			fmt.Println()
		}
		if i == 1 {
			inputBarang(&inventori)
		} else if i == 2 {
			tambahBarang(&inventori)
		} else if i == 3 {
			barangKeluarGudang(&inventori)
		}
	}
}

func barangKeluarGudang(A *tabBarang) {
	//menjalankan algoritma pengurangan stok pad barang yang dicari
	var namabarang string
	var ketemu bool = false
	var j, tambah int
	fmt.Print("Masukkan nama Barang Yang Ingin Dikurangi : ")
	fmt.Scan(&namabarang)
	fmt.Println()
	for i := 0; i < MaxBarang; i++ {
		if A[i].Nama == namabarang {
			j = i
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Print("Masukkan Jumlah Yang Ingin Dikurangi : ")
		fmt.Scan(&tambah)
		fmt.Println()
		A[j].Jumlah = A[j].Jumlah - tambah
		fmt.Println("Stok Barang Telah Dikurangi")
	} else {
		fmt.Println("Nama yang anda masukkan tidak valid")
	}
}

func tambahBarang(A *tabBarang) {
	//menjalankan algoritma penambahan stok pad barang yang dicari
	var namabarang string
	var ketemu bool = false
	var j, tambah int
	fmt.Print("Masukkan Nama Barang Yang Ingin Ditambahkan : ")
	fmt.Scan(&namabarang)
	fmt.Println()
	for i := 0; i < MaxBarang; i++ {
		if A[i].Nama == namabarang {
			j = i
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Print("Masukkan Jumlah Yang Ingin Ditambahakan : ")
		fmt.Scan(&tambah)
		fmt.Println()
		A[j].Jumlah = A[j].Jumlah + tambah
		fmt.Println("Stok Barang Telah Ditambahkan")
	} else {
		fmt.Println("Nama yang anda masukkan tidak valid")
	}
}

func inputBarang(A *tabBarang) {
	//input barang pada array
	var i_ID, i_jumlah, i_tanggal, i_bulan, i_tahun int
	var i_nama, i_kategori string
	var ketemu bool = false
	var ketemu2 bool = false
	fmt.Print("Masukkan ID Barang (Maksimal 3 digit angka) : ")
	fmt.Scan(&i_ID)
	fmt.Println()
	for i := 0; i < MaxBarang; i++ {
		if A[i].ID == i_ID {
			ketemu = true
		}
	}
	if (i_ID > 0 && i_ID < 1000) && ketemu == false {
		fmt.Print("Masukkan Nama Barang : ")
		fmt.Scan(&i_nama)
		fmt.Println()
		for i := 0; i < MaxBarang; i++ {
			if A[i].Nama == i_nama {
				ketemu2 = true
			}
		}
		if ketemu2 == false {
			fmt.Print("Masukkan Jumlah Barang : ")
			fmt.Scan(&i_jumlah)
			fmt.Println()
			fmt.Print("Masukkan Kategori Barang : ")
			fmt.Scan(&i_kategori)
			fmt.Println()
			fmt.Print("Masukkan Tanggal Input Barang (dd mm yyyy) : ")
			fmt.Scan(&i_tanggal, &i_bulan, &i_tahun)
			fmt.Println()
			if (i_tanggal > 0 && i_tanggal < 32) && (i_bulan > 0 && i_bulan < 13) && (i_tahun > 0 && i_tahun < 2025) {
				A[jumlahBarang].ID = i_ID
				A[jumlahBarang].Nama = i_nama
				A[jumlahBarang].Jumlah = i_jumlah
				A[jumlahBarang].Kategori = i_kategori
				A[jumlahBarang].tanggal = i_tanggal
				A[jumlahBarang].bulan = i_bulan
				A[jumlahBarang].tahun = i_tahun
				jumlahBarang++
			} else {
				fmt.Println("Data Yang Anda Masukkan Tidak Sesuai Dengan Kriteria, Harap Coba Kembali")
				fmt.Println()
			}
		} else {
			fmt.Println("Data Yang Anda Masukkan Terdeteksi Sama Dengan Data Lain, Pastikan Input Anda Benar")
			fmt.Println()
		}
	} else if ketemu == true {
		fmt.Println("Data Yang Anda Masukkan Terdeteksi Sama Dengan Data Lain, Pastikan Input Anda Benar")
		fmt.Println()
	} else if i_ID <= 0 || i_ID >= 1000 {
		fmt.Println("Data Yang Anda Masukkan Tidak Sesuai Dengan Kriteria, Harap Coba Kembali")
		fmt.Println()
	}
}

func menu_daftar_barang() {
	//menampilkan menu daftar dan list barang
	i := "11"
	for i != "0" {
		fmt.Println()
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|              ~~ APLIKASI INVENTORI BARANG ~~             |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  Daftar & List Barang :                                  |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  1. Tampilkan Daftar Barang                              |")
		fmt.Println("|  2. Cari Barang Berdasarkan...                           |")
		fmt.Println("|     a. Nama Barang                                       |")
		fmt.Println("|     b. ID Barang                                         |")
		fmt.Println("|     c. Waktu input barang                                |")
		fmt.Println("|     d. Kategori Barang                                   |")
		fmt.Println("|  3. Urutkan Berdasarkan...                               |")
		fmt.Println("|     a. Jumlah Stok Tersedia                              |")
		fmt.Println("|     b. ID Barang                                         |")
		fmt.Println("|     c. Waktu input barang                                |")
		fmt.Println("|  0. Kembali ke Menu                                      |")
		fmt.Println("------------------------------------------------------------")
		fmt.Print("Masukkan Pilihan (1/2a/2b/2c/2d/3a/3b/3c/0) : ")
		fmt.Scan(&i)
		fmt.Println()

		for i != "1" && i != "2a" && i != "2b" && i != "2c" && i != "2d" && i != "3a" && i != "3b" && i != "3c" && i != "0" {
			fmt.Println("Pastikan Pilihan Anda Benar")
			fmt.Print("Masukkan Pilihan (1/2a/2b/2c/2d/3a/3b/3c/0) : ")
			fmt.Scan(&i)
			fmt.Println()
		}
		if i == "1" {
			printBarang(inventori, jumlahBarang)
		} else if i == "2a" {
			caribarangNama(inventori)
		} else if i == "2b" {
			caribarangID(inventori)
		} else if i == "2c" {
			caribarangWaktu(inventori)
		} else if i == "2d" {
			caribarangKategori(inventori)
		} else if i == "3a" {
			urutkan_stok(&inventori, jumlahBarang) // ascending
		} else if i == "3b" {
			urutkan_ID(&inventori, jumlahBarang) // ascending
		} else if i == "3c" {
			urutkan_waktu(&inventori, jumlahBarang) // ascending
		}
	}
}

func urutkan_waktu(A *tabBarang, n int) {
	//mengurutkan array berdasarkan waktu
	var i, j, idx_min int
	var t Barang
	i = 1
	if n > 0 {
		for i <= n-1 {
			idx_min = i - 1
			j = i
			for j < n {
				if A[idx_min].tahun > A[j].tahun ||
					(A[idx_min].tahun == A[j].tahun && A[idx_min].bulan > A[j].bulan) ||
					(A[idx_min].tahun == A[j].tahun && A[idx_min].bulan == A[j].bulan && A[idx_min].tanggal > A[j].tanggal) {
					idx_min = j
				}
				j = j + 1
			}
			t = A[idx_min]
			A[idx_min] = A[i-1]
			A[i-1] = t
			i = i + 1
		}
		fmt.Println("Data Telah Terurut Berdasarkan Waktu Masuk Barang")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Tidak Ada Data Barang")
		fmt.Println()
	}
}

func urutkan_ID(A *tabBarang, n int) {
	//mengurutkan berdasarkan id
	var i, j, idx_min int
	var t Barang
	i = 1
	if n > 0 {
		for i <= n-1 {
			idx_min = i - 1
			j = i
			for j < n {
				if A[idx_min].ID > A[j].ID {
					idx_min = j
				}
				j = j + 1
			}
			t = A[idx_min]
			A[idx_min] = A[i-1]
			A[i-1] = t
			i = i + 1
		}
		fmt.Println("Data Telah Terurut Berdasarkan ID Barang")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Tidak Ada Data Barang")
		fmt.Println()
	}
}

func urutkan_stok(A *tabBarang, n int) {
	//mengurutkan berdasarkan jumlah barang
	var i, j, idx_min int
	var t Barang
	i = 1
	if n > 0 {
		for i <= n-1 {
			idx_min = i - 1
			j = i
			for j < n {
				if A[idx_min].Jumlah > A[j].Jumlah {
					idx_min = j
				}
				j = j + 1
			}
			t = A[idx_min]
			A[idx_min] = A[i-1]
			A[i-1] = t
			i = i + 1
		}
		fmt.Println("Data Telah Terurut Berdasarkan Stok Barang")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Tidak Ada Data Barang")
		fmt.Println()
	}
}

func caribarangWaktu(A tabBarang) {
	//searching barang berdasarkan waktu
	var tanggal, bulan, tahun int
	var ketemu bool = false
	fmt.Print("Masukkan Waktu Input Barang (TT MM YYYY) (input '0' jika ingin dikosongkan) : ")
	fmt.Scan(&tanggal, &bulan, &tahun)
	fmt.Println()

	if tanggal <= 0 && bulan <= 0 && tahun <= 0 { //000
		fmt.Println("Tidak Ada Barang Yang Akan Dicari")
	} else if tanggal == 0 && bulan == 0 && tahun != 0 { //001
		for i := 0; i < MaxBarang; i++ {
			if A[i].tahun == tahun {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].tahun == tahun {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	} else if tanggal == 0 && bulan != 0 && tahun == 0 { //010
		for i := 0; i < MaxBarang; i++ {
			if A[i].bulan == bulan {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].bulan == bulan {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	} else if tanggal == 0 && bulan != 0 && tahun != 0 { // 011
		for i := 0; i < MaxBarang; i++ {
			if A[i].bulan == bulan && A[i].tahun == tahun {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].bulan == bulan && A[i].tahun == tahun {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	} else if tanggal != 0 && bulan == 0 && tahun == 0 { // 100
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	} else if tanggal != 0 && bulan == 0 && tahun != 0 { // 101
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal && A[i].tahun == tahun {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal && A[i].tahun == tahun {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	} else if tanggal != 0 && bulan != 0 && tahun == 0 { // 110
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal && A[i].bulan == bulan {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal && A[i].bulan == bulan {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	} else if tanggal != 0 && bulan != 0 && tahun != 0 { // 111
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal && A[i].bulan == bulan && A[i].tahun == tahun {
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
			fmt.Println()
			fmt.Print("----------------------------------------------------------------------------------")
			fmt.Println()
		}
		for i := 0; i < MaxBarang; i++ {
			if A[i].tanggal == tanggal && A[i].bulan == bulan && A[i].tahun == tahun {
				fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
				fmt.Println()
			}
		}
		if ketemu == true {
			fmt.Println("----------------------------------------------------------------------------------")
		}
		if ketemu == false {
			fmt.Println("Nama barang yang anda cari tidak ada")
		}
		fmt.Println()
	}
}

func caribarangKategori(A tabBarang) {
	//searching barang berdasarkan kategori
	var kategori string
	var ketemu bool = false
	fmt.Print("Masukkan Kategori Barang : ")
	fmt.Scan(&kategori)
	fmt.Println()
	for i := 0; i < MaxBarang; i++ {
		if A[i].Kategori == kategori {
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
		fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
	}
	for i := 0; i < MaxBarang; i++ {
		if A[i].Kategori == kategori {
			fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
			fmt.Println()
		}
	}
	if ketemu == true {
		fmt.Println("----------------------------------------------------------------------------------")
	}
	if ketemu == false {
		fmt.Println("Nama barang yang anda cari tidak ada")
	}
	fmt.Println()
}

func caribarangID(A tabBarang) {
	//searching barang berdasarkan ID
	var id int
	var ketemu bool = false
	fmt.Print("Masukkan ID Barang : ")
	fmt.Scan(&id)
	fmt.Println()
	for i := 0; i < MaxBarang; i++ {
		if A[i].ID == id {
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
		fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
	}
	for i := 0; i < MaxBarang; i++ {
		if A[i].ID == id {
			fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
			fmt.Println()
		}
	}
	if ketemu == true {
		fmt.Println("----------------------------------------------------------------------------------")
	}
	if ketemu == false {
		fmt.Println("Nama barang yang anda cari tidak ada")
	}
	fmt.Println()
}

func caribarangNama(A tabBarang) {
	//searching barang berdasarkan nama
	var namabarang string
	var ketemu bool = false
	fmt.Print("Masukkan nama Barang : ")
	fmt.Scan(&namabarang)
	fmt.Println()
	for i := 0; i < MaxBarang; i++ {
		if A[i].Nama == namabarang {
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
		fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
	}
	for i := 0; i < MaxBarang; i++ {
		if A[i].Nama == namabarang {
			fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
			fmt.Println()
		}
	}
	if ketemu == true {
		fmt.Println("----------------------------------------------------------------------------------")
	}
	if ketemu == false {
		fmt.Println("Nama barang yang anda cari tidak ada")
	}
	fmt.Println()
}

func printBarang(A tabBarang, n int) {
	//print
	if n > 0 {
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
		fmt.Printf("%-10v %-20v %-15v %-20v %-14v", "ID BARANG", "NAMA", "STOK BARANG", "KATEGORI", "WAKTU INPUT")
		fmt.Println()
		fmt.Print("----------------------------------------------------------------------------------")
		fmt.Println()
		for i := 0; i < jumlahBarang; i++ {
			fmt.Printf("%-10v %-20v %-15v %-20v %-2v/%-2v/%-4v", A[i].ID, A[i].Nama, A[i].Jumlah, A[i].Kategori, A[i].tanggal, A[i].bulan, A[i].tahun)
			fmt.Println()
		}
		fmt.Println("----------------------------------------------------------------------------------")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Tidak Ada Data Barang")
		fmt.Println()
	}
}

func menu_ubah_barang() {
	//menu ubah barang
	i := 11
	for i != 0 {
		fmt.Println()
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|              ~~ APLIKASI INVENTORI BARANG ~~             |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  Ubah Data Barang :                                      |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("|  1. Ubah Data Barang                                     |")
		fmt.Println("|  2. Hapus Data Barang                                    |")
		fmt.Println("|  0. Kembali ke Menu                                      |")
		fmt.Println("------------------------------------------------------------")
		fmt.Print("Masukkan Pilihan (1/2/0) : ")
		fmt.Scan(&i)
		fmt.Println()

		for i != 1 && i != 2 && i != 0 {
			fmt.Println("Pastikan Pilihan Anda Benar")
			fmt.Print("Masukkan Pilihan (1/2/0) : ")
			fmt.Scan(&i)
			fmt.Println()
		}
		if i == 1 {
			ubah_data_barang()
		} else if i == 2 {
			hapus_data_barang()
		}
	}
}

func ubah_data_barang() {
	//ubah data barang berdasarkan...
	var i string
	i = "null"
	for i != "ID" && i != "NamaBarang" && i != "Id" && i != "namabarang" && i != "id" && i != "Namabarang" {
		fmt.Print("Ubah Data Berdasarkan... (ID / NamaBarang) : ")
		fmt.Scan(&i)
		fmt.Println()
		if !(i == "ID" || i == "NamaBarang" || i == "Id" || i == "namabarang" || i == "id" || i == "Namabarang") {
			fmt.Println("Pastikan Pilihan Anda Benar")
		}
	}
	if i == "ID" || i == "Id" || i == "id" {
		Ubah_data_barang_id(&inventori)
	} else if i == "NamaBarang" || i == "Namabarang" || i == "namabarang" {
		Ubah_data_barang_namaBarang(&inventori)
	}
}

func Ubah_data_barang_id(A *tabBarang) {
	//algoritma merubah array dengan meng-scan array tersebut kembali jika id yang dicari sama dengan array tersebut
	var id int
	var i_ID, i_jumlah, i_tanggal, i_bulan, i_tahun int
	var i_nama, i_kategori string
	fmt.Print("Masukkan ID Barang Yang Ingin Diubah : ")
	fmt.Scan(&id)
	fmt.Println()

	jumlahBarang := len(*A)
	found := false
	for x := 0; x < jumlahBarang; x++ {
		if (*A)[x].ID == id {
			found = true

			fmt.Print("Masukkan ID Barang (Maksimal 3 digit angka) : ")
			fmt.Scan(&i_ID)
			fmt.Println()
			if i_ID > 0 && i_ID <= 1000 {
				fmt.Print("Masukkan Nama Barang : ")
				fmt.Scan(&i_nama)
				fmt.Println()
				fmt.Print("Masukkan Jumlah Barang : ")
				fmt.Scan(&i_jumlah)
				fmt.Println()
				fmt.Print("Masukkan Kategori Barang : ")
				fmt.Scan(&i_kategori)
				fmt.Println()
				fmt.Print("Masukkan Tanggal Input Barang (dd mm yyyy) : ")
				fmt.Scan(&i_tanggal, &i_bulan, &i_tahun)
				if (i_tanggal > 0 && i_tanggal < 32) && (i_bulan > 0 && i_bulan < 13) && (i_tahun > 0 && i_tahun < 2025) {
					A[x].ID = i_ID
					A[x].Nama = i_nama
					A[x].Jumlah = i_jumlah
					A[x].Kategori = i_kategori
					A[x].tanggal = i_tanggal
					A[x].bulan = i_bulan
					A[x].tahun = i_tahun
				} else {
					fmt.Println("Data Yang Anda Masukkan Tidak Sesuai Dengan Kriteria, Harap Coba Kembali")
				}
			} else {
				fmt.Println("Data Yang Anda Masukkan Tidak Sesuai Dengan Kriteria, Harap Coba Kembali")
			}
		}
	}

	if !found {
		fmt.Printf("Barang dengan ID %d tidak ditemukan.\n", id)
	}
}

func Ubah_data_barang_namaBarang(A *tabBarang) {
	//algoritma merubah array dengan meng-scan array tersebut kembali jika nama yang dicari sama dengan array tersebut
	var nama string
	var i_ID, i_jumlah, i_tanggal, i_bulan, i_tahun int
	var i_nama, i_kategori string
	fmt.Print("Masukkan Nama Barang Yang Ingin Diubah : ")
	fmt.Scan(&nama)
	fmt.Println()

	jumlahBarang := len(*A)
	found := false
	for x := 0; x < jumlahBarang; x++ {
		if (*A)[x].Nama == nama {
			found = true

			fmt.Print("Masukkan ID Barang (Maksimal 3 digit angka) : ")
			fmt.Scan(&i_ID)
			fmt.Println()
			if i_ID > 0 && i_ID <= 1000 {
				fmt.Print("Masukkan Nama Barang : ")
				fmt.Scan(&i_nama)
				fmt.Println()
				fmt.Print("Masukkan Jumlah Barang : ")
				fmt.Scan(&i_jumlah)
				fmt.Println()
				fmt.Print("Masukkan Kategori Barang : ")
				fmt.Scan(&i_kategori)
				fmt.Println()
				fmt.Print("Masukkan Tanggal Input Barang (dd mm yyyy) : ")
				fmt.Scan(&i_tanggal, &i_bulan, &i_tahun)
				if (i_tanggal > 0 && i_tanggal < 32) && (i_bulan > 0 && i_bulan < 13) && (i_tahun > 0 && i_tahun < 2025) {
					A[x].ID = i_ID
					A[x].Nama = i_nama
					A[x].Jumlah = i_jumlah
					A[x].Kategori = i_kategori
					A[x].tanggal = i_tanggal
					A[x].bulan = i_bulan
					A[x].tahun = i_tahun
				} else {
					fmt.Println("Data Yang Anda Masukkan Tidak Sesuai Dengan Kriteria, Harap Coba Kembali")
				}

			} else {
				fmt.Println("Data Yang Anda Masukkan Tidak Sesuai Dengan Kriteria, Harap Coba Kembali")
			}
		}
	}

	if !found {
		fmt.Printf("Barang dengan Nama %s tidak ditemukan.\n", nama)
	}
}

func hapus_data_barang() {
	//hapus array berdasarkan...
	var i string
	i = "null"
	for i != "ID" && i != "NamaBarang" && i != "Id" && i != "namabarang" && i != "id" && i != "Namabarang" {
		fmt.Print("Hapus Data Berdasarkan... (ID/NamaBarang) : ")
		fmt.Scan(&i)
		fmt.Println()
		if !(i == "ID" || i == "NamaBarang" || i == "Id" || i == "namabarang" || i == "id" || i == "Namabarang") {
			fmt.Println("Pastikan Pilihan Anda Benar")
		}
	}
	if i == "ID" || i == "Id" || i == "id" {
		hapus_data_barang_id(&inventori, &jumlahBarang)
	} else if i == "NamaBarang" || i == "Namabarang" || i == "namabarang" {
		hapus_data_barang_namaBarang(&inventori, &jumlahBarang)
	}
}

func hapus_data_barang_id(A *tabBarang, n *int) {
	//algoritma menghapus array dengan mencari data id yang akan dihapus dan menimpa data tersebut
	var id int
	fmt.Print("Masukkan ID barang Yang Ingin Dihapus : ")
	fmt.Scan(&id)
	fmt.Println()
	idx := -1
	for i := 0; i < *n; i++ {
		if A[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Data Yang Anda Cari Tidak Ada")
	} else {
		for i := idx; i < *n; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Printf("Data dengan ID %d Telah Berhasil Dihapus", id)
		fmt.Println()
	}
}

func hapus_data_barang_namaBarang(A *tabBarang, n *int) {
	//algoritma menghapus array dengan mencari data barang yang akan dihapus dan menimpa data tersebut
	var nama string
	fmt.Print("Masukkan Nama Barang Yang Ingin Dihapus : ")
	fmt.Scan(&nama)
	fmt.Println()
	idx := -1
	for i := 0; i < *n; i++ {
		if A[i].Nama == nama {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Data Yang Anda Cari Tidak Ada")
	} else {
		for i := idx; i < *n; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Printf("Data Dengan Nama %s Telah Berhasil Dihapus", nama)
		fmt.Println()
	}
}
