package main

import (
	"github.com/planning_poker/go/common/constant"
	"github.com/planning_poker/go/room"
	"log"
	"net/http"
	"os"
)

func main() {
	// パスを設定する
	http.HandleFunc(constant.WarmupPath, initialize)
	http.HandleFunc("/", handler)

	http.HandleFunc(constant.RoomRootPath, room.Dispatch)

	// ポートを設定する。
	port := os.Getenv("PORT")
	if port == "" {
		port = "8801"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// initialize
func initialize(w http.ResponseWriter, r *http.Request) {
	log.Printf("warm up")
}

// handler : デフォルト動作
func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)
}
