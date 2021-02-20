package main

import (
  "encoding/json"
  "fmt"
  "os"
  "time"
  "strings"
  tm "github.com/buger/goterm"
)

var user = []byte(`[
  {
    "id": 323,
    "username": "rinood30",
    "profile": {
      "full_name": "Shabrina Fauzan",
      "birthday": "1988-10-30",
      "phones": [
        "08133473821",
        "082539163912"
      ]
    },
    "articles:": [
      {
        "id": 3,
        "title": "Tips Berbagi Makanan",
        "published_at": "2019-01-03T16:00:00"
      },
      {
        "id": 7,
        "title": "Cara Membakar Ikan",
        "published_at": "2019-01-07T14:00:00"
      }
    ]
  },
  {
    "id": 201,
    "username": "norisa",
    "profile": {
      "full_name": "Noor Annisa",
      "birthday": "1986-08-14",
      "phones": []
    },
    "articles:": [
      {
        "id": 82,
        "title": "Cara Membuat Kue Kering",
        "published_at": "2019-10-08T11:00:00"
      },
      {
        "id": 91,
        "title": "Cara Membuat Brownies",
        "published_at": "2019-11-11T13:00:00"
      },
      {
        "id": 31,
        "title": "Cara Membuat Brownies",
        "published_at": "2019-11-11T13:00:00"
      }
    ]
  },
  {
    "id": 42,
    "username": "karina",
    "profile": {
      "full_name": "Karina Triandini",
      "birthday": "1986-04-14",
      "phones": [
        "06133929341"
      ]
    },
    "articles:": []
  },
  {
    "id": 201,
    "username": "icha",
    "profile": {
      "full_name": "Annisa Rachmawaty",
      "birthday": "1987-12-30",
      "phones": []
    },
    "articles:": [
      {
        "id": 39,
        "title": "Tips Berbelanja Bulan Tua",
        "published_at": "2019-04-06T07:00:00"
      },
      {
        "id": 43,
        "title": "Cara Memilih Permainan di Steam",
        "published_at": "2019-06-11T05:00:00"
      },
      {
        "id": 58,
        "title": "Cara Membuat Brownies",
        "published_at": "2019-09-12T04:00:00"
      }
    ]
  }
]`)
var garis = "------------------------------"

type UserList struct {
	ID       int    `json:"id"`
	Username string `json:"username"`

	Profile  struct {
		FullName string   `json:"full_name"`
		Birthday string   `json:"birthday"`
		Phones   []string `json:"phones"`
	} `json:"profile"`

	Articles []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		PublishedAt string `json:"published_at"`
	} `json:"articles:"`

}

var detailUser []UserList

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

  if err := json.Unmarshal([]byte(user), &detailUser); err != nil {
    fmt.Println(err)
    return
  }

  switch s {
    case "noPhone":

      fmt.Println(`Find users who don't have any phone numbers.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

      	if len(user.Profile.Phones) <= 0 {
          fmt.Printf("Username : %v, Fullname : %v\n", user.Username, user.Profile.FullName)
        }

      }

    case "haveArticles":

      fmt.Println(`Find users who have articles.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

      	if len(user.Articles) > 0 {
          fmt.Printf("Username : %v, Fullname : %v\n", user.Username, user.Profile.FullName)
        }

      }

    case "haveAnnis":

      fmt.Println(`Find users who have "annis" on their name.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

        fullname := strings.ToLower(user.Profile.FullName)
    		if strings.Contains(fullname, "annis") {
          fmt.Printf("Username : %v, Fullname : %v\n", user.Username, user.Profile.FullName)
    		}

      }

    case "haveArticles2020":

      fmt.Println(`Find users who have articles on the year 2020.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

        have := 0

        for _, article := range user.Articles{

          pub, err := time.Parse("2006-01-02T15:04:05.000Z", article.PublishedAt+".000Z")
    			if err != nil {
    				fmt.Println(err.Error())
    				return
    			}
    			if pub.Year() == 2020 {
            have = 1
    			}

        }

        if have > 0 {
          fmt.Printf( "Username : %v, Fullname : %v\n", user.Username, user.Profile.FullName )
        }

      }

    case "bornIn1986":

      fmt.Println(`Find users who are born in 1986.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

        birth, err := time.Parse("2006-01-02", user.Profile.Birthday)
        if err != nil {
          fmt.Println(err.Error())
          return
        }
        if birth.Year() == 1986 {
          fmt.Printf("Username : %v, Fullname : %v\n", user.Username, user.Profile.FullName)
        }


      }

    case "containTips":

      fmt.Println(`Find articles that contain "tips" on the title.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

        for _, article := range user.Articles{

          containTips := strings.ToLower(article.Title)
      		if strings.Contains(containTips, "tips") {
            fmt.Printf("Title : %v, Published At : %v\n", article.Title, article.PublishedAt)
      		}

        }

      }

    case "publishedBeforeAugust2019":

      fmt.Println(`Find articles published before August 2019.`)
      fmt.Println( center( garis ) )

      for _, user  := range detailUser{

        for _, article := range user.Articles{

          befAug2019, err := time.Parse("2006-01-02", "2019-08-01")
      		if err != nil {
      			fmt.Println(err.Error())
      			return
      		}
      		pub, err := time.Parse("2006-01-02T15:04:05.000Z", article.PublishedAt+".000Z")
      		if err != nil {
      			fmt.Println(err.Error())
      			return
      		}
      		if pub.Before(befAug2019) {
            fmt.Printf("Title : %v, Published At : %v\n", article.Title, article.PublishedAt)
      		}

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
      check("noPhone")
    case 2:
      check("haveArticles")
    case 3:
      check("haveAnnis")
    case 4:
      check("haveArticles2020")
    case 5:
      check("bornIn1986")
    case 6:
      check("containTips")
    case 7:
      check("publishedBeforeAugust2019")
    default:
      fmt.Println( center( "Angka menu tidak valid" ) )
      time.Sleep(1 * time.Second)
      clearScreen()
      menuUtama()
  }
}

func menuUtama(){
  fmt.Println( center("Q2. JSON Manipulation") )
  fmt.Println( center( garis ) )

  fmt.Println(`1. Find users who don't have any phone numbers.`)
  fmt.Println(`2. Find users who have articles.`)
  fmt.Println(`3. Find users who have "annis" on their name.`)
  fmt.Println(`4. Find users who have articles on the year 2020.`)
  fmt.Println(`5. Find users who are born in 1986.`)
  fmt.Println(`6. Find articles that contain "tips" on the title.`)
  fmt.Println(`7. Find articles published before August 2019.`)

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
