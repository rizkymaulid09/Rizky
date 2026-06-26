package main

import "fmt"

type Film struct {
	ID        int
	Judul     string
	Genre     string
	Tahun     int
	Deskripsi string
	Rating    float64
}

var data [100]Film
var total int = 0
var nextID int = 1

func kecil(s string) string {
	out := ""
	for i := 0; i < len(s); i++ {
		h := s[i]
		if h >= 'A' && h <= 'Z' {
			out += string(h + 32)
		} else {
			out += string(h)
		}
	}
	return out
}

func ada(s, sub string) bool {
	if len(sub) > len(s) {
		return false
	}
	for i := 0; i <= len(s)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func bandingString(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func bacaString(prompt string) string {
	fmt.Print(prompt)
	var s string
	fmt.Scan(&s)
	return s
}

func bacaInt(prompt string) int {
	var n int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&n)
		if err == nil {
			return n
		}
		var buang string
		fmt.Scanln(&buang)
		fmt.Println("  Angka dulu ya!")
	}
}

func bacaFloat(prompt string) float64 {
	var f float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&f)
		if err == nil && f >= 0 && f <= 10 {
			return f
		}
		var buang string
		fmt.Scanln(&buang)
		fmt.Println("  Rating harus 0.0 - 10.0!")
	}
}

func cetakFilm(f Film) {
	fmt.Printf("  [%d] %s (%d) | %s | Rating: %.1f/10\n", f.ID, f.Judul, f.Tahun, f.Genre, f.Rating)
	fmt.Printf("      %s\n", f.Deskripsi)
	fmt.Println("  ---")
}

func lihatSemua() {
	fmt.Println("\n=== Koleksi Film ===")
	if total == 0 {
		fmt.Println("  Belum ada film nih.")
		return
	}
	for i := 0; i < total; i++ {
		cetakFilm(data[i])
	}
}

func gantiUnderscore(s string) string {
	out := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '_' {
			out += " "
		} else {
			out += string(s[i])
		}
	}
	return out
}

func tambah() {
	fmt.Println("\n=== Tambah Film ===")
	if total >= 100 {
		fmt.Println("  Udah penuh!")
		return
	}
	fmt.Println("  (gunakan underscore _ untuk spasi, contoh: The_Dark_Knight)")
	judul := bacaString("  Judul    : ")
	genre := bacaString("  Genre    : ")
	tahun := bacaInt("  Tahun    : ")
	desk := bacaString("  Deskripsi: ")
	rate := bacaFloat("  Rating   : ")

	judul = gantiUnderscore(judul)
	genre = gantiUnderscore(genre)
	desk = gantiUnderscore(desk)

	data[total] = Film{nextID, judul, genre, tahun, desk, rate}
	total++
	nextID++
	fmt.Println("  Berhasil ditambahkan!")
}

func cariIndex(id int) int {
	for i := 0; i < total; i++ {
		if data[i].ID == id {
			return i
		}
	}
	return -1
}

func ubah() {
	fmt.Println("\n=== Ubah Film ===")
	lihatSemua()
	id := bacaInt("  ID film yang mau diubah: ")
	idx := cariIndex(id)
	if idx == -1 {
		fmt.Println("  Film tidak ditemukan!")
		return
	}
	fmt.Println("  (ketik 0 untuk skip, gunakan _ untuk spasi)")
	judul := bacaString("  Judul baru    : ")
	genre := bacaString("  Genre baru    : ")
	tahun := bacaInt("  Tahun baru    : ")
	desk := bacaString("  Deskripsi baru: ")
	rate := bacaFloat("  Rating baru   : ")

	if judul != "0" {
		data[idx].Judul = gantiUnderscore(judul)
	}
	if genre != "0" {
		data[idx].Genre = gantiUnderscore(genre)
	}
	if tahun != 0 {
		data[idx].Tahun = tahun
	}
	if desk != "0" {
		data[idx].Deskripsi = gantiUnderscore(desk)
	}
	if rate != 0 {
		data[idx].Rating = rate
	}
	fmt.Println("  Data diperbarui!")
}

func hapus() {
	fmt.Println("\n=== Hapus Film ===")
	if total == 0 {
		fmt.Println("  Belum ada film nih.")
		return
	}
	lihatSemua()
	for {
		id := bacaInt("  ID film yang mau dihapus: ")
		idx := cariIndex(id)
		if idx == -1 {
			fmt.Println("  Film tidak ditemukan, coba ID lain!")
			continue
		}
		for i := idx; i < total-1; i++ {
			data[i] = data[i+1]
		}
		total--

		for i := 0; i < total; i++ {
			data[i].ID = i + 1
		}
		nextID = total + 1

		fmt.Println("  Film dihapus!")
		break
	}
}

