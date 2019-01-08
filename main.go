package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mlabouardy/dialogflow-go-client"
	"github.com/mlabouardy/dialogflow-go-client/models"
	"github.com/oayomide/messenger"
)

var verifyT, accessT, appS, dialogflowToken = messenger.GetTokens()

var (
	verifyToken = flag.String("verify-token", verifyT, "The token used to verify facebook (required)")
	verify      = flag.Bool("should-verify", false, "Whether or not the app should verify itself")
	pageToken   = flag.String("page-token", accessT, "The token that is used to verify the page on facebook")
	appSecret   = flag.String("app-secret", appS, "The app secret from the facebook developer portal (required)")
	host        = flag.String("host", "localhost", "The host used to serve the messenger bot")
	port        = flag.Int("port", 4000, "The port used to serve the messenger bot")
)

var bot = messenger.New(messenger.Options{
	Verify:      *verify,
	AppSecret:   *appSecret,
	VerifyToken: *verifyToken,
	Token:       *pageToken,
	WebhookURL:  "/api/webhook",
})

func main() {
	flag.Parse()

	if *verifyToken == "" || *appSecret == "" || *pageToken == "" {
		fmt.Println("There seem to be fields that are empty. Please make sure none is empty")
		fmt.Println()
		flag.Usage()

		os.Exit(-1)
	}

	//we want to initialize our DialogFlow here
	err, dialogflowClient := dialogflow.NewDialogFlowClient(models.Options{
		AccessToken: dialogflowToken,
	})

	if err != nil {
		fmt.Println("Error creating an instance of dialogflow", err)
	}
	//we want to trigger something when out bot receives a message(ing) (event)
	bot.HandleMessage(func(msg messenger.Message, r *messenger.Response) {
		fmt.Printf("Message %v received at %v\n", msg.Text, msg.Time.Format(time.UnixDate))

		//user, err := bot.ProfileByID(msg.Sender.ID)

		if err != nil {
			fmt.Println("Oops! Error here:", err)
		}
		//we want to get the intent of the message of the user
		qr := models.Query{Query: msg.Text}
		info, err := dialogflowClient.QueryFindRequest(qr)

		if err != nil {
			fmt.Printf("Error processing NLP on the message. Got the error: %s", err)
		}

		vd, _ := json.Marshal(info)
		fmt.Printf("This is the info of the text processed: %v", string(vd))
		r.Text(info.Result.Fulfillment.Speech, messenger.ResponseType)
		//HandleIntents(info.Result.Contexts)
	})

	servingURL := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Bot up and running on: ", servingURL)
	log.Fatal(http.ListenAndServe(servingURL, bot.Handler()))
}
