package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"

	"google.golang.org/api/option"
)

type FirebaseStorage struct {
	app         *firebase.App
	stateRef    *db.Ref
	messagesRef *db.Ref
}

var _ wakepc.PcStateStorage = &FirebaseStorage{}

func NewFirebaseStorage() FirebaseStorage {

	credentialsFile := os.Getenv("FIREBASE_CREDENTIALS_FILE")
	if credentialsFile == "" {
		log.Fatalf("could not read credentials file, please inform FIREBASE_CREDENTIALS_FILE environment variable")
	}

	databaseURL := os.Getenv("FIREBASE_DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("could not find database url, please inform FIREBASE_DATABASE_URL environment variable")
	}

	opt := option.WithCredentialsFile(credentialsFile)
	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("failed to start firebase backend %v", err)
	}
	client, err := app.Database(context.Background())
	ref := client.NewRef("state/")
	messages := client.NewRef("messages/")

	return FirebaseStorage{app: app, stateRef: ref, messagesRef: messages}
}

func (p *FirebaseStorage) Save(ctx context.Context, state wakepc.PcState) error {
	return p.stateRef.Child(state.MacAddress).Set(ctx, state)
}

func (p *FirebaseStorage) Find(ctx context.Context, mac string) (wakepc.PcState, error) {

	var state wakepc.PcState
	if err := p.stateRef.Child(mac).Get(ctx, &state); err != nil {
		return state, err
	}
	return state, nil
}

func (p *FirebaseStorage) FindAll(ctx context.Context) ([]wakepc.PcState, error) {

	l := make([]wakepc.PcState, 0)

	results, err := p.stateRef.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	for _, r := range results {
		var d wakepc.PcState
		if err := r.Unmarshal(&d); err != nil {
			return nil, fmt.Errorf("error unmarshaling results %w", err)
		}
		l = append(l, d)

	}
	return l, nil
}

func (p *FirebaseStorage) Listen(ctx context.Context, mac string, listen chan wakepc.PcCommandEvent) error {
	log.Printf("starting listeners for %v", mac)
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case <-time.After(1 * time.Second):
				var message wakepc.PcCommandEvent

				if err := p.messagesRef.Child(mac).Get(ctx, &message); err != nil {
					log.Printf("error reading messages %v", err)
					continue
				}

				if message.Command == "" {
					continue
				}
				p.messagesRef.Child(mac).Delete(ctx)

				if message.Time < time.Now().Add(-30*time.Second).Unix() {
					log.Printf("ignoring old event %v\n", message)
					continue
				}

				listen <- message
			}
		}
	}()
	return nil

}

func (p *FirebaseStorage) Push(ctx context.Context, mac string, event wakepc.PcCommandEvent) error {
	event.Time = time.Now().Unix()
	return p.messagesRef.Child(mac).Set(ctx, event)
}
