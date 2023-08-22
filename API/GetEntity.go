package API

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
	"net/http"
)

func GetEntity(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"List Entities under projects/%s/agent:\n")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	var entities []*dialogflowpb.EntityType_Entity
	var err error
	entities, err = GetEntityFromID(params[ProjectID], params[ProjectCred], params[EntityTypeID])
	if err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Fprintf(w,"Value: %s\n", entity.GetValue())
		fmt.Fprintf(w,"Synonyms:\n")
		for _, synonym := range entity.GetSynonyms() {
			fmt.Fprintf(w, "\t- %s\n", synonym)
		}
		fmt.Fprintln(w, "\n")
	}
}

func GetEntityFromID(projectID, permissionsJSONPath, entityTypeID string) ([]*dialogflowpb.EntityType_Entity, error) {
	ctx := context.Background()

	entityTypesClient, clientErr := dialogflow.NewEntityTypesClient(ctx, option.WithCredentialsFile(permissionsJSONPath))

	if clientErr != nil {
		return nil, clientErr
	}
	defer entityTypesClient.Close()

	if projectID == "" || entityTypeID == "" {
		return nil, errors.New(fmt.Sprintf("Received empty project (%s) or entity type (%s)", projectID, entityTypeID))
	}


	parent := fmt.Sprintf("projects/%s/agent/entityTypes/%s", projectID, entityTypeID)
	request := dialogflowpb.GetEntityTypeRequest{Name: parent}

	entityType, err := entityTypesClient.GetEntityType(ctx, &request)
	if err != nil {
		return []*dialogflowpb.EntityType_Entity{}, err
	}

	return entityType.GetEntities(), nil
}