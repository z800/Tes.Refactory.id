package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
  "log"
  "os"
)

type Message struct {
	message string `json:"message"`
}

type RequestCounter struct {
	Counter int    `json:"counter"`
	Header  string `json:"header"`
}

const (
	LOG_FORMAT = "[%s] Success: POST http://%s %s"
)

func LogToFile(dateTime time.Time, ip string, data RequestCounter) {
	f, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	timeFormattedString := dateTime.Format(time.RFC3339)
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)
	json, _ := json.Marshal(data)
	dataString := string(json)
	log.Printf(LOG_FORMAT, timeFormattedString, ip, dataString)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := RequestCounter{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Header = r.Header.Get("X-RANDOM")

		LogToFile(time.Now(), GetLocalIP(), data)
		w.WriteHeader(http.StatusCreated)

	} else {
		w.Header().Set("Content-Type", "application/json")
		message := Message{message: "Bad Request"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
	}

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{message: "Server is Up"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthCheck)
	fmt.Println("SERVER LISTEN AT PORT: 8090")
	http.ListenAndServe(":8090", nil)
}
