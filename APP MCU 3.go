package main

import "fmt"

const NMAX int = 1000

type tLayanan struct {
	tingkat string
	nominal int
}

type tPasien struct {
	nama, hasil_check_up string
	id                   int
	waktu_check_up       waktu_check_up
	jenis_check_up       tLayanan
}

type waktu_check_up struct {
	tanggal, bulan, tahun int
}

type tabPasien [NMAX]tPasien
type tabLayanan [NMAX]tLayanan

func main() {
	var data_pasien tabPasien
	var data_layanan tabLayanan
	data_layanan = [NMAX]tLayanan{
		{tingkat: "Umum", nominal: 100000},
		{tingkat: "BPJS", nominal: 200000},
		{tingkat: "Asuransi", nominal: 300000},
		{tingkat: "VIP", nominal: 400000},
		{tingkat: "VVIP", nominal: 500000},
	}
	data_pasien = [NMAX]tPasien{
		{nama: "shania", id: 12, waktu_check_up: waktu_check_up{tahun: 2021, bulan: 9, tanggal: 13}, jenis_check_up: data_layanan[0], hasil_check_up: "buta"},
		{nama: "rahmalia", id: 15, waktu_check_up: waktu_check_up{tahun: 2022, bulan: 12, tanggal: 25}, jenis_check_up: data_layanan[1], hasil_check_up: "lumpuh"},
		{nama: "galuh", id: 17, waktu_check_up: waktu_check_up{tahun: 2023, bulan: 3, tanggal: 9}, jenis_check_up: data_layanan[2], hasil_check_up: "HIV"},
		{nama: "ajeng", id: 11, waktu_check_up: waktu_check_up{tahun: 2024, bulan: 10, tanggal: 3}, jenis_check_up: data_layanan[3], hasil_check_up: "diare"},
		{nama: "shany", id: 10, waktu_check_up: waktu_check_up{tahun: 2019, bulan: 1, tanggal: 31}, jenis_check_up: data_layanan[4], hasil_check_up: "diabetes"},
		{nama: "jey", id: 13, waktu_check_up: waktu_check_up{tahun: 2018, bulan: 8, tanggal: 17}, jenis_check_up: data_layanan[4], hasil_check_up: "cacar air"},
	}
	tampilan_home(&data_pasien, &data_layanan)
}

func tampilan_home(A *tabPasien, B *tabLayanan) {
	var opsi int
	var nData, mData int = 6, 5
	for opsi != 9 {
		fmt.Println("---------------------------------------------")
		fmt.Println("||                                         ||")
		fmt.Println("||    # Shania Rahmalia # 103032300018 #   ||")
		fmt.Println("||      # Galuh Ajeng # 103032300087 #     ||")
		fmt.Println("||      # Aplikasi Medical Check Up #      ||")
		fmt.Println("||                                         ||")
		fmt.Println("---------------------------------------------")
		fmt.Println("1. Penambahan Data Pasien")
		fmt.Println("2. Penghapusan Data Pasien")
		fmt.Println("3. Pengeditan Data Pasien")
		fmt.Println("4. Pencarian Data Pasien")
		fmt.Println("5. Penambahan Paket Layanan")
		fmt.Println("6. Penghapusan Paket Layanan")
		fmt.Println("7. Pengeditan Paket Layanan")
		fmt.Println("8. Menampilkan Data")
		fmt.Println("9. Keluar")
		fmt.Println("---------------------------------------------")
		fmt.Print("Masukkan Opsi: ")
		fmt.Scan(&opsi)
		if opsi == 1 {
			main_tambah_pasien(A, *B, &nData, mData)
		} else if opsi == 2 {
			main_hapus_pasien(A, *B, &nData)
		} else if opsi == 3 {
			main_edit_pasien(A, B, nData, mData)
		} else if opsi == 4 {
			main_cari_pasien(*A, *B, nData, mData)
		} else if opsi == 5 {
			main_tambah_paket(B, &mData)
		} else if opsi == 6 {
			main_hapus_paket(A, B, &nData, &mData)
		} else if opsi == 7 {
			main_edit_layanan(A, B, nData, mData)
		} else if opsi == 8 {
			main_display(*A, *B, nData, mData)
		} else if opsi < 1 || opsi > 9 {
			fmt.Println("Opsi Invalid")
		}
	}
}

