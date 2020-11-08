package api

import (
	"encoding/json"
	"net/http"

	"github.com/PichuChen/go-bbs/bbs"
)

func GetFavBoards(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.Form.Get("userID")
	if userID == "" {
		http.Error(w, "userID not specified", http.StatusInternalServerError)
	}

	fav, err := bbs.FavLoad(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(fav)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
