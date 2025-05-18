package main

import "fmt"

type PerSahaman struct {
	kode  string
	nama  string
	harga int
	naik  float64
	jumlahpembeli int
}

type Portobeli struct {
	kode  string
	nama  string
	lot   int
	total int
	naik  int
	persentase float64
}

var Portofolio [100]Portobeli
var Saham [10]PerSahaman
var BesarPorto int
var jmlhUang int
var KodeSaham string

func StockSaham() {
	Saham = [10]PerSahaman{
		{"Saham 1", "Saham GOGO Curry", 10000, 3.92, 1000},
		{"Saham 2", "Saham Mak Gembus", 5000, 500000, 1000},
		{"Saham 3", "Saham Pak Gembus", 8000, 800000, 786228},
		{"Saham 4", "Saham Pak Gembus", 8000, 800000, 76876852},
		{"Saham 5", "Saham Pak Gembus", 8000, 800000, 78789},
		{"Saham 6", "Saham Pak Gembus", 8000, 800000, 87},
		{"Saham 7", "Saham Pak Gembus", 8000, 800000, 9},
		{"Saham 8", "Saham Pak Gembus", 8000, 800000, 987},
		{"Saham 9", "Saham Pak Gembus", 8000, 800000, 71},
		{"Saham 10", "Saham Pak Gembus", 8000, 800000, 99},
	}
}

func Menu() {
	fmt.Println("\n Simulasi Pasar Saham Virtual ")
	fmt.Printf("Saldo: %d\n", jmlhUang)
	fmt.Println("1. Lihat Daftar Saham")
	fmt.Println("2. Beli Saham")
	fmt.Println("3. Jual Saham")
	fmt.Println("4. Lihat Portofolio")
	fmt.Println("5. Lihat Perkembangan")
	fmt.Println("6. Statistik Akhir")
	fmt.Println("7. Lihat Riwayat Transaksi")
	fmt.Println("8. Tambah Saldo Virtual")
	fmt.Println("9. Exit")
	fmt.Print("Pilih opsi: ")
}

func tampilkanSaham() {
	var i int
	fmt.Println("Daftar Saham:")
	fmt.Printf("%-6s %-25s %-15s %-15s\n", "Kode", "Nama", "Harga", "Kenaikan")
	for i = 0; i < 10; i++ {
		fmt.Printf("%-6s %-25s %-15d %-15.2f\n",
			Saham[i].kode, Saham[i].nama, Saham[i].harga, Saham[i].naik)
	}
}

func MenuSaham() {
	var pilih, idxnama, idxkode int
	var tipeUrut, cari, nama, jenisurut string

	for {
		
		tampilkanSaham()
		fmt.Println("1. Urutkan Saham ")
		fmt.Println("2. Kembali ke Menu")
		fmt.Println("3. Beli Saham")
		fmt.Println("4. Cek Detail Saham")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Diurutkan Berdasarkan (Harga / Jumlah Pembeli): ")
			fmt.Scan(&jenisurut)
			if jenisurut == "Harga" {
				fmt.Println("DIurutkan menurut (Termahal / Termurah): ")
				fmt.Scan(&tipeUrut)
				if tipeUrut == "Termurah" {
					UrutTermurahHarga()
				} else if tipeUrut == "Termahal" {
					UrutTermahalHarga()
				} else {
					fmt.Println("Metode pengurutan tidak valid!")
				}
			} else if jenisurut == "Jumlah Pembeli" {
				fmt.Println("DIurutkan menurut (Terbanyak / Terdikit): ")
				fmt.Scan(&tipeUrut)
				if tipeUrut == "Terdikit" {
					UrutSedikitPembelian()
				} else if tipeUrut == "Terbanyak" {
					UrutBanyakPembelian()
				} else {
					fmt.Println("Metode pengurutan tidak valid!")
				}
			}
			
		case 2:
			return
		case 3:
			BeliSaham()
			return
		case 4:
			fmt.Print("Cari Berdasarkan (Kode / Nama Perusahaan): ")
			fmt.Scan(&cari)
			if cari == "Kode" {
				fmt.Println("Masukkan Kode: ")
				fmt.Scan(&KodeSaham)
				idxkode = CariKode(KodeSaham)
				if idxkode != -1 {
					fmt.Printf("%s %s %d %2.f", Saham[idxkode].kode, Saham[idxkode].nama, Saham[idxkode].harga, Saham[idxkode].naik)
				} else {
					fmt.Println("Kode Tidak Ditemukan (Invalid)")
				}
			} else if cari == "Nama Perusahaan" {
				fmt.Println("Masukkan Nama Perusahaan: ")
				fmt.Scan(&nama)
				idxnama = CariNama(nama)
				if idxnama == -1 {
					fmt.Printf("%s %s %d %2.f", Saham[idxkode].kode, Saham[idxkode].nama, Saham[idxkode].harga, Saham[idxkode].naik)
				} else {
					fmt.Println("Nama Perusahaan Tidak Ditemukan (Invalid)")
				}
			} else {
				fmt.Println("Pencarian error")
			}
		default:
			fmt.Println("Opsi tidak ada (Invalid)!")
		}
	}
}

