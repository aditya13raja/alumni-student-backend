package configs

import (
	"log"
	"os"

	"github.com/pusher/pusher-http-go/v5"
)

// Pusher client instance
var PusherClient *pusher.Client

// Initalize pusher -> set app keys to start it
func InitPusher() {
	app_id := os.Getenv("PUSHER_APP_ID")
	key := os.Getenv("PUSHER_KEY")
	cluster := os.Getenv("PUSHER_CLUSTER")
	secret := os.Getenv("PUSHER_SECRET")

	if app_id == "" || key == "" || cluster == "" || secret == "" {
		log.Fatal("Pusher configuration is missing in .env file")
	}

	PusherClient = &pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	log.Println("Pusher client initalized successfully")
}
