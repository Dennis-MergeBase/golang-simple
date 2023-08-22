package API

import (
	. "CreateConversationBackend_NLP/ErrorHandling"
	"cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"net/http"
	"strings"
)

func CreateEntity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	projectID := params[ProjectID]
	synonymsRaw := params[EntitySynonyms]
	entityValue := params[EntityTypeValue]
	entityTypeID := params[EntityTypeID]
	synonyms := strings.Split(synonymsRaw, ",")

	fmt.Printf("Creating entity %s...\n", entityValue)
	err := CreateAnEntity(projectID, "", entityValue, synonyms, entityTypeID)
	if err != nil {
		ReturnGeneralError(w, http.StatusBadRequest, err.Error())
	}
	fmt.Fprintf(w, "Entity type %s created under type %s\n", entityValue, entityTypeID)
}

// [START dialogflow_create_entity]
func CreateAnEntity(projectID, permissionsJSONPath string, entityValue string, synonyms []string, entityTypeID string) error {
	ctx := context.Background()


	entityTypesClient, clientErr := dialogflow.NewEntityTypesClient(ctx, option.WithCredentialsFile(permissionsJSONPath))
	if clientErr != nil {
		return clientErr
	}
	defer entityTypesClient.Close()

	if projectID == "" || entityTypeID == "" {
		return errors.New(fmt.Sprintf("Received empty project (%s) or entity type (%s)", projectID, entityTypeID))
	}

	parent := fmt.Sprintf("projects/%s/agent/entityTypes/%s", projectID, entityTypeID)
	entity := dialogflowpb.EntityType_Entity{Value: entityValue, Synonyms: synonyms}
	entities := []*dialogflowpb.EntityType_Entity{&entity}

	request := dialogflowpb.BatchCreateEntitiesRequest{Parent: parent, Entities: entities}

	creationOp, requestErr := entityTypesClient.BatchCreateEntities(ctx, &request)
	if requestErr != nil {
		return requestErr
	}

	err := creationOp.Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}