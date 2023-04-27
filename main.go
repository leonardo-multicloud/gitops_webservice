package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Webservice API - ArgoCD K8s"))
	})
	http.ListenAndServe(":8081", nil)
}
