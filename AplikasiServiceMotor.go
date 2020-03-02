/*

DATA PEMBUAT
dickyadi44@gmail.com
Judul	:	Aplikasi Service Motor
Dicky Adi Naufal Farhansyah		(1301194041)

*/
package main

import "fmt"

	type DetailSparepart struct{
		Nama string
		Jumlah int
		Pabrik string
		Tahun int
		Harga int
	}

	type historyPelanggan struct {
		nama string
		jenis string
		tipe string
		harga int
	}

	type history[n]historyPelanggan

	type sparepart[n]DetailSparepart

	var arrSparepart sparepart
	var arrHistory history
	var iSparepart int //i temp buat nampung iterasi
	var iHistory int //i temp buat nampung history pelanggan
	const n = 10000//jumlah untuk array

//STARTING SORTING HISTORY
func sortingHistory(arr *history, total int) {
	var i, pass int
	var temp historyPelanggan
		pass = 0
			for pass < total - 1 {
				i = pass + 1
				temp = arr[i]
					for i >= 1 && temp.harga > arr[i - 1].harga {
						arr[i] = arr[i - 1]
						i--
					}
				arr[i] = temp
				pass++
			}
}
//ENDING SORTING HISTORY
//STARTING SORTING SPAREPART
func sortingSparepartHarga(arr *sparepart, total int) {
	var i, pass int
	var temp DetailSparepart
		pass = 0
			for pass < total {
				i = pass + 1
				temp = arr[i]
					for i >= 1 && temp.Harga > arr[i - 1].Harga {
						arr[i] = arr[i - 1]
						i--
					}
				arr[i] = temp
				pass++
			}
}
//ENDING SORTING SPAREPART
//STARTING SORTING SPAREPART.JUMLAH
func sortingSparepartJumlah(arr *sparepart, total int) {
	var i, pass int
	var temp DetailSparepart
		pass = 0 
			for pass < total {
				i = pass + 1
				temp = arr[i]
					for i >= 1 && temp.Jumlah > arr[i - 1].Jumlah {
						arr[i] = arr[i - 1]
						i--
					}
				arr[i] = temp
				pass++
			}
}
//ENDING SORTING SPAREPART.JUMLAH
//STARTING DELETE SPAREPART
func deleteSparepart(arr *sparepart, total *int) {
	var i, idx int
	fmt.Println("Masukkan Index data yang ingin di delete")
	fmt.Scan(&idx)
	i = idx + 1
		for idx < n && i <= *total {
			arr[i - 1] = arr[i]
			i++
		}
	*total = *total - 1
}
//ENDING DELETE SPAREPART
//STARTING SEARCHING HISTORY
func searchHistory(arr history, total *int, key string, jenis string) historyPelanggan {
	var i, found int
	var status bool
	status = false
	i = 0
		for i < *total && !status {
			if arr[i].nama == key && arr[i].jenis == jenis {
				status = true
				found = i
			}
		i++
		}
	return arr[found]
}
//ENDING SEARCHING HISTORY
//STARTING SEARCHING SPAREPART
func searchSparepart(arr sparepart, total *int, key string, merk string) DetailSparepart {
	var i, found int
	var status bool
	status = false
	i = 0
	found = 9999
		for i < *total && !status {
			if arr[i].Nama == key && arr[i].Pabrik == merk {
				status = true
				found = i
			}
		i++
		}
			if status == false {
				fmt.Println("PASTIKAN BAHWA NAMA SPAREPART DAN MERKNYA BENAR!")
			}
	return arr[found]
}
//ENDING SEARCHING SPAREPART
//STARTING EDIT SPAREPART
func editSparepart(arr *sparepart, total *int) {
	var id int
	fmt.Println("Masukkan ID sparepart yang ingin di edit(dimulai dari 0)")
	fmt.Scan(&id)
		for id != 99999 {
			fmt.Println("Masukkan nama sparepart ")
			fmt.Scan(&arr[id].Nama)
			fmt.Println("Masukkan Jumlah sparepart")
			fmt.Scan(&arr[id].Jumlah)
			fmt.Println("Masukkan nama pabrik ")
			fmt.Scan(&arr[id].Pabrik)
			fmt.Println("Masukkan tahun ")
			fmt.Scan(&arr[id].Tahun)
			fmt.Println("Masukkan harga ")
			fmt.Scan(&arr[id].Harga)
			
			fmt.Println("Masukkan ID selanjutnya, jika tidak ada ketikkan 99999")
			fmt.Scan(&id)
		}
}
//ENDING EDIT SPAREPART
//STARTING PROCCALL
	func procCall(c *int, totalSP *int, totalHS *int) {
	var katkunSparepart, katkunHistory, katkunSparepartMerk, katkunHistoryJenis string
		fmt.Println("")//pembatas
	*c = 999999999
	
	for *c != 0 {
		fmt.Println("Masukkan command yang ingin digunakan!")
		fmt.Println("1. Input Sparepart\n2. Input Service\n3. View Tabel Sparepart\n4. History Pelanggan\n5. Cari Pelanggan\n6. Cari Sparepart\n7. Edit Sparepart\n8. Delete Sparepart")
		fmt.Println("Ketikkan 0 jika ingin berhenti!")
		fmt.Scan(c)
			if *c == 2 {
				procService(&arrHistory, totalHS)
			} else if *c == 1 {
				procSparepart(&arrSparepart, totalSP)
			} else if *c == 3 {
				viewSparepart(arrSparepart, *totalSP)
			} else if *c == 4 {
				viewHistory(arrHistory, *totalHS)
			} else if *c == 5 {
				fmt.Println("Masukkan Nama dan jenis kendaraan dari pelanggan yang ingin di cari ")
				fmt.Scan(&katkunHistory, &katkunHistoryJenis)
				fmt.Println(searchHistory(arrHistory, totalHS, katkunHistory, katkunHistoryJenis))
			} else if *c == 7 {
				editSparepart(&arrSparepart, totalSP)
			} else if *c == 6 {
				fmt.Println("Masukkan Nama Sparepart dan merk dari sparepart tersebut ")
				fmt.Scan(&katkunSparepart, &katkunSparepartMerk)
				fmt.Println(searchSparepart(arrSparepart, totalSP, katkunSparepart, katkunSparepartMerk))
			} else if *c == 8 {
				deleteSparepart(&arrSparepart, totalSP)
			} else if *c == 0 {
				fmt.Println("Apps Berakhir")
			} else {
				fmt.Println("Command tidak VALID!")
			}
	}
}
//ENDING PROCCALL
//STARTING PROCSPAREPART
func procSparepart(arr *sparepart, barisSP *int ) {
		var i int
		var logic string
	fmt.Println("")//Pembatas
	fmt.Println("Selamat datang di menu input sparepart")
	fmt.Println("Silahkan masukkan nama, jumlah, pabrik, tahun, dan harga sparepart")
	fmt.Println("")//pembatas
	
	i = iSparepart
		for logic != "n" {
			fmt.Println("Masukkan nama sparepart ")
			fmt.Scan(&arr[i].Nama)
			fmt.Println("Masukkan Jumlah sparepart")
			fmt.Scan(&arr[i].Jumlah)
			fmt.Println("Masukkan nama pabrik ")
			fmt.Scan(&arr[i].Pabrik)
			fmt.Println("Masukkan tahun ")
			fmt.Scan(&arr[i].Tahun)
			fmt.Println("Masukkan harga ")
			fmt.Scan(&arr[i].Harga)
			i++
			fmt.Println("Input baru?")
			fmt.Print("y / n?? ")
			fmt.Scan(&logic)
			fmt.Println("")
		}
	iSparepart = i
	*barisSP = iSparepart
}
//ENDING PROCSPAREPART
//STARTING VIEWSPAREPART
func viewSparepart(arrSparepart sparepart,barisSP int) {
	var i int
	var perintah string
	i = 0
	fmt.Println("Jumlah\t\t","Nama\t\t","Pabrik\t\t","Tahun\t\t","Harga")
	for i < barisSP && barisSP < 10000 {
		fmt.Println(arrSparepart[i].Jumlah,"\t\t",arrSparepart[i].Nama,"\t\t",arrSparepart[i].Pabrik,"\t\t",arrSparepart[i].Tahun,"\t\t",arrSparepart[i].Harga)
		i++
	}
	fmt.Println("harga\njumlah")
	fmt.Println("Sort by . . .")
	fmt.Scan(&perintah)
		for perintah != "n" {
			if perintah == "harga" {
				sortingSparepartHarga(&arrSparepart, barisSP)
					i = 0 
					fmt.Println("Sorted by Harga Sparepart")
					fmt.Println("Jumlah\t\t","Nama\t\t","Pabrik\t\t","Tahun\t\t","Harga")
					for i < barisSP && barisSP < 10000 {
						fmt.Println(arrSparepart[i].Jumlah,"\t\t",arrSparepart[i].Nama,"\t\t",arrSparepart[i].Pabrik,"\t\t",arrSparepart[i].Tahun,"\t\t",arrSparepart[i].Harga)
						i++
					}
			} else if perintah == "jumlah" {
				sortingSparepartJumlah(&arrSparepart, barisSP)
					i = 0
					fmt.Println("Sorted by Jumlah Sparepart")
					fmt.Println("Jumlah\t\t","Nama\t\t","Pabrik\t\t","Tahun\t\t","Harga")
					for i < barisSP && barisSP < 10000 {
						fmt.Println(arrSparepart[i].Jumlah,"\t\t",arrSparepart[i].Nama,"\t\t",arrSparepart[i].Pabrik,"\t\t",arrSparepart[i].Tahun,"\t\t",arrSparepart[i].Harga)
						i++
					}
			} else {
				fmt.Println("TIDAK ADA PERINTAH SORT BY SESUAI DENGAN YANG ANDA MASUKKAN")
			}
			fmt.Println("Ketikkan n jika ingin berhenti")
			fmt.Scan(&perintah)
		}
}
//ENDING VIEWSPAREPART
//STARTING SERVICE
func procService(arrHistory *history, barisHS *int) {
	var i int
	var x, y, z string

    i = iHistory
        fmt.Println("Silahkan masukkan nama pelanggan, jenis kendaraan, dan tipe service")
        fmt.Scan(&x, &y, &z)
            for x != "n" && y != "n" && z != "n" {
                arrHistory[i].nama = x
                arrHistory[i].jenis = y
                arrHistory[i].tipe = z
                    if arrHistory[i].jenis == "motor" && arrHistory[i].tipe == "parah" {
                        arrHistory[i].harga = (50000 * 100/100)
                    } else if arrHistory[i].jenis == "motor" && arrHistory[i].tipe == "sedang"{
                        arrHistory[i].harga = (50000 * 75/100)
                    } else if arrHistory[i].jenis == "motor" && arrHistory[i].tipe == "ringan"{
                        arrHistory[i].harga = (50000 * 60/100)
                    } else if arrHistory[i].jenis == "mobil" && arrHistory[i].tipe == "parah" {
                        arrHistory[i].harga = (150000 * 100/100)
                    } else if arrHistory[i].jenis == "mobil" && arrHistory[i].tipe == "sedang"{
                        arrHistory[i].harga = (150000 * 75/100)
                    } else if arrHistory[i].jenis == "mobil" && arrHistory[i].tipe == "ringan"{
                        arrHistory[i].harga = (150000 * 60/100)
                    } else if arrHistory[i].jenis == "berat" && arrHistory[i].tipe == "parah"{
                        arrHistory[i].harga = (500000 * 100/100)
                    } else if arrHistory[i].jenis == "berat" && arrHistory[i].tipe == "sedang"{
                        arrHistory[i].harga = (500000 * 75/100)
                    } else if arrHistory[i].jenis == "berat" && arrHistory[i].tipe  == "ringan"{
                        arrHistory[i].harga = (500000 * 60/100)
                    }
				fmt.Println(arrHistory[i].harga)
				i++
				fmt.Println("Ketikkan n n n jika ingin berhenti")
                fmt.Scan(&x, &y, &z)
			}
		iHistory = i
		*barisHS = iHistory
}
//ENDING SERVICE
//STARTING HISTORY SERVICE
func viewHistory(arrHistory history, barisHS int) {
	var i int
	i = 0
	sortingHistory(&arrHistory, barisHS)
	fmt.Println("Tabel dibawah sudah terurut sesuai dengan biaya yang paling besar")
	fmt.Println("Nama		Jenis		Tipe		Biaya")
	for i < barisHS && barisHS < 10000 {
		fmt.Println(arrHistory[i].nama,"		",arrHistory[i].jenis,"			",arrHistory[i].tipe,"		",arrHistory[i].harga)
		i++
	}
}
//ENDING HISTORY SERVICE
//STARTING PROCLOGIN
func procLogin(uid string, pass string) {
	var logCondition bool

	logCondition = false

	fmt.Println("")//pembatas
	fmt.Println("Silahkan Masukkan ID dan Password yang benar!")
	fmt.Print("Silahkan masukkan ID ")
	fmt.Scanln(&uid)
	fmt.Print("Silahkan masukkan Password ")
	fmt.Scanln(&pass)

		for logCondition != true {
			if uid == "admin" && pass == "admin" {
				fmt.Println("Login Berhasil")
				logCondition = true
			} else {
				fmt.Println("")//pembatas
				fmt.Println("Login GAGAL!")
				fmt.Println("Silahkan Masukkan ID dan Password yang benar!")
				fmt.Print("Silahkan masukkan ID ")
				fmt.Scanln(&uid)
				fmt.Print("Silahkan masukkan Password ")
				fmt.Scanln(&pass)
			}
		}
}
//ENDING PROCLOGIN
//STARTING MAIN CODE
func main () {
	var barisSP int
	var barisHS int
	var uid, pass string //variabel yang dipake buat navigasi utama
	var command int
	iHistory = 0
	iSparepart = 0
	fmt.Println("Selamat Datang dalam aplikasi sistem informasi service motor!")
	fmt.Println("Silahkan login terlebih dahulu!")
	procLogin(uid, pass)//prosedur login
	fmt.Println("Masukkan perintah sesuai yang ingin anda lakukan! (case sensitive!)")
	procCall(&command, &barisSP, &barisHS)//prosedur panggil command dengan mark "n"
}
//ENDING MAIN CODE