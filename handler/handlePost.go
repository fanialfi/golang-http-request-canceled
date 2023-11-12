package handler

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)

	go func() {
		// khusus untuk request dengan http method yang ada datanya / request body-nya (payload)
		// maka channel r.Context().Done() tidak akan menerima data hingga proses read pada body dilakukan
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(body)

		time.Sleep(time.Second * 5)
		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occured", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}
