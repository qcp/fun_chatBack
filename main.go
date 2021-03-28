package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	writer.Header().Set("Access-Control-Allow-Methods", "GET")

	if request.Method == "GET" {
		resp, err := http.Get("http://chat21.std-400.ist.mospolytech.ru/get.ajax.php")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatalln("Response not OK")
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(writer, string(bodyBytes))
	}
}

type Message struct {
	Name string
	Text string
}

func add(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")

	if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		var m Message
		err := decoder.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Get("http://chat21.std-400.ist.mospolytech.ru/add.ajax.php?name="+m.Name+"&text="+m.Text)

		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatalln("Response not OK")
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(writer, string(bodyBytes))
	}
}

func main() {
	fmt.Println("http server up!")
	http.HandleFunc("/get", get)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":0", nil)
}
