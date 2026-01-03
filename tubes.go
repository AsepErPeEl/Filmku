package main

import "fmt"

const NMAX int = 100

type Film struct {
	Judul  string
	Genre  string
	Rating float64
	Status string
}

type tabFilm [NMAX]Film

func bacaData(A *tabFilm, N *int) {
	fmt.Print("Jumlah film: ")
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Print("Judul  : ")
		fmt.Scan(&A[i].Judul)   
		fmt.Print("Genre  : ")
		fmt.Scan(&A[i].Genre)   
		fmt.Print("Rating : ")
		fmt.Scan(&A[i].Rating)  
		fmt.Print("Status (Ditonton/Belum): ")
		fmt.Scan(&A[i].Status)  
	}
}

func cetakData(A tabFilm, N int) {
	fmt.Println("\n+----------------------+------------+--------+-----------+")
	fmt.Println("|        Judul         |   Genre    | Rating |  Status   |")
	fmt.Println("+----------------------+------------+--------+-----------+")
	for i := 0; i < N; i++ {
		fmt.Printf("| %-20s | %-10s | %6.1f | %-9s |\n",
			A[i].Judul, A[i].Genre, A[i].Rating, A[i].Status)
	}
	fmt.Println("+----------------------+------------+--------+-----------+")
}

// sequen
func cariGenre(A tabFilm, N int, genre string) {
	fmt.Println("\nHasil Pencarian:")
	fmt.Println("+----------------------+------------+--------+-----------+")
	fmt.Println("|        Judul         |   Genre    | Rating |  Status   |")
	fmt.Println("+----------------------+------------+--------+-----------+")
	for i := 0; i < N; i++ {
		if A[i].Genre == genre {
			fmt.Printf("| %-20s | %-10s | %6.1f | %-9s |\n",
				A[i].Judul, A[i].Genre, A[i].Rating, A[i].Status)
		}
	}
	fmt.Println("+----------------------+------------+--------+-----------+")
}

// selection sort
func sortJudulAsc(A *tabFilm, N int) {
	for i := 0; i < N-1; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if A[j].Judul < A[min].Judul {
				min = j
			}
		}
		A[i], A[min] = A[min], A[i]
	}
}

// insertion sort
func sortRatingDesc(A *tabFilm, N int) {
	for i := 1; i < N; i++ {
		temp := A[i]
		j := i - 1
		for j >= 0 && A[j].Rating < temp.Rating {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

// binsearch
func binarySearchJudul(A tabFilm, N int, judul string) int {
	kiri := 0
	kanan := N - 1
	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if A[mid].Judul == judul {
			return mid
		} else if A[mid].Judul < judul {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}
	return -1
}

// ngedit data film
func editFilm(A *tabFilm, N int, judul string) {
	sortJudulAsc(A, N)
	idx := binarySearchJudul(*A, N, judul)
	if idx != -1 {
		fmt.Print("Judul baru: ")
		fmt.Scan(&A[idx].Judul)
		fmt.Print("Genre baru: ")
		fmt.Scan(&A[idx].Genre)
		fmt.Print("Rating baru: ")
		fmt.Scan(&A[idx].Rating)
		fmt.Print("Status baru: ")
		fmt.Scan(&A[idx].Status)
		fmt.Println("Data film berhasil diperbarui.")
	} else {
		fmt.Println("Film tidak ditemukan.")
	}
}

// hapus data film
func hapusFilm(A *tabFilm, N *int, judul string) {
	sortJudulAsc(A, *N)
	idx := binarySearchJudul(*A, *N, judul)
	if idx != -1 {
		for i := idx; i < *N-1; i++ {
			A[i] = A[i+1]
		}
		*N = *N - 1
		fmt.Println("Film berhasil dihapus.")
	} else {
		fmt.Println("Film tidak ditemukan.")
	}
}

// tampilan menu
func menu() {
	fmt.Println("\n=== MENU FILM ===")
	fmt.Println("1. Tambah Film")
	fmt.Println("2. Tampilkan Film")
	fmt.Println("3. Cari Film Berdasarkan Genre")
	fmt.Println("4. Urut Judul ASC")
	fmt.Println("5. Urut Rating DESC")
	fmt.Println("6. Edit Film")
	fmt.Println("7. Hapus Film")
	fmt.Println("0. Keluar")
}

func main() {
	var data tabFilm
	var N, pilih int
	var input string

	for {
		menu()
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			bacaData(&data, &N)
		case 2:
			cetakData(data, N)
		case 3:
			fmt.Print("Genre: ")
			fmt.Scan(&input)
			cariGenre(data, N, input)
		case 4:
			sortJudulAsc(&data, N)
			cetakData(data, N)
		case 5:
			sortRatingDesc(&data, N)
			cetakData(data, N)
		case 6:
			fmt.Print("Judul: ")
			fmt.Scan(&input)
			editFilm(&data, N, input)
		case 7:
			fmt.Print("Judul: ")
			fmt.Scan(&input)
			hapusFilm(&data, &N, input)
		case 0:
			fmt.Println("Keluar program. Terima kasih!")
		default:
			fmt.Println("Pilihan tidak tersedia, coba lagi.")
		}
	}
}