func main_tambah_pasien(A *tabPasien, B tabLayanan, n *int, m int) {
	var opsi, tahun, bulan, tanggal int
	var id int
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Tambah Pasien")
	fmt.Print("Masukkan Nama Pasien: ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Masukkan ID Pasien: ")
	fmt.Scan(&id)
	for !cek_id_pasien(*A, *n, id) {
		fmt.Print("Masukkan ID Pasien yang valid: ")
		fmt.Scan(&id)
	}
	A[*n].id = id
	fmt.Print("Masukkan Rekap Pasien: ")
	fmt.Scan(&A[*n].hasil_check_up)
	fmt.Print("Masukkan Waktu Check Up Pasien (YYYY/MM/DD): ")
	fmt.Scan(&tahun, &bulan, &tanggal)
	for !cek_waktu_pasien(tahun, bulan, tanggal) {
		fmt.Print("Masukkan Waktu Check Up Pasien yang valid (YYYY/MM/DD): ")
		fmt.Scan(&tahun, &bulan, &tanggal)
	}
	A[*n].waktu_check_up.tahun = tahun
	A[*n].waktu_check_up.bulan = bulan
	A[*n].waktu_check_up.tanggal = tanggal
	list_paket(B, m)
	fmt.Print("Masukkan Jenis Paket Pasien Berdasarkan List Diatas: ")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > m {
		fmt.Println("Input Invalid")
		fmt.Print("Masukkan Jenis Paket Pasien Berdasarkan List Diatas: ")
		fmt.Scan(&opsi)
	}
	A[*n].jenis_check_up = B[opsi-1]
	*n++
	fmt.Println("---------------------------------------------")
	fmt.Println("Data Pasien Berhasil Ditambahkan")
	fmt.Println("---------------------------------------------")
}

func list_paket(B tabLayanan, m int) {
	fmt.Println("Jenis Paket:")
	for i := 0; i < m; i++ {
		fmt.Printf("%d. %s %d\n", i+1, B[i].tingkat, B[i].nominal)
	}
}

func main_hapus_pasien(A *tabPasien, B tabLayanan, n *int) {
	var opsi, y int
	var x string
	var idx int = -1
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Hapus Data Pasien")
	fmt.Println("Cari Data Pasien yang akan Dihapus berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Opsi Invalid")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Print("Masukkan nama pasien: ")
		fmt.Scan(&x)
		idx = cari_nama(*A, *n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID pasien: ")
		fmt.Scan(&y)
		ascend_insertion_sort(A, *n)
		idx = cari_id(*A, *n, y)
	}
	hapus_pasien(A, B, n, idx)
	fmt.Println("---------------------------------------------")
}

func cari_nama(A tabPasien, n int, x string) int {
	var idx int = -1
	for i := 0; i < n && idx == -1; i++ {
		if A[i].nama == x {
			idx = i
		}
	}
	return idx
}

func cari_id(A tabPasien, n, y int) int {
	var left, mid, right int
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if A[mid].id == y {
			return mid
		} else if y < A[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func ascend_insertion_sort(A *tabPasien, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if A[j-1].id > A[j].id {
				A[j-1], A[j] = A[j], A[j-1]
			}
			j = j - 1
		}
	}
}

func hapus_pasien(A *tabPasien, B tabLayanan, n *int, idx int) {
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data pasien berhasil dihapus")
	}

}

