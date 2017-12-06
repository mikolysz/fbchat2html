package main

import (
	"html/template"
	"io"
	"time"
)

// MessagesTree contains messages sorted by date to make displaying them as structured html easier.
// It's a map with years as keys and embedded maps as values.
// Those embedded maps contain months as keys and other embedded maps as values.
// Those contain days as keys and slices of messages for one particular day as values.
type MessagesTree map[int]year

type month map[int][]Message
type year map[time.Month]month

// ToMessagesTree converts a Conversation, as unmarshalled from json, to a structure easy to use when  converting messages to  html format.
func (t *Thread) ToMessagesTree() *MessagesTree {
	tree := make(MessagesTree)
	for _, v := range t.Messages {
		tree.Insert(v)
	}
	return &tree
}

// Inserts one message into the tree, creating the needed embedded maps as needed.
func (t MessagesTree) Insert(m Message) {
	d := m.Date
	if t[d.Year()] == nil {
		t[d.Year()] = make(year)
	}
	if t[d.Year()][d.Month()] == nil {
		t[d.Year()][d.Month()] = make(month)
	}
	t[d.Year()][d.Month()][d.Day()] = append(t[d.Year()][d.Month()][d.Day()], m)
}

//The template we use to create the html documents. Probably should reside in a separate file and get some indentation.
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

// HTML Converts a MessagesTree to html.
func (t *MessagesTree) HTML(w io.Writer) error {
	return tmpl.Execute(w, t)
}
