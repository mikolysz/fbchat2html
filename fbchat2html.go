package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

var outputDir string

func main() {
	flag.StringVar(&outputDir, "output", "output", "Output")
	flag.StringVar(&outputDir, "o", "output", "Output") // Allow -o instead of --output
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: fbchat2html <filename>")
		return
	}

	data := Conversations{} // Data read from the json-formatted file.
	fmt.Println("Reading file")
	if f, err := ioutil.ReadFile(flag.Arg(0)); err != nil {
		fmt.Printf("Error opening file, exiting\n%s\n", err)
		return
	} else {
		if err := json.Unmarshal(f, &data); err != nil {
			fmt.Println("Error unmarshalling ", err)
			return
		}
	}
	/*
		//I was removing the output directory just for safety but it seems there's no need.
		// If something goes wrong, you can uncomment.
		if err := os.RemoveAll(outputDir); err != nil {
			fmt.Println("Error removing output directory", err)
			return
		}
	*/
	if err := os.Mkdir(outputDir, 0); err != nil && !os.IsExist(err) {
		fmt.Println("Error creating output directory", err)
		return
	}
	wg := sync.WaitGroup{} // Track all the operations running simultaneously.
	wg.Add(1)
	go handleStatistics(&data, &wg)
	wg.Add(1)
	go handleArchive(&data, &wg)
	wg.Wait()

}

// Output some interesting statistics to output_dir/stats.txt.
// TODO: Allow changing the path.
func handleStatistics(data *Conversations, wg *sync.WaitGroup) {
	stats, err := os.Create(path.Join(outputDir, "stats.txt"))
	if err != nil {
		fmt.Println("Can't create stats file, exiting.", err)
		return
	}
	normalThreads, groupThreads := 0, 0
	received, sent := 0, 0
	receivedFromGroups, sentToGroups := 0, 0
	for _, t := range data.Threads {
		r, s := t.CountMessages(data.User)
		received += r
		sent += s
		if t.IsGroupThread() {
			groupThreads++
			receivedFromGroups += r
			sentToGroups += s
		} else {
			normalThreads++
		}
	}
	fmt.Fprintln(stats,
		"Here are some statistics:\n",
		"Threads:\n",
		"normal: ", normalThreads, "\n",
		"group: ", groupThreads, "\n",
		"all: ", normalThreads+groupThreads, "\n",
		"messages:", "\n",
		"total received: ", received, "\n",
		"total sent", sent, "\n",
		"all: ", received+sent, "\n",
		"received from group conversationns: ", receivedFromGroups, "\n",
		"sent to group Conversations", sentToGroups, "\n",
		"received from normal conversations: ", received-receivedFromGroups, "\n",
		"sent to normal conversations: ", sent-sentToGroups, "\n",
	)
	wg.Done()
}

// Actually output the html documents.
func handleArchive(data *Conversations, wg *sync.WaitGroup) {
	for _, t := range data.Threads {
		if len(t.Participants) >= 5 { // Ugly hack to make the names short enough so windows doesn't complain.
			//TODO: Change "John Smit, John Doe, Allice, Bob, Others" to "John Smith, John Doe, Allice, Bob and  others" or use (...) at the end of the name.
			// TODO: Allow creating a separate directory for group threads.
			t.Participants = append(t.Participants[0:4], "others")
		}
		filename := strings.Join(t.Participants, ",") + ".html"
		if f, err := os.Create(path.Join(outputDir, filename)); err != nil {
			panic(err.Error())
		} else {
			t.ToMessagesTree().HTML(f)
		}
	}

	wg.Done()
}
