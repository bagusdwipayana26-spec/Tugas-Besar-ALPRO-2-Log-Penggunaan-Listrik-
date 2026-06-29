package main

import "fmt"

const Nmax = 200

type Perangkat struct {
	nama     string
	watt     float64
	durasi   float64
	ruangan  string
	konsumsi float64
	aktif    bool
	id       int
}

type ArrayPerangkat [Nmax]Perangkat

//fungsi utama
func main() {
	var A ArrayPerangkat
	var n int
	display(A, n)
}
func display(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan jumlah data n terdefinisi (bisa kosong).
		F.S. : Menampilkan menu utama ke layar dan mengeksekusi sub-program sesuai dengan nomor fitur yang dipilih pengguna.
	*/
	var a string
	fmt.Println("---------------------------------------------------")
	fmt.Println("| SELAMAT DATANG DI SISTEM PENGECEKAN LOG LISTRIK |")
	fmt.Println("---------------------------------------------------")
	fmt.Println("silahkan memilih salah satu fitur dibawah ini:")
	fmt.Println("1 : Edit Data Perangkat")
	fmt.Println("2 : Cari Perangkat")
	fmt.Println("3 : Urutkan Perangkat")
	fmt.Println("4 : Tampilkan Data Perangkat")
	fmt.Println("5 : Tampilkan Statistik Perangkat Paling Boros")
	fmt.Println("6 : Cek Peringatan Durasi")
	fmt.Println("7 : Evaluasi Batas Aman Bulanan")
	fmt.Println("8 : Hitung Budget biaya perbulan")
	fmt.Println("9 : EXIT")
	fmt.Print("Silahkan Masukkan Nomor Fitur (1/2/3/4/5/6/7/8/9): ")
	fmt.Scan(&a)
	if a == "1" {
		edt(A, n)
	} else if a == "2" {
		cari(A, n)
	} else if a == "3" {
		urut(A, n)
	} else if a == "4" {
		tampil(A, n)
	} else if a == "5" {
		TampilkanStatistik(A, n)
	} else if a == "6" {
		CekPeringatandurasi(A, n)
	} else if a == "7" {
		EvaluasiBatasAman(A, n)
	} else if a == "8" {
		hitungbiaya(A, n)
	} else if a == "9" {

	} else {
		fmt.Println()
		fmt.Println("---INPUT TIDAK VALID!---")
		fmt.Println()
		display(A, n)
	}
}