func UrutTermahalHarga() {
	var i, idx, pass int
    var temp PerSahaman

    pass = 1
    for pass < 10 {
        idx = pass - 1
        i = pass
        for i < 10 {
            if Saham[i].harga > Saham[idx].harga {
                idx = i
            }
            i++
        }
        temp = Saham[pass-1]
        Saham[pass-1] = Saham[idx]
        Saham[idx] = temp
        pass++
    }

    fmt.Println("\nHasil pengurutan (Harga Termahal):")
    fmt.Printf("%-6s %-25s %-15s %-15s %-15s\n", "Kode", "Nama", "Harga", "Kenaikan", "Jumlah Pembeli")
    for i = 0; i < 10; i++ {
        fmt.Printf("%-6s %-25s %-15d %-15.2f %-15d\n",
            Saham[i].kode, Saham[i].nama, Saham[i].harga, Saham[i].naik, Saham[i].jumlahpembeli)
    }
}

func UrutTermurahHarga() {
	var i, idx, pass int
    var temp PerSahaman

    pass = 1
    for pass < 10 {
        idx = pass - 1
        i = pass
        for i < 10 {
            if Saham[i].harga < Saham[idx].harga {
                idx = i
            }
            i++
        }
        temp = Saham[pass-1]
        Saham[pass-1] = Saham[idx]
        Saham[idx] = temp
        pass++
    }

    fmt.Println("\nHasil pengurutan (Harga Termurah):")
    fmt.Printf("%-6s %-25s %-15s %-15s %-15s\n", "Kode", "Nama", "Harga", "Kenaikan", "Jumlah Pembeli")
    for i = 0; i < 10; i++ {
        fmt.Printf("%-6s %-25s %-15d %-15.2f %-15d\n",
            Saham[i].kode, Saham[i].nama, Saham[i].harga, Saham[i].naik, Saham[i].jumlahpembeli)
    }
}

func UrutBanyakPembelian() {
	var i, pass int
	var temp PerSahaman

	for pass = 1; pass < 10; pass++ {
		temp = Saham[pass]
		i = pass
		for i > 0 && Saham[i-1].jumlahpembeli < temp.jumlahpembeli {
			Saham[i] = Saham[i-1]
			i--
		}
		Saham[i] = temp
	}

	fmt.Println("\nHasil pengurutan (Jumlah Pembelian Terbanyak):")
	fmt.Printf("%-6s %-25s %-15s %-15s %-15s\n", "Kode", "Nama", "Harga", "Kenaikan", "Jumlah Pembeli")
	for i = 0; i < 10; i++ {
		fmt.Printf("%-6s %-25s %-15d %-15.2f %-15d\n",
			Saham[i].kode, Saham[i].nama, Saham[i].harga, Saham[i].naik, Saham[i].jumlahpembeli)
	}
}

func UrutSedikitPembelian() {
	var i, pass int
	var temp PerSahaman

	for pass = 1; pass < 10; pass++ {
		temp = Saham[pass]
		i = pass
		for i > 0 && Saham[i-1].jumlahpembeli > temp.jumlahpembeli {
			Saham[i] = Saham[i-1]
			i--
		}
		Saham[i] = temp
	}

	fmt.Println("\nHasil pengurutan (Jumlah Pembelian Terbanyak):")
	fmt.Printf("%-6s %-25s %-15s %-15s %-15s\n", "Kode", "Nama", "Harga", "Kenaikan", "Jumlah Pembeli")
	for i = 0; i < 10; i++ {
		fmt.Printf("%-6s %-25s %-15d %-15.2f %-15d\n",
			Saham[i].kode, Saham[i].nama, Saham[i].harga, Saham[i].naik, Saham[i].jumlahpembeli)
	}

}

