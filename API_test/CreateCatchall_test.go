package API_test

import (
	. "CreateConversationBackend_NLP/ErrorHandling"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMalformedEndpointHandler(t *testing.T) {
	testingStruct := []struct {
		pathVar    string
		shouldPass bool
	}{
		{ "/", true},
		{ "/intent/frozen/", false},
		{ "/confusion/reigns", false},
	}
	fmt.Println("#####TestMalformedEndpointHandler#####")

	for _, testingVar := range testingStruct {
		path := testingVar.pathVar
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()

		// CATCH ERRORS
		router.NotFoundHandler = http.HandlerFunc(ReturnEndpointError)
		router.HandleFunc("/", helloGirls)

		router.ServeHTTP(rr, req)

		// some of these should pass, and some should not
		fmt.Println("path: ", path, "has a statusCode of: ", rr.Code, ", and a message of: ", rr.Body)
		if rr.Code == http.StatusOK  && !testingVar.shouldPass {
			t.Errorf("handler should have failed on routeVariable: got %v want %v", rr.Code, http.StatusBadRequest)
		}
		message, err := rr.Body.ReadString(200)
		if rr.Code == http.StatusOK && message != "Hello beautiful..."{
			t.Errorf("handler should have returned message 'Hello beautiful...': got %v", rr.Body)
		}
		if rr.Code != http.StatusOK && message != "400 : Sorry, we don't support that endpoint."{
			t.Errorf("handler should have returned message '400 : Sorry, we don't support that endpoint.': got %v", rr.Body)
		}
	}
}

func helloGirls(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello beautiful..."))
}