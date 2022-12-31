package backend

import (
	"context"
	"log"
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

	opt := option.WithCredentialsFile("firebase-credentials.json")
	conf := &firebase.Config{
		DatabaseURL: "https://wake-my-pc-bbc7c-default-rtdb.firebaseio.com",
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

				listen <- message
			}
		}
	}()
	return nil

}
