package main

import (
	"encoding/json"
	"flag"
	"github.com/tkstorm/go-long-polling/server/lyric"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// server addr
var addr = ":9102"

// request count
var cnt struct {
	num int
	mux sync.Mutex
}

func reqCount() {
	cnt.mux.Lock()
	cnt.num++
	cnt.mux.Unlock()
	log.Println("request count:", cnt.num)
}

// lyric
var lyc *lyric.Lyric
var filename, musicName *string

func init() {
	filename = flag.String("filename", "", "music lyric file to open")
	musicName = flag.String("musicName", "", "music name")
}

func main() {
	flag.Parse()

	lyc = lyric.New(*musicName, *filename)
	if err := lyc.Parse(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listen on %s...", addr)

	http.HandleFunc("/music/lyric", lyricSentence)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func lyricSentence(w http.ResponseWriter, r *http.Request) {
	reqCount()
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	scaleId, _ := strconv.Atoi(r.FormValue("ScaleId"))

	// get lyric sentence
	prev := lyc.NextSentence(scaleId - 1)
	next := lyc.NextSentence(scaleId)

	d, _ := json.Marshal(next)
	_, _ = w.Write(d)

	time.Sleep(lyric.SubStrTime(prev.Scale, next.Scale))
}
