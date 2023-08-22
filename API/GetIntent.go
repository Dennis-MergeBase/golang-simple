package API

import (
	errorMessage "CreateConversationBackend_NLP/ErrorHandling"
	"cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetIntent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	sessionID := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000))

	_, err := GetResponseFromTextReq( params[ProjectID], params[ProjectCred], sessionID, params[SpeakerInput], "en")
	if err != nil {
		errorMessage.ReturnGeneralError(w, http.StatusBadRequest, err.Error())
	}
}

// [START dialogflow_detect_intent_text]
func GetResponseFromTextReq(projectID, permissionsJSONPath, sessionID, text, languageCode string) (string, error) {
	ctx := context.Background()
	sessionClient, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsFile(permissionsJSONPath))
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)

	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}
// [END dialogflow_detect_intent_text]














// [START dialogflow_detect_intent_audio]
func GetResponseFromAudioReq(projectID, sessionID, audioFile, languageCode string) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)

	// In this example, we hard code the encoding and sample rate for simplicity.
	audioConfig := dialogflowpb.InputAudioConfig{AudioEncoding: dialogflowpb.AudioEncoding_AUDIO_ENCODING_LINEAR_16, SampleRateHertz: 16000, LanguageCode: languageCode}

	queryAudioInput := dialogflowpb.QueryInput_AudioConfig{AudioConfig: &audioConfig}

	audioBytes, err := ioutil.ReadFile(audioFile)
	if err != nil {
		return "", err
	}

	queryInput := dialogflowpb.QueryInput{Input: &queryAudioInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput, InputAudio: audioBytes}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}
// [END dialogflow_detect_intent_audio]

// [START dialogflow_detect_intent_streaming]
func GetResponseFromStreamReq(projectID, sessionID, audioFile, languageCode string) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)

	// In this example, we hard code the encoding and sample rate for simplicity.
	audioConfig := dialogflowpb.InputAudioConfig{AudioEncoding: dialogflowpb.AudioEncoding_AUDIO_ENCODING_LINEAR_16, SampleRateHertz: 16000, LanguageCode: languageCode}

	queryAudioInput := dialogflowpb.QueryInput_AudioConfig{AudioConfig: &audioConfig}

	queryInput := dialogflowpb.QueryInput{Input: &queryAudioInput}

	streamer, err := sessionClient.StreamingDetectIntent(ctx)
	if err != nil {
		return "", err
	}

	f, err := os.Open(audioFile)
	if err != nil {
		return "", err
	}

	defer f.Close()

	go func() {
		audioBytes := make([]byte, 1024)

		request := dialogflowpb.StreamingDetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}
		err = streamer.Send(&request)
		if err != nil {
			log.Fatal(err)
		}

		for {
			_, err := f.Read(audioBytes)
			if err == io.EOF {
				streamer.CloseSend()
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			request = dialogflowpb.StreamingDetectIntentRequest{InputAudio: audioBytes}
			err = streamer.Send(&request)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	var queryResult *dialogflowpb.QueryResult

	for {
		response, err := streamer.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		recognitionResult := response.GetRecognitionResult()
		transcript := recognitionResult.GetTranscript()
		log.Printf("Recognition transcript: %s\n", transcript)

		queryResult = response.GetQueryResult()
	}

	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}

// [END dialogflow_detect_intent_streaming]



