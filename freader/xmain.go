package freader

import (
	"flag"
)

var dir *string = flag.String("file_path", ".,e:/1", "The dir of the file which we need to process. you can use ',' to seperate different dir")
var statusFile *string = flag.String("status", "e:/status.txt", "The status file which holds the processing status")
var priorityLevel *int = flag.Int("priority_level", 0, "The max priority level of the file handler. 0 means that it don't has priorty")


func Run() {
	flag.Parse()

	dispatcher, err := NewDispatcher(*dir)
	if err != nil {
		return
	}
	dispatcher.Run()
}