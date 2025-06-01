package main

import "fmt"

// Konstanta untuk ukuran array statis
const MAX_PESANAN = 100 //Maksimal 100 pesanan.
const MAX_ITEMS_PER_PESANAN = 10 //Maksimal 10 item makanan per pesanan.
const MAX_MENU = 10 //Maksimal 10 makanan tersedia di daftar menu.

// Struktur data
type MenuMakanan struct {
	Nomor int
	Nama  string
	Harga int
}

type ItemMakanan struct {
	Nama        string
	Jumlah      int
	HargaSatuan int
}

type Pesanan struct {
	ID          int
	Items       [MAX_ITEMS_PER_PESANAN]ItemMakanan
	JumlahItems int // Untuk melacak jumlah item aktual dalam pesanan
}

// Variabel global
var daftarPesanan [MAX_PESANAN]Pesanan
var jumlahPesananSaatIni int // Untuk melacak jumlah pesanan aktual

var menu = [MAX_MENU]MenuMakanan{
	{1, "Nasi Goreng", 15000},
	{2, "Mie Ayam", 12000},
	{3, "Bakso", 13000},
	{4, "Soto Ayam", 14000},
	{5, "Es Teh", 5000},
	{6, "Es Jeruk", 7000},
	// Tambahkan menu lain jika perlu, hingga MAX_MENU
}
var jumlahMenuAktual int = 6 

// Fungsi untuk mencetak teks dengan bingkai
func cetakDenganBingkai(teks []string) {
	maxPanjang := 0
	for i := 0; i < len(teks); i++ {
		baris := teks[i]
		if len(baris) > maxPanjang {
			maxPanjang = len(baris)
		}
	}

	fmt.Print("┌")
	for i := 0; i < maxPanjang+2; i++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	for i := 0; i < len(teks); i++ {
		var baris string = teks[i]
		fmt.Printf("│ %-*s │\n", maxPanjang, baris)
	}

	fmt.Print("└")
	for i := 0; i < maxPanjang+2; i++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

// Fungsi untuk membaca input integer dari pengguna
func bacaInt() int {
	var input int
	var err error
	var n int
	var valid bool = false

	for !valid {
		fmt.Print("Masukkan angka: ")
		n, err = fmt.Scanln(&input)

		if err != nil || n != 1 {
			fmt.Println("Input tidak valid, silakan coba lagi.")
			// Membersihkan buffer input jika terjadi kesalahan
			var buang string
			fmt.Scanln(&buang) // Membaca sisa input yang salah
		} else {
			valid = true
		}
	}
	return input
}

// Fungsi utama program
func main() {
	for {
		menuUtama := []string{
			"Created by Alya dan Azizah",
			"========== MENU ==========",
			"1. Input ID Pesanan",
			"2. Pilih Makanan",
			"3. Cetak Struk Belanja",
			"4. Urutkan Harga Termahal",
			"5. Hapus Data Pesanan",
			"6. Cari Nilai Ekstrim", 
			"7. Keluar",           
		}
		cetakDenganBingkai(menuUtama)
		fmt.Print("Pilih menu (1-7): ") 
		pilihan := bacaInt()
		fmt.Println()

		switch pilihan {
		case 1:
			inputIDPesanan()
		case 2:
			pilihMakanan()
		case 3:
			cetakStruk()
		case 4:
			urutkanHargaTermahal()
		case 5:
			hapusPesanan()
		case 6: // Sekarang case 6 memanggil findExtremeOrderValues
			findExtremeOrderValues()
		case 7: // Sekarang case 7 memanggil Keluar
			cetakDenganBingkai([]string{"Terima kasih! Program selesai."})
			return
		default:
			cetakDenganBingkai([]string{"Pilihan tidak valid."})
		}
	}
}

// Fungsi untuk menginput ID Pesanan baru
func inputIDPesanan() {
	if jumlahPesananSaatIni >= MAX_PESANAN {
		cetakDenganBingkai([]string{"Kapasitas pesanan penuh."})
		return
	}

	fmt.Print("Masukkan ID Pesanan: ")
	id := bacaInt()

	idSudahAda := false
	for i := 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == id {
			idSudahAda = true
			break
		}
	}

	if idSudahAda {
		cetakDenganBingkai([]string{"ID Pesanan sudah ada. Silakan gunakan ID lain."})
	} else {
		daftarPesanan[jumlahPesananSaatIni] = Pesanan{ID: id, JumlahItems: 0}
		jumlahPesananSaatIni++
		cetakDenganBingkai([]string{"ID disimpan. Sekarang pilih makanan di menu 2."})
	}
}

