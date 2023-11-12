package handler

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)

	go func() {
		// do the process here
		// simulate a long time request by putting 10 second sleep
		time.Sleep(time.Second * 5)
		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occured.", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}
