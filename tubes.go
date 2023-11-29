package main

import "fmt"

const NMAX = 100

type tabMHS struct {
	mhs  [NMAX]mahasiswa
	nMhs int
}

type mahasiswa struct {
	id        int
	nim, name string
	matkul    [NMAX]matakuliah
	nMatkul   int
	totalsks  int
}

type matakuliah struct {
	name                  string
	uts, uas, quiz, total int
	grade                 string
	sks                   int
}

func login(username, password string) bool {
	validUsername := "admin"
	validPassword := "123"

	if username == validUsername && password == validPassword {
		return true
	}

	return false
}

func menu(T *tabMHS) {
	var pilih int
	fmt.Println("MENU")
	fmt.Println("========================================================================")
	fmt.Println("||Silahkan pilih:                                                     ||")
	fmt.Println("||1. Edit Data Mahasiswa                                              ||")
	fmt.Println("||2. Edit  Matakuliah dan nilai                                       ||")
	fmt.Println("||3. Cari daftar mahasiswa dan matakuliah                             ||")
	fmt.Println("||4. Cari daftar mahasiswa terurut berdasarkan nilai dan jumlah sks   ||")
	fmt.Println("||5. Cari transkip nilai mahasiswa                                    ||")
	fmt.Println("||6. Keluar                                                           ||")
	fmt.Println("========================================================================")
	fmt.Print("Silahkan pilih :")
	fmt.Scanln(&pilih)
	if pilih == 6 {
		main()
	}
	if pilih == 1 {
		editData(T)
		menu(T)
	} else if pilih == 2 {
		editMatakuliahdanNilai(T)
		menu(T)
	} else if pilih == 3 {
		cariMahasiswa(T) 
		menu(T)
	} else if pilih == 4 {
		cariNilaiTerurutDanSks(T)
		menu(T)
	} else if pilih == 5 {
		transkip(T)
		menu(T)
	}

}

func editData(T *tabMHS) {
	var pilih int
	fmt.Println(" ")
	fmt.Println("EDIT DATA")
	fmt.Println("------------------------------")
	fmt.Println("1. Penambahan Data Mahasiswa")
	fmt.Println("2. Pengubahan Data Mahasiswa")
	fmt.Println("3. Penghapusan Data Mahasiswa")
	fmt.Println("4. Menu")
	fmt.Println("------------------------------")
	fmt.Print("Silahkan Pilih :")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		tambahDataMahasiswa(T)
		menu(T)
	} else if pilih == 2 {
		ubahDataMahasiswa(T)
		menu(T)
	} else if pilih == 3 {
		hapusDataMahasiswa(T)
		menu(T)
	} else if pilih == 4 {
		menu(T)
	}
}

func tambahDataMahasiswa(T *tabMHS) {
	var mhs mahasiswa
	fmt.Println(" ")
	fmt.Println("----------------------------------- ")
	fmt.Print("NIM: ")
	fmt.Scanln(&mhs.nim)
	fmt.Print("NAMA: ")
	fmt.Scanln(&mhs.name)
	mhs.id = T.nMhs + 1
	T.mhs[T.nMhs] = mhs
	T.nMhs++

	fmt.Println("Data mahasiswa berhasil ditambahkan")
	fmt.Println("----------------------------------- ")
	fmt.Println(" ")
}

func ubahDataMahasiswa(T *tabMHS) {
	var nim string
	fmt.Println(" ")
	fmt.Print("NIM Mahasiswa yang ingin diubah : ")
	fmt.Scanln(&nim)

	for i := 0; i < T.nMhs; i++ {
		if T.mhs[i].nim == nim {
			fmt.Println(" ")
			fmt.Println("----------------------------------- ")
			fmt.Print("Nama Baru: ")
			fmt.Scanln(&T.mhs[i].name)
			fmt.Println("Data Mahasiswa berhasil diubah")
			fmt.Println(" ")
			fmt.Println("----------------------------------- ")
			return
		}
	}

	fmt.Println("NIM Mahasiswa tidak ditemukan")
	fmt.Println(" ")
}

