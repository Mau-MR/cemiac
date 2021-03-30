package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"time"
)

func CreateClient() *firestore.Client {
	projectID := "cemiac"
	ctx, _ := context.WithTimeout(context.Background(),5*time.Second)


	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