// Fungsi untuk memilih makanan dan menambahkannya ke pesanan
func pilihMakanan() {
	if jumlahPesananSaatIni == 0 {
		cetakDenganBingkai([]string{"Harap masukkan ID Pesanan terlebih dahulu (menu 1)."})
		return
	}

	pesananTargetIndex := -1
	fmt.Print("Masukkan ID Pesanan yang akan ditambahkan item: ")
	idTarget := bacaInt()

	foundTargetPesanan := false
	for i := 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == idTarget {
			pesananTargetIndex = i
			foundTargetPesanan = true
			break
		}
	}

	if !foundTargetPesanan {
		cetakDenganBingkai([]string{fmt.Sprintf("ID Pesanan %d tidak ditemukan.", idTarget)})
		return
	}

	pesananSaatIni := &daftarPesanan[pesananTargetIndex]

	if pesananSaatIni.JumlahItems >= MAX_ITEMS_PER_PESANAN {
		cetakDenganBingkai([]string{"Jumlah item maksimum untuk pesanan ini telah tercapai."})
		return
	}

	teksMenu := make([]string, 1+jumlahMenuAktual)
	idxTeksMenu := 0

	teksMenu[idxTeksMenu] = "Pilih Makanan:"
	idxTeksMenu++

	for k := 0; k < jumlahMenuAktual; k++ {
		item := menu[k]
		if item.Nomor != 0 {
			teksMenu[idxTeksMenu] = fmt.Sprintf("%d. %s - Rp%d", item.Nomor, item.Nama, item.Harga)
			idxTeksMenu++
		}
	}
	cetakDenganBingkai(teksMenu[:idxTeksMenu])

	fmt.Print("Masukkan nomor makanan: ")
	pilihanMenu := bacaInt()

	var itemDipilih *MenuMakanan
	foundItemMenu := false
	for j := 0; j < jumlahMenuAktual; j++ {
		if menu[j].Nomor == pilihanMenu {
			itemDipilih = &menu[j]
			foundItemMenu = true
			break
		}
	}

	if !foundItemMenu || itemDipilih == nil || itemDipilih.Nomor == 0 {
		cetakDenganBingkai([]string{"Menu tidak tersedia."})
		return
	}

	fmt.Print("Masukkan jumlah: ")
	jumlahItem := bacaInt()
	if jumlahItem <= 0 {
		cetakDenganBingkai([]string{"Jumlah item tidak valid."})
		return
	}

	pesananSaatIni.Items[pesananSaatIni.JumlahItems] = ItemMakanan{
		Nama:        itemDipilih.Nama,
		Jumlah:      jumlahItem,
		HargaSatuan: itemDipilih.Harga,
	}
	pesananSaatIni.JumlahItems++

	cetakDenganBingkai([]string{"Makanan berhasil ditambahkan ke pesanan."})
}