func hapusDataMahasiswa(T *tabMHS) {
	var nim string
	fmt.Println(" ")
	fmt.Print("NIM mahasiswa yang ingin dihapus : ")
	fmt.Scanln(&nim)

	for i := 0; i < T.nMhs; i++ {
		if T.mhs[i].nim == nim {
			for j := i; j < T.nMhs-1; j++ {
				T.mhs[j] = T.mhs[j+1]
			}
			T.nMhs--
			fmt.Println("Data Mahasiswa berhasil dihapus")
			return
		}
	}

	fmt.Println("NIM Mahasiswa tidak ditemukan")
	fmt.Println(" ")
}

func cariIndexmhs(T *tabMHS, nim string) int {
	for i := 0; i < T.nMhs; i++ {
		if T.mhs[i].nim == nim {
			return i
		}

	}
	return -1
}
func binarySearchMhs(T *tabMHS, nim string) int {
	kiri := 0
	kanan := T.nMhs - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if T.mhs[tengah].nim == nim {
			return tengah
		} else if T.mhs[tengah].nim < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
    
	return -1
}

func editMatakuliahdanNilai(T *tabMHS) {
	var pilih int
	fmt.Println(" ")
	fmt.Println("EDIT MATAKULIAH DAN NILAI")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. Penambahan Matakuliah dan Nilai Mahasiswa")
	fmt.Println("2. Pengubahan Matakuliah dan Nilai Mahasiswa")
	fmt.Println("3. Penghapusan Matakuliah dan Nilai Mahasiswa")
	fmt.Println("4. Menu")
	fmt.Println("----------------------------------------------")
	fmt.Print("Silahkan Pilih :")
	fmt.Scanln(&pilih)
	fmt.Println(" ")
	if pilih == 1 {
		tambahMatakuliahdanNilai(T)
		menu(T)
	} else if pilih == 2 {
		ubahMatakuliahdanNilai(T)
		menu(T)
	} else if pilih == 3 {
		hapusMatakuliahdanNilai(T)
		menu(T)
	} else if pilih == 4 {
		menu(T)
	}
}

func tambahMatakuliahdanNilai(T *tabMHS) {
	fmt.Print("Masukan NIM Mahasiswa:")
	var nim string
	var total int 
	fmt.Scanln(&nim)
	index := cariIndexmhs(T, nim)
	if index == -1 {
		fmt.Println("NIM Tidak Ditemukan")
		return
	}
	fmt.Println("-------------------------------------- ")
	fmt.Print("Silahkan masukan nama mata kuliah:")
	fmt.Scanln(&T.mhs[index].matkul[T.mhs[index].nMatkul].name)
	fmt.Println(" ")
	fmt.Print("Masukan Nilai UTS  :")
	fmt.Scanln(&T.mhs[index].matkul[T.mhs[index].nMatkul].uts)
	fmt.Println(" ")
	fmt.Print("Masukan Nilai UAS  :")
	fmt.Scanln(&T.mhs[index].matkul[T.mhs[index].nMatkul].uas)
	fmt.Println(" ")
	fmt.Print("Masukan Nilai QUIZ :")
	fmt.Scanln(&T.mhs[index].matkul[T.mhs[index].nMatkul].quiz)
	fmt.Println(" ")
	fmt.Print("Nilai TOTAL        :")
	T.mhs[index].matkul[T.mhs[index].nMatkul].total = (T.mhs[index].matkul[T.mhs[index].nMatkul].uts + T.mhs[index].matkul[T.mhs[index].nMatkul].uas + T.mhs[index].matkul[T.mhs[index].nMatkul].quiz)/3 
	fmt.Print(T.mhs[index].matkul[T.mhs[index].nMatkul].total)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Print("Nilai GRADE        :")
	total=T.mhs[index].matkul[T.mhs[index].nMatkul].total
	if total >= 90 {
		T.mhs[index].matkul[T.mhs[index].nMatkul].grade="A"
	}else if total >= 80 {
		T.mhs[index].matkul[T.mhs[index].nMatkul].grade="B"
    }else if total >= 70 {
		T.mhs[index].matkul[T.mhs[index].nMatkul].grade="C"
	}else if total >= 60 {
		T.mhs[index].matkul[T.mhs[index].nMatkul].grade="D"
	}else {
		T.mhs[index].matkul[T.mhs[index].nMatkul].grade="E"
	}
	fmt.Println(T.mhs[index].matkul[T.mhs[index].nMatkul].grade)
	fmt.Println(" ")
	fmt.Print("Masukan Juamlah sks:")
	fmt.Scanln(&T.mhs[index].matkul[T.mhs[index].nMatkul].sks)
	T.mhs[index].nMatkul++
	T.mhs[index].totalsks++
	fmt.Println("------------------------------------- ")
	fmt.Println(" ")
}