func main_edit_pasien(A *tabPasien, B *tabLayanan, n, m int) {
	var opsi int
	var id, tahun, bulan, tanggal int
	var idx int = -1
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Edit Data Pasien")
	fmt.Print("Masukkan ID pasien: ")
	fmt.Scan(&id)
	ascend_insertion_sort(A, n)
	idx = cari_id(*A, n, id)
	display_pasien(*A, idx)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan")
		fmt.Println("---------------------------------------------")

	} else {
		for opsi != 6 {
			fmt.Println("1. Nama")
			fmt.Println("2. Hasil Check Up")
			fmt.Println("3. Waktu Check Up")
			fmt.Println("4. Jenis Paket")
			fmt.Println("5. Tampilkan Data Pasien")
			fmt.Println("6. Kembali")
			fmt.Print("Pilih opsi data yang ingin diubah: ")
			fmt.Scan(&opsi)
			if opsi == 1 {
				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scan(&A[idx].nama)
			} else if opsi == 2 {
				fmt.Print("Masukkan Hasil Check Up Baru: ")
				fmt.Scan(&A[idx].hasil_check_up)
			} else if opsi == 3 {
				fmt.Print("Masukkan waktu_check_up Check Up Pasien (YYYY/MM/DD): ")
				fmt.Scan(&tahun, &bulan, &tanggal)
				for !cek_waktu_pasien(tahun, bulan, tanggal) {
					fmt.Print("Masukkan waktu_check_up Check Up Pasien yang valid (YYYY/MM/DD): ")
					fmt.Scan(&tahun, &bulan, &tanggal)
				}
				A[idx].waktu_check_up.tahun = tahun
				A[idx].waktu_check_up.bulan = bulan
				A[idx].waktu_check_up.tanggal = tanggal
			} else if opsi == 4 {
				list_paket(*B, m)
				fmt.Print("Masukkan Jenis Paket Pasien Berdasarkan List Diatas: ")
				fmt.Scan(&opsi)
				for opsi < 1 || opsi > m {
					fmt.Println("Input Invalid")
					fmt.Print("Masukkan Jenis Paket Pasien Berdasarkan List Diatas: ")
					fmt.Scan(&opsi)
				}
				A[idx].jenis_check_up = (*B)[opsi-1]
			} else if opsi == 5 {
				fmt.Printf("ID: %d\n", A[idx].id)
				fmt.Printf("Nama: %s\n", A[idx].nama)
				fmt.Printf("Hasil Check Up: %s\n", A[idx].hasil_check_up)
				fmt.Printf("Waktu Check Up: %d/%d/%d\n", A[idx].waktu_check_up.tahun, A[idx].waktu_check_up.bulan, A[idx].waktu_check_up.tanggal)
				fmt.Printf("Jenis Paket: %s\n", A[idx].jenis_check_up.tingkat)
				fmt.Printf("Nominal: %d\n", A[idx].jenis_check_up.nominal)
			}
		}
		fmt.Println("---------------------------------------------")
	}

}

func main_cari_pasien(A tabPasien, B tabLayanan, n, m int) {
	var opsi, y int
	var x string
	var idx int = -1
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Cari Data Pasien")
	fmt.Println("Cari Data Pasien berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Opsi Invalid")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Print("Masukkan nama pasien: ")
		fmt.Scan(&x)
		idx = cari_nama(A, n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID pasien: ")
		fmt.Scan(&y)
		ascend_insertion_sort(&A, n)
		idx = cari_id(A, n, y)
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		fmt.Println("---------------------------------------------")
	} else {
		fmt.Printf("ID: %d\n", A[idx].id)
		fmt.Printf("Nama: %s\n", A[idx].nama)
		fmt.Printf("Hasil Check Up: %s\n", A[idx].hasil_check_up)
		fmt.Printf("Waktu Check Up: %d/%d/%d\n", A[idx].waktu_check_up.tahun, A[idx].waktu_check_up.bulan, A[idx].waktu_check_up.tanggal)
		fmt.Printf("Jenis Paket: %s\n", A[idx].jenis_check_up.tingkat)
		fmt.Printf("Nominal: %d\n", A[idx].jenis_check_up.nominal)
		fmt.Println("---------------------------------------------")
	}
}

func main_tambah_paket(B *tabLayanan, m *int) {
	var nominal int
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Tambah Paket Layanan")
	fmt.Print("Masukkan Nama Paket Baru: ")
	fmt.Scan(&B[*m].tingkat)
	fmt.Print("Masukkan Nominal Paket Baru: ")
	fmt.Scan(&nominal)
	B[*m].nominal = nominal
	*m++
	fmt.Println("---------------------------------------------")
	fmt.Println("Data Paket Berhasil Ditambahkan")
	fmt.Println("---------------------------------------------")
}

