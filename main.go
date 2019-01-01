package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/oayomide/messenger"
)

var verifyT, accessT, appS = messenger.GetTokens()

var (
	verifyToken = flag.String("verify-token", verifyT, "The token used to verify facebook (required)")
	verify      = flag.Bool("should-verify", false, "Whether or not the app should verify itself")
	pageToken   = flag.String("page-token", accessT, "The token that is used to verify the page on facebook")
	appSecret   = flag.String("app-secret", appS, "The app secret from the facebook developer portal (required)")
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

		//if the text contains "Hello or Hi" which signifies greetings, we want to reply and greet too (with an emoji probably)
		if strings.Contains(msg.Text, "Hello") {
			r.Text(fmt.Sprintf("Hey, %v. I am Max and the personal chatbot to Ayomide Onigbinde. Dont worry, I get better 😉", user.FirstName), messenger.ResponseType)
		}
		r.Text(fmt.Sprintf("Hello, %v", user.FirstName), messenger.ResponseType)
	})

	servingURL := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Bot up and running on: ", servingURL)
	log.Fatal(http.ListenAndServe(servingURL, bot.Handler()))
}

// func HandleKeyText(text string, func()) bool {
// 	if
// }