func ubahMatakuliahdanNilai(T *tabMHS) {
	var nim string
	fmt.Print("Masukan NIM Mahasiswa : ")
	fmt.Scanln(&nim)

	index := cariIndexmhs(T, nim)
	if index == -1 {
		fmt.Println("NIM mahasiswa tidak ditemukan")
		return
	}

	fmt.Print("Nama MataKuliah yang ingin diubah: ")
	var namaMatakuliah string
	fmt.Scanln(&namaMatakuliah)
	fmt.Println(" ")
	for i := 0; i < T.mhs[index].nMatkul; i++ {
		if T.mhs[index].matkul[i].name == namaMatakuliah {
			fmt.Println("------------------------------------------- ")
			fmt.Print("Masukkan nilai UTS baru: ")
			fmt.Scanln(&T.mhs[index].matkul[i].uts)
			fmt.Println(" ")
			fmt.Print("Masukkan nilai UAS baru: ")
			fmt.Scanln(&T.mhs[index].matkul[i].uas)
			fmt.Println(" ")
			fmt.Print("Masukkan nilai QUIZ baru: ")
			fmt.Scanln(&T.mhs[index].matkul[i].quiz)
			fmt.Println(" ")
			fmt.Print("Masukkan nilai TOTAL baru: ")
			fmt.Scanln(&T.mhs[index].matkul[i].total)
			fmt.Println(" ")
			fmt.Print("Masukkan nilai GRADE baru: ")
			fmt.Scanln(&T.mhs[index].matkul[i].grade)
			fmt.Println(" ")
			fmt.Print("Masukan Juamlah sks:")
			fmt.Scanln(&T.mhs[index].matkul[i].sks)
			fmt.Println(" ")
			fmt.Println("Data mata kuliah dan nilai berhasil diubah")
			fmt.Println("--------------------------------------------- ")
			fmt.Println(" ")
			return
		}
	}

	fmt.Println("Mata kuliah tidak ditemukan")
	fmt.Println(" ")
}

// func binarySearch(matkul []matakuliah, namaMatakuliah string) int {
// 	low := 0
// 	high := len(matkul) - 1

// 	for low <= high {
// 		mid := (low + high) / 2
// 		if matkul[mid].name == namaMatakuliah {
// 			return mid
// 		} else if matkul[mid].name < namaMatakuliah {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}

// 	return -1
// }

// func hapusMatakuliahdanNilai(T *tabMHS) {
// 	// Get input from user
// 	fmt.Print("Masukkan Nim mahasiswa: ")
// 	index := -1
// 	fmt.Scanln(&index)

// 	fmt.Print("Masukkan nama matakuliah yang ingin dihapus: ")
// 	namaMatakuliah := ""
// 	fmt.Scanln(&namaMatakuliah)

