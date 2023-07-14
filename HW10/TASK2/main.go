package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TranslationRequest struct {
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
	Text           string `json:"text"`
}

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
	myRouter.HandleFunc("/translate", GetTranslateResponse).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetTranslateResponse(w http.ResponseWriter, r *http.Request) {
	var translationRequest TranslationRequest

	err := json.NewDecoder(r.Body).Decode(&translationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	translation := GetGoogleTranslateResponse(translationRequest.SourceLanguage, translationRequest.TargetLanguage, translationRequest.Text)

	// Encode the translation response as JSON
	jsonResponse, err := json.Marshal(translation.Data.Translations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetGoogleTranslateResponse(sourceLang string, targetLang string, payloadText string) TranslationResponse {
	urlOfUser := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	requestBody := struct {
		Q      string `json:"q"`
		Format string `json:"format"`
		Target string `json:"target"`
		Source string `json:"source"`
	}{
		Q:      payloadText,
		Format: "html",
		Target: targetLang,
		Source: sourceLang,
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", urlOfUser, ioutil.NopCloser(bytes.NewReader(requestBytes)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("X-RapidAPI-Key"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("X-RapidAPI-Host"))

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
