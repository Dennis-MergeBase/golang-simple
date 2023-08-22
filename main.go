package main

import (
	. "CreateConversationBackend_NLP/API"
	. "CreateConversationBackend_NLP/ErrorHandling"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

const EDIT_SERVER_PORT = "4200"
const FETCH_SERVER_PORT = "4300"
/////////////////
/////////////////


func main() {
	var waitGroup sync.WaitGroup

	// we are running two servers here
	// one to edit intents and entities
	editServer := mux.NewRouter()
	// one to support the resulting chatbots
	fetchServer := mux.NewRouter()

	// FOR SERVER TESTING PURPOSES
	editServer.HandleFunc("/", helloBoys)
	fetchServer.HandleFunc("/", helloBoys)

	// CATCH ERRORS
	editServer.NotFoundHandler = http.HandlerFunc(ReturnEndpointError)
	fetchServer.NotFoundHandler = http.HandlerFunc(ReturnEndpointError)

	// CREATE
	editServer.HandleFunc("/entity/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + EntityTypeID + "}/{" + EntityTypeValue + "}/{" + EntitySynonyms + "}", CreateEntity).Methods("POST")


	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}", ReturnCreateIntentError).Methods("POST")
	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + DisplayName + "}", CreateAnIntent).Methods("POST")
	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + DisplayName + "}/{" + TrainingPhraseParts + "}", ReturnCreateIntentError).Methods("POST")
	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " + DisplayName + "}/{" + TrainingPhraseParts + "}/{" + MessageTexts + "}", CreateAnIntent).Methods("POST")

	// READ
	// TODO: entity
	fetchServer.HandleFunc("/entities/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + EntityTypeID + "}", ListAllEntitiesFromID).Methods("GET")
	fetchServer.HandleFunc("/intents/{" + ProjectID + "}/{" + ProjectCred + ":.*}", ListAllIntents).Methods("GET")

	fetchServer.HandleFunc("/entity/{" + ProjectID + "}/{" + ProjectCred + ":.*}", GetEntity).Methods("GET")
	fetchServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + SpeakerInput + "}", GetIntent).Methods("GET")

	// UPDATE
	// TODO: entity
	editServer.HandleFunc("/entity/{ + " + ProjectID + "}/{" + ProjectCred + ":.*}", UpdateEntity).Methods("PATCH")
	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{ " +IntentID+ "}/{ " + DisplayName + "}/{" + TrainingPhraseParts + "}/{" + MessageTexts + "}", UpdateAnIntent).Methods("PATCH")

	// DELETE
	editServer.HandleFunc("/entity/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + EntityTypeID + "}/{" + EntityTypeValue + "}", DeleteAnEntity).Methods("DELETE")
	editServer.HandleFunc("/intent/{" + ProjectID + "}/{" + ProjectCred + ":.*}/{" + IntentPath + "}", DeleteAnIntent).Methods("DELETE")

	// EXPORT
	// TODO
	editServer.HandleFunc("/agent", ExportAgentToZip).Methods("GET")

	// Run the web servers.
	waitGroup.Add(1)
	// one for our edit server
	go ListenNow(EDIT_SERVER_PORT, editServer)
	// one for our fetch server
	waitGroup.Add(2)
	go ListenNow(FETCH_SERVER_PORT, fetchServer)

	// this will wait forever as it should... prevents main routine from exiting
	waitGroup.Wait()
}
func ListenNow(port string, server http.Handler){
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})


	log.Fatal(http.ListenAndServe(":" + port, server),  handlers.CORS(originsOk, headersOk, methodsOk)(server))
}

func helloBoys(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello beautiful..."))
}