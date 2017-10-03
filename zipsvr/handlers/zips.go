package handlers

import "github.com/info344-a17/info344-in-class/zipsvr/models"
import "net/http"
import "strings"
import "encoding/json"

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
