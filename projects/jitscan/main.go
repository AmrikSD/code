package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/validate", validate)

	log.Println("🚀 Starting basic admission webhook on :8443")
	log.Println(http.ListenAndServeTLS(":8443", "/etc/certs/tls.crt", "/etc/certs/tls.key", mux))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func validate(w http.ResponseWriter, r *http.Request) {
	// Read the request
	if r.Body == nil {
		log.Println("No body in request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body provided"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}

	// Parse admission review
	var review admissionv1.AdmissionReview
	if err := json.Unmarshal(body, &review); err != nil {
		log.Printf("Failed to unmarshal admission review: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid admission review JSON"))
		return
	}

	// Check if request is valid
	if review.Request == nil {
		log.Println("No request in admission review")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No request in admission review"))
		return
	}

	// Log that we received a request
	log.Printf("🎯 Pod deployment detected! Resource: %s/%s\n", review.Request.Namespace, review.Request.Name)

	// Always allow - create response
	response := &admissionv1.AdmissionResponse{
		UID:     review.Request.UID,
		Allowed: true,
		Result:  &metav1.Status{Message: "All good! 👍"},
	}

	// Send response back
	reviewResponse := admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Response: response,
	}

	respBytes, _ := json.Marshal(reviewResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
