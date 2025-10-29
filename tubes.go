package main

import "fmt"

const NMAX int = 100

type film struct {
	Nama, Waktu, Kategori string
	Harga, Stok           int
}
type riwayat struct {
	Nama, Kategori, Waktu string
	Harga                 int
}
type tabFilm [NMAX]film
type tabRiwayat [NMAX]riwayat

func main() {
	var DaftarFilm tabFilm
	var RiwayatPemesanan tabRiwayat
	var nFilm, nRiwayat int

	DaftarFilm[0] = film{"Avengers:Endgame", "20:00", "Action", 50000, 10}
	DaftarFilm[1] = film{"Final Destination", "18:30", "Thriller", 45000, 8}
	DaftarFilm[2] = film{"Interstellar", "21:00", "Sci-Fi", 60000, 5}
	DaftarFilm[3] = film{"Jumbo", "17:00", "Drama", 40000, 12}
	nFilm = 4

	for {
		var pilihRole int
		fmt.Println("\n====== TIKS.ID ======")
		fmt.Println("1. Admin")
		fmt.Println("2. Pelanggan")
		fmt.Print("Siapa Anda?")
		fmt.Scan(&pilihRole)

		switch pilihRole {
		case 1:
			admin(&DaftarFilm, &nFilm)
		case 2:
			pelanggan(&DaftarFilm, &RiwayatPemesanan, &nFilm, &nRiwayat)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func admin(DaftarFilm *tabFilm, n *int) {
	for {
		var pilihAdmin int
		fmt.Println("\n====== Menu Admin ======")
		fmt.Println("1. Lihat Daftar Film")
		fmt.Println("2. Tambah Film")
		fmt.Println("3. Hapus Film")
		fmt.Println("4. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihAdmin)

		switch pilihAdmin {
		case 1:
			tampilkanFilm(*DaftarFilm, *n)
		case 2:
			tambahFilm(DaftarFilm, n)
		case 3:
			hapusFilm(DaftarFilm, n)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func pelanggan(Daftarfilm *tabFilm, RiwayatPemesanan *tabRiwayat, nFilm *int, nRiwayat *int) {
	for {
		var pilih int
		fmt.Println("\n====== Menu Pelanggan ======")
		fmt.Println("1. Lihat Daftar Film")
		fmt.Println("2. Cari Film")
		fmt.Println("3. Pesan Tiket")
		fmt.Println("4. Tampilkan Riwayat Pemesanan")
		fmt.Println("5. Kembali")
		fmt.Println("6. Cari Film Berdasarkan Kategori")
		fmt.Println("7. Urutkan Film Berdasarkan Harga (Termurah)")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tampilkanFilm(*Daftarfilm, *nFilm)
		case 2:
			cariFilm(*Daftarfilm, *nFilm)
		case 3:
			pesanTiket(Daftarfilm, RiwayatPemesanan, nFilm, nRiwayat)
		case 4:
			tampilkanRiwayat(*RiwayatPemesanan, *nRiwayat)
		case 5:
			return
		case 6:
			cariBerdasarkanKategori(*Daftarfilm, *nFilm)
		case 7:
			urutkanFilmBerdasarkanHarga(Daftarfilm, *nFilm)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanFilm(A tabFilm, n int) {
	fmt.Printf("==================================================================\n")
	fmt.Printf("%-25s %-10s %-15s %-10s %-5s\n", "Nama", "Waktu", "Kategori", "Harga", "Stok")
	fmt.Println("==================================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("%-25s %-10s %-15s Rp%-9d %-5d\n", A[i].Nama, A[i].Waktu, A[i].Kategori, A[i].Harga, A[i].Stok)
	}
}

func tambahFilm(A *tabFilm, n *int) {
	if *n >= NMAX {
		fmt.Println("Maaf, kapasitas penuh.")
		return
	}

	var nama, waktu, kategori string
	var harga, stok, pilihKategori int

	fmt.Print("Nama Film: ")
	fmt.Scan(&nama)
	fmt.Print("Waktu: ")
	fmt.Scan(&waktu)

	fmt.Print("Kategori: ")
	fmt.Println("1. Action")
	fmt.Println("2. Drama")
	fmt.Println("3. Horror")
	fmt.Println("4. Romance")
	fmt.Println("5. Sci-fi")
	fmt.Println("6. Thriller")
	fmt.Scan(&pilihKategori)
	switch pilihKategori {
	case 1:
		kategori = "Action"
	case 2:
		kategori = "Drama"
	case 3:
		kategori = "Horror"
	case 4:
		kategori = "Romance"
	case 5:
		kategori = "Thriller"
	default:
		fmt.Println("Pilihan tidak valid.")
	}

	fmt.Print("Harga: ")
	fmt.Scan(&harga)
	fmt.Print("Stok: ")
	fmt.Scan(&stok)

	(*A)[*n] = film{nama, waktu, kategori, harga, stok}
	*n++
	fmt.Println("Film berhasil ditambahkan.")
}

func hapusFilm(A *tabFilm, n *int) {
	var nama string
	var found int

	fmt.Print("Masukkan nama film yang ingin dihapus: ")
	fmt.Scan(&nama)

	found = -1
	for i := 0; i < *n; i++ {
		if (*A)[i].Nama == nama {
			found = 1
			(*A)[i] = (*A)[i+1]
		}
	}

	if found == -1 {
		fmt.Println("Film tidak ditemukan.")
		return
	}
	*n--
	fmt.Println("Film berhasil dihapus.")
}

func cariFilm(A tabFilm, n int) {
	var nama string
	var found int

	fmt.Print("Masukkan nama film yang ingin Anda cari(Case-Sensitive): ")
	fmt.Scan(&nama)

	found = -1
	fmt.Println("================")
	fmt.Println("Hasil Pencarian:")
	fmt.Println("================")

	for i := 0; i < n; i++ {
		if cocokSubstr(A[i].Nama, nama) {
			fmt.Printf("%-25s %-10s %-15s Rp%-9d %-5d\n", A[i].Nama, A[i].Waktu, A[i].Kategori, A[i].Harga, A[i].Stok)
			found = 1
		}
	}

	if found == -1 {
		fmt.Println("Film tidak ditemukan.")
	}
}

func cocokSubstr(kata string, substring string) bool {
	var nSubstr, nKata int

	nSubstr = len(substring)
	nKata = len(kata)

	if nSubstr > nKata {
		return false
	}
	for i := 0; i < nSubstr; i++ {
		if kata[i] != substring[i] {
			return false
		}
	}
	return true
}

/*
	var x string
	var found int

	fmt.Print("Masukkan nama film yang ingin Anda cari: ")
	fmt.Scan(&x)
	found = -1
	fmt.Println("Hasil Pencarian:")
	for i := 0; i < n; i++ {
		if A[i].Nama == x {
			fmt.Printf("%-25s %-10s %-15s Rp%-9d %-5d\n", A[i].Nama, A[i].Waktu, A[i].Kategori, A[i].Harga, A[i].Stok)
			found = 1
		}
	}
	if found == -1 {
		fmt.Println("Film tidak ditemukan.")
	}
*/

func cariBerdasarkanKategori(A tabFilm, n int) {
	var pilihKategori, found int
	var kategori string

	fmt.Println("Masukkan kategori yang ingin dicari: ")
	fmt.Println("1. Action")
	fmt.Println("2. Drama")
	fmt.Println("3. Horror")
	fmt.Println("4. Romance")
	fmt.Println("5. Sci-fi")
	fmt.Println("6. Thriller")

	fmt.Scan(&pilihKategori)
	switch pilihKategori {
	case 1:
		kategori = "Action"
	case 2:
		kategori = "Drama"
	case 3:
		kategori = "Horror"
	case 4:
		kategori = "Romance"
	case 5:
		kategori = "Thriller"
	default:
		fmt.Println("Pilihan tidak valid.")
	}

	found = -1
	fmt.Println("=====================================")
	fmt.Println("Hasil Pencarian Berdasarkan Kategori:")
	fmt.Println("=====================================")

	for i := 0; i < n; i++ {
		if A[i].Kategori == kategori {
			fmt.Printf("%-25s %-10s %-15s Rp%-9d %-5d\n", A[i].Nama, A[i].Waktu, A[i].Kategori, A[i].Harga, A[i].Stok)
			found = 1
		}
	}
	if found == -1 {
		fmt.Println("Tidak ada film dengan kategori tersebut.")
	}
}

func urutkanFilmBerdasarkanHarga(A *tabFilm, n int) {
	var i, idx, pass int
	var temp film

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if (*A)[i].Harga < (*A)[idx].Harga {
				idx = i
			}
		}
		// Tukar posisi
		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp
	}
	fmt.Println("Film berhasil diurutkan berdasarkan harga (termurah ke termahal):")
	tampilkanFilm(*A, n)
}

func pesanTiket(A *tabFilm, B *tabRiwayat, nFilm *int, nRiwayat *int) {
	var nama string

	fmt.Print("Masukkan nama film yang ingin Anda pesan: ")
	fmt.Scan(&nama)

	for i := 0; i < *nFilm; i++ {
		if (*A)[i].Nama == nama {
			if (*A)[i].Stok > 0 {
				(*A)[i].Stok--
				(*B)[*nRiwayat] = riwayat{
					Nama:     (*A)[i].Nama,
					Kategori: (*A)[i].Kategori,
					Waktu:    (*A)[i].Waktu,
					Harga:    (*A)[i].Harga,
				}
				*nRiwayat++
				fmt.Println("Tiket berhasil dipesan.")
				return
			} else {
				fmt.Println("Maaf, stok tiket habis.")
				return
			}
		}
	}
	fmt.Println("Film tidak ditemukan.")
}

func tampilkanRiwayat(B tabRiwayat, n int) {
	if n == 0 {
		fmt.Println("Belum ada pemesanan.")
		return
	}

	fmt.Println("====== Riwayat Pemesanan ======")
	fmt.Printf("%-25s %-15s %-10s Rp%-9s\n", "Nama", "Kategori", "Waktu", "Harga")
	fmt.Println("----------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%-25s %-15s %-10s Rp%-9d\n", B[i].Nama, B[i].Kategori, B[i].Waktu, B[i].Harga)
	}
}
