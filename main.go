package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Webservice API - ArgoCD Kubernetes"))
	})
	http.ListenAndServe(":8081", nil)
}
