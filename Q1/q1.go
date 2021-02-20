package main

import (
  "encoding/json"
  "fmt"
  "os"
  "time"
  "bufio"
  "strings"
  tm "github.com/buger/goterm"
  humanize "github.com/dustin/go-humanize"
  "strconv"
)

var namaResto, namaKasir string
var garis = "------------------------------"
var listMakananan = []byte(`[
    {"Nama": "Nasi", "Harga": 5000},
    {"Nama": "Lauk", "Harga": 10000},
    {"Nama": "Minum", "Harga": 5000}
]`)

var listOrder = []byte(`[]`)

type menuMakanan struct {
  Nama      string      `json:"nama"`
  Harga     int         `json:"harga"`
}

type menuOrder struct {
  Nama      string      `json:"nama"`
  Harga     int         `json:"harga"`
}

var dataOrder []menuOrder

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func dataRestoran() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Printf("Masukkan Nama Restoran : ")
  scanner.Scan()
  namaResto = scanner.Text()
  clearScreen()
  menuUtama()
}

func dataKasir() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Printf("Masukkan Nama Kasir : ")
  scanner.Scan()
  namaKasir = scanner.Text()
  clearScreen()
  menuUtama()
}

func cetakStruk() {

  clearScreen()

  var tanggal = time.Now()

  fmt.Println( center( namaResto ) )
  fmt.Println()
  fmt.Printf("%-9s %v %18s\n", "Tanggal :", "", tanggal.Format("2006/01/02 15:04:05"))
  fmt.Println()
  fmt.Printf("%-12s %v %15s\n", "Nama Kasir :", "", namaKasir)
  fmt.Println()

  fmt.Println( center( "==============================" ) )
  fmt.Println()

  total := 0

  for index := range dataOrder{

    str2 := dataOrder[index].Nama
    str3 := FormatRupiah(float64(dataOrder[index].Harga))
    rpt  := strings.Repeat(".", 30-(len(str2)+len(str3)))
    total += dataOrder[index].Harga

    fmt.Printf("%v%v%v\n", str2, rpt, str3)
    fmt.Println()

    if len(dataOrder) == index+1 {

      rpt  := FormatRupiah(float64(total))

      fmt.Printf("%v%v%v\n", "Total", strings.Repeat(".", 30-(5+len(rpt))), rpt)

    }

  }

  listOrder = []byte(`[]`)

  var pilih int
  fmt.Println()

  fmt.Printf("'1' Input ulang pesanan & '2' Ke menu awal : ")
  fmt.Scanf("%d", &pilih)

  clearScreen()

  switch pilih {
    case 1:
      dataPembelian()
    case 2:
      menuUtama()
    default:
      fmt.Println( center( "Angka tidak valid" ) )
      time.Sleep(1 * time.Second)
  }

}

func dataPembelian() {

  clearScreen()
  fmt.Println( center("Menu Makanan") )
  fmt.Println( center( garis ) )

  var data []menuMakanan

  var err = json.Unmarshal([]byte(listMakananan), &data)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  if err := json.Unmarshal([]byte(listOrder), &dataOrder); err != nil {
    fmt.Println(err)
    return
  }

  var num = 1

  for index := range data{

    str1 := num+index
    str2 := data[index].Nama
    str3 := FormatRupiah(float64(data[index].Harga))
    rpt  := strings.Repeat(".", 30-(len(strconv.Itoa(str1))+len(str2)+len(str3)+2))

    fmt.Printf("%v. %v%v%v\n", str1, str2, rpt, str3)

  }

  fmt.Printf("%v. %v\n", len(data)+1, "Cetak Struk")

  fmt.Println( center( garis ) )

  if len( dataOrder) > 0 {

    fmt.Println( center("Menu Makanan Di Input") )
    fmt.Println()

    for index := range dataOrder{

      str2 := dataOrder[index].Nama
      str3 := FormatRupiah(float64(dataOrder[index].Harga))
      rpt  := strings.Repeat(".", 30-(len(str2)+len(str3)))

      fmt.Printf("%v%v%v\n", str2, rpt, str3)

    }

    fmt.Println( center( garis ) )

  }

  var pilih int
  fmt.Printf("Pilih Menu Makanan: ")
  fmt.Scanf("%d", &pilih)
  fmt.Println()

  if pilih < len(data)+1 {

  	dataOrder = append(dataOrder, menuOrder{Nama: data[pilih-1].Nama, Harga: data[pilih-1].Harga})

  	result, err := json.Marshal(dataOrder)
  	if err != nil {
  		fmt.Println(err)
  	}

    listOrder = []byte(string(result))

    dataPembelian()

  } else {
    cetakStruk()
  }

}

func pilihanMenu() {
  var pilih int
  fmt.Printf("Pilih menu : ")
  fmt.Scanf("%d", &pilih)
  fmt.Println()
  switch pilih {
    case 1:
      dataRestoran()
    case 2:
      dataKasir()
    case 3:
      if len(namaResto) < 1  {
        fmt.Println( "Silahkan Masukkan Nama Resto Terlebih Dahulu." )
        dataRestoran()
      } else if len(namaKasir) < 1 {
        fmt.Println( "Silahkan Masukkan Nama Kasir Terlebih Dahulu." )
        dataKasir()
      } else {
        dataPembelian()
      }
    case 4:
      fmt.Println( center( "Keluar dalam 1 detik..." ) )
      time.Sleep(1 * time.Second)
      os.Exit(0)
    default:
      fmt.Println( center( "Angka menu tidak valid" ) )
      time.Sleep(1 * time.Second)
      clearScreen()
      menuUtama()
  }
}

func center(s string) string {
  w := 30
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w + len(s))/2, s))
}

func clearScreen() {
  tm.Clear()
  tm.MoveCursor(1, 1)
  tm.Flush()
}

func menuUtama(){
  fmt.Println( center("Q1. Aplikasi Restoran") )
  fmt.Println( center( garis ) )

  if len(namaResto) > 0  {
    fmt.Println( "1. Nama Resto :", namaResto )
  } else {
    fmt.Println("1. Masukan Nama Resto")
  }

  if len(namaKasir) > 0  {
    fmt.Println( "2. Nama Kasir :", namaKasir )
  } else {
    fmt.Println("2. Masukan Nama Kasir")
  }

  fmt.Println("3. Masukan Data Pembelian")
  fmt.Println("4. Keluar")
  fmt.Println( center( garis ) )
  pilihanMenu()
}

func main() {
  clearScreen()
  menuUtama()
}