//fungsi edit perangkat
func edt(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menampilkan sub-menu edit dan memanggil fungsi manipulasi data (tambah, edit, hapus, nonaktif) sesuai pilihan pengguna.
	*/
	var b string
	fmt.Println("--------------------------------------")
	fmt.Println("| Anda Memasuki Fitur Edit Perangkat |")
	fmt.Println("--------------------------------------")
	fmt.Println("Silahkan Memilih Salah Satu Fitur Dibawah Ini:")
	fmt.Println("1 : Tambah Perangkat")
	fmt.Println("2 : Edit Perangkat")
	fmt.Println("3 : Hapus Perangkat")
	fmt.Println("4 : Nonaktifkan Perangkat")
	fmt.Println("5 : EXIT")
	fmt.Print("Silahkan Masukkan Nomor Fitur (1/2/3/4/5): ")
	fmt.Scan(&b)
	if b == "1" {
		tambahperangkat(A, n)
	} else if b == "2" {
		editperangkat(A, n)
	} else if b == "3" {
		hapusperangkat(A, n)
	} else if b == "4" {
		nonaktif(A, n)
	} else if b == "5" {
		display(A, n)
	} else {
		fmt.Println("INPUT TIDAK VALID!")
		edt(A, n)
	}
}
func tambahperangkat(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi, n > 0 (jika n = 0, tampil pesan kosong).
		F.S. : Nilai nama, ruangan, watt, durasi, dan konsumsi dari perangkat dengan ID yang dicari berhasil diperbarui di dalam Array A.
	*/
	var b string
	fmt.Println("\n---Masukkan Data Perangkat---")
	fmt.Print("Nama Perangkat              : ")
	fmt.Scan(&A[n].nama)
	fmt.Print("Daya (Watt)                 : ")
	fmt.Scan(&A[n].watt)
	fmt.Print("Lama Pemakaian Harian (Jam) : ")
	fmt.Scan(&A[n].durasi)
	fmt.Print("Ruangan                     : ")
	fmt.Scan(&A[n].ruangan)
	A[n].konsumsi = A[n].watt * A[n].durasi
	A[n].aktif = true
	A[n].id = n + 1
	n++
	fmt.Println("---Data Berhasil Ditambahkan---")
	fmt.Println("Masih ingin menambahkan data perangkat? (1: IYA / 2: TIDAK)")
	fmt.Print("Masukkan Angka(1/2): ")
	fmt.Scan(&b)
	if b == "1" {
		tambahperangkat(A, n)
	} else if b == "2" {
		edt(A, n)
	} else {
		fmt.Println("INPUT TIDAK VALID!")
		fmt.Println()
		edt(A, n)
	}

}
func editperangkat(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi, n > 0 (jika n = 0, tampil pesan kosong).
		F.S. : Nilai nama, ruangan, watt, durasi, dan konsumsi dari perangkat dengan ID yang dicari berhasil diperbarui di dalam Array A.
	*/
	var cariid int
	var idx int
	if n == 0 {
		fmt.Println("Data log masih kosong!")
	} else {
		fmt.Println("\n---Edit Data Perangkat---")
		fmt.Print("Masukkan ID Perangkat yang ingin diedit : ")
		fmt.Scan(&cariid)
		insertionID(&A, n)
		idx = Cariid(A, n, cariid)
		if idx != -1 {
			fmt.Println("Data ditemukan!")
			fmt.Print("Masukkan Nama Perangkat Baru  :")
			fmt.Scan(&A[idx].nama)
			fmt.Print("Masukkan Ruang Perangkat Baru :")
			fmt.Scan(&A[idx].ruangan)
			fmt.Print("Masukkan Daya Baru (Watt)     : ")
			fmt.Scan(&A[idx].watt)
			fmt.Print("Masukkan Durasi Baru (Jam)    : ")
			fmt.Scan(&A[idx].durasi)
			A[idx].konsumsi = A[idx].watt * A[idx].durasi
			fmt.Println("Data berhasil diperbarui!")
		} else {
			fmt.Println("Data perangkat tidak ditemukan atau sudah dihapus.")
		}
	}
	edt(A, n)
}
func insertionID(A *ArrayPerangkat, n int) {
	var pass, j int
	var temp Perangkat
	for pass = 1; pass < n; pass++ {
		temp = (*A)[pass]
		j = pass - 1
		for j >= 0 && (*A)[j].id > temp.id {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = temp
	}
}
func hapusperangkat(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Jika ID ditemukan, elemen array digeser untuk menimpa data yang dihapus. (Catatan: jangan lupa tambahkan *n-- di kodenya agar jumlah total data ikut berkurang).
	*/
	var cariid int
	var idxhapus, i int
	fmt.Println("\n---Hapus Data Perangkat---")
	fmt.Print("Masukkan ID Perangkat yang ingin dihapus: ")
	fmt.Scan(&cariid)
	if n == 0 {
		fmt.Println("Data log masih kosong!")
	}
	idxhapus = Cariid(A, n, cariid)
	if idxhapus == -1 {
		fmt.Println("ID perangkat tidak ditemukan.")
	} else {
		for i = idxhapus; i < n-1; i++ {
			A[i] = A[i+1]
			A[i].id = i + 1
		}
		n = n - 1
		fmt.Println("---Data Berhasil Dihapus---")
	}
	edt(A, n)
}
func nonaktif(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Status (boolean 'aktif') dari perangkat dengan ID yang dicari berubah menjadi false (Tidak Aktif).
	*/
	fmt.Println()
	var id, idx int
	fmt.Print("Masukkan ID Perangkat yang ingin dinonaktifkan: ")
	fmt.Scan(&id)
	idx = Cariid(A, n, id)
	A[idx].aktif = false
	fmt.Println("----Perangkat Berhasil Dinonaktifkan")
	edt(A, n)
}

//fungsi mencari perangkat (Sequential Search & Binary Search)
func cari(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menampilkan sub-menu pencarian dan memanggil fungsi pencarian spesifik sesuai pilihan pengguna.
	*/
	var b string
	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Println("| Anda Memasuki Fitur Cari Perangkat |")
	fmt.Println("--------------------------------------")
	fmt.Println("Silahkan Memilih Salah Satu Fitur Dibawah Ini:")
	fmt.Println("1 : Cari Perangkat Berdasarkan Ruangan")
	fmt.Println("2 : Cari Perangkat Berdasarkan Nama Perangkat")
	fmt.Println("3 : EXIT")
	fmt.Print("Silahkan Masukkan Nomor Fitur (1/2/3): ")
	fmt.Scan(&b)
	if b == "1" {
		cariberdasarkanruangan(A, n)
	} else if b == "2" {
		cariberdasarkannama(A, n)
	} else if b == "3" {
		display(A, n)
	} else {
		fmt.Println("INPUT TIDAK VALID!")
		cari(A, n)
	}
}
func Cariid(A ArrayPerangkat, n int, cari int) int {
	/*
		I.S. : Array A, jumlah data n, dan integer 'cari' (ID) terdefinisi. Data di array sudah harus terurut berdasarkan ID jika menggunakan Binary Search.
		F.S. : Mengembalikan nilai indeks (idx) dari array di mana ID tersebut ditemukan. Mengembalikan -1 jika tidak ditemukan.
	*/
	var mid, left, right, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A[mid].id == cari {
			idx = mid
		} else if A[mid].id > cari {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}
func cariberdasarkanruangan(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menampilkan semua data perangkat yang berstatus aktif dan memiliki atribut 'ruangan' sesuai dengan input pengguna.
	*/
	var cariruang string
	var ketemu bool = false
	var i, j int
	fmt.Println("\n---Pencarian Berdasarkan Ruangan---")
	fmt.Print("Masukkan Nama Ruangan : ")
	fmt.Scan(&cariruang)
	fmt.Printf("\nDaftar Perangkat di Ruangan %s:\n", cariruang)
	j = 0
	for i = 0; i < n; i++ {
		if A[i].aktif && A[i].ruangan == cariruang {
			j++
			if j == 1 {
				fmt.Println("-------------------------------------------------------------------------------------------------")
				fmt.Println("| ID | Nama Perangkat     | Ruangan           | Watt     | Durasi(JAM) | Konsumsi | Status      |")
				fmt.Println("-------------------------------------------------------------------------------------------------")
			}
			if A[i].aktif == true {
				fmt.Printf("| %-3d| %-19s| %-18s| %-9.0f| %-12.1f| %-9.0f| AKTIF       |\n", A[i].id, A[i].nama, A[i].ruangan, A[i].watt, A[i].durasi, A[i].konsumsi)
			} else {
				fmt.Printf("| %-3d| %-19s| %-18s| %-9.0f| %-12.1f| %-9.0f| TIDAK AKTIF |\n", A[i].id, A[i].nama, A[i].ruangan, A[i].watt, A[i].durasi, A[i].konsumsi)
			}
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("-------------------------------------------------------------------------------------------------")
	} else {
		fmt.Println("Tidak ada data perangkat di ruangan tersebut.")
	}
	cari(A, n)
}
func cariberdasarkannama(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menampilkan semua data perangkat yang berstatus aktif dan memiliki atribut 'nama' sesuai dengan input pengguna.
	*/
	var carinama string
	var ketemu bool = false
	var i, j int
	fmt.Println("\n---Pencarian Berdasarkan Nama Perangkat---")
	fmt.Print("Masukkan Nama Perangkat : ")
	fmt.Scan(&carinama)
	fmt.Printf("\nDaftar Perangkat dengan Nama '%s':\n", carinama)
	j = 0
	for i = 0; i < n; i++ {
		if A[i].aktif && A[i].nama == carinama {
			j++
			if j == 1 {
				fmt.Println("-------------------------------------------------------------------------------------------------")
				fmt.Println("| ID | Nama Perangkat     | Ruangan           | Watt     | Durasi(JAM) | Konsumsi | Status      |")
				fmt.Println("-------------------------------------------------------------------------------------------------")
			}
			if A[i].aktif == true {
				fmt.Printf("| %-3d| %-19s| %-18s| %-9.0f| %-12.1f| %-9.0f| AKTIF       |\n", A[i].id, A[i].nama, A[i].ruangan, A[i].watt, A[i].durasi, A[i].konsumsi)
			} else {
				fmt.Printf("| %-3d| %-19s| %-18s| %-9.0f| %-12.1f| %-9.0f| TIDAK AKTIF |\n", A[i].id, A[i].nama, A[i].ruangan, A[i].watt, A[i].durasi, A[i].konsumsi)
			}
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("-------------------------------------------------------------------------------------------------")
	} else {
		fmt.Println("Tidak ada data perangkat dengan nama tersebut.")
	}
	cari(A, n)
}

//fungsi mengurutkan data (Insertion Sort)
func urut(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menampilkan sub-menu urutkan dan memanggil fungsi sorting spesifik sesuai pilihan pengguna.
	*/
	var b, c string
	fmt.Println("-----------------------------------------")
	fmt.Println("| Anda Memasuki Fitur Urutkan Perangkat |")
	fmt.Println("-----------------------------------------")
	fmt.Println("Silahkan Memilih Salah Satu Fitur Dibawah Ini:")
	fmt.Println("1 : Urutkan Berdasarkan ID Perangkat")
	fmt.Println("2 : Urutkan Berdasarkan Nama Perangkat")
	fmt.Println("3 : Urutkan Berdasarkan Ruangan Perangkat")
	fmt.Println("4 : EXIT")
	fmt.Print("Silahkan Masukkan Nomor Fitur (1/2/3/4): ")
	fmt.Scan(&b)
	fmt.Println()
	if b == "1" {
		fmt.Println("Urutkan ID Perangkat Secara:")
		fmt.Println("1 : Ascending")
		fmt.Println("2 : Descending")
		fmt.Print("Masukkan Angka (1/2): ")
		fmt.Scan(&c)
		if c == "1" {
			insertionSortIDASC(A, n)
		} else if c == "2" {
			insertionSortIDDESC(A, n)
		} else {
			fmt.Println("INPUT TIDAK VALID!")
			urut(A, n)
		}
	} else if b == "2" {
		fmt.Println("Urutkan Nama Perangkat Secara:")
		fmt.Println("1 : Ascending")
		fmt.Println("2 : Descending")
		fmt.Print("Masukkan Angka (1/2): ")
		fmt.Scan(&c)
		if c == "1" {
			insertionSortNamaASC(A, n)
		} else if c == "2" {
			insertionSortNamaDESC(A, n)
		} else {
			fmt.Println("INPUT TIDAK VALID!")
			urut(A, n)
		}
	} else if b == "3" {
		fmt.Println("Urutkan Ruangan Perangkat Secara:")
		fmt.Println("1 : Ascending")
		fmt.Println("2 : Descending")
		fmt.Print("Masukkan Angka (1/2): ")
		fmt.Scan(&c)
		if c == "1" {
			insertionSortRuangASC(A, n)
		} else if c == "2" {
			insertionSortRuangDESC(A, n)
		} else {
			fmt.Println("INPUT TIDAK VALID!")
			urut(A, n)
		}
	} else if b == "4" {
		display(A, n)
	} else {
		fmt.Println("INPUT TIDAK VALID!")
		urut(A, n)
	}
}
func insertionSortIDASC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].id > temp.id {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan ID (Ascending).")
	urut(A, n)
}
func insertionSortIDDESC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].id < temp.id {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan ID (Descending).")
	urut(A, n)
}
func insertionSortNamaASC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].nama > temp.nama {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan abjad nama (A-Z).")
	urut(A, n)
}
func insertionSortNamaDESC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].nama < temp.nama {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan abjad nama (Z-A).")
	urut(A, n)
}
func insertionSortRuangASC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].ruangan > temp.ruangan {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan abjad nama (A-Z).")
	urut(A, n)
}
func insertionSortRuangDESC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].ruangan < temp.ruangan {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan abjad nama (Z-A).")
	urut(A, n)
}
func insertionSortKonsumsiASC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].konsumsi > temp.konsumsi {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan banyaknya konsumsi energi (Ascending).")
	urut(A, n)
}
func insertionSortKonsumsiDESC(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan [ID / Nama / Ruangan / Konsumsi] secara [Ascending / Descending] menggunakan metode Insertion Sort.
	*/
	var pass, j int
	var temp Perangkat

	for pass = 1; pass < n; pass++ {
		temp = (A)[pass]
		j = pass - 1
		for j >= 0 && (A)[j].konsumsi < temp.konsumsi {
			(A)[j+1] = (A)[j]
			j--
		}
		(A)[j+1] = temp
	}
	fmt.Println("\n[SUKSES] Data berhasil diurutkan berdasarkan banyaknya konsumsi energi (Descending).")
	urut(A, n)
}

