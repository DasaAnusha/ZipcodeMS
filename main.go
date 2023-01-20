package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")

	vars := mux.Vars(r)

	ZipCode := vars["zip_code"]
	CountryCode := vars["country_code"]

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://api.zippopotam.us/" + CountryCode + "/" + ZipCode)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(body))
	fmt.Fprintf(w, string(body))

}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running ")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{country_code}/{zip_code}", ZipCodeHandler)

	r.HandleFunc("/", rootHandler)
	log.Println("Web server has started!!!")
	http.ListenAndServe(":80", r)
}