func cariJudul(keyword string) {
	keyword = kecil(gantiUnderscore(keyword))
	ketemu := false
	fmt.Println("\n  [Sequential - Judul]")
	for i := 0; i < total; i++ {
		if ada(kecil(data[i].Judul), keyword) {
			cetakFilm(data[i])
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("  Tidak ketemu.")
	}
}

func cariGenre(keyword string) {
	keyword = kecil(gantiUnderscore(keyword))
	ketemu := false
	fmt.Println("\n  [Sequential - Genre]")
	for i := 0; i < total; i++ {
		if ada(kecil(data[i].Genre), keyword) {
			cetakFilm(data[i])
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("  Tidak ketemu.")
	}
}

func salin() [100]Film {
	var tmp [100]Film
	for i := 0; i < total; i++ {
		tmp[i] = data[i]
	}
	return tmp
}

func binaryJudul(keyword string) {
	tmp := salin()
	for i := 1; i < total; i++ {
		k := tmp[i]
		j := i - 1
		for j >= 0 && bandingString(kecil(tmp[j].Judul), kecil(k.Judul)) > 0 {
			tmp[j+1] = tmp[j]
			j--
		}
		tmp[j+1] = k
	}
	keyword = kecil(gantiUnderscore(keyword))
	lo, hi := 0, total-1
	ketemu := false
	fmt.Println("\n  [Binary Search - Judul]")
	for lo <= hi {
		mid := (lo + hi) / 2
		cmp := bandingString(kecil(tmp[mid].Judul), keyword)
		if cmp == 0 {
			cetakFilm(tmp[mid])
			ketemu = true
			for i := mid - 1; i >= 0 && bandingString(kecil(tmp[i].Judul), keyword) == 0; i-- {
				cetakFilm(tmp[i])
			}
			for j := mid + 1; j < total && bandingString(kecil(tmp[j].Judul), keyword) == 0; j++ {
				cetakFilm(tmp[j])
			}
			break
		} else if cmp < 0 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if !ketemu {
		fmt.Println("  Tidak ketemu.")
	}
}

func menuCari() {
	fmt.Println("\n=== Cari Film ===")
	fmt.Println("  1. Sequential - Judul")
	fmt.Println("  2. Sequential - Genre")
	fmt.Println("  3. Binary     - Judul")
	pilih := bacaInt("  Pilih: ")
	keyword := bacaString("  Kata kunci: ")
	switch pilih {
	case 1:
		cariJudul(keyword)
	case 2:
		cariGenre(keyword)
	case 3:
		binaryJudul(keyword)
	default:
		fmt.Println("  Pilihan salah.")
	}
}

func sortRating() {
	tmp := salin()
	for i := 0; i < total-1; i++ {
		maks := i
		for j := i + 1; j < total; j++ {
			if tmp[j].Rating > tmp[maks].Rating {
				maks = j
			}
		}
		tmp[i], tmp[maks] = tmp[maks], tmp[i]
	}
	fmt.Println("\n  [Rating Tertinggi ke Terendah]")
	for i := 0; i < total; i++ {
		cetakFilm(tmp[i])
	}
}

func sortTahun() {
	tmp := salin()
	for i := 1; i < total; i++ {
		k := tmp[i]
		j := i - 1
		for j >= 0 && tmp[j].Tahun > k.Tahun {
			tmp[j+1] = tmp[j]
			j--
		}
		tmp[j+1] = k
	}
	fmt.Println("\n  [Tahun Terlama ke Terbaru]")
	for i := 0; i < total; i++ {
		cetakFilm(tmp[i])
	}
}

func sortTahunTerbaru() {
	tmp := salin()
	for i := 1; i < total; i++ {
		k := tmp[i]
		j := i - 1
		for j >= 0 && tmp[j].Tahun < k.Tahun {
			tmp[j+1] = tmp[j]
			j--
		}
		tmp[j+1] = k
	}
	fmt.Println("\n  [Tahun Terbaru ke Terlama]")
	for i := 0; i < total; i++ {
		cetakFilm(tmp[i])
	}
}

func menuUrut() {
	fmt.Println("\n=== Urutkan Film ===")
	fmt.Println("  1. Rating tertinggi ke terendah")
	fmt.Println("  2. Tahun terlama ke terbaru")
	fmt.Println("  3. Tahun terbaru ke terlama")
	pilih := bacaInt("  Pilih: ")
	switch pilih {
	case 1:
		sortRating()
	case 2:
		sortTahun()
	case 3:
		sortTahunTerbaru()
	default:
		fmt.Println("  Pilihan salah.")
	}
}

func statistik() {
	fmt.Println("\n=== Statistik ===")
	if total == 0 {
		fmt.Println("  Belum ada data.")
		return
	}
	var gList [100]string
	var gCount [100]int
	jg := 0
	totalRate := 0.0

	for i := 0; i < total; i++ {
		totalRate += data[i].Rating
		g := data[i].Genre
		found := false
		for j := 0; j < jg; j++ {
			if gList[j] == g {
				gCount[j]++
				found = true
				break
			}
		}
		if !found {
			gList[jg] = g
			gCount[jg] = 1
			jg++
		}
	}

	fmt.Println("  Film per Genre:")
	for i := 0; i < jg; i++ {
		fmt.Printf("    %-20s: %d film\n", gList[i], gCount[i])
	}
	fmt.Printf("\n  Total Film      : %d\n", total)
	fmt.Printf("  Rata-rata Rating: %.2f/10\n", totalRate/float64(total))
}

func rekomendasiTertinggi() {
	fmt.Println("\n  [Top 5 Film Rating Tertinggi]")
	if total == 0 {
		fmt.Println("  Belum ada film nih.")
		return
	}
	tmp := salin()

	for i := 0; i < total-1; i++ {
		maks := i
		for j := i + 1; j < total; j++ {
			if tmp[j].Rating > tmp[maks].Rating {
				maks = j
			}
		}
		tmp[i], tmp[maks] = tmp[maks], tmp[i]
	}

	tampil := total
	if tampil > 5 {
		tampil = 5
	}
	for i := 0; i < tampil; i++ {
		fmt.Printf("  #%d\n", i+1)
		cetakFilm(tmp[i])
	}
}

func rekomendasiGenre() {
	fmt.Println("\n  [Rekomendasi Berdasarkan Genre]")
	if total == 0 {
		fmt.Println("  Belum ada film nih.")
		return
	}

	var gList [100]string
	jg := 0
	for i := 0; i < total; i++ {
		g := data[i].Genre
		found := false
		for j := 0; j < jg; j++ {
			if gList[j] == g {
				found = true
				break
			}
		}
		if !found {
			gList[jg] = g
			jg++
		}
	}

	fmt.Println("  Genre yang tersedia:")
	for i := 0; i < jg; i++ {
		fmt.Printf("    - %s\n", gList[i])
	}

	keyword := bacaString("  Masukkan genre: ")
	keyword = kecil(gantiUnderscore(keyword))

	var hasil [100]Film
	jumlahHasil := 0
	for i := 0; i < total; i++ {
		if ada(kecil(data[i].Genre), keyword) {
			hasil[jumlahHasil] = data[i]
			jumlahHasil++
		}
	}

	if jumlahHasil == 0 {
		fmt.Println("  Tidak ada film dengan genre tersebut.")
		return
	}

	for i := 0; i < jumlahHasil-1; i++ {
		maks := i
		for j := i + 1; j < jumlahHasil; j++ {
			if hasil[j].Rating > hasil[maks].Rating {
				maks = j
			}
		}
		hasil[i], hasil[maks] = hasil[maks], hasil[i]
	}

	fmt.Printf("\n  Rekomendasi film genre \"%s\" (rating tertinggi):\n", gantiUnderscore(keyword))
	for i := 0; i < jumlahHasil; i++ {
		fmt.Printf("  #%d\n", i+1)
		cetakFilm(hasil[i])
	}
}

func menuRekomendasi() {
	fmt.Println("\n=== Rekomendasi Film ===")
	fmt.Println("  1. Top 5 rating tertinggi")
	fmt.Println("  2. Rekomendasi berdasarkan genre")
	pilih := bacaInt("  Pilih: ")
	switch pilih {
	case 1:
		rekomendasiTertinggi()
	case 2:
		rekomendasiGenre()
	default:
		fmt.Println("  Pilihan salah.")
	}
}

func main() {
	total = 0
	nextID = 1

	for {
		fmt.Println(" ============================== ")
		fmt.Println("           CineReview           ")
		fmt.Println(" ============================== ")
		fmt.Println("  1. Lihat Semua")
		fmt.Println("  2. Tambah Film")
		fmt.Println("  3. Ubah Film")
		fmt.Println("  4. Hapus Film")
		fmt.Println("  5. Cari Film")
		fmt.Println("  6. Urutkan Film")
		fmt.Println("  7. Statistik")
		fmt.Println("  8. Rekomendasi Film")
		fmt.Println("  0. Keluar")
		fmt.Println(" ============================== ")
		pilih := bacaInt("  Pilih: ")

		switch pilih {
		case 1:
			lihatSemua()
		case 2:
			tambah()
		case 3:
			ubah()
		case 4:
			hapus()
		case 5:
			menuCari()
		case 6:
			menuUrut()
		case 7:
			statistik()
		case 8:
			menuRekomendasi()
		case 0:
			fmt.Println(" ============================== ")
			fmt.Println("     Terima kasih, arigatou!    ")
			fmt.Println(" ============================== ")
			return
		default:
			fmt.Println("  Pilihan salah, coba lagi.")
		}

		fmt.Print("\n  Enter untuk lanjut...")
		fmt.Scanln()
		fmt.Scanln()
	}
}