// 	// Find the mahasiswa in the tabMHS struct
// 	if index >= 0 && index < len(T.mhs) {
// 		// Find the matakuliah using binary search
// 		i := binarySearch(T.mhs[index].matkul[:], namaMatakuliah)
// 		if i != -1 {
// 			// Shift the remaining elements
// 			for j := i; j < T.mhs[index].nMatkul-1; j++ {
// 				T.mhs[index].matkul[j] = T.mhs[index].matkul[j+1]
// 			}
// 			T.mhs[index].nMatkul--
// 			fmt.Println("Data mata kuliah dan nilai berhasil dihapus")
// 			return
// 		}
// 		fmt.Println("Mata kuliah tidak ditemukan")
// 	} else {
// 		fmt.Println("Indeks mahasiswa tidak valid")
// 	}
// }

func hapusMatakuliahdanNilai(T *tabMHS) {
	var nim string
	fmt.Print("Masukan NIM Mahasiswa : ")
	fmt.Scanln(&nim)

	index := binarySearchMhs(T, nim)
	if index == -1 {
		fmt.Println("NIM Mahasiswa tidak ditemukan")
		return
	}

	fmt.Print("Nama MataKuliah yang ingin dihapus : ")
	var namaMatakuliah string
	fmt.Scanln(&namaMatakuliah)

	for i := 0; i < T.mhs[index].nMatkul; i++ {
		if T.mhs[index].matkul[i].name == namaMatakuliah {
			for j := i; j < T.mhs[index].nMatkul-1; j++ {
				T.mhs[index].matkul[j] = T.mhs[index].matkul[j+1]
			}
			T.mhs[index].nMatkul--
			fmt.Println("Data MataKuliah dan nilai berhasil dihapus")
			fmt.Println(" ")
			return
		}
	}

	fmt.Println("MataKuliah tidak ditemukan")
}

func cariMahasiswa(T *tabMHS) {
	var pilih int
	fmt.Println(" ")
	fmt.Println("CARI MAHASISWA")
	fmt.Println("-----------------------------")
	fmt.Println("1.Cari Bedasarkan NIM")
	fmt.Println("2.Cari Bedasarkan MataKuliah")
	fmt.Println("3.Menu")
	fmt.Println("-----------------------------")
	fmt.Print("Silahkan Pilih : ")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		cariMahasiswaBerdasarakanNIM(T)
	} else if pilih == 2 {
		cariMahasiswaBerdasarakaMatkul(T)
	} else if pilih == 3 {
		menu(T)
	}

}

