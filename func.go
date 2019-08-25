package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"
	_ "os"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))

	// ------- local development ---------
	// reader := os.Stdin
	// writer := os.Stdout
	// myHandler(context.TODO(), reader, writer)
}

//EventsInput test
type EventsInput struct {
	CloudEventsVersion string      `json:"cloudEventsVersion"`
	EventID            string      `json:"eventID"`
	EventType          string      `json:"eventType"`
	Source             string      `json:"source"`
	EventTypeVersion   string      `json:"eventTypeVersion"`
	EventTime          time.Time   `json:"eventTime"`
	SchemaURL          interface{} `json:"schemaURL"`
	ContentType        string      `json:"contentType"`
	Extensions         struct {
		CompartmentID string `json:"compartmentId"`
	} `json:"extensions"`
	Data struct {
		CompartmentID      string `json:"compartmentId"`
		CompartmentName    string `json:"compartmentName"`
		ResourceName       string `json:"resourceName"`
		ResourceID         string `json:"resourceId"`
		AvailabilityDomain string `json:"availabilityDomain"`
		FreeFormTags       struct {
			Department string `json:"Department"`
		} `json:"freeFormTags"`
		DefinedTags struct {
			Operations struct {
				CostCenter string `json:"CostCenter"`
			} `json:"Operations"`
		} `json:"definedTags"`
		AdditionalDetails struct {
			Namespace        string `json:"namespace"`
			PublicAccessType string `json:"publicAccessType"`
			ETag             string `json:"eTag"`
		} `json:"additionalDetails"`
	} `json:"data"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	input := &EventsInput{}
	json.NewDecoder(in).Decode(input)

	outputJSON, _ := json.Marshal(&input)
	fmt.Println(string(outputJSON))
	out.Write([]byte("text message wo kaku"))
}
