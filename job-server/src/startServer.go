package main

import (
	"log"
	"runtime"
	"os"
	"net/http"
	"controller"
	"encoding/json"
	"dto"
)


var serverVersion string

func main() {
	log.Println("StarsMurasaki MPS System. Compiled version: ",serverVersion)
	log.Println("You are running on the following OS: ",runtime.GOOS)

	readPropertiesFile()

	//Main page
	http.HandleFunc("/status/",controller.StatController)

	if dto.ServerProperties.SSLCertChain != "" && dto.ServerProperties.SSLCertKey != "" {
		log.Println("Server started with SSL encryption")
		go http.ListenAndServe(":80", http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
			target := "https://" + req.Host + req.URL.Path
			if len(req.URL.RawQuery) > 0 {
				target += "?" + req.URL.RawQuery
			}
			log.Printf("Redirect to: %s", target)
			http.Redirect(w, req, target,
			http.StatusTemporaryRedirect)
		}))

		http.ListenAndServeTLS(dto.ServerProperties.ServerStartPort, dto.ServerProperties.SSLCertChain, dto.ServerProperties.SSLCertKey, nil)
	} else {
		log.Println("Server started.")
		http.ListenAndServe(dto.ServerProperties.ServerStartPort, nil)
	}

}

func readPropertiesFile(){

	dto.ServerProperties.ServerStartPort = ":80"
	dto.ServerProperties.ServerStartUrl = "http://127.0.0.1"

	log.Println("Check for dir: "+"./conf")
	if _, err := os.Stat("./conf"); os.IsNotExist(err) {
		os.Mkdir("./conf", 0755)
	}

	file, err := os.Open("./conf/server.properties")
	if err != nil {
		log.Fatal("Opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&dto.ServerProperties); err != nil {
		log.Fatal("Parsing config file", err.Error())
	}

	log.Print("Server config loaded: ",dto.ServerProperties)

	return
}
