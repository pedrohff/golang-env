package main

import (
	"fmt"
	"github.com/pedrohff/golang-env/pkg"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	_ = http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("fileName")
	dir, err := pkg.FormatFilePathForWorkingDir(fileName)
	if err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"`, err.Error())))
		return
	}
	w.WriteHeader(200)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"path": "%s"`, dir)))
	return
}
