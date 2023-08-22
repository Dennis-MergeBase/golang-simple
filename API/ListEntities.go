package API

import "net/http"

//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/gorilla/mux"
//	"log"
//	"net/http"
//	dialogflow "cloud.google.com/go/dialogflow/apiv2"
//	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
//)
//
func ListAllEntitiesFromID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

}
//
//func ListEntities(projectID, entityTypeID string) ([]*dialogflowpb.EntityType_Entity, error) {
//	ctx := context.Background()
//
//	entityTypesClient, clientErr := dialogflow.NewEntityTypesClient(ctx)
//	if clientErr != nil {
//		return nil, clientErr
//	}
//	defer entityTypesClient.Close()
//
//	if projectID == "" || entityTypeID == "" {
//		return nil, errors.New(fmt.Sprintf("Received empty project (%s) or entity type (%s)", projectID, entityTypeID))
//	}
//
//
//	parent := fmt.Sprintf("projects/%s/agent/entityTypes/%s", projectID, entityTypeID)
//	request := dialogflowpb.ListEntityTypesRequest{}
//
//	err := entityTypesClient.ListEntityTypes(ctx, &request)
//	if err != nil {
//		return []*dialogflowpb.EntityType_Entity{}, err
//	}
//
//	return entityType.List(), nil
//}