package main

// Holds the data types needed to unmarshal the json and their associated methods.
import "time"

// Conversations is the top-level type holding data unmarshalled from the json object.
type Conversations struct {
	User    string   // The name of the user the archive was generated for, useful when you need to see which messages were send and which were received.
	Threads []Thread // List of conversations.
}

// Holds one conversation, as unmarshalled from the Json.
type Thread struct {
	Participants []string  // List of people participating in this conversation, not including the user the archive was generated for.
	Messages     []Message // List of messages in that conversation.
}

// IsGroupThread determines if this conversation is a group conversation.
func (t Thread) IsGroupThread() bool {
	if len(t.Participants) > 1 {
		return true
	}
	return false
}

// CountMessages counts messages received and sent by one particular user in this conversation.
// It accepts one parameter, "user" which should be the name of the user who we consider as the onesending messages.
// Usually it will be the Username field from the main Conversations struct, meaning the user whose archive we're dealing with.
func (t Thread) CountMessages(user string) (received, sent int) {
	for _, m := range t.Messages {
		if m.Sender == user {
			sent++
		} else {
			received++
		}
	}
	return
}

// A single message.
type Message struct {
	Date    timestamp // The date this message was sent on.
	Sender  string    // The name of the sender.
	Message string    // Contents of the message.
}

// timestamp is a type needed because the json has  timestamps in a peculiar format and json.Unmarshal can't deal with it when a normal time.Time is used.
type timestamp struct{ time.Time }

// UnmarshalJSON unmarshals that peculiar time format to a proper Time value.

func (t *timestamp) UnmarshalJSON(data []byte) error {
	// Copied almost verbatim from the source of the standard library, with one minor modification to the string passed to time.Parse.
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	result, err := time.Parse(`"2006-01-02T15:04-07:00"`, string(data))
	t.Time = result
	return err
}
