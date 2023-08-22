package ErrorHandling

import (
	"net/http"
)

func ReturnEndpointError(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 : Sorry, we don't support that endpoint."))
}

func ReturnCreateIntentError(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(CREATE_INTENT_ERROR))
}

func ReturnGeneralError(w http.ResponseWriter, status int, response string){
	w.WriteHeader(status)
	returnString := response
	w.Write([]byte(returnString))
}
