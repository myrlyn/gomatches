// WebServer
package main

import (
	"log"
	"net/http"
	"path/filepath"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func configureWebServer() {
	herepath, _ := filepath.Abs(".")
	staticpath := filepath.Join(herepath, "ui")
	staticpath = filepath.Join(staticpath, "SlutMatching")
	staticpath = filepath.Join(staticpath, "dist")
	staticpath = filepath.Join(staticpath, "slut-matching")
	staticpath = filepath.Join(staticpath, "browser")

	r := mux.NewRouter()
	sttc := StaticHandler{StaticPath: staticpath}
	r.PathPrefix("/").PathPrefix(sttc.StaticPath)
	if cfg.HasSection("server") {
		port := "8080"
		if cfg.Section("server").HasKey("port") {
			port = cfg.Section("server").Key("port").String()
		}
		log.Fatal(http.ListenAndServe(":"+port, nil))
	} else {
		log.Println("NO WEB SERVER INFO DEFINED, RUNNING UNSECURED")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