func main_hapus_paket(A *tabPasien, B *tabLayanan, n, m *int) {
	var x string
	var idx, i int = -1, 0
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Hapus Paket Layanan")
	fmt.Print("Masukkan Nama Paket yang ingin Dihapus: ")
	fmt.Scan(&x)
	for i < *m && idx == -1 {
		if (*B)[i].tingkat == x {
			idx = i
		}
		i++
	}
	if idx == -1 {
		fmt.Println("Paket tidak ditemukan")
		fmt.Println("---------------------------------------------")
	} else {
		hapus_paket(A, B, n, m, idx)
		fmt.Println("Paket Berhasil Dihapus")
		fmt.Println("---------------------------------------------")
	}

}

func hapus_paket(A *tabPasien, B *tabLayanan, n, m *int, idx int) {
	var ketemu bool = true
	for i := 0; i < *n; i++ {
		if (*A)[i].jenis_check_up.tingkat == (*B)[idx].tingkat {
			fmt.Printf("Paket %s tidak bisa dihapus karena sedang digunakan pasien %s\n", (*B)[idx].tingkat, (*A)[i].nama)
			ketemu = false
		}
	}
	if ketemu {
		for i := idx; i < *m-1; i++ {
			(*B)[i] = (*B)[i+1]
		}
		*m--
	}

}

func main_edit_layanan(A *tabPasien, B *tabLayanan, n, m int) {
	var x, nama string
	var idx, opsi, nominal int
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu Edit Paket Layanan")
	fmt.Print("Masukkan Nama Paket yang ingin Diedit: ")
	fmt.Scan(&x)
	idx = -1
	for i := 0; i < m; i++ {
		if (*B)[i].tingkat == x {
			idx = i
			i += m
		}
	}
	if idx == -1 {
		fmt.Println("Paket tidak ditemukan")
		fmt.Println("---------------------------------------------")
	} else {
		for opsi != 3 {
			fmt.Println("1. Nama Paket")
			fmt.Println("2. Nominal Paket")
			fmt.Println("3. Kembali")
			fmt.Print("Pilih opsi data yang ingin diubah: ")
			fmt.Scan(&opsi)
			if opsi == 1 {
				fmt.Print("Masukkan Nama Paket Baru: ")
				fmt.Scan(&nama)
				update_nama_layanan(&*A, *B, n, idx, nama)
				(*B)[idx].tingkat = nama
			} else if opsi == 2 {
				fmt.Print("Masukkan Nominal Paket Baru: ")
				fmt.Scan(&nominal)
				update_harga_layanan(&*A, *B, n, idx, nominal)
				(*B)[idx].nominal = nominal
			}
		}
	}
	fmt.Println("---------------------------------------------")
}
func update_nama_layanan(A *tabPasien, B tabLayanan, n, idx int, x string) {
	for i := 0; i < n; i++ {
		if A[i].jenis_check_up.tingkat == B[idx].tingkat {
			A[i].jenis_check_up.tingkat = x
		}
	}
}
func update_harga_layanan(A *tabPasien, B tabLayanan, n, idx, nominal int) {
	for i := 0; i < n; i++ {
		if A[i].jenis_check_up.tingkat == B[idx].tingkat {
			A[i].jenis_check_up.nominal = nominal
		}
	}
}

