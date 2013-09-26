package gogo

import (
	"fmt"
	"github.com/adeven/goenv"
	"log"
	"net/http"
)

func Gogo(goenv *goenv.Goenv) {
	log.Println("entering Gogo func")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(
			w,
			"this is env %s and branch %s",
			goenv.Get("env", ""),
			goenv.Get("branch", ""),
		)
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println(err)
	}
}