//fungsi menampilkan data perangkat
func tampil(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menampilkan tabel berisi seluruh data log perangkat (aktif maupun tidak aktif) beserta konsumsi listriknya dalam satuan kWh.
	*/
	fmt.Println()
	var i int
	fmt.Println("------------------------------------------------------------------------------------------------------")
	fmt.Println("| ID | Nama Perangkat     | Ruangan           | Watt     | Durasi(JAM) | Konsumsi(kWh) | Status      |")
	fmt.Println("------------------------------------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		if A[i].aktif == true {
			fmt.Printf("| %-3d| %-19s| %-18s| %-9.0f| %-12.1f| %-14.2f| AKTIF       |\n", A[i].id, A[i].nama, A[i].ruangan, A[i].watt, A[i].durasi, A[i].konsumsi/1000)
		} else {
			fmt.Printf("| %-3d| %-19s| %-18s| %-9.0f| %-12.1f| %-14.2f| TIDAK AKTIF |\n", A[i].id, A[i].nama, A[i].ruangan, A[i].watt, A[i].durasi, A[i].konsumsi/1000)
		}
	}
	if n < 1 {
		fmt.Println("| Data log masih kosong                                                                              |")
	}
	fmt.Println("------------------------------------------------------------------------------------------------------\n")
	display(A, n)
}

