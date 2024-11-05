package main

type user struct {
	name   string
	number int
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

/* Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead. */

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

