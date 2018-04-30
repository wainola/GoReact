package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var configDb map[string]interface{}

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func GetApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api endpoint")
}

func init() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("dat es %s\n", configDb)
	if err := json.Unmarshal(config, &configDb); err != nil {
		panic(err)
	}
	fmt.Println(configDb)
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "./src")
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/api", GetApi)

	log.Println("Listening on port 4500")
	http.ListenAndServe(":4500", nil)
}
