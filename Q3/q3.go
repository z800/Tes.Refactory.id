package main

import (
  "encoding/json"
  "fmt"
  "os"
  "time"
  "strings"
  tm "github.com/buger/goterm"
)

var inventaris = []byte(`[
  {
    "inventory_id": 9382,
    "name": "Brown Chair",
    "type": "furniture",
    "tags": [
      "chair",
      "furniture",
      "brown"
    ],
    "purchased_at": 1579190471,
    "placement": {
      "room_id": 3,
      "name": "Meeting Room"
    }
  },
  {
    "inventory_id": 9380,
    "name": "Big Desk",
    "type": "furniture",
    "tags": [
      "desk",
      "furniture",
      "brown"
    ],
    "purchased_at": 1579190642,
    "placement": {
      "room_id": 3,
      "name": "Meeting Room"
    }
  },
  {
    "inventory_id": 2932,
    "name": "LG Monitor 50 inch",
    "type": "electronic",
    "tags": [
      "monitor"
    ],
    "purchased_at": 1579017842,
    "placement": {
      "room_id": 3,
      "name": "Lavender"
    }
  },
  {
    "inventory_id": 232,
    "name": "Sharp Pendingin Ruangan 2PK",
    "type": "electronic",
    "tags": [
      "ac"
    ],
    "purchased_at": 1578931442,
    "placement": {
      "room_id": 5,
      "name": "Dhanapala"
    }
  },
  {
    "inventory_id": 9382,
    "name": "Alat Makan",
    "type": "tableware",
    "tags": [
      "spoon",
      "fork",
      "tableware"
    ],
    "purchased_at": 1578672242,
    "placement": {
      "room_id": 10,
      "name": "Rajawali"
    }
  }
]`)
var garis = "------------------------------"

type Invent struct {
	InventoryID int      `json:"inventory_id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	PurchasedAt int      `json:"purchased_at"`
	Placement   struct {
		RoomID int    `json:"room_id"`
		Name   string `json:"name"`
	} `json:"placement"`
}

var detailInvent []Invent

func back() {

  var pilih int
  fmt.Printf(`0 to Exit, 1 Back :`)
  fmt.Scanf("%d", &pilih)
  fmt.Println()
  switch pilih {
    case 0:
      fmt.Println( center( "Keluar dalam 1 detik..." ) )
      time.Sleep(1 * time.Second)
      os.Exit(0)
    case 1:
      main()
    default:
      fmt.Println( center( "Angka menu tidak valid, kembali ke menu utama" ) )
      time.Sleep(1 * time.Second)
      main()
  }

}

func check(s string) {

  clearScreen()

  if err := json.Unmarshal([]byte(inventaris), &detailInvent); err != nil {
    fmt.Println(err)
    return
  }

  switch s {
    case "FindItems":

      fmt.Println(`Find items in the Meeting Room.`)
      fmt.Println( center( garis ) )

      for _, ivt := range detailInvent{

      	if ivt.Placement.Name == "Meeting Room" {
          fmt.Printf("Name : %v, Type : %v\n", ivt.Name, ivt.Type)
        }

      }

    case "FindElectronic":

      fmt.Println(`Find all electronic devices.`)
      fmt.Println( center( garis ) )

      for _, ivt := range detailInvent{

      	if ivt.Type == "electronic" {
          fmt.Printf("Name : %v, Type : %v\n", ivt.Name, ivt.Type)
        }

      }

    case "FindFurniture":

      fmt.Println(`Find all the furniture.`)
      fmt.Println( center( garis ) )

      for _, ivt := range detailInvent{

      	if ivt.Type == "furniture" {
          fmt.Printf("Name : %v, Type : %v\n", ivt.Name, ivt.Type)
        }

      }

    case "FindPurchased":

      fmt.Println(`Find all items were purchased on 16 Januari 2020.`)
      fmt.Println( center( garis ) )

      for _, ivt := range detailInvent{

        pItem := time.Unix(int64(ivt.PurchasedAt), 0)

        dateParams, err := time.Parse("2006-01-02", "2020-01-16")
        if err != nil {
          fmt.Println(err.Error())
          return
        }

        if pItem.Day() == dateParams.Day() && pItem.Month() == dateParams.Month() && pItem.Year() == dateParams.Year() {
          fmt.Printf("Name : %v, Type : %v\n", ivt.Name, ivt.Type)
        }

      }

    case "FindBrown":

      fmt.Println(`Find all items with brown color.`)
      fmt.Println( center( garis ) )

      for _, ivt := range detailInvent{

        check := 0

        for _, tag := range ivt.Tags {
          if strings.Contains(tag, "brown") {
            check = 1
          }
    		}

        if check > 0 {
            fmt.Printf("Name : %v, Type : %v\n", ivt.Name, ivt.Type)
        }

      }

  }

  fmt.Println( center( garis ) )

  back()

}

func pilihanMenu() {
  var pilih int
  fmt.Printf("Pilih menu : ")
  fmt.Scanf("%d", &pilih)
  fmt.Println()
  switch pilih {
    case 1:
      check("FindItems")
    case 2:
      check("FindElectronic")
    case 3:
      check("FindFurniture")
    case 4:
      check("FindPurchased")
    case 5:
      check("FindBrown")
    default:
      fmt.Println( center( "Angka menu tidak valid" ) )
      time.Sleep(1 * time.Second)
      clearScreen()
      menuUtama()
  }
}

func menuUtama(){
  fmt.Println( center("Q3. JSON Manipulation") )
  fmt.Println( center( garis ) )

  fmt.Println(`1. Find items in the Meeting Room.`)
  fmt.Println(`2. Find all electronic devices.`)
  fmt.Println(`3. Find all the furniture.`)
  fmt.Println(`4. Find all items were purchased on 16 Januari 2020.`)
  fmt.Println(`5. Find all items with brown color.`)

  fmt.Println( center( garis ) )
  pilihanMenu()
}

func clearScreen() {
  tm.Clear()
  tm.MoveCursor(1, 1)
  tm.Flush()
}

func center(s string) string {
  w := 30
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w + len(s))/2, s))
}

func main() {
  clearScreen()
  menuUtama()
}