func main_display(A tabPasien, B tabLayanan, n, m int) {
	var opsi int
	var y1, m1, d1, y2, m2, d2 int
	fmt.Println("Menu Menampilkan Data")
	fmt.Println("1. Display Pemasukkan Berdasarkan Periode")
	fmt.Println("2. Ascending Waktu")
	fmt.Println("3. Descending Waktu")
	fmt.Println("4. Ascending Paket")
	fmt.Println("5. Descending Paket")
	fmt.Println("6. Cetak Data Layanan")
	fmt.Println("7. Kembali")
	fmt.Println("===================================================")
	fmt.Print("Masukkan Opsi (1/2/3/4/5/6/7): ")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > 7 {
		fmt.Println("Opsi Invalid")
		fmt.Print("Masukkan Opsi (1/2/3/4/5/6/7): ")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Menu Cari Pasien Periodik")
		fmt.Println("Masukkan Waktu Awal (YYYY/MM/DD): ")
		fmt.Scan(&y1, &m1, &d1)
		for !cek_waktu_pasien(y1, m1, d1) {
			fmt.Scan(&y1, &m1, &d1)
		}
		fmt.Println("Masukkan Waktu Akhir (YYYY/MM/DD): ")
		fmt.Scan(&y2, &m2, &d2)
		for !cek_waktu_pasien(y2, m2, d2) || y1 > y2 || (y1 == y2 && m2 < m1) || (y1 == y2 && m2 == m1 && d2 < d1) {
			fmt.Scan(&y2, &m2, &d2)
		}
		hitung_pemasukkan(A, n, y1, m1, d1, y2, m2, d2)
	} else if opsi == 2 {
		descending_waktu(A, n)
	} else if opsi == 3 {
		ascending_waktu(A, n)
	} else if opsi == 4 {
		descending_paket(A, n)
	} else if opsi == 5 {
		ascending_paket(A, n)
	} else if opsi == 6 {
		list_paket(B, m)
	}
}

func hitung_pemasukkan(A tabPasien, n, y1, m1, d1, y2, m2, d2 int) {
	var i, total int
	var hari1, hari2, hari_cek int
	hari1 = y1*360 + bulan_hari(m1, y1) + d1
	hari2 = y2*360 + bulan_hari(m2, y2) + d2
	fmt.Println(hari1, hari2, hari_cek)

	for i = 0; i < n; i++ {
		hari_cek = A[i].waktu_check_up.tahun*360 + bulan_hari(A[i].waktu_check_up.bulan, A[i].waktu_check_up.tahun) + A[i].waktu_check_up.tanggal
		if hari2 >= hari_cek && hari1 <= hari_cek {
			total += A[i].jenis_check_up.nominal
		}
	}
	fmt.Printf("Total Pemasukkan Mulai Dari %d/%d/%d hingga %d/%d/%d adalah sebesar Rp. %d Rupiah \n", y1, m1, d1, y2, m2, d2, total)
}

func ascending_waktu(A tabPasien, n int) {
	sort_tahun_descend(&A, n)
	sort_bulan_descend(&A, n)
	sort_tanggal_descend(&A, n)
	for i := 0; i < n; i++ {
		display_pasien(A, i)
	}
}

func sort_tahun_descend(A *tabPasien, n int) {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu_check_up.tahun > A[idx_max].waktu_check_up.tahun {
				idx_max = i
			}
		}
		A[pass], A[idx_max] = A[idx_max], A[pass]
	}
}

func sort_bulan_descend(A *tabPasien, n int) {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu_check_up.bulan > A[idx_max].waktu_check_up.bulan && A[i].waktu_check_up.tahun >= A[idx_max].waktu_check_up.tahun {
				idx_max = i
			}
		}
		A[pass], A[idx_max] = A[idx_max], A[pass]
	}
}

func sort_tanggal_descend(A *tabPasien, n int) {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu_check_up.tanggal > A[idx_max].waktu_check_up.tanggal && A[i].waktu_check_up.tahun >= A[idx_max].waktu_check_up.tahun && A[i].waktu_check_up.bulan >= A[idx_max].waktu_check_up.bulan {
				idx_max = i
			}
		}
		A[pass], A[idx_max] = A[idx_max], A[pass]
	}
}

func descending_waktu(A tabPasien, n int) {
	sort_tahun_ascend(&A, n)
	sort_bulan_ascend(&A, n)
	sort_tanggal_ascend(&A, n)
	for i := 0; i < n; i++ {
		display_pasien(A, i)
	}
}

func sort_tahun_ascend(A *tabPasien, n int) {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu_check_up.tahun < A[idx_min].waktu_check_up.tahun {
				idx_min = i
			}
		}
		A[pass], A[idx_min] = A[idx_min], A[pass]
	}
}

func sort_bulan_ascend(A *tabPasien, n int) {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu_check_up.bulan < A[idx_min].waktu_check_up.bulan && A[i].waktu_check_up.tahun <= A[idx_min].waktu_check_up.tahun {
				idx_min = i
			}
		}
		A[pass], A[idx_min] = A[idx_min], A[pass]
	}
}

