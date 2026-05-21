package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NMAX int = 100
const DATA_DIR = "data"

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
	var nUser, nBarang, nPengguna, nTransaksi int
	var pilih int

	// Load data dari disk supaya state persist antar sesi
	nUser = LoadUsers(&regist)
	nBarang = LoadBarang(&barang)
	nPengguna = LoadPengguna(&pengguna)
	nTransaksi = LoadTransaksi(&transaksi)

	for pilih != 3 {
		fmt.Println("---------------------------------")
		fmt.Println("    Aplikasi Service Motor X    ")
		fmt.Println("---------------------------------")
		fmt.Println("Pengguna terdaftar:", nUser)
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			Regist(&regist, &nUser)
		} else if pilih == 2 {
			username, ok := Login(&regist, nUser)
			if ok {
				MenuUtama(username, &barang, &nBarang, &pengguna, &nPengguna, &transaksi, &nTransaksi)
			}
		} else if pilih == 3 {
			fmt.Println("Terimakasih")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

// =========================================================
// REGISTRASI & LOGIN (MULTI-USER)
// =========================================================
func Regist(regist *TabRegist, n *int) {
	if *n >= NMAX {
		fmt.Println("Slot registrasi penuh.")
		return
	}
	var username, email, password, confirm string

	fmt.Println("---------------------------------")
	fmt.Println("Silahkan melakukan registrasi")
	fmt.Print("Username: ")
	fmt.Scan(&username)

	if CariIndexUser(regist, *n, username) != -1 {
		fmt.Println("Username", username, "sudah terdaftar. Silahkan login.")
		return
	}

	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	fmt.Print("Konfirmasi Password: ")
	fmt.Scan(&confirm)

	for confirm != password {
		fmt.Println("Konfirmasi Password Tidak Valid")
		fmt.Print("Password: ")
		fmt.Scan(&password)
		fmt.Print("Konfirmasi Password: ")
		fmt.Scan(&confirm)
	}

	regist[*n] = registration{username, email, password, confirm}
	*n++
	SaveUsers(regist, *n)

	fmt.Println("---------------------------------")
	fmt.Println("   Akun", username, "berhasil dibuat   ")
}

func Login(regist *TabRegist, n int) (string, bool) {
	if n == 0 {
		fmt.Println("Belum ada user terdaftar. Silahkan registrasi dulu.")
		return "", false
	}

	for percobaan := 0; percobaan < 3; percobaan++ {
		var username, password string
		fmt.Println("---------------------------------")
		fmt.Println("Silahkan Login")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		idx := CariIndexUser(regist, n, username)
		if idx != -1 && regist[idx].password == password {
			fmt.Println("---------------------------------")
			fmt.Print("  Halo ", username)
			fmt.Println(", Selamat Datang   ")
			fmt.Println("---------------------------------")
			return username, true
		}
		fmt.Println("Username atau password anda salah (percobaan", percobaan+1, "dari 3)")
	}
	fmt.Println("Login dibatalkan, terlalu banyak percobaan salah.")
	return "", false
}

func CariIndexUser(regist *TabRegist, n int, username string) int {
	for i := 0; i < n; i++ {
		if regist[i].username == username {
			return i
		}
	}
	return -1
}

// =========================================================
// MENU UTAMA
// =========================================================
func MenuUtama(username string, barang *TabBarang, nBarang *int, pengguna *TabPengguna, nPengguna *int, transaksi *TabTransaksi, nTransaksi *int) {
	var opsiPilihan int

	for opsiPilihan != 5 {
		fmt.Println("---------------------------------")
		fmt.Println("    Menu Utama -", username)
		fmt.Println("---------------------------------")
		fmt.Println("1. Data Spare Part   (", *nBarang, ")")
		fmt.Println("2. Data Pelanggan    (", *nPengguna, ")")
		fmt.Println("3. Data Transaksi    (", *nTransaksi, ")")
		fmt.Println("4. Laporan Transaksi")
		fmt.Println("5. Logout")
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
			MenuLaporan(barang, *nBarang, pengguna, *nPengguna, transaksi, *nTransaksi)
		} else if opsiPilihan == 5 {
			fmt.Println("Logout. Sampai jumpa,", username)
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
	SaveBarang(barang, *n)

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
	SaveBarang(barang, *n)
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
	SaveBarang(barang, n)

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
	SaveBarang(barang, n)
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

	if CariIndexPengguna(pengguna, *n, id) != -1 {
		fmt.Println("ID Pelanggan", id, "sudah terdaftar.")
		return
	}

	fmt.Print("Nama Pelanggan: ")
	fmt.Scan(&nama)

	pengguna[*n].idPengguna = id
	pengguna[*n].nama = nama
	*n++
	SavePengguna(pengguna, *n)

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
	SavePengguna(pengguna, *n)
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
	SavePengguna(pengguna, n)

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
	SavePengguna(pengguna, n)
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
	SaveTransaksi(transaksi, *n)

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
	SaveTransaksi(transaksi, *n)
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
	SaveTransaksi(transaksi, n)

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
	SaveTransaksi(transaksi, n)
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
// MENU LAPORAN
// =========================================================
func MenuLaporan(barang *TabBarang, nBarang int, pengguna *TabPengguna, nPengguna int, transaksi *TabTransaksi, nTransaksi int) {
	var opsi int
	for opsi != 5 {
		fmt.Println("---------------------------------")
		fmt.Println("       Laporan Transaksi       ")
		fmt.Println("---------------------------------")
		fmt.Println("1. Ringkasan Pendapatan")
		fmt.Println("2. Laporan Per Pelanggan")
		fmt.Println("3. Laporan Per Spare Part")
		fmt.Println("4. Riwayat Transaksi Pelanggan")
		fmt.Println("5. Kembali")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opsi)

		if opsi == 1 {
			LaporanRingkasan(transaksi, nTransaksi)
		} else if opsi == 2 {
			LaporanPerPelanggan(transaksi, nTransaksi, pengguna, nPengguna)
		} else if opsi == 3 {
			LaporanPerBarang(transaksi, nTransaksi, barang, nBarang)
		} else if opsi == 4 {
			LaporanRiwayatPelanggan(transaksi, nTransaksi, pengguna, nPengguna)
		} else if opsi == 5 {
			fmt.Println("Kembali ke Menu Utama.")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func LaporanRingkasan(t *TabTransaksi, n int) {
	if n == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	total := 0
	totalItem := 0
	idxMax, idxMin := 0, 0
	for i := 0; i < n; i++ {
		total += t[i].total
		totalItem += t[i].jumlah
		if t[i].total > t[idxMax].total {
			idxMax = i
		}
		if t[i].total < t[idxMin].total {
			idxMin = i
		}
	}
	fmt.Println("---------------------------------")
	fmt.Println("        Ringkasan Penjualan       ")
	fmt.Println("---------------------------------")
	fmt.Println("Jumlah Transaksi   :", n)
	fmt.Println("Total Item Terjual :", totalItem)
	fmt.Println("Total Pendapatan   : Rp", total)
	fmt.Println("Rata-rata per Trans: Rp", total/n)
	fmt.Println("Transaksi Tertinggi: Rp", t[idxMax].total, "(", t[idxMax].idTransaksi, "-", t[idxMax].namaBarang, ")")
	fmt.Println("Transaksi Terendah : Rp", t[idxMin].total, "(", t[idxMin].idTransaksi, "-", t[idxMin].namaBarang, ")")
	fmt.Println("---------------------------------")
}

func LaporanPerPelanggan(t *TabTransaksi, nT int, p *TabPengguna, nP int) {
	if nT == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	fmt.Println("---------------------------------")
	fmt.Println("      Laporan Per Pelanggan      ")
	fmt.Println("---------------------------------")
	ada := false
	for i := 0; i < nP; i++ {
		jumlahTrans := 0
		totalBelanja := 0
		for j := 0; j < nT; j++ {
			if t[j].idPelanggan == p[i].idPengguna {
				jumlahTrans++
				totalBelanja += t[j].total
			}
		}
		if jumlahTrans > 0 {
			fmt.Println(p[i].idPengguna, "-", p[i].nama, ":", jumlahTrans, "transaksi, total Rp", totalBelanja)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Belum ada pelanggan dengan transaksi.")
	}
	fmt.Println("---------------------------------")
}

func LaporanPerBarang(t *TabTransaksi, nT int, b *TabBarang, nB int) {
	if nT == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	fmt.Println("---------------------------------")
	fmt.Println("     Laporan Per Spare Part     ")
	fmt.Println("---------------------------------")
	ada := false
	for i := 0; i < nB; i++ {
		totalJual := 0
		totalRevenue := 0
		for j := 0; j < nT; j++ {
			if t[j].namaBarang == b[i].sparePart {
				totalJual += t[j].jumlah
				totalRevenue += t[j].total
			}
		}
		if totalJual > 0 {
			fmt.Println(b[i].sparePart, ": terjual", totalJual, "unit, revenue Rp", totalRevenue)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Belum ada spare part yang terjual.")
	}
	fmt.Println("---------------------------------")
}

func LaporanRiwayatPelanggan(t *TabTransaksi, nT int, p *TabPengguna, nP int) {
	if nP == 0 {
		fmt.Println("Belum ada pelanggan.")
		return
	}
	TampilPengguna(p, nP)
	var id string
	fmt.Print("ID Pelanggan: ")
	fmt.Scan(&id)
	idx := CariIndexPengguna(p, nP, id)
	if idx == -1 {
		fmt.Println("Pelanggan tidak terdaftar.")
		return
	}
	fmt.Println("---------------------------------")
	fmt.Println("Riwayat", p[idx].idPengguna, "-", p[idx].nama, ":")
	fmt.Println("---------------------------------")
	total := 0
	count := 0
	for i := 0; i < nT; i++ {
		if t[i].idPelanggan == id {
			fmt.Println(" ", t[i].idTransaksi, "|", t[i].namaBarang, "x", t[i].jumlah, "| Rp", t[i].total)
			total += t[i].total
			count++
		}
	}
	if count == 0 {
		fmt.Println("  (belum ada transaksi)")
	} else {
		fmt.Println("---------------------------------")
		fmt.Println("Jumlah Transaksi:", count)
		fmt.Println("Total Belanja   : Rp", total)
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

// =========================================================
// FILE PERSISTENCE (TSV di folder data/)
// =========================================================
func ensureDataDir() {
	os.MkdirAll(DATA_DIR, 0755)
}

func SaveUsers(regist *TabRegist, n int) {
	ensureDataDir()
	f, err := os.Create(DATA_DIR + "/users.txt")
	if err != nil {
		fmt.Println("Gagal simpan users:", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%s\t%s\t%s\n", regist[i].username, regist[i].email, regist[i].password)
	}
	w.Flush()
}

func LoadUsers(regist *TabRegist) int {
	f, err := os.Open(DATA_DIR + "/users.txt")
	if err != nil {
		return 0
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	n := 0
	for s.Scan() && n < NMAX {
		parts := strings.Split(s.Text(), "\t")
		if len(parts) < 3 {
			continue
		}
		regist[n] = registration{
			username:    parts[0],
			email:       parts[1],
			password:    parts[2],
			confirmPass: parts[2],
		}
		n++
	}
	return n
}

func SaveBarang(barang *TabBarang, n int) {
	ensureDataDir()
	f, err := os.Create(DATA_DIR + "/barang.txt")
	if err != nil {
		fmt.Println("Gagal simpan barang:", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%s\t%d\n", barang[i].sparePart, barang[i].harga)
	}
	w.Flush()
}

func LoadBarang(barang *TabBarang) int {
	f, err := os.Open(DATA_DIR + "/barang.txt")
	if err != nil {
		return 0
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	n := 0
	for s.Scan() && n < NMAX {
		parts := strings.Split(s.Text(), "\t")
		if len(parts) < 2 {
			continue
		}
		harga, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		barang[n].sparePart = parts[0]
		barang[n].harga = harga
		n++
	}
	return n
}

func SavePengguna(pengguna *TabPengguna, n int) {
	ensureDataDir()
	f, err := os.Create(DATA_DIR + "/pelanggan.txt")
	if err != nil {
		fmt.Println("Gagal simpan pelanggan:", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%s\t%s\n", pengguna[i].idPengguna, pengguna[i].nama)
	}
	w.Flush()
}

func LoadPengguna(pengguna *TabPengguna) int {
	f, err := os.Open(DATA_DIR + "/pelanggan.txt")
	if err != nil {
		return 0
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	n := 0
	for s.Scan() && n < NMAX {
		parts := strings.Split(s.Text(), "\t")
		if len(parts) < 2 {
			continue
		}
		pengguna[n].idPengguna = parts[0]
		pengguna[n].nama = parts[1]
		n++
	}
	return n
}

func SaveTransaksi(transaksi *TabTransaksi, n int) {
	ensureDataDir()
	f, err := os.Create(DATA_DIR + "/transaksi.txt")
	if err != nil {
		fmt.Println("Gagal simpan transaksi:", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\n",
			transaksi[i].idTransaksi,
			transaksi[i].idPelanggan,
			transaksi[i].namaBarang,
			transaksi[i].jumlah,
			transaksi[i].total)
	}
	w.Flush()
}

func LoadTransaksi(transaksi *TabTransaksi) int {
	f, err := os.Open(DATA_DIR + "/transaksi.txt")
	if err != nil {
		return 0
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	n := 0
	for s.Scan() && n < NMAX {
		parts := strings.Split(s.Text(), "\t")
		if len(parts) < 5 {
			continue
		}
		jumlah, err1 := strconv.Atoi(parts[3])
		total, err2 := strconv.Atoi(parts[4])
		if err1 != nil || err2 != nil {
			continue
		}
		transaksi[n].idTransaksi = parts[0]
		transaksi[n].idPelanggan = parts[1]
		transaksi[n].namaBarang = parts[2]
		transaksi[n].jumlah = jumlah
		transaksi[n].total = total
		n++
	}
	return n
}
