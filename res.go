package main

import (
	"github.com/oayomide/messenger"
)

func HandleIntents(intent string, msg messenger.Message, r *messenger.Response) {
	switch intent {
	case 'fun':
		user, err := bot.ProfileByID(msg.Sender.ID)

		//want to send the image here
		r.Image("dfkdf")

	case 'about':
		r.TextWithReplies("What do you want to know?", []messenger.QuickReply{}, "msg_tpe")
	}
}