func cariMahasiswaBerdasarakanNIM(T *tabMHS) {
	var id, i int
	var nim string
	var found = false
	fmt.Print("Masukan NIM Mahasiswa :")
	fmt.Scanln(&nim)
	for i < T.nMhs && !found {
		if T.mhs[i].nim == nim {
			id = i
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("NIM tidak di temmukan")
	} else {
		fmt.Println(" ")
		fmt.Println("----------------------------------- ")
		fmt.Println("")
		fmt.Println("Data Mahasiswa")
		fmt.Println("-------------------------")
		fmt.Println("NIM :", T.mhs[id].nim)
		fmt.Println("Nama Mahasiswa :", T.mhs[id].name)
		fmt.Println("-------------------------")
		fmt.Println("")
		for i := 0; i < T.mhs[id].nMatkul; i++ {
			fmt.Println("-------------------------")
			fmt.Println("Nama MataKuliah :", T.mhs[id].matkul[i].name)
			fmt.Println("SKS :", T.mhs[id].matkul[i].sks)
			fmt.Println("-------------------------")
			fmt.Println("")
			fmt.Println("----------------------------------- ")

			fmt.Println(" ")
		}
		fmt.Println(" ")
	}
}
func cariMahasiswaBerdasarakaMatkul(T *tabMHS) {
	var i int
	var matkul string
	var tabid [NMAX]int
	var totalId int
	var found = false
	fmt.Print("Masukan Matakuliah Mahasiswa yang ingin dicari: ")
	fmt.Scanln(&matkul)
	for i < T.nMhs {
		for j := 0; j < T.mhs[i].nMatkul; j++ {
			if T.mhs[i].matkul[j].name == matkul {
				tabid[totalId] = T.mhs[i].id
				totalId++
				found = true
			}
		}
		i++
	}
	if !found {
		fmt.Println("Matakuliah tidak ditemukan")
	} else {
		for i := 0; i < totalId; i++ {
			fmt.Println("-----------------------------------")
			fmt.Println("ID           :", tabid[i])
			fmt.Println("NIM          :", T.mhs[tabid[i]-1].nim)
			fmt.Println("Nama Mahasiswa:", T.mhs[tabid[i]-1].name)
			fmt.Println("-----------------------------------")
		}
	}
	fmt.Println()
}

// func cariMahasiswaBerdasarakaMatkul(T *tabMHS) {
// 	var i int
// 	var matkul string
// 	var tabid [NMAX]int
// 	var totalId int
// 	var found = false 
// 	fmt.Print("Masukan Matakuliah Mahasiswa yang ingin di cari :")
// 	fmt.Println(" ")
// 	fmt.Scanln(&matkul)
// 	for i < T.nMhs {
// 		for j := 0; j < T.mhs[i].nMatkul; j++ {
// 			if T.mhs[i].matkul[j].name == matkul {
// 				tabid[totalId] = T.mhs[i].id
// 				totalId++
// 			}
// 		}
// 		i++
// 	}
// 	if !found {
// 		fmt.Println("MataKulih tidak ditemukan")
// 	}else{
// 	for i := 0; i < totalId; i++ {
// 		fmt.Println(" ")
// 		fmt.Println("----------------------------------- ")
// 		fmt.Println("ID :", tabid[i])
// 		fmt.Println("NIM :", T.mhs[tabid[i]-1].nim)
// 		fmt.Println("Nama Mahasiswa :", T.mhs[tabid[i]-1].name)
// 		fmt.Println("----------------------------------- ")

// 	}

// 	fmt.Println(" ")
// }
// }

func cariNilaiTerurutDanSks(T *tabMHS) {
	var pilih int
	fmt.Println(" ")
	fmt.Println("CARI NILAI DAN SKS")
	fmt.Println("-----------------------")
	fmt.Println("1.Berdasarkan Nilai ")
	fmt.Println("2.Bedasarkan Sks")
	fmt.Println("3.Menu")
	fmt.Println("-----------------------")
	fmt.Print("Silahkan Pilih :")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		sortNilai(T)
	} else if pilih == 2 {
		sortSks(T)

	} else if pilih == 3 {
		menu(T)
	}

}

func sortNilai(T *tabMHS) {
	var tab tabMHS

	for i := 0; i < T.nMhs; i++ {
		tab.mhs[i].nim = T.mhs[i].nim
		tab.mhs[i].name = T.mhs[i].name
		tab.mhs[i].nMatkul = T.mhs[i].nMatkul
		for j := 0; j < T.mhs[i].nMatkul; j++ {
			tab.mhs[i].matkul[j].name = T.mhs[i].matkul[j].name
			tab.mhs[i].matkul[j].total = T.mhs[i].matkul[j].total
		}
		tab.nMhs++
	}
	for i := 0; i < tab.nMhs-1; i++ {
		maxIndex := i
		for j := i + 1; j < tab.nMhs; j++ {
			if tab.mhs[j].matkul[0].total > tab.mhs[maxIndex].matkul[0].total {
				maxIndex = j
			}
		}
		tab.mhs[i], tab.mhs[maxIndex] = tab.mhs[maxIndex], tab.mhs[i]
	}
	fmt.Println(" ")
	fmt.Println("Berikut adalah daftar nilai mahasiswa tertinggi :")
	for i := 0; i < tab.nMhs; i++ {
		fmt.Println("---------------------------------------------")
		fmt.Println("Data Mahasiswa")
		fmt.Println("-------------------------")
		fmt.Println("NIM :", tab.mhs[i].nim)
		fmt.Println("Nama Mahasiswa :", tab.mhs[i].name)
		fmt.Println("-------------------------")
		fmt.Println("Data Nilai")
		for j := 0; j < tab.mhs[i].nMatkul; j++ {
			fmt.Println("-------------------------")
			fmt.Println("Nama MataKuliah :", tab.mhs[i].matkul[j].name)
			fmt.Println("Total Nilai     :", tab.mhs[i].matkul[j].total)
			fmt.Println("-------------------------")
		}
		fmt.Println("---------------------------------------------")
		fmt.Println(" ")
	}
}


