package models

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataDir = "data"
const dataFile = "data/data.json"

type DataStore struct {
	Barang    []Barang    `json:"barang"`
	Transaksi []Transaksi `json:"transaksi"`
}

// SaveData menyimpan semua data ke file JSON
func SaveData() error {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("gagal membuat direktori data: %w", err)
	}

	store := DataStore{
		Barang:    DataBarang,
		Transaksi: DataTransaksi,
	}

	bytes, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("gagal marshal data: %w", err)
	}

	if err := os.WriteFile(dataFile, bytes, 0644); err != nil {
		return fmt.Errorf("gagal menulis file: %w", err)
	}

	return nil
}

// LoadData memuat data dari file JSON saat server start
func LoadData() error {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		// File belum ada, mulai dengan data kosong
		DataBarang = []Barang{}
		DataTransaksi = []Transaksi{}
		return nil
	}

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		return fmt.Errorf("gagal membaca file: %w", err)
	}

	var store DataStore
	if err := json.Unmarshal(bytes, &store); err != nil {
		return fmt.Errorf("gagal unmarshal data: %w", err)
	}

	if store.Barang == nil {
		store.Barang = []Barang{}
	}
	if store.Transaksi == nil {
		store.Transaksi = []Transaksi{}
	}

	DataBarang = store.Barang
	DataTransaksi = store.Transaksi
	CekStokKritis()

	return nil
}