//fungsi menampilkan statistik perangkat paling boros
func TampilkanStatistik(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Mencetak pesan peringatan beserta saran AI untuk setiap perangkat aktif yang durasi penggunaannya lebih dari 12 jam.
	*/
	var totalWH float64
	var i int
	fmt.Println("\n+------------------------------------------+")
	fmt.Println("|            STATISTIK KONSUMSI            |")
	fmt.Println("+------------------------------------------+")
	for i = 0; i < n; i++ {
		if A[i].aktif {
			totalWH += A[i].konsumsi
		}
	}
	fmt.Printf("Total Konsumsi Energi: %.2f kWh/Hari\n", totalWH/1000)

	fmt.Println("\n+-----------------------------------------------------------+")
	fmt.Println("|            Daftar Paling Boros (Top 5 Teratas)            |")
	fmt.Println("+-----------------------------------------------------------+")

	selectionSortKonsumsiASC(&A, n)
	for i = 0; i < 5 && i < n; i++ {
		if A[i].aktif {
			fmt.Printf("- ID: %-3d | %-15s | Ruang: %-20s | Konsumsi: %.2f kWh\n", A[i].id, A[i].nama, A[i].ruangan, A[i].konsumsi/1000)
		}
	}
	display(A, n)
}
func selectionSortKonsumsiASC(A *ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi. Data di dalam array masih acak.
		F.S. : Elemen di dalam Array A berhasil diurutkan berdasarkan atribut Konsumsi Energi secara Ascending menggunakan metode Selection Sort.
	*/
	var pass, j int
	var temp Perangkat
	var max int

	for pass = 0; pass < n-1; pass++ {
		max = pass
		for j = pass + 1; j < n; j++ {
			if (*A)[j].konsumsi > (*A)[max].konsumsi {
				max = j
			}
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[max]
		(*A)[max] = temp
	}
}

//fungsi cek peringatan durasi penggunaan perangkat
func CekPeringatandurasi(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Mencetak pesan peringatan beserta saran AI untuk setiap perangkat aktif yang durasi penggunaannya lebih dari 12 jam.
	*/
	var adaPerangkat bool
	var i int

	fmt.Println("\n+--------------------------+")
	fmt.Println("| SISTEM PERINGATAN DURASI |")
	fmt.Println("+--------------------------+")

	for i = 0; i < n; i++ {
		if A[i].aktif && A[i].durasi > 12 {
			adaPerangkat = true
			fmt.Printf("\n[PERINGATAN] Perangkat '%s' di ruang '%s' menyala selama %.1f jam.\n", A[i].nama, A[i].ruangan, A[i].durasi)
			fmt.Println("Saran AI: Durasi ini terlalu lama untuk penggunaan harian. Harap matikan jika tidak digunakan atau pasang timer otomatis.")
		}
	}
	if !adaPerangkat {
		fmt.Println("\n Semua perangkat menyala dalam batas durasi yang wajar (di bawah 12 jam).")
	}
	fmt.Println("--------------------------------------------")
	display(A, n)
}

//fungsi evaluasi batas aman penggunaan listrik bulanan
func EvaluasiBatasAman(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Mengevaluasi dan mencetak nama perangkat aktif yang estimasi konsumsi bulanannya melebihi batas (threshold) yang diinputkan pengguna.
	*/
	var i int
	var konsumsiBulananKWh float64
	var thresholdKWh float64
	var adaPelanggaran bool
	fmt.Println("\n+---------------------------------+")
	fmt.Println("|ANALISIS BATAS AMAN BULANAN (kWh)|")
	fmt.Println("+---------------------------------+")
	fmt.Print("Masukkan batas aman bulanan perangkat (kWh): ")
	fmt.Scan(&thresholdKWh)
	fmt.Println("Memindai perangkat...")
	fmt.Println("--------------------------------------------")
	for i = 0; i < n; i++ {
		if A[i].aktif {
			konsumsiBulananKWh = (A[i].konsumsi * 30) / 1000
			if konsumsiBulananKWh > thresholdKWh {
				adaPelanggaran = true
				fmt.Printf("\n[EVALUASI] Perangkat '%s' di '%s' terdeteksi boros!\n", A[i].nama, A[i].ruangan)
				fmt.Printf("Estimasi bulanan: %.2f kWh (Melebihi batas aman %.2f kWh).\n", konsumsiBulananKWh, thresholdKWh)
				fmt.Println("Rekomendasi: Perangkat ini wajib dievaluasi! Pertimbangkan untuk mengurangi durasi harian atau ganti ke perangkat yang lebih hemat.")
				fmt.Println("--------------------------------------------")
			}
		}
	}
	if !adaPelanggaran {
		fmt.Printf("\n Semua perangkat aktif masih berada di bawah batas aman %.2f kWh/bulan.\n", thresholdKWh)
		fmt.Println("--------------------------------------------")
	}
	display(A, n)
}

//fungsi menghitung biaya penggunaan listrik dalam sebulan
func hitungbiaya(A ArrayPerangkat, n int) {
	/*
		I.S. : Array A dan n terdefinisi.
		F.S. : Menghitung estimasi total konsumsi listrik dalam sebulan (dalam kWh) dan mencetak taksiran biaya dengan tarif Rp.1444/kWh.
	*/
	var i, watt, kWH int
	for i = 0; i < n; i++ {
		if A[i].aktif == true {
			watt = watt + int(A[i].konsumsi)
		}
	}
	kWH = watt * 30 / 1000
	fmt.Println("\nTotal Konsumsi seluruh perangkat elektronik dalam sebulan adalah sebesar", kWH, "kWH")
	fmt.Printf("Atau dibutuhkan budget sebesar Rp.%d dalam sebulan\n\n", kWH*1444)
	display(A, n)
}

/*
Isi Data Demo:
1
1
AC 800 10 Kamar_Utama 1
TV 150 4 Kamar_Utama 1
Lampu 20 8 Kamar_Utama 1
Lampu 20 8 Kamar_Utama 1
Lampu_Tidur 10 8 Kamar_Utama 1
Kipas_Angin 50 4 Kamar_Utama 1
Charger 15 5 Kamar_Utama 1
Water_Heater 350 2 Kamar_Mandi_Utama 1
Exhaust_Fan 30 2 Kamar_Mandi_Utama 1
Lampu 15 3 Kamar_Mandi_Utama 1
AC 400 8 Kamar_Anak_1 1
Lampu 15 8 Kamar_Anak_1 1
PC 450 6 Kamar_Anak_1 1
Monitor 40 6 Kamar_Anak_1 1
Lampu_Belajar 10 4 Kamar_Anak_1 1
AC 400 8 Kamar_Anak_2 1
Lampu 15 8 Kamar_Anak_2 1
TV 100 4 Kamar_Anak_2 1
Konsol_Game 200 4 Kamar_Anak_2 1
Kipas_Angin 50 6 Kamar_Anak_2 1
Water_Heater 350 2 Kamar_Mandi_Anak 1
Exhaust_Fan 30 2 Kamar_Mandi_Anak 1
Lampu 15 3 Kamar_Mandi_Anak 1
AC 1000 6 Ruang_Keluarga 1
TV 200 6 Ruang_Keluarga 1
Home_Theater 250 4 Ruang_Keluarga 1
Router_WiFi 15 24 Ruang_Keluarga 1
Lampu 25 6 Ruang_Keluarga 1
Lampu 25 6 Ruang_Keluarga 1
Kulkas 150 24 Dapur 1
Microwave 800 0.5 Dapur 1
Rice_Cooker 350 2 Dapur 1
Air_Fryer 900 0.5 Dapur 1
Dispenser 250 24 Dapur 1
Lampu 15 6 Dapur 1
Lampu 20 6 Ruang_Makan 1
Lampu 20 6 Ruang_Makan 1
Kipas_Angin 60 4 Ruang_Tamu 1
Lampu 20 6 Ruang_Tamu 1
Lampu 20 6 Ruang_Tamu 1
AC 400 6 Ruang_Kerja 1
PC 450 8 Ruang_Kerja 1
Printer 30 1 Ruang_Kerja 1
Lampu 15 8 Ruang_Kerja 1
Mesin_Cuci 400 2 Ruang_Cuci 1
Setrika 350 2 Ruang_Cuci 1
Pompa_Air 250 3 Garasi 1
Lampu 15 12 Teras 1
Lampu 15 12 Teras 1
CCTV 10 24 Halaman 2
5
*/
