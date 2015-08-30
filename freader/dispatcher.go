package freader

import (
    "github.com/howeyc/fsnotify"
    "log"
    "sync"
    "github.com/golang/glog"
)

type Dispatcher struct {
    dir    string
    watcher *fsnotify.Watcher
    status *ProcessStatus
    h *FilesHandler
}

func NewDispatcher(dir string) (d *Dispatcher, err error) {
    glog.Infof("NewDispatcher")
    d = &Dispatcher{}
    d.watcher, err = fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    d.dir = dir
    d.status, err = NewProcessStatus(*statusFile)
    if err != nil {
        log.Fatal(err)
    }

    d.h, err = NewFilesHandler(dir)
    if err != nil {
        log.Fatal(err)
    }

    return d, err
}


func (d *Dispatcher) onCreate(ev *fsnotify.FileEvent) {
    if IsDir(ev.Name) {
        d.watcher.Watch(ev.Name)
        //Ignore this : FIXME if we renamed ev.Name latterly, we should add the new name to the watching list.
    } else {
        d.h.OnFileCreated(ev.Name)
    }
}

func (d *Dispatcher) onDelete(ev *fsnotify.FileEvent) {
    d.status.OnFileDeleted(ev.Name)
}

func (d *Dispatcher) onModify(ev *fsnotify.FileEvent) {
    d.h.OnFileModified(ev.Name)
}

func (d *Dispatcher) watchEvent(wg *sync.WaitGroup) {
    wg.Done()
    for {
        select {
        case ev := <-d.watcher.Event:
            glog.Info("event:", ev, " name=", ev.Name)
            if ev.IsCreate() {
                d.onCreate(ev)
            } else if ev.IsDelete() {
                d.onDelete(ev)
            } else if ev.IsModify() {
                d.onModify(ev)
            } else {
                log.Printf("don't care")
            }
        case err := <-d.watcher.Error:
            log.Println("error:", err)
        }
    }
}

func (d *Dispatcher) Run() {
    glog.Infof("Watching <%v>", dir)
    err := d.watcher.Watch(d.dir)
    if err != nil {
        log.Fatal("Watch event of " + d.dir + " FAILED: " + err.Error())
    }

    //start to watch the file event and wait the goroutine started
    var wg sync.WaitGroup
    wg.Add(1)
    go d.watchEvent(&wg)
    wg.Wait()

    //start file handler to run
    d.h.Run()

    d.Close()
}

func (d *Dispatcher) Close() {
    d.watcher.Close()
}
