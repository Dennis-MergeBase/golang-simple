package API

const (
	ProjectID           string = "project_id"
	ProjectCred         string = "project_cred"
	SpeakerInput        string = "speaker_input"
	DisplayName         string = "display_name"
	TrainingPhraseParts string = "training_phrase_parts"
	MessageTexts        string = "message_texts"
	IntentID            string = "intent_name"
	IntentPath          string = "intent_path"
	EntityName          string = "entity_name"
	EntityTypeID        string = "entity_type_id"
	EntityTypeValue     string = "entity_type_value"
	EntitySynonyms      string = "entity_synonyms"
)

const (
	Error200 string = "200: Success is yours"
	Error400 string = "400: Endpoint not supported. Potentially malformed."
	Error401 string = "401: Unauthorized user, does not permissions."
	Error404 string = "404: No response available for utterance"
	Error500 string = "500: Something went sideways. Sorry."
)