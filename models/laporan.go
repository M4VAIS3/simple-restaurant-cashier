package models

// BarangLaris digunakan untuk laporan barang terlaris
type BarangLaris struct {
	KodeBarang     string
	NamaBarang     string
	TotalTerjual   int
	TotalPendapatan int
}