// Fungsi untuk mencetak struk belanja semua pesanan
func cetakStruk() {
	if jumlahPesananSaatIni == 0 {
		cetakDenganBingkai([]string{"Belum ada pesanan."})
		return
	}

	maxLinesPerOrder := 1 + MAX_ITEMS_PER_PESANAN + 1 + 1
	estimatedMaxLines := 1 + jumlahPesananSaatIni*maxLinesPerOrder
	struk := make([]string, estimatedMaxLines)
	strukIdx := 0

	struk[strukIdx] = "Struk Belanja:"
	strukIdx++

	for i := 0; i < jumlahPesananSaatIni; i++ {
		p := daftarPesanan[i]
		struk[strukIdx] = fmt.Sprintf("ID Pesanan: %d", p.ID)
		strukIdx++
		totalPesanan := 0
		if p.JumlahItems == 0 {
			struk[strukIdx] = "  (Belum ada item makanan)"
			strukIdx++
		} else {
			for j := 0; j < p.JumlahItems; j++ {
				item := p.Items[j]
				subtotal := item.Jumlah * item.HargaSatuan
				struk[strukIdx] = fmt.Sprintf("  - %s x%d @ Rp%d = Rp%d", item.Nama, item.Jumlah, item.HargaSatuan, subtotal)
				strukIdx++
				totalPesanan += subtotal
			}
		}
		// Calculate total using the recursive function
		totalPesananRecursive := totalHargaRecursive(p, 0)
		struk[strukIdx] = fmt.Sprintf("  Total untuk ID %d: Rp%d (Rekursif: Rp%d)", p.ID, totalPesanan, totalPesananRecursive)
		strukIdx++
		struk[strukIdx] = ""
		strukIdx++
	}
	cetakDenganBingkai(struk[:strukIdx])
}

// Fungsi untuk menghitung total harga item dalam satu pesanan (menggunakan array statis)
func hitungTotalHargaPesanan(items [MAX_ITEMS_PER_PESANAN]ItemMakanan, jumlahItemAktual int) int {
	total := 0
	for i := 0; i < jumlahItemAktual; i++ {
		total += items[i].Jumlah * items[i].HargaSatuan
	}
	return total
}

// Fungsi totalHargaRecursive: Menghitung total harga pesanan secara rekursif
// Parameter:
//   pesanan: Struktur Pesanan yang akan dihitung total harganya.
//   currentIndex: Indeks item saat ini yang sedang diproses dalam array Items.
//
// Cara kerja:
//   - Basis Kasus: Jika currentIndex sudah mencapai atau melebihi JumlahItems,
//     berarti tidak ada item lagi yang perlu dijumlahkan, maka kembalikan 0.
//   - Langkah Rekursif: Jumlahkan harga subtotal dari item saat ini
//     dengan hasil pemanggilan rekursif untuk item berikutnya (currentIndex + 1).
func totalHargaRecursive(pesanan Pesanan, currentIndex int) int {
	// Basis Kasus: Jika tidak ada item lagi untuk dijumlahkan
	if currentIndex >= pesanan.JumlahItems {
		return 0
	}

	// Langkah Rekursif: Hitung subtotal item saat ini dan tambahkan dengan total sisa item
	currentItem := pesanan.Items[currentIndex]
	subtotal := currentItem.Jumlah * currentItem.HargaSatuan
	return subtotal + totalHargaRecursive(pesanan, currentIndex+1)
}

