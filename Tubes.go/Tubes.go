package main

import "fmt"

const NMAX int = 100

type registration struct {
	username, email, password, confirmPass string
}

type Barang struct {
	sparePart string
	harga     int
}

type Pengguna struct {
	idPengguna, nama string
}

type Transaksi struct {
	idTransaksi string
	idPelanggan string
	namaBarang  string
	jumlah      int
	total       int
}

type TabRegist [NMAX]registration
type TabBarang [NMAX]Barang
type TabPengguna [NMAX]Pengguna
type TabTransaksi [NMAX]Transaksi

func main() {
	var regist TabRegist
	var barang TabBarang
	var pengguna TabPengguna
	var transaksi TabTransaksi
	var nBarang, nPengguna, nTransaksi int
	var pilih int

	for pilih != 2 {
		fmt.Println("---------------------------------")
		fmt.Println("    Aplikasi Service Motor X    ")
		fmt.Println("---------------------------------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			Regist(&regist, 0)
			if Login(&regist, 0) {
				MenuUtama(&barang, &nBarang, &pengguna, &nPengguna, &transaksi, &nTransaksi)
			}
		} else if pilih == 2 {
			fmt.Println("Terimakasih")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func Regist(regist *TabRegist, n int) {
	fmt.Println("---------------------------------")
	fmt.Println("Silahkan melakukan registrasi")
	fmt.Print("Username: ")
	fmt.Scan(&regist[n].username)
	fmt.Print("Email: ")
	fmt.Scan(&regist[n].email)
	fmt.Print("Password: ")
	fmt.Scan(&regist[n].password)
	fmt.Print("Konfirmasi Password: ")
	fmt.Scan(&regist[n].confirmPass)

	for regist[n].confirmPass != regist[n].password {
		fmt.Println("---------------------------------")
		fmt.Println("Konfirmasi Password Tidak Valid")
		fmt.Print("Password: ")
		fmt.Scan(&regist[n].password)
		fmt.Print("Konfirmasi Password: ")
		fmt.Scan(&regist[n].confirmPass)
	}
	fmt.Println("---------------------------------")
	fmt.Println("   Akun anda berhasil dibuat   ")
}

func Login(regist *TabRegist, n int) bool {
	var username, password string

	for username != regist[n].username || password != regist[n].password {
		fmt.Println("---------------------------------")
		fmt.Println("Silahkan Login")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		if username == regist[n].username && password == regist[n].password {
			fmt.Println("---------------------------------")
			fmt.Print("  Halo ", username)
			fmt.Println(", Selamat Datang   ")
			fmt.Println("---------------------------------")
			return true
		}
		fmt.Println("Username atau password anda salah")
	}
	return false
}

func MenuUtama(barang *TabBarang, nBarang *int, pengguna *TabPengguna, nPengguna *int, transaksi *TabTransaksi, nTransaksi *int) {
	var opsiPilihan int

	for opsiPilihan != 4 {
		fmt.Println("---------------------------------")
		fmt.Println("1. Data Spare Part")
		fmt.Println("2. Data Pelanggan")
		fmt.Println("3. Data Transaksi")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opsiPilihan)

		if opsiPilihan == 1 {
			MenuBarang(barang, nBarang)
		} else if opsiPilihan == 2 {
			MenuPelanggan(pengguna, nPengguna)
		} else if opsiPilihan == 3 {
			MenuTransaksi(transaksi, nTransaksi, barang, *nBarang, pengguna, *nPengguna)
		} else if opsiPilihan == 4 {
			fmt.Println("Terimakasih telah menggunakan aplikasi")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

// =========================================================
// MENU SPARE PART (BARANG)
// =========================================================
func MenuBarang(barang *TabBarang, n *int) {
	var input int
	for input != 6 {
		fmt.Println("---------------------------------")
		fmt.Println("        Data Spare Part        ")
		fmt.Println("---------------------------------")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Hapus  Data")
		fmt.Println("3. Edit   Data")
		fmt.Println("4. Cari   Data")
		fmt.Println("5. Urutkan Data")
		fmt.Println("6. Submit")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&input)

		if input == 1 {
			TambahBarang(barang, n)
		} else if input == 2 {
			HapusBarang(barang, n)
		} else if input == 3 {
			EditBarang(barang, *n)
		} else if input == 4 {
			CariBarang(barang, *n)
		} else if input == 5 {
			UrutkanBarang(barang, *n)
		} else if input == 6 {
			fmt.Println("Data Spare Part disimpan.")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func TambahBarang(barang *TabBarang, n *int) {
	var namaBarang string
	var hargaBarang int

	if *n >= NMAX {
		fmt.Println("Data spare part sudah penuh.")
		return
	}

	fmt.Print("Spare Part: ")
	fmt.Scan(&namaBarang)
	fmt.Print("Harga:      ")
	fmt.Scan(&hargaBarang)

	barang[*n].sparePart = namaBarang
	barang[*n].harga = hargaBarang
	*n++

	fmt.Println("Anda menambahkan", namaBarang, "dengan harga Rp", hargaBarang)
	TampilBarang(barang, *n)
}

func HapusBarang(barang *TabBarang, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data spare part.")
		return
	}
	TampilBarang(barang, *n)

	var delIdx int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&delIdx)
	delIdx--

	if delIdx < 0 || delIdx >= *n {
		fmt.Println("Index tidak valid")
		return
	}

	dihapus := barang[delIdx]
	for i := delIdx; i < *n-1; i++ {
		barang[i] = barang[i+1]
	}
	*n--
	fmt.Println("Data", dihapus.sparePart, "dengan harga Rp", dihapus.harga, "berhasil dihapus.")
	TampilBarang(barang, *n)
}

func EditBarang(barang *TabBarang, n int) {
	if n == 0 {
		fmt.Println("Belum ada data spare part.")
		return
	}
	TampilBarang(barang, n)

	var editIdx int
	fmt.Print("Masukkan nomor data yang ingin diedit: ")
	fmt.Scan(&editIdx)

	if editIdx <= 0 || editIdx > n {
		fmt.Println("Index tidak valid")
		return
	}

	var namaBarang string
	var hargaBarang int
	fmt.Print("Spare Part: ")
	fmt.Scan(&namaBarang)
	fmt.Print("Harga:      ")
	fmt.Scan(&hargaBarang)

	barang[editIdx-1].sparePart = namaBarang
	barang[editIdx-1].harga = hargaBarang

	fmt.Println("Data berhasil diedit")
	TampilBarang(barang, n)
}

// Sequential search pada nama spare part
func CariBarang(barang *TabBarang, n int) {
	if n == 0 {
		fmt.Println("Belum ada data spare part.")
		return
	}
	var key string
	fmt.Print("Cari nama spare part: ")
	fmt.Scan(&key)

	ditemukan := false
	for i := 0; i < n; i++ {
		if barang[i].sparePart == key {
			fmt.Println("Ditemukan pada nomor", i+1, "->", barang[i].sparePart, "Rp", barang[i].harga)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data", key, "tidak ditemukan.")
	}
}

// Selection sort harga ascending
func UrutkanBarang(barang *TabBarang, n int) {
	if n == 0 {
		fmt.Println("Belum ada data spare part.")
		return
	}
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if barang[j].harga < barang[idxMin].harga {
				idxMin = j
			}
		}
		if idxMin != i {
			barang[i], barang[idxMin] = barang[idxMin], barang[i]
		}
	}
	fmt.Println("Data diurutkan berdasarkan harga (ascending).")
	TampilBarang(barang, n)
}

func TampilBarang(barang *TabBarang, n int) {
	fmt.Println("---------------------------------")
	fmt.Println("Daftar Spare Part:")
	if n == 0 {
		fmt.Println("(kosong)")
	}
	for i := 0; i < n; i++ {
		fmt.Println(i+1, ".", barang[i].sparePart, "- Rp", barang[i].harga)
	}
	fmt.Println("---------------------------------")
}

// =========================================================
// MENU PELANGGAN
// =========================================================
func MenuPelanggan(pengguna *TabPengguna, n *int) {
	var input int
	for input != 6 {
		fmt.Println("---------------------------------")
		fmt.Println("         Data Pelanggan         ")
		fmt.Println("---------------------------------")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Hapus  Data")
		fmt.Println("3. Edit   Data")
		fmt.Println("4. Cari   Data")
		fmt.Println("5. Urutkan Data")
		fmt.Println("6. Submit")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&input)

		if input == 1 {
			TambahPengguna(pengguna, n)
		} else if input == 2 {
			HapusPengguna(pengguna, n)
		} else if input == 3 {
			EditPengguna(pengguna, *n)
		} else if input == 4 {
			CariPengguna(pengguna, *n)
		} else if input == 5 {
			UrutkanPengguna(pengguna, *n)
		} else if input == 6 {
			fmt.Println("Data Pelanggan disimpan.")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func TambahPengguna(pengguna *TabPengguna, n *int) {
	if *n >= NMAX {
		fmt.Println("Data pelanggan sudah penuh.")
		return
	}
	var id, nama string
	fmt.Print("ID Pelanggan:   ")
	fmt.Scan(&id)
	fmt.Print("Nama Pelanggan: ")
	fmt.Scan(&nama)

	pengguna[*n].idPengguna = id
	pengguna[*n].nama = nama
	*n++

	fmt.Println("Anda menambahkan pelanggan", id, "-", nama)
	TampilPengguna(pengguna, *n)
}

func HapusPengguna(pengguna *TabPengguna, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	TampilPengguna(pengguna, *n)

	var delIdx int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&delIdx)
	delIdx--

	if delIdx < 0 || delIdx >= *n {
		fmt.Println("Index tidak valid")
		return
	}

	dihapus := pengguna[delIdx]
	for i := delIdx; i < *n-1; i++ {
		pengguna[i] = pengguna[i+1]
	}
	*n--
	fmt.Println("Data pelanggan", dihapus.idPengguna, "-", dihapus.nama, "berhasil dihapus.")
	TampilPengguna(pengguna, *n)
}

func EditPengguna(pengguna *TabPengguna, n int) {
	if n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	TampilPengguna(pengguna, n)

	var editIdx int
	fmt.Print("Masukkan nomor data yang ingin diedit: ")
	fmt.Scan(&editIdx)

	if editIdx <= 0 || editIdx > n {
		fmt.Println("Index tidak valid")
		return
	}

	var id, nama string
	fmt.Print("ID Pelanggan:   ")
	fmt.Scan(&id)
	fmt.Print("Nama Pelanggan: ")
	fmt.Scan(&nama)

	pengguna[editIdx-1].idPengguna = id
	pengguna[editIdx-1].nama = nama

	fmt.Println("Data pelanggan berhasil diedit")
	TampilPengguna(pengguna, n)
}

func CariPengguna(pengguna *TabPengguna, n int) {
	if n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	var key string
	fmt.Print("Cari ID Pelanggan: ")
	fmt.Scan(&key)

	ditemukan := false
	for i := 0; i < n; i++ {
		if pengguna[i].idPengguna == key {
			fmt.Println("Ditemukan pada nomor", i+1, "->", pengguna[i].idPengguna, "-", pengguna[i].nama)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Pelanggan dengan ID", key, "tidak ditemukan.")
	}
}

// Selection sort nama pelanggan ascending
func UrutkanPengguna(pengguna *TabPengguna, n int) {
	if n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if pengguna[j].nama < pengguna[idxMin].nama {
				idxMin = j
			}
		}
		if idxMin != i {
			pengguna[i], pengguna[idxMin] = pengguna[idxMin], pengguna[i]
		}
	}
	fmt.Println("Data diurutkan berdasarkan nama (A-Z).")
	TampilPengguna(pengguna, n)
}

func TampilPengguna(pengguna *TabPengguna, n int) {
	fmt.Println("---------------------------------")
	fmt.Println("Daftar Pelanggan:")
	if n == 0 {
		fmt.Println("(kosong)")
	}
	for i := 0; i < n; i++ {
		fmt.Println(i+1, ".", pengguna[i].idPengguna, "-", pengguna[i].nama)
	}
	fmt.Println("---------------------------------")
}

// =========================================================
// MENU TRANSAKSI
// =========================================================
func MenuTransaksi(transaksi *TabTransaksi, n *int, barang *TabBarang, nBarang int, pengguna *TabPengguna, nPengguna int) {
	var input int
	for input != 6 {
		fmt.Println("---------------------------------")
		fmt.Println("         Data Transaksi         ")
		fmt.Println("---------------------------------")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Hapus  Data")
		fmt.Println("3. Edit   Data")
		fmt.Println("4. Cari   Data")
		fmt.Println("5. Urutkan Data")
		fmt.Println("6. Submit")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&input)

		if input == 1 {
			TambahTransaksi(transaksi, n, barang, nBarang, pengguna, nPengguna)
		} else if input == 2 {
			HapusTransaksi(transaksi, n)
		} else if input == 3 {
			EditTransaksi(transaksi, *n, barang, nBarang)
		} else if input == 4 {
			CariTransaksi(transaksi, *n)
		} else if input == 5 {
			UrutkanTransaksi(transaksi, *n)
		} else if input == 6 {
			fmt.Println("Data Transaksi disimpan.")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func TambahTransaksi(transaksi *TabTransaksi, n *int, barang *TabBarang, nBarang int, pengguna *TabPengguna, nPengguna int) {
	if *n >= NMAX {
		fmt.Println("Data transaksi sudah penuh.")
		return
	}
	if nBarang == 0 {
		fmt.Println("Belum ada spare part. Tambahkan spare part terlebih dahulu.")
		return
	}
	if nPengguna == 0 {
		fmt.Println("Belum ada pelanggan. Tambahkan pelanggan terlebih dahulu.")
		return
	}

	var idTrans, idPel, nama string
	var jumlah int

	fmt.Print("ID Transaksi: ")
	fmt.Scan(&idTrans)

	TampilPengguna(pengguna, nPengguna)
	fmt.Print("ID Pelanggan: ")
	fmt.Scan(&idPel)
	if CariIndexPengguna(pengguna, nPengguna, idPel) == -1 {
		fmt.Println("ID Pelanggan tidak terdaftar.")
		return
	}

	TampilBarang(barang, nBarang)
	fmt.Print("Nama Spare Part: ")
	fmt.Scan(&nama)
	idxBarang := CariIndexBarang(barang, nBarang, nama)
	if idxBarang == -1 {
		fmt.Println("Spare part", nama, "tidak terdaftar.")
		return
	}

	fmt.Print("Jumlah: ")
	fmt.Scan(&jumlah)
	if jumlah <= 0 {
		fmt.Println("Jumlah harus lebih dari 0.")
		return
	}

	total := barang[idxBarang].harga * jumlah
	transaksi[*n].idTransaksi = idTrans
	transaksi[*n].idPelanggan = idPel
	transaksi[*n].namaBarang = nama
	transaksi[*n].jumlah = jumlah
	transaksi[*n].total = total
	*n++

	fmt.Println("Transaksi", idTrans, "berhasil dicatat. Total: Rp", total)
	TampilTransaksi(transaksi, *n)
}

func HapusTransaksi(transaksi *TabTransaksi, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}
	TampilTransaksi(transaksi, *n)

	var delIdx int
	fmt.Print("Masukkan nomor transaksi yang ingin dihapus: ")
	fmt.Scan(&delIdx)
	delIdx--

	if delIdx < 0 || delIdx >= *n {
		fmt.Println("Index tidak valid")
		return
	}

	dihapus := transaksi[delIdx]
	for i := delIdx; i < *n-1; i++ {
		transaksi[i] = transaksi[i+1]
	}
	*n--
	fmt.Println("Transaksi", dihapus.idTransaksi, "berhasil dihapus.")
	TampilTransaksi(transaksi, *n)
}

func EditTransaksi(transaksi *TabTransaksi, n int, barang *TabBarang, nBarang int) {
	if n == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}
	TampilTransaksi(transaksi, n)

	var editIdx int
	fmt.Print("Masukkan nomor transaksi yang ingin diedit: ")
	fmt.Scan(&editIdx)

	if editIdx <= 0 || editIdx > n {
		fmt.Println("Index tidak valid")
		return
	}

	var nama string
	var jumlah int
	fmt.Print("Nama Spare Part baru: ")
	fmt.Scan(&nama)
	idxBarang := CariIndexBarang(barang, nBarang, nama)
	if idxBarang == -1 {
		fmt.Println("Spare part", nama, "tidak terdaftar.")
		return
	}

	fmt.Print("Jumlah baru: ")
	fmt.Scan(&jumlah)
	if jumlah <= 0 {
		fmt.Println("Jumlah harus lebih dari 0.")
		return
	}

	transaksi[editIdx-1].namaBarang = nama
	transaksi[editIdx-1].jumlah = jumlah
	transaksi[editIdx-1].total = barang[idxBarang].harga * jumlah

	fmt.Println("Transaksi berhasil diedit.")
	TampilTransaksi(transaksi, n)
}

func CariTransaksi(transaksi *TabTransaksi, n int) {
	if n == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}
	var key string
	fmt.Print("Cari ID Transaksi: ")
	fmt.Scan(&key)

	for i := 0; i < n; i++ {
		if transaksi[i].idTransaksi == key {
			fmt.Println("Ditemukan pada nomor", i+1, ":")
			fmt.Println("  ID Transaksi :", transaksi[i].idTransaksi)
			fmt.Println("  ID Pelanggan :", transaksi[i].idPelanggan)
			fmt.Println("  Spare Part   :", transaksi[i].namaBarang)
			fmt.Println("  Jumlah       :", transaksi[i].jumlah)
			fmt.Println("  Total        : Rp", transaksi[i].total)
			return
		}
	}
	fmt.Println("Transaksi dengan ID", key, "tidak ditemukan.")
}

// Selection sort total transaksi descending (terbesar dulu)
func UrutkanTransaksi(transaksi *TabTransaksi, n int) {
	if n == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if transaksi[j].total > transaksi[idxMax].total {
				idxMax = j
			}
		}
		if idxMax != i {
			transaksi[i], transaksi[idxMax] = transaksi[idxMax], transaksi[i]
		}
	}
	fmt.Println("Data diurutkan berdasarkan total (terbesar dulu).")
	TampilTransaksi(transaksi, n)
}

func TampilTransaksi(transaksi *TabTransaksi, n int) {
	fmt.Println("---------------------------------")
	fmt.Println("Daftar Transaksi:")
	if n == 0 {
		fmt.Println("(kosong)")
	}
	for i := 0; i < n; i++ {
		fmt.Println(i+1, ".", transaksi[i].idTransaksi,
			"| Pel:", transaksi[i].idPelanggan,
			"|", transaksi[i].namaBarang, "x", transaksi[i].jumlah,
			"| Total: Rp", transaksi[i].total)
	}
	fmt.Println("---------------------------------")
}

// =========================================================
// HELPER LOOKUP
// =========================================================
func CariIndexBarang(barang *TabBarang, n int, nama string) int {
	for i := 0; i < n; i++ {
		if barang[i].sparePart == nama {
			return i
		}
	}
	return -1
}

func CariIndexPengguna(pengguna *TabPengguna, n int, id string) int {
	for i := 0; i < n; i++ {
		if pengguna[i].idPengguna == id {
			return i
		}
	}
	return -1
}
