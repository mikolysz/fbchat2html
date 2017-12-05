package main
//  MessagesTree is a map indexed by year, containing embedded maps of messages for that year.
import (
	"bytes"
	"html/template"
	"time"
)

// Those maps contain embedded maps for months, and those contain slices of messages written on a single day.
// This is the tree we need to display a conversation as html.
type MessagesTree map[int]year

type month map[int][]Message
type year map[time.Month]month

// New converts a Conversation, as unmarshalled from json, to a structure easy to use when  converting messages to  html format.
// TODO: Rename to something that makes more sense.
func New(t *Thread) *MessagesTree {
	tree := make(MessagesTree)
	for _, v := range t.Messages {
		tree.Insert(v)
	}
	return &tree
}

// Inserts one message into the tree, creating the needed embedded maps as needed.
func (t MessagesTree) Insert(m Message) {
	d := time.Time(m.Date)
	if t[d.Year()] == nil {
		t[d.Year()] = make(year)
	}
	if t[d.Year()][d.Month()] == nil {
		t[d.Year()][d.Month()] = make(month)
	}
	t[d.Year()][d.Month()][d.Day()] = append(t[d.Year()][d.Month()][d.Day()], m)
}

//The template we use to create the html documents. Proably should reside in a separate file and get some indentation.
const templateString = `
<html>
<head>
<title>Conversation</title>
<meta charset="UTF-8">
</head>
<body>
<h1>Messages</h1>
{{range $year, $yearElem := .}}
<h2>{{$year}}</h2>
{{range $month, $monthElem :=$yearElem}}
<h3>{{$month}}</h3>
{{range $day, $messages := $monthElem}}
<h4>{{$day}}</h4>
{{range $messages}}
<h5>{{.Sender}}</h5>
{{.Message}}
{{end}}
{{end}}
{{end}}
{{end}}
</body>
</html>
`

var tmpl = template.New("conversation")

func init() {
	template.Must(tmpl.Parse(templateString))
}

// Converts a MessagesTree to html.
// It's called String for simplicity (it satisfies the Stringer interface) but it probably should be renamed to something more appropriate.
func (t *MessagesTree) String() string {
	w := bytes.NewBuffer([]byte{})
	tmpl.Execute(w, t)
	return w.String()
	//I know, that was ugly.
}