func sort_tanggal_ascend(A *tabPasien, n int) {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu_check_up.tanggal < A[idx_min].waktu_check_up.tanggal && A[i].waktu_check_up.tahun <= A[idx_min].waktu_check_up.tahun && A[i].waktu_check_up.bulan <= A[idx_min].waktu_check_up.bulan {
				idx_min = i
			}
		}
		A[pass], A[idx_min] = A[idx_min], A[pass]
	}
}

func ascending_paket(A tabPasien, n int) {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].jenis_check_up.nominal > A[idx_max].jenis_check_up.nominal {
				idx_max = i
			}
		}
		A[pass], A[idx_max] = A[idx_max], A[pass]
	}
	for i := 0; i < n; i++ {
		display_pasien(A, i)
	}
}

func descending_paket(A tabPasien, n int) {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].jenis_check_up.nominal < A[idx_min].jenis_check_up.nominal {
				idx_min = i
			}
		}
		A[pass], A[idx_min] = A[idx_min], A[pass]
	}
	for i := 0; i < n; i++ {
		display_pasien(A, i)
	}
}
func display_pasien(A tabPasien, idx int) {
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		fmt.Println("===================================================")
		fmt.Println("Nama Pasien          : ", A[idx].nama)
		fmt.Println("ID Pasien            : ", A[idx].id)
		fmt.Println("Rekap Pasien         : ", A[idx].hasil_check_up)
		fmt.Printf("Waktu Check Up Pasien: %d/%d/%d \n", A[idx].waktu_check_up.tahun, A[idx].waktu_check_up.bulan, A[idx].waktu_check_up.tanggal)
		fmt.Println("Jenis Layanan        : ", A[idx].jenis_check_up.tingkat)
	}
}

func cek_id_pasien(A tabPasien, n, x int) bool {
	for i := 0; i < n; i++ {
		if A[i].id == x {
			return false
		}
	}
	return true
}

func cek_waktu_pasien(tahun, bulan, tanggal int) bool {
	if bulan < 1 || bulan > 12 {
		return false
	}
	if tanggal < 1 || tanggal > 31 {
		return false
	}
	if (bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11) && tanggal > 30 {
		return false
	}
	if bulan == 2 {
		if tahun%4 == 0 {
			if tanggal > 29 {
				return false
			}
		} else if tanggal > 28 {
			return false
		}
	}
	return true
}

func bulan_hari(x, y int) int {
	var hasil int
	if x == 1 {
		hasil = 0
	} else if x == 2 {
		hasil = 31
	} else if x == 3 {
		if cek_tahun_kabisat(y) {
			hasil = 60
		}
		hasil = 59
	} else if x == 4 {
		if cek_tahun_kabisat(y) {
			hasil = 91
		}
		hasil = 90
	} else if x == 5 {
		if cek_tahun_kabisat(y) {
			hasil = 121
		}
		hasil = 120
	} else if x == 6 {
		if cek_tahun_kabisat(y) {
			hasil = 152
		}
		hasil = 151
	} else if x == 7 {
		if cek_tahun_kabisat(y) {
			hasil = 182
		}
		hasil = 181
	} else if x == 8 {
		if cek_tahun_kabisat(y) {
			hasil = 213
		}
		hasil = 212
	} else if x == 9 {
		if cek_tahun_kabisat(y) {
			hasil = 244
		}
		hasil = 243
	} else if x == 10 {
		if cek_tahun_kabisat(y) {
			hasil = 274
		}
		hasil = 273
	} else if x == 11 {
		if cek_tahun_kabisat(y) {
			hasil = 305
		}
		hasil = 304
	} else if x == 12 {
		if cek_tahun_kabisat(y) {
			hasil = 335
		}
		hasil = 334
	}
	return hasil
}
func cek_tahun_kabisat(x int) bool {
	var valid bool = false
	if x%400 == 0 {
		valid = true
	} else if x%100 == 0 {
		valid = false
	} else if x%4 == 0 {
		valid = true
	} else {
		valid = false
	}
	return valid
}
