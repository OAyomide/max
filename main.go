package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/paked/messenger"
)

var (
	verifyToken = flag.String("verify-token", "token_verify", "The token used to verify facebook (required)")
	verify      = flag.Bool("should-verify", false, "Whether or not the app should verify itself")
	pageToken   = flag.String("page-token", "EAAILYOXCJxEBAMtT6ZC53b4hDxifOb7c2u3Kv8ejj1Vn50OcH7zWX6nEc9DBWpgzZAEFOMmu2AmfjzA7fcZBqr2yFM8UF3xdzhHC6satDaiacvnqYZAu0PRPqAs2JQ4eEH9ZAHLcSiD46dPh2TxRgIZB2QTkYZBrZBUiWoktrxqgkQZDZD", "The token that is used to verify the page on facebook")
	appSecret   = flag.String("app-secret", "07040c1b74220f2a0d53eb65cb9f46a6", "The app secret from the facebook developer portal (required)")
	host        = flag.String("host", "localhost", "The host used to serve the messenger bot")
	port        = flag.Int("port", 4000, "The port used to serve the messenger bot")
)

func main() {
	flag.Parse()

	if *verifyToken == "" || *appSecret == "" || *pageToken == "" {
		fmt.Println("There seem to be fields that are empty. Please make sure none is empty")
		fmt.Println()
		flag.Usage()

		os.Exit(-1)
	}

	bot := messenger.New(messenger.Options{
		Verify:      *verify,
		AppSecret:   *appSecret,
		VerifyToken: *verifyToken,
		Token:       *pageToken,
		WebhookURL:  "/api/webhook",
	})

	//we want to trigger something when out bot receives a message(ing) (event)
	bot.HandleMessage(func(msg messenger.Message, r *messenger.Response) {
		fmt.Printf("Message %v received at %v\n", msg.Text, msg.Time.Format(time.UnixDate))

		user, err := bot.ProfileByID(msg.Sender.ID)

		if err != nil {
			fmt.Println("Oops! Error here:", err)
		}

		r.Text(fmt.Sprintf("Hello, %v", user.FirstName), messenger.ResponseType)
	})

	servingURL := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Bot up and running on: ", servingURL)
	log.Fatal(http.ListenAndServe(servingURL, bot.Handler()))
}
