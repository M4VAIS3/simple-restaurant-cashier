# Simple Restaurant Cashier

> A web-based restaurant cashier application built with **Go (net/http)**. No external frameworks, no external database, and no third-party Go dependencies beyond bundled UI libraries.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Application Architecture](#application-architecture)
- [Data Models](#data-models)
- [Routing & Endpoints](#routing--endpoints)
- [Algorithms Used](#algorithms-used)
- [Getting Started](#getting-started)
- [Production Build](#production-build)
- [Data Storage](#data-storage)
- [Sample JSON Data](#sample-json-data)

---

## Overview

**Simple Restaurant Cashier** is a lightweight point-of-sale application for small restaurants. It runs entirely server-side using Go's standard library (`net/http`) and renders pages via Go HTML Templates. All data is stored locally in a JSON file — no MySQL, PostgreSQL, or any other database engine is required.

This project also demonstrates the manual implementation of **sorting and searching algorithms** (without using `sort.Slice` from the standard library) as a learning exercise.

---

## Features

| Module | Description |
|---|---|
| **Dashboard** | Summary of total items, total transactions, daily revenue, and low-stock count |
| **Item Management** | Add, edit, and delete items; filter by category; sort by code, price, or stock |
| **Transactions** | Record multi-item sales transactions with automatic change calculation |
| **Live Search** | Real-time item search via a JSON API endpoint |
| **Reports** | Daily, weekly, and monthly reports; 7-day sales chart; best-selling items; low-stock alerts |
| **Low Stock Alert** | Automatic flag for items with stock of 5 units or fewer |

---

## Technology Stack

| Category | Technology |
|---|---|
| **Backend** | Go 1.23+ (standard library: `net/http`, `encoding/json`, `html/template`) |
| **Frontend** | HTML5, CSS3 (Vanilla CSS), JavaScript |
| **Charts** | [Chart.js](https://www.chartjs.org/) (bundled locally — no CDN required) |
| **Icons** | [Lucide Icons](https://lucide.dev/) (bundled locally) |
| **Storage** | Local JSON file (`data/data.json`) |

> **Note:** There is no `go.sum` file because this project has **zero external Go dependencies**.

---

## Project Structure

```
simple-restaurant-cashier/
├── main.go                     # Entry point & routing configuration
├── go.mod                      # Go module definition (simple-restaurant-cashier)
│
├── controllers/                # HTTP handlers — business logic per page
│   ├── dashboard.go            # Dashboard handler (/)
│   ├── barang.go               # Item CRUD handler (/barang/*)
│   ├── transaksi.go            # Transaction handler (/transaksi/*)
│   └── laporan.go              # Daily report handler (/laporan)
│
├── models/                     # Data struct definitions & storage
│   ├── barang.go               # Barang struct & global variable DataBarang
│   ├── transaksi.go            # Transaksi, ItemTransaksi structs & DataTransaksi
│   ├── laporan.go              # BarangLaris struct (for best-seller reports)
│   └── storage.go              # LoadData() & SaveData() to/from data.json
│
├── utils/                      # Helper functions
│   ├── search.go               # Search algorithms (Sequential & Binary Search)
│   ├── sort.go                 # Sorting algorithms (Selection, Insertion, Bubble Sort)
│   └── template.go             # Template FuncMap (formatRupiah, inc, shortDate)
│
├── views/                      # HTML templates (Go html/template)
│   ├── index.html              # Dashboard page
│   ├── barang.html             # Item list page
│   ├── tambah_barang.html      # Add item form
│   ├── edit_barang.html        # Edit item form
│   ├── transaksi.html          # Transaction list page
│   ├── tambah_transaksi.html   # New transaction form
│   ├── detail_transaksi.html   # Transaction receipt / detail view
│   └── laporan.html            # Reports & charts page
│
├── static/                     # Static assets served to the browser
│   ├── css/
│   │   └── styles.css          # Main application stylesheet
│   └── js/
│       ├── chart.umd.min.js    # Chart.js library (bundled locally)
│       ├── lucide.min.js       # Lucide Icons library (bundled locally)
│       └── sidebar.js          # Sidebar navigation toggle logic
│
└── data/
    └── data.json               # Data storage file (created automatically)
```

---

## Application Architecture

```
Browser (HTTP Request)
        │
        ▼
   main.go (Router — net/http)
        │
        ▼
  controllers/
  ├── dashboard.go  →  views/index.html
  ├── barang.go     →  views/barang.html
  │                    views/tambah_barang.html
  │                    views/edit_barang.html
  ├── transaksi.go  →  views/transaksi.html
  │                    views/tambah_transaksi.html
  │                    views/detail_transaksi.html
  └── laporan.go    →  views/laporan.html
        │
        ▼
     models/                          utils/
  ├── DataBarang   []Barang       ├── CariBarangSequential()  (Sequential Search)
  ├── DataTransaksi []Transaksi   ├── CariBarangBinary()      (Binary Search)
  └── storage.go                  ├── CariBarangByNama()      (case-insensitive filter)
       ├── LoadData()              ├── FilterTransaksiByTime() (filter by time)
       └── SaveData()             ├── UrutkanKodeBarang()     (Selection Sort)
                │                 ├── UrutkanHargaBarang()    (Insertion Sort)
                ▼                 ├── UrutkanStokBarang()     (Insertion Sort)
          data/data.json          └── BarangTerlaris()        (Bubble Sort + tally)
```

---

## Data Models

### `models.Barang`

```go
type Barang struct {
    Kode       string  // Unique item code, e.g. "NSG001"
    Nama       string  // Item name
    Kategori   string  // Category, e.g. "Makanan" (Food), "Minuman" (Drink)
    Harga      int     // Unit price (in Indonesian Rupiah)
    Stok       int     // Available stock quantity
    StokKritis bool    // true if Stok <= 5
}
```

### `models.ItemTransaksi`

```go
type ItemTransaksi struct {
    KodeBarang string  // Item code at time of purchase
    NamaBarang string  // Item name at time of purchase
    Harga      int     // Unit price at time of purchase
    Jumlah     int     // Number of units purchased
    Subtotal   int     // Harga * Jumlah
}
```

### `models.Transaksi`

```go
type Transaksi struct {
    ID        string           // Format: "TRX-YYYYMMDD-NNN"
    Waktu     time.Time        // Transaction timestamp
    Items     []ItemTransaksi  // List of purchased items
    Total     int              // Total price of all items
    Bayar     int              // Amount tendered by the customer
    Kembalian int              // Change returned (Bayar - Total)
}
```

### `models.BarangLaris`

```go
type BarangLaris struct {
    KodeBarang      string  // Item code
    NamaBarang      string  // Item name
    TotalTerjual    int     // Cumulative units sold
    TotalPendapatan int     // Cumulative revenue from this item
}
```

---

## Routing & Endpoints

| Method | Path | Handler | Description |
|---|---|---|---|
| `GET` | `/` | `Dashboard` | Main dashboard & summary page |
| `GET` | `/barang` | `TampilkanBarang` | Item list (supports `?sort=`, `?q=`, `?kategori=`) |
| `GET/POST` | `/barang/tambah` | `TambahBarang` | Add new item form & action |
| `GET/POST` | `/barang/edit` | `EditBarang` | Edit item form & action (`?kode=`) |
| `GET` | `/barang/hapus` | `HapusBarang` | Delete item by code (`?kode=`) |
| `GET` | `/barang/cari` | `CariBarangJSON` | **JSON API** — live item search (`?q=`) |
| `GET` | `/transaksi` | `TampilkanTransaksi` | Transaction list (supports `?filter=`) |
| `GET/POST` | `/transaksi/tambah` | `TambahTransaksi` | New transaction form & action |
| `GET` | `/transaksi/detail` | `DetailTransaksi` | Transaction receipt / detail (`?id=`) |
| `GET` | `/laporan` | `LaporanHarian` | Full reports page |
| `GET` | `/static/` | `FileServer` | Serves static files (CSS, JS) |

### Query Parameters for `/barang`

| Parameter | Value | Description |
|---|---|---|
| `sort` | `kode-asc`, `kode-desc` | Sort by item code |
| `sort` | `harga-asc`, `harga-desc` | Sort by price |
| `sort` | `stok-asc`, `stok-desc` | Sort by stock level |
| `q` | `<keyword>` | Search items (name, code, or category) |
| `kategori` | `<category name>` | Filter by category |

### Query Parameters for `/transaksi`

| Parameter | Format | Example |
|---|---|---|
| `filter` | `DD-MM-YYYY` | `?filter=21-07-2026` |
| `filter` | `MM-YYYY` | `?filter=07-2026` |
| `filter` | `YYYY` | `?filter=2026` |
| `filter` | `HH:MM` | `?filter=20:54` |

---

## Algorithms Used

This project implements sorting and searching algorithms manually as a learning demonstration:

### Searching (`utils/search.go`)

| Function | Algorithm | Purpose |
|---|---|---|
| `CariBarangSequential(kode)` | **Sequential Search** | Finds an item by exact code match; used in all edit and delete operations |
| `CariBarangBinary(kode)` | **Binary Search** | Alternative code lookup (requires data to be sorted in ascending order) |
| `CariBarangByNama(query)` | Linear filter | Case-insensitive search across item name, code, and category |
| `FilterTransaksiByTime(query)` | Linear filter | Filters transactions by time precision (minute, hour, day, month, or year) |

### Sorting (`utils/sort.go`)

| Function | Algorithm | Complexity |
|---|---|---|
| `UrutkanKodeBarang(asc)` | **Selection Sort** | O(n²) |
| `UrutkanHargaBarang(asc)` | **Insertion Sort** | O(n²) worst-case, O(n) best-case |
| `UrutkanStokBarang(asc)` | **Insertion Sort** | O(n²) worst-case, O(n) best-case |
| `BarangTerlaris(topN)` | **Bubble Sort** (descending) | O(n²) — applied after building the tally map |

---

## Getting Started

### Prerequisites

- **Go 1.23** or later must be installed
- Run `go version` to verify

```bash
go version
# go version go1.23.x ...
```

### Steps

**1. Open the project directory**

```bash
cd "simple-restaurant-cashier"
```

**2. Start the server**

```bash
go run main.go
```

**3. Open in your browser**

```
http://localhost:8080
```

Expected terminal output:

```
Data berhasil dimuat: 0 barang, 0 transaksi
Server berjalan di http://localhost:8080
```

> **Note:** If `data/data.json` does not yet exist, the application will start with an empty dataset. The JSON file is created automatically the first time any data is saved.

---

## Production Build

Compile to a standalone executable binary:

```bash
# Windows
go build -o kasir-restoran.exe main.go

# Linux / macOS
go build -o kasir-restoran main.go
```

Run the binary:

```bash
# Windows
.\kasir-restoran.exe

# Linux / macOS
./kasir-restoran
```

> The binary does not require a Go runtime to run. Simply ensure the `views/`, `static/`, and `data/` directories are in the same folder as the binary.

---

## Data Storage

All data is stored in a single JSON file: `data/data.json`.

- **On server start** → `models.LoadData()` reads the JSON file into memory (global variables `DataBarang` and `DataTransaksi`)
- **On any change** → `models.SaveData()` rewrites the entire dataset to the JSON file
- **If the file does not exist** → the application starts with an empty dataset (no error is thrown)

Data is held **entirely in memory (RAM)** whilst the server is running. Changes are written to disc after every add, edit, delete, or new transaction operation.

---

## Sample JSON Data

The following is an example of the `data/data.json` format:

```json
{
  "barang": [
    {
      "Kode": "NSG001",
      "Nama": "Nasi Goreng Biasa",
      "Kategori": "Makanan",
      "Harga": 15000,
      "Stok": 50,
      "StokKritis": false
    },
    {
      "Kode": "EST001",
      "Nama": "Es Teh Manis",
      "Kategori": "Minuman",
      "Harga": 5000,
      "Stok": 3,
      "StokKritis": true
    }
  ],
  "transaksi": [
    {
      "ID": "TRX-20260721-001",
      "Waktu": "2026-07-21T20:54:26.591+07:00",
      "Items": [
        {
          "KodeBarang": "NSG001",
          "NamaBarang": "Nasi Goreng Biasa",
          "Harga": 15000,
          "Jumlah": 2,
          "Subtotal": 30000
        },
        {
          "KodeBarang": "EST001",
          "NamaBarang": "Es Teh Manis",
          "Harga": 5000,
          "Jumlah": 2,
          "Subtotal": 10000
        }
      ],
      "Total": 40000,
      "Bayar": 50000,
      "Kembalian": 10000
    }
  ]
}
```

---

## Template Helpers

Custom functions available in all HTML templates (`utils/template.go`):

| Function | Description | Example |
|---|---|---|
| `formatRupiah` | Formats an integer as Indonesian Rupiah with dot separators | `15000` → `15.000` |
| `inc` | Increments an integer by one (used for row numbering in tables) | `inc 0` → `1` |
| `shortDate` | Trims whitespace from a date string | `" 21 July 2026 "` → `"21 July 2026"` |

---

## Limitations

- **No authentication** — anyone with access to the URL can use the application
- **In-memory data** — if the server crashes before `SaveData()` completes, the most recent changes may be lost
- **Single-user** — there is no concurrent write safety mechanism (not suitable for multiple cashiers operating simultaneously)
- **No automatic backups** — all data exists in a single JSON file

---

## Licence

This project was created for personal portfolio purposes. However, you are free to use it as long as you credit me as the original author.

---