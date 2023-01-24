package web

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//go:embed static
var FrontendFS embed.FS

func StartWebserver(port int) {
	StartTimer() //move to other location

	host := fmt.Sprintf("0.0.0.0:%v", port)

	//rerender frontend
	go func() {
		for {
			UpdateGui()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/ws", BuildWebsocket())
	r.HandleFunc("/frame", FrameHandler)
	r.PathPrefix("/").Handler(getStaticFilesHandler(FrontendFS))
	log.Fatal(http.ListenAndServe(host, r))
}

var cur time.Time

func StartTimer() {
	duration, _ := time.ParseDuration("-1s")
	reset_duration, _ := time.ParseDuration("+45m")

	go func() {
		for {
			if cur.IsZero() {
				cur = cur.Add(reset_duration)
			}
			cur = cur.Add(duration)
			time.Sleep(time.Second * 1)
		}
	}()

}

func FrameHandler(w http.ResponseWriter, r *http.Request) {
	if cur.IsZero() {
		fmt.Fprintf(w, "END")
	} else {
		timestr := fmt.Sprintf("{\"time\":\"%02d:%02d\"}", cur.Minute(), cur.Second())
		fmt.Fprintf(w, "%s\n", timestr)
	}
}

func getStaticFilesHandler(fefiles embed.FS) http.Handler {
	matches, _ := fs.Glob(fefiles, "static")
	if len(matches) != 1 {
		panic("unable to find frontend build files in FrontendFS")
	}
	feRoot, _ := fs.Sub(fefiles, matches[0])
	buildHandler := http.FileServer(http.FS(feRoot))
	return buildHandler
}
