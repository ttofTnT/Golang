package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"./versions"
	"log"
	"net/http"
	"os"
)

/**
 * @Description:	起点，处理各个请求
 */
func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
