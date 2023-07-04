package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TranslationResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func main() {
	handleRequest()
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/translate/{sourceLanguage}-{targetLanguage}/{text}", GetTranslateResponse).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetTranslateResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceLanguage := vars["sourceLanguage"]
	targetLanguage := vars["targetLanguage"]
	text := vars["text"]

	translation := GetGoogleTranslateResponse(sourceLanguage, targetLanguage, text)

	// Encode the translation response as JSON
	jsonResponse, err := json.Marshal(translation.Data.Translations)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func GetGoogleTranslateResponse(sourceLang string, targetLang string, payloadText string) TranslationResponse {
	urlOfUser := "https://google-translate1.p.rapidapi.com/language/translate/v2"
	encodedPayloadText := html.EscapeString(payloadText)

	payload := fmt.Sprintf("q=%s&format=html&target=%s&source=%s", encodedPayloadText, targetLang, sourceLang)

	req, err := http.NewRequest("POST", urlOfUser, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "93d0bc0245msh36f35f4080c8ee0p173769jsn21bbdc0408a9")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var translationResponse TranslationResponse
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(string(body))

	return translationResponse
}
