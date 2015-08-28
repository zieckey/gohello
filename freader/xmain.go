package freader

import (
	"flag"
)

var dir *string = flag.String("f", ".,e:/1", "The dir of the file which we need to process. you can use ',' to seperate different dir")



func Run() {
	flag.Parse()

	dispatcher, err := NewDispatcher(*dir)
	if err != nil {
		return
	}
	dispatcher.Run()
}