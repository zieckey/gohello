package freader

import (
	"flag"
	"log"
)

var dir *string = flag.String("file_path", "e:/1/1", "The dir of the data file which we need to process")
var statusFile *string = flag.String("status", "e:/1/status.txt", "The status file which holds the processing status")
var priorityLevel *int = flag.Int("priority_level", 0, "The max priority level of the file handler. 0 means that it don't has any priorty")
var filePattern *string = flag.String("file_pattern", "inc_*.gz", "The pattern of the name which we need to process")
//var bufferBlockSize *int = flag.Int("buffer_block_size", 1*1024*1024, "We use this size to specify the buffer size when read from a data file")
var reader_type *string = flag.String("reader_type", "PTailReader", "The type of the file reader, options supported now : GzipReader, PTailReader")
/*

 */

var dispatcher *Dispatcher
func Run() {
	flag.Parse()

	var err error
	dispatcher, err = NewDispatcher(*dir)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	dispatcher.Run()
}