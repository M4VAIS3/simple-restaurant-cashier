package models

type Barang struct {
	Kode      string
	Nama      string
	Kategori  string
	Harga     int
	Stok      int
	StokKritis bool // true jika stok <= 5
}

var DataBarang []Barang

// CekStokKritis memperbarui flag StokKritis untuk semua barang
func CekStokKritis() {
	for i := range DataBarang {
		DataBarang[i].StokKritis = DataBarang[i].Stok <= 5
	}
}
