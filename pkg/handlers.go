package pkg

import "net/http"

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world Yerkanat!"))
}
