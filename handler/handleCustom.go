package handler

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func HandleCustom(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)

	go func() {
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
			log.Println("request done")
		}
	}()

	w.Write([]byte("custom handler"))
	time.Sleep(time.Second * 5)
	done <- true
}
