package API_test

import (
. "CreateConversationBackend_NLP/API"
"fmt"
"github.com/gorilla/mux"
"net/http"
"net/http/httptest"
"testing"
)

// CREATE CASES TO COVER
//editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}", ReturnCreateIntentError).Methods("POST")
//editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + IntentID + "}", ReturnCreateIntentError).Methods("POST")
//editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + IntentID + "}/{ " + DisplayName + "}", CreateAnIntent).Methods("POST")
//editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + IntentID + "}/{ " + DisplayName + "}/{" + TrainingPhraseParts + "}", CreateAnIntent).Methods("POST")
//editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + IntentID + "}/{ " + DisplayName + "}/{" + TrainingPhraseParts + "}/{" + MessageTexts + "}", CreateAnIntent).Methods("POST")

func TestCreateDeleteHandler(t *testing.T) {
	testingStruct := []struct {
		projectVar string
		pathVar    string
		shouldPass bool
	}{
		{"modusmea", "Users/vee/Desktop/ModusMea-84ada1644657.json/", true},
		{"createagents", "Users/vee/Desktop/CreateAgents-1ea24b69d927.json", true},
		{"fail", "Users/vee/Desktop/ModusMea-84ada1644657.json", false},
		{"modusmea", "Users/vee/Desktop/ModusMea.json", false},
	}
	fmt.Println("#####TestGetIntentsHandlerPass#####")

	for _, testingVar := range testingStruct {
		path := fmt.Sprintf("/intent/%s/%s", testingVar.projectVar, testingVar.pathVar)
		req, err := http.NewRequest("POST", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()

		router.HandleFunc("/intents/{"+ProjectID+"}/{"+ProjectCred+":.*}", CreateAnIntent)
		router.ServeHTTP(rr, req)

		// some of these should pass, and some should not
		fmt.Println("path: ", path, "has a statusCode of: ", rr.Code, ", and an error message of: ", rr.Body)
		if rr.Code == http.StatusOK && !testingVar.shouldPass {
			t.Errorf("handler should have failed on routeVariable %s: got %v want %v", testingVar.projectVar, rr.Code, http.StatusOK)
		}


	}
}

