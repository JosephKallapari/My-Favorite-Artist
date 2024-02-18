package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetArtistByCountry() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		countryName := vars["countryName"]
		url := fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=geo.gettopartists&country=%s&api_key=8b57c9978bbbd8af6ef70e284419aa74&format=json", countryName)
		response, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var dat map[string]interface{}
		if err := json.Unmarshal(data, &dat); err != nil {
			panic(err)
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusFound)
		w.Write(data)
	}
}

func GetSimilarTracks() http.HandlerFunc {
	//Example : track:believe , artist:cher
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		trackName := vars["trackName"]
		artistName := vars["artistName"]
		url := fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=track.getsimilar&artist=%s&track=%s&api_key=8b57c9978bbbd8af6ef70e284419aa74&format=json", artistName, trackName)
		response, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var dat map[string]interface{}
		if err := json.Unmarshal(data, &dat); err != nil {
			panic(err)
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusFound)
		w.Write(data)
	}
}