func BeliSaham() {
	var konfirmasi string
	var i, n, idxSaham, Belilot, lembar, TotalSemua int

	fmt.Println("1. Masukkan kode saham")
	fmt.Println("2. Tampilkan daftar saham")
	fmt.Println("3. Kembali ke menu sebelumnya")
	fmt.Scan("Pilihan: ", &n)
	switch n {
	case 1:
		fmt.Print("Kode Saham: ")
		fmt.Scan(&KodeSaham)
		if !CekKodeSaham(KodeSaham) {
			fmt.Println("Kode saham tidak valid!")
			return
		}
		// Sequential Search
		idxSaham = 0
		for i = 0; i < 10; i++ {
			if Saham[i].kode == KodeSaham {
				idxSaham = i
			}
		}
		//Perubahan dari lembar ke lot
		fmt.Printf("Saldo Anda: %d\n", jmlhUang)
		fmt.Println("Harga yang ingin dibeli: ")
		fmt.Scan(&lembar)
		if lembar < Saham[i].harga {
			fmt.Println("Tidak memenuhi minimal pembelian")
			return
		}
		fmt.Println("Berapa Lot: ")
		fmt.Scan(&Belilot)
		if Belilot <= 0 {
			fmt.Print("Invalid!")
		}
		TotalSemua = (Belilot * 100) * lembar
		if TotalSemua > jmlhUang {
			fmt.Println("Uang anda tidak mencukupi!")
			return
		}
		fmt.Printf("Pembelian Saham %s (%s)\n", Saham[idxSaham].nama, Saham[idxSaham].kode)
		fmt.Printf("Dengan Total %d sebanyak %d\n", TotalSemua, Belilot)
		fmt.Print("Konfirmasi pembelian? (y/n): ")
		fmt.Scan(&konfirmasi)
		if konfirmasi != "y" {
			fmt.Println("Pembelian dibatalkan.")
			return
		}
		Portofolio[BesarPorto] = Portobeli{Saham[idxSaham].kode, Saham[idxSaham].nama, Belilot, TotalSemua, Saham[idxSaham].naik}
		BesarPorto++
		jmlhUang -= TotalSemua
		fmt.Printf("Berhasil membeli saham %s (%s) dengan total %d lot seharga %d\n", Saham[idxSaham].nama, Saham[idxSaham].kode, Belilot, TotalSemua)

	case 2:
		tampilkanSaham()
		fmt.Print("Masukkan kode saham: ")
		fmt.Scan(&KodeSaham)
		if !CekKodeSaham(KodeSaham) {
			fmt.Println("Kode saham tidak valid!")
			return
		}
		idxSaham = -1
		i = 0
		for i < 10 && idxSaham == -1 {
			if Saham[i].kode == KodeSaham {
				idxSaham = i
			}
			i++
		}
		//Perubahan dari lembar ke lot
		fmt.Printf("Saldo Anda: %d\n", jmlhUang)
		fmt.Println("Harga yang ingin dibeli: ")
		fmt.Scan(&lembar)
		if lembar < Saham[i].harga {
			fmt.Println("Tidak memenuhi minimal pembelian")
			return
		}
		fmt.Println("Berapa Lot: ")
		fmt.Scan(&Belilot)
		if Belilot <= 0 {
			fmt.Print("Invalid! Minimum pembelian 1 lot!")
			return
		}
		TotalSemua = (Belilot * 100) * lembar
		if TotalSemua > jmlhUang {
			fmt.Println("Uang anda tidak mencukupi!")
			return
		}
		fmt.Printf("Pembelian Saham %s (%s)\n", Saham[idxSaham].nama, Saham[idxSaham].kode)
		fmt.Printf("Dengan Total %d sebanyak %d\n", TotalSemua, Belilot)
		fmt.Print("Konfirmasi pembelian? (y/n): ")
		fmt.Scan(&konfirmasi)
		if konfirmasi != "y" {
			fmt.Println("Pembelian dibatalkan.")
			return
		}
		Portofolio[BesarPorto] = Portobeli{Saham[idxSaham].kode, Saham[idxSaham].nama, Belilot, TotalSemua, Saham[idxSaham].naik}
		BesarPorto++
		jmlhUang -= TotalSemua
		fmt.Printf("Berhasil membeli saham %s (%s) dengan total %d lot seharga %d\n", Saham[idxSaham].nama, Saham[idxSaham].kode, Belilot, TotalSemua)
	case 3:
		return
	default:
		fmt.Println("Opsi tidak ada (Invalid)!")
	}
}

