package languageProcessing

import (
	"context"
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1beta2"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1beta2"
)

func AnaliseEntitySentiment( text string) error {
	ctx := context.Background()
	c, err := language.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Unable to conect to the natural language Processing API")
	}
	req := &languagepb.AnalyzeEntitySentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	}
	resp, err := c.AnalyzeEntitySentiment(ctx, req)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Unable to process the sentiment of the text")
	}
	log.Println(resp)
	return nil
}