package models

import "time"

type ItemTransaksi struct {
	KodeBarang string
	NamaBarang string
	Harga      int
	Jumlah     int
	Subtotal   int
}

type Transaksi struct {
	ID     string
	Waktu  time.Time
	Items  []ItemTransaksi
	Total  int
	Bayar  int
	Kembalian int
}

var DataTransaksi []Transaksi
