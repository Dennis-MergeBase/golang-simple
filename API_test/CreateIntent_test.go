package API_test

//func TestCreateIntentHandler(w http.ResponseWriter, r *http.Request) {
//	// A very simple health check.
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//
//	// In the future we could report back on the status of our DB, or our cache
//	// (e.g. Redis) by performing a simple PING, and include them in the response.
//	io.WriteString(w, `{"alive": true}`)
//}
//
////	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + DisplayName + "}", CreateAnIntent).Methods("POST")
//
//
//func TestHelloWorldHandler(t *testing.T) {
//	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//	// pass 'nil' as the third parameter.
//	req, err := http.NewRequest("GET", "/intents/createagent/Users%2Fvee%2FDesktop%2FCreateAgents-1ea24b69d927.json", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(TestCreateIntentHandler)
//
//	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//	// directly and pass in our Request and ResponseRecorder.
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code is what we expect.
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	// Check the response body is what we expect.
//	expected := `{"alive": true}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//}