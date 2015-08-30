package freader
import "github.com/golang/glog"

type TextModule interface {
    OnRecord([]byte)
}

type DefaultTextModule struct {
}

func (m *DefaultTextModule) OnRecord(line []byte) {
    glog.Infof("Read a new line, len=%v <%s> ", len(line), string(line))
}