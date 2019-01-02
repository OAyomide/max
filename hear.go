package main

import "github.com/oayomide/messenger"

//Hear listens for the text that is coming from the user then trigger, instead of using the
//
const (
	greeting = "dfdf"
	enquiry  = "sdlfkslf"
	jokes    = "fdkjfkdg"
	games    = "dfkldkdf"
)

func Hear(word string) func(mes messenger.Messenger, res *messenger.Response) {

	if word == mes.Text {
		return
	}
	return
}