func CekKodeSaham(Kode string) bool {
	var i int
	for i = 0; i < 10; i++ {
		if Saham[i].kode == Kode {
			return true
		}
	}
	return false
}

func CariKode(kode string) int {
	var i int
	for i = 0; i < 10; i++ {
		if Saham[i].kode == kode {
			return i
		}
	}
	return -1
}

func CariNama(nama string) int {
	var i int
	for i = 0; i < 10; i++ {
		if Saham[i].nama == nama {
			return i
		}
	}
	return -1
}

func JualSaham() {
	var i, idxPorto, lotjual, hargajual, dapat int
	var konfirmasi string

	idxPorto = 0
	if BesarPorto == 0 {
		fmt.Println("Portofolio kosong, tidak ada saham yang bisa dijual")
		return
	}
	MenuPortofolio()
	fmt.Print("Masukkan Kode atau Nama saham yang ingin dijual: ")
	fmt.Scan(&KodeSaham)
	if KodeSaham == "kembali" || KodeSaham == "cancel" {
		return
	}
	for i = 0; i < BesarPorto; i++ {
		if Portofolio[i].kode == KodeSaham || Portofolio[i].nama == KodeSaham {
			idxPorto = i
		}
	}
	if idxPorto == 0 {
		fmt.Println("Saham tidak ada di Portofolio (Invalid)")
		return
	}
	fmt.Printf("Penjualan Saham %s (%s) dengan total %d (%d lot)", Portofolio[idxPorto].kode, Portofolio[idxPorto].nama, Portofolio[idxPorto].total, Portofolio[idxPorto].lot)
	fmt.Println("Harga yang ingin dijual: ")
	fmt.Scan(&hargajual)
	if hargajual <= 0 {
		fmt.Println("Invalid!")
	}
	fmt.Println("Berapa lot: ")
	fmt.Scan(&lotjual)
	if lotjual <= 0 {
		fmt.Println("Invalid! Minimum penjualan 1 lot!")
		return
	}
	dapat = (lotjual * 100) * hargajual
	if dapat > Portofolio[idxPorto].total {
		fmt.Println("Request melebihi total saham yang dimiliki!")
		return
	}
	fmt.Println("Konfirmasi Penjualan (y/n): ")
	fmt.Scan(&konfirmasi)
	if konfirmasi != "y" {
		fmt.Println("Penjualan dibatalkan.")
		return
	}
	Portofolio[idxPorto].lot -= lotjual
	jmlhUang += dapat
	fmt.Printf("Berhasil menjual saham %s (%s) dengan total %d lot seharga %d\n", Portofolio[idxPorto].nama, Portofolio[idxPorto].kode, lotjual, dapat)
	if Portofolio[idxPorto].lot == 0 {
		for i = idxPorto; i < BesarPorto-1; i++ {
			Portofolio[i] = Portofolio[i+1]
		}
		BesarPorto--
	}

}

func MenuPortofolio() {
	var i, j, totalporto int
    if BesarPorto == 0 {
        fmt.Println("\nPortofolio kosong!")
        return
    }
	totalporto = 0
    fmt.Printf("\nPortofolio (Saldo: %d):\n", jmlhUang)
    fmt.Printf("%-6s %-25s %-15s %-15s %-15s\n", "Kode", "Nama", "Jumlah Lot", "Total Pembelian", "Kenaikan")
    totalKeuntungan := 0.0
    for i = 0; i < BesarPorto; i++ {
        for j = 0; j < 10; j++ {
            if Saham[j].kode == Portofolio[i].kode {
				totalporto = (Portofolio[i].lot * 100) * Saham[j].harga
                Portofolio[i].total = Saham[j].harga
                kenaikan = Saham[j].naik
                break
            }
        }
        fmt.Printf("%-6s %-25s %-15d %-15d %-15.2f%\n", Portofolio[i].kode, Portofolio[i].nama, Portofolio[i].lot, Portofolio[i].total, Portofolio[i].naik)
    }
    fmt.Println("Total Saldo dengan Keuntungan: %.2f\n", float64(jmlhUang)+totalKeuntungan)
}
