package freader
import (
    "sync"
    "container/list"
    "io"
    "github.com/golang/glog"
    "time"
)

type PathReader struct {
    dir string


    fr FileReader

    waiting bool    //FIXME use atomic variable ??
    wakeup chan int

    currentReadingFile string

    mutex sync.Mutex
    files *list.List // The files to be reading
}

func NewPathReader(dir string) (*PathReader, error) {
    r := &PathReader{}
    r.dir = dir
    r.waiting = false
    r.files = list.New()
    r.wakeup = make(chan int)
    r.fr = createReader()
    return r, nil
}

//func (r *PathReader) ReadLine() (line string, err error)  {
//
//    return line, err
//}
//
//func (r *PathReader) Read(p []byte) (n int, err error) {
//    return 0, nil
//}

func (r *PathReader) Append(file string) (err error) {
    r.mutex.Lock()
    defer r.mutex.Unlock()
    r.files.PushBack(file)
    return err
}

const (
    kModify int = 1
    kCreate int = 2
)

func (r *PathReader) OnFileModified(file string) (err error) {
    if r.currentReadingFile == file && r.waiting {
        glog.Infof("send kModify signal")
        r.wakeup <- kModify
    } else {
        glog.Infof("do not need to send kModify signal")
    }
    return nil
}

func (r *PathReader) OnFileCreated(file string) (err error) {
    r.Append(file)
    if r.waiting && r.files.Len() == 1 {
        /*
        r.waiting : we will send a signal only if the goroutine is waiting
        r.files.Len() == 0 : when we create more than 2 files in the same time, the waiting goroutine may be still waiting when we try to send the second signal
         */
        glog.Infof("send kCreate signal")
        r.wakeup <- kCreate
    } else {
        glog.Infof("do not need to send kCreate signal")
    }
    return nil
}

func createReader() FileReader {
    //GzipReader, PTailReader
    if *reader_type == "PTailReader" {
        return NewPTailReader()
    } else if *reader_type == "GzipReader" {
        //TODO
    }

    return nil
}

func (r *PathReader) StartToRead() (err error) {
    glog.Infof("Starting to read files ...")
    startTime := time.Now()
    for {
        if r.files.Len() == 0 {
            glog.Infof("No files. Waiting ...")
            r.Wait()
            if r.files.Len() == 0 {
                glog.Errorf("logic ERROR, but ignore it now, we should review the code logic.")
                continue
            }
        }

        r.mutex.Lock()
        e := r.files.Front()
        r.files.Remove(e)
        r.mutex.Unlock()

        file, ok := e.Value.(string)
        glog.Infof("Processing file %v", file)
        if !ok {
            glog.Errorf("Get element from file List failed.")
            continue
        }

        r.fr.ReadFile(file, 0)
        if len(r.currentReadingFile) > 0 {
            dispatcher.status.OnFileProcessingFinished(r.currentReadingFile, startTime)
        }
        startTime = time.Now()
        r.currentReadingFile = file

        for {
            line, err := r.fr.ReadLine()
            if err == io.EOF {
                // there are still files which are ready to be processed
                if r.files.Len() > 0 {
                    break
                }

                // no more files. we wait this file to be updated or wait new file created
                glog.Infof("no more files, we wait this file <%v> to be updated. Waiting ...", file)
                r.Wait()
            } else if err != nil {
                glog.Errorf("Read data from <%s> failed : %v", file, err.Error())
                break
            }

            glog.Infof("Read a new line:<%s>", string(line))
            //TODO process the new line reading from data file
        }
    }
}

func (r *PathReader) Wait() int {
    r.waiting = true
    event := <-r.wakeup
    r.waiting = false
    return event
}