package main

import "fmt"

const NMAX int = 100

type registration struct {
	username, email, password, confirpackage main
	
	import "fmt"
	
	const NMAX int = 100
	
	type registration struct {
		username, email, password, confirmPass string
	}
	
	type TabRegist [NMAX]registration
	
	type Barang struct {
		sparePart string
		transaksi int
	}
	
	type Pengguna struct {
		nama, idPengguna string
	}
	
	type TabBarang [NMAX]Barang
	type TabPengguna [NMAX]Pengguna
	
	func main() {
		var regist TabRegist
		var n int
		var pilih int
		var barang TabBarang
	
		fmt.Println("---------------------------------")
		fmt.Println("    Aplikasi Service Motor X    ")
		fmt.Println("---------------------------------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Keluar")
	
		for pilih!= 1 && pilih!= 2 {
			fmt.Scan(&pilih)
			if pilih == 1 {
				Regist(&regist, n)
				Login(&regist, n)
				PenggunaUser(&barang, n)
			} else if pilih == 2 {
				fmt.Println("Terimakasih")
			} else {
				fmt.Println("Pilihan Tidak Valid")
			}
		}
	}
	
	func Regist(regist *TabRegist, n int) {
		fmt.Println("Silahkan melakukan regitrasi")
		fmt.Print("Username: ")
		fmt.Scan(&regist[n].username)
		fmt.Print("Password: ")
		fmt.Scan(&regist[n].password)
		fmt.Print("Konfirmasi Password: ")
		fmt.Scan(&regist[n].confirmPass)
	
		for regist[n].confirmPass!= regist[n].password {
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
	
	func Login(regist *TabRegist, n int) {
		var username, password string
	
		for username!= regist[n].username || password!= regist[n].password {
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
				return
			} else {
				fmt.Println("Username atau password anda salah")
			}
		}
	}
	
	func PenggunaUser(Barang *TabBarang, n int) {
		var inputBarang int
		var jumlah int
		var namaBarang string
		var hargaBarang int
		var opsiPilihan int
		var namaPengguna string
		var idPengguna string
		var pengguna TabPengguna
	
		for opsiPilihan!= 4 {
			fmt.Println("1. Data Spare Part")
			fmt.Println("2. Data Pelanggan")
			fmt.Println("3. Data Transaksi")
			fmt.Println("4. Keluar")
			fmt.Println("---------------------------------")
			fmt.Scan(&opsiPilihan)
	
			if opsiPilihan == 1 {
				fmt.Println("---------------------------------")
				fmt.Println("     Ada yang Bisa Dibantu   ")
				fmt.Println("---------------------------------")
				for inputBarang!= 4 {
					fmt.Println("1. Tambah Data")
					fmt.Println("2. Hapus  Data")
					fmt.Println("3. Edit   Data")
					fmt.Println("4. Submit")
					fmt.Println("---------------------------------")
					fmt.Scan(&inputBarang)
					//tambah data dan hapus
					if inputBarang == 1 {
						fmt.Print("Spare Part: ")
						fmt.Scan(&namaBarang)
						fmt.Print("Harga:  ")
						fmt.Scan(&hargaBarang)
	
						Barang[jumlah].sparePart = namaBarang
						Barang[jumlah].transaksi = hargaBarang
						jumlah++
	
						fmt.Println("Anda menambahkan", namaBarang, "dengan harga Rp ", hargaBarang)
						fmt.Println("Anda data menambahkan data:")
	
						for i := 0; i < jumlah; i++ {
							fmt.Println(i+1, ".", Barang[i].sparePart, "dengan harga Rp ", Barang[i].transaksi)
						}
						fmt.Println("---------------------------------")
					} else if inputBarang == 2 {
						//... (continue with the rest of the code)
					} else if inputBarang == 3 {
						//... (continue with the rest of the code)
					}
				}
			} else if opsiPilihan == 2 {
				//... (continue with the rest of the code)
			}
		}
	
		fmt.Println("Anda menambahkan data:")
		for i := 0; i < jumlah; i++ {
			fmt.Println(i+1, ".", Barang[i].sparePart, "dengan harga Rp ", Barang[i].transaksi)
		}
	}mPass string
}

type TabRegist [NMAX]registration

type Barang struct {
	sparePart string
	transaksi int
}

type Pengguna struct {
	nama, idPengguna string
}

type TabBarang [NMAX]Barang
type TabPengguna [NMAX]Pengguna

func main() {
	var regist TabRegist
	var n int
	var pilih int
	var barang TabBarang

	fmt.Println("---------------------------------")
	fmt.Println("    Aplikasi Service Motor X    ")
	fmt.Println("---------------------------------")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Keluar")

	for pilih != 1 && pilih != 2 {
		fmt.Scan(&pilih)
		if pilih == 1 {
			Regist(&regist, n)
			Login(&regist, n)
			PenggunaUser(&barang, n)
		} else if pilih == 2 {
			fmt.Println("Terimakasih")
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func Regist(regist *TabRegist, n int) {
	fmt.Println("Silahkan melakukan regitrasi")
	fmt.Print("Username: ")
	fmt.Scan(&regist[n].username)
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

func Login(regist *TabRegist, n int) {
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
			return
		} else {
			fmt.Println("Username atau password anda salah")
		}
	}
}

func PenggunaUser(Barang *TabBarang, n int) {
	var inputBarang int
	var jumlah int
	var namaBarang string
	var hargaBarang int
	var opsiPilihan int
	var namaPengguna string
	var idPengguna string
	var pengguna TabPengguna

	for opsiPilihan != 4 {
		fmt.Println("1. Data Spare Part")
		fmt.Println("2. Data Pelanggan")
		fmt.Println("3. Data Transaksi")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------------------")
		fmt.Scan(&opsiPilihan)

		if opsiPilihan == 1 {
			fmt.Println("---------------------------------")
			fmt.Println("     Ada yang Bisa Dibantu   ")
			fmt.Println("---------------------------------")
			for inputBarang != 4 {
				fmt.Println("1. Tambah Data")
				fmt.Println("2. Hapus  Data")
				fmt.Println("3. Edit   Data")
				fmt.Println("4. Submit")
				fmt.Println("---------------------------------")
				fmt.Scan(&inputBarang)
				//tambah data dan hapus
				if inputBarang == 1 {
					fmt.Print("Spare Part: ")
					fmt.Scan(&namaBarang)
					fmt.Print("Harga:  ")
					fmt.Scan(&hargaBarang)

					Barang[jumlah].sparePart = namaBarang
					Barang[jumlah].transaksi = hargaBarang
					jumlah++

					fmt.Println("Anda menambahkan", namaBarang, "dengan harga Rp ", hargaBarang)
					fmt.Println("Anda data menambahkan data:")

					for i := 0; i < jumlah; i++ {
						fmt.Println(i+1, ".", Barang[i].sparePart, "dengan harga Rp ", Barang[i].transaksi)
					}
					fmt.Println("---------------------------------")
				} else if inputBarang == 2 {
					fmt.Print("Masukkan data yang ingin di hapus: ")
					var delIdx int
					fmt.Scan(&delIdx)
					delIdx--

					for i := delIdx; i < jumlah-1; i++ {
						Barang[i] = Barang[i+1]
					}
					jumlah--
					fmt.Println("Data ", Barang[delIdx].sparePart, "dengan harga Rp ", Barang[delIdx].transaksi, "berhasil dihapus.")
					fmt.Println("Anda data menambahkan data:")
					for i := 0; i < jumlah; i++ {
						fmt.Println(i+1, ".", Barang[i].sparePart, "dengan harga Rp ", Barang[i].transaksi)
					}
					fmt.Println("---------------------------------")
				} else if inputBarang == 3 {
					var editIdx int
					fmt.Println()
					fmt.Println("Silahkan edit barang")

					fmt.Scan(&editIdx)
					if editIdx > 0 && editIdx <= jumlah {
						fmt.Print("Spare Part: ")
						fmt.Scan(&namaBarang)
						fmt.Print("Harga: ")
						fmt.Scan(&hargaBarang)

						Barang[editIdx-1].sparePart = namaBarang
						Barang[editIdx-1].transaksi = hargaBarang

						fmt.Println()
						fmt.Println("Data berhasil diedit")

						fmt.Println("Barang setelah diedit:")
						for i := 0; i < jumlah; i++ {
							fmt.Println(i+1, ".", Barang[i].sparePart, "dengan harga Rp ", Barang[i].transaksi)
						}
					} else {
						fmt.Println("Index tidak valid")
					}
					fmt.Println("---------------------------------")
				}
			}
		} else if opsiPilihan == 2 {
			fmt.Println("---------------------------------")
			fmt.Println("     Ada yang Bisa Dibantu   ")
			fmt.Println("---------------------------------")
			for namaPengguna != "4" {
				fmt.Println("1. Tambah Data")
				fmt.Println("2. Hapus  Data")
				fmt.Println("3. Edit   Data")
				fmt.Println("4. Submit")
				fmt.Println("---------------------------------")
				fmt.Scan(&namaPengguna)

				if namaPengguna == "1" {
					fmt.Print("ID Pelanggan: ")
					fmt.Scan(&idPengguna)
					fmt.Print("Nama Pelanggan:  ")
					fmt.Scan(&namaPengguna)

					pengguna[n].idPengguna = idPengguna
					pengguna[n].nama = namaPengguna
					n++

					fmt.Println("Anda menambahkan Data Pengguna ", idPengguna, namaPengguna)
					fmt.Println("---------------------------------")
				}
			}
		}
	}

	fmt.Println("Anda menambahkan data:")
	for i := 0; i < jumlah; i++ {
		fmt.Println(i+1, ".", Barang[i].sparePart, "dengan harga Rp ", Barang[i].transaksi)
	}
}
