package freader
import (
    "os"
    "github.com/golang/glog"
    "bufio"
)

type PTailReader struct {
    path string
    pos int
    fp *os.File
    r* bufio.Reader // The reader of os.File fp
}

func NewPTailReader() *PTailReader {
    br := &PTailReader{
        path : "",
        pos:0,
        fp:nil,
    }

    return br
}

func (r *PTailReader) ReadFile(file string, pos int) (err error) {
    if r.fp != nil {
        glog.Infof("Finished to process file %v", r.path)
        r.fp.Close()
        r.fp = nil
    }

    r.path = file
    r.fp, err = os.OpenFile(file, os.O_RDONLY, 0644)
    if err != nil {
        glog.Errorf("OpenFile <%s> failed : %v\n", file, err.Error())
        return err
    }
    glog.Infof("OpenFile %v OK", file)

    if pos > 0 {
        r.fp.Seek(int64(pos), os.SEEK_SET)
    }

    if r.r == nil {
        r.r = bufio.NewReader(r.fp)
    } else {
        r.r.Reset(r.fp)
    }

    return nil
}

func (r *PTailReader) ReadLine() (line string, err error) {
    //glog.Infof("ReadString from %v", r.path)
    line, err = r.r.ReadString('\n')
    if len(line) > 0 && line[len(line) - 1] == '\n' {
        line = line[:len(line) - 1]
    }
    return line, err
}

func (r *PTailReader) GetPos() int {
    return r.pos
}