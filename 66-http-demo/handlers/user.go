package handlers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"io"
	v2 "math/rand/v2"
	"net/http"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("not implemented"))
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := new(models.User)
	err = json.Unmarshal(bytes, user) //
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Id = uint(v2.IntN(100000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()
	// here store in the file
	err = utils.SaveToFile("data.txt", user)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("user successfully saved in the file"))
	w.WriteHeader(201)
}