// Fungsi untuk mengurutkan pesanan berdasarkan total harga termahal (Selection Sort)
func urutkanHargaTermahal() {
	if jumlahPesananSaatIni == 0 {
		cetakDenganBingkai([]string{"Belum ada pesanan untuk diurutkan."})
		return
	}

	dataUrut := make([]Pesanan, jumlahPesananSaatIni)
	copy(dataUrut, daftarPesanan[:jumlahPesananSaatIni])

	for i := 0; i < jumlahPesananSaatIni-1; i++ {
		maxIdx := i
		totalMaxIdx := hitungTotalHargaPesanan(dataUrut[maxIdx].Items, dataUrut[maxIdx].JumlahItems)
		for j := i + 1; j < jumlahPesananSaatIni; j++ {
			totalJ := hitungTotalHargaPesanan(dataUrut[j].Items, dataUrut[j].JumlahItems)
			if totalJ > totalMaxIdx {
				maxIdx = j
				totalMaxIdx = totalJ
			}
		}
		dataUrut[i], dataUrut[maxIdx] = dataUrut[maxIdx], dataUrut[i]
	}

	hasilUrut := make([]string, 1+jumlahPesananSaatIni)
	hasilUrutIdx := 0

	hasilUrut[hasilUrutIdx] = "ID Pesanan urut dari total harga tertinggi:"
	hasilUrutIdx++

	for i := 0; i < jumlahPesananSaatIni; i++ {
		p := dataUrut[i]
		total := hitungTotalHargaPesanan(p.Items, p.JumlahItems)
		hasilUrut[hasilUrutIdx] = fmt.Sprintf("ID: %d - Total: Rp%d", p.ID, total)
		hasilUrutIdx++
	}
	cetakDenganBingkai(hasilUrut[:hasilUrutIdx])
}

// Fungsi untuk menghapus data pesanan berdasarkan ID
func hapusPesanan() {
	if jumlahPesananSaatIni == 0 {
		cetakDenganBingkai([]string{"Belum ada pesanan untuk dihapus."})
		return
	}
	fmt.Print("Masukkan ID pesanan yang ingin dihapus: ")
	idHapus := bacaInt()

	indexDitemukan := -1
	foundPesananHapus := false
	for i := 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == idHapus {
			indexDitemukan = i
			foundPesananHapus = true
			break
		}
	}

	if foundPesananHapus {
		for k := indexDitemukan; k < jumlahPesananSaatIni-1; k++ {
			daftarPesanan[k] = daftarPesanan[k+1]
		}
		daftarPesanan[jumlahPesananSaatIni-1] = Pesanan{}
		jumlahPesananSaatIni--
		cetakDenganBingkai([]string{"Pesanan berhasil dihapus."})
	} else {
		cetakDenganBingkai([]string{fmt.Sprintf("ID Pesanan %d tidak ditemukan.", idHapus)})
	}
}

// Fungsi untuk mencari nilai ekstrem (maksimum dan minimum) dari total harga pesanan
func findExtremeOrderValues() {
	if jumlahPesananSaatIni == 0 {
		cetakDenganBingkai([]string{"Belum ada pesanan untuk mencari nilai ekstrem."})
		return
	}

	// Inisialisasi dengan nilai pesanan pertama
	minTotal := totalHargaRecursive(daftarPesanan[0], 0)
	maxTotal := totalHargaRecursive(daftarPesanan[0], 0)
	minID := daftarPesanan[0].ID
	maxID := daftarPesanan[0].ID

	for i := 1; i < jumlahPesananSaatIni; i++ {
		currentTotal := totalHargaRecursive(daftarPesanan[i], 0)

		if currentTotal < minTotal {
			minTotal = currentTotal
			minID = daftarPesanan[i].ID
		}
		if currentTotal > maxTotal {
			maxTotal = currentTotal
			maxID = daftarPesanan[i].ID
		}
	}

	output := []string{
		"Analisis Nilai Ekstrim:",
		fmt.Sprintf("Pesanan Termahal: ID %d dengan total Rp%d", maxID, maxTotal),
		fmt.Sprintf("Pesanan Termurah: ID %d dengan total Rp%d", minID, minTotal),
		"", // Just for spacing in output
		"Demonstrasi Rekursi (Detail untuk setiap pesanan):",
	}

	for i := 0; i < jumlahPesananSaatIni; i++ {
		p := daftarPesanan[i]
		totalRecursive := totalHargaRecursive(p, 0)
		output = append(output, fmt.Sprintf("ID %d Total Rekursif: Rp%d", p.ID, totalRecursive))
	}

	cetakDenganBingkai(output)
}

