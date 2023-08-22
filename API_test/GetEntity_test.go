package API_test

import (
	. "CreateConversationBackend_NLP/API"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIntentHandler(t *testing.T) {
	testingStruct := []struct {
		projectVar string
		pathVar    string
		shouldPass bool
	}{
		{"modusmea", "intent/modusmea/Users/vee/Desktop/ModusMea-84ada1644657.json/hello", true},
		{"modusmea", "intent/modusmea/Users/vee/Desktop/ModusMea-84ada1644657.json/start", true},
		{"modusmea", "intent/modusmea/Users/vee/Desktop/ModusMea-84ada1644657.json/nothing", false},
	}
	fmt.Println("#####TestGetIntentHandler#####")

	for _, testingVar := range testingStruct {
		path := fmt.Sprintf("/intents/%s/%s", testingVar.projectVar, testingVar.pathVar)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()

		router.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + SpeakerInput + "}", GetIntent).Methods("GET")
		router.ServeHTTP(rr, req)

		// some of these should pass, and some should not
		fmt.Println("path: ", path, "has a statusCode of: ", rr.Code, ", and an error message of: ", rr.Body)
		if rr.Code == http.StatusOK && !testingVar.shouldPass {
			t.Errorf("handler should have failed on routeVariable %s: got %v want %v", testingVar.projectVar, rr.Code, http.StatusOK)
		}
	}
}
