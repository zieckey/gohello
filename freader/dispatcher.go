package freader

import (
    "github.com/howeyc/fsnotify"
    "log"
    "strings"
)

type Dispatcher struct {
    dirs    []string
    watcher *fsnotify.Watcher
    status *ProcessStatus
}

func NewDispatcher(dir string) (d *Dispatcher, err error) {
    d = &Dispatcher{}
    d.watcher, err = fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    dirs := strings.Split(dir, ",")
    for _, v := range dirs {
        d.dirs = append(d.dirs, GetAbsPath(v))
    }
    log.Println(d.dirs)
    d.status, err = NewProcessStatus(*statusFile)
    if err != nil {
        log.Fatal(err)
    }
    return d, err
}


func (d *Dispatcher) OnCreate(ev *fsnotify.FileEvent) {
    if IsDir(ev.Name) {
        d.watcher.Watch(ev.Name)
        //TODO if we renamed ev.Name latterly, we should add the new name to the watching list
    }
}

func (d *Dispatcher) OnDelete(ev *fsnotify.FileEvent) {

}

func (d *Dispatcher) OnModify(ev *fsnotify.FileEvent) {

}

func (d *Dispatcher) WatchEvent() {
    for {
        select {
        case ev := <-d.watcher.Event:
            log.Println("event:", ev, " name=", ev.Name)
            if ev.IsCreate() {
                d.OnCreate(ev)
            } else if ev.IsDelete() {
                d.OnDelete(ev)
            } else if ev.IsModify() {
                d.OnModify(ev)
            } else {
                log.Printf("don't care")
            }
        case err := <-d.watcher.Error:
            log.Println("error:", err)
        }
    }
}

func (d *Dispatcher) Run() {
    for _, f := range d.dirs {
        err := d.watcher.Watch(f)
        if err != nil {
            log.Fatal("Watch event of " + f + " FAILED: " + err.Error())
        }
    }

    d.WatchEvent()
    d.Close()
}

func (d *Dispatcher) Close() {
    d.watcher.Close()
}