func sortSks(T *tabMHS) {
	var tab tabMHS

	for i := 0; i < T.nMhs; i++ {
		tab.mhs[i].nim = T.mhs[i].nim
		tab.mhs[i].name = T.mhs[i].name
		tab.mhs[i].nMatkul = T.mhs[i].nMatkul
		for j := 0; j < T.mhs[i].nMatkul; j++ {
			tab.mhs[i].matkul[j].name = T.mhs[i].matkul[j].name
			tab.mhs[i].matkul[j].sks = T.mhs[i].matkul[j].sks
		}
		tab.nMhs++
	}
	for i := 0; i < tab.nMhs-1; i++ {
		maxIndex := i
		for j := i + 1; j < tab.nMhs; j++ {
			if tab.mhs[j].matkul[0].total > tab.mhs[maxIndex].matkul[0].total {
				maxIndex = j
			}
		}
		tab.mhs[i], tab.mhs[maxIndex] = tab.mhs[maxIndex], tab.mhs[i]
	}
	fmt.Println(" ")
	fmt.Println("Berikut adalah daftar SKS mahasiswa :")
	for i := 0; i < tab.nMhs; i++ {
		fmt.Println("---------------------------------------------")
		fmt.Println("Data Mahasiswa")
		fmt.Println("------------------------")
		fmt.Println("NIM :", tab.mhs[i].nim)
		fmt.Println("Nama Mahasiswa :", tab.mhs[i].name)
		fmt.Println("------------------------")
		fmt.Println("Data Nilai")
		for j := 0; j < tab.mhs[i].nMatkul; j++ {
			fmt.Println("-------------------------")
			fmt.Println("Nama MataKuliah :", tab.mhs[i].matkul[j].name)
			fmt.Println("SKS :", tab.mhs[i].matkul[j].sks)
			fmt.Println("-------------------------")
			fmt.Println(" ")
		}
		fmt.Println("---------------------------------------------")
		fmt.Println(" ")
	}
}

func transkip(T *tabMHS) {
	var nim string
	fmt.Println(" ")
	fmt.Println("TRANSKIP MAHASISWA")
	fmt.Println("-----------------------------------")
	fmt.Print("Silahkan Masukan NIM :")
	fmt.Println(" ")
	fmt.Scanln(&nim)
	id := cariIndexmhs(T, nim)
	if id != -1 {
		fmt.Println("-------------------------")
		fmt.Println("NIM Mahasiswa :", T.mhs[id].nim)
		fmt.Println("Nama Mahasiswa :", T.mhs[id].name)
		fmt.Println("-------------------------")
		for i := 0; i < T.mhs[id].nMatkul; i++ {
			fmt.Println("----------------------------")
			fmt.Println("Nama MataKuliah : ", T.mhs[id].matkul[i].name)
			fmt.Println("Jumlah SKS      : ", T.mhs[id].matkul[i].sks)
			fmt.Println("Grade           : ", T.mhs[id].matkul[i].grade)
			fmt.Println("Total nilai     : ", T.mhs[id].matkul[i].total)
			fmt.Println("----------------------------")
			fmt.Println(" ")
			fmt.Println("-----------------------------------")
		}
	}
}

func main() {
	var username, password string

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if login(username, password) {
		fmt.Println("Login berhasil!")
		fmt.Println(" ")
		var mhs tabMHS
		menu(&mhs)
	} else {
		fmt.Println("Login gagal!")
	}
}
