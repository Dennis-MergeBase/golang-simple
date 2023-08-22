package API

import (
	"cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
	"net/http"
)

func DeleteAnEntity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	fmt.Printf("Deleting values %s under projects/%s/agent/entityTypes/%s...\n", params[EntityTypeValue], params[ProjectID], params[EntityTypeID])
	err := DeleteEntity(params[ProjectID], params[ProjectCred], params[EntityTypeID], params[EntityTypeValue])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w,"Done!\n")
}


// [START dialogflow_delete_entity]
func DeleteEntity(projectID, permissionsJSONPath, entityTypeID, entityValue string) error {
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
	entityValues := []string{entityValue}
	request := dialogflowpb.BatchDeleteEntitiesRequest{Parent: parent, EntityValues: entityValues}

	deletionOp, requestErr := entityTypesClient.BatchDeleteEntities(ctx, &request)
	if requestErr != nil {
		return requestErr
	}

	err := deletionOp.Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}