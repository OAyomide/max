package main

import (
	"github.com/oayomide/messenger"
)

//HandleIntents checks the type of intent the bot receives and sends formatted messages based on that
func HandleIntents(intent string, msg messenger.Message, r *messenger.Response) {
	switch intent {
	case "unknown":
		//user, err := bot.ProfileByID(msg.Sender.ID)

		//want to send the image here
		r.Attachment("image", "http://www.messenger-rocks.com/image.jpg", "RESPONSE")

	case "who_is_this":
		img := []messenger.QuickReply{messenger.QuickReply{Title: "Whatchu wanna knw", Payload: "qr1"}}
		r.TextWithReplies("What do you want to know?", img, "image")
	}
}
