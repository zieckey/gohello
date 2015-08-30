package freader
import (
    "os"
    "github.com/golang/glog"
    "bufio"
    "bytes"
    "compress/gzip"
)

type GzipReader struct {
    path string
    pos int
    fp *os.File
    r* bufio.Reader // The reader of os.File fp
    gr* gzip.Reader
}

func NewGzipReader() *GzipReader {
    br := &GzipReader{
        path : "",
        pos:0,
        fp:nil,
    }

    return br
}

func (r *GzipReader) ReadFile(file string, pos int) (err error) {
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
        r.gr, err = gzip.NewReader(r.fp)
        if err != nil {
            return err
        }
        r.r = bufio.NewReader(r.gr)
    } else {
        r.gr.Reset(r.fp)
        r.r.Reset(r.gr)
    }

    return nil
}

func (r *GzipReader) ReadLine() (line []byte, err error) {
    line, err = r.r.ReadBytes('\n')
    //glog.Infof("len(line)=%v %v", len(line), base64.StdEncoding.EncodeToString(line))
    line = bytes.TrimRight(line, "\r\n")
    //glog.Infof("len(line)=%v %v after trim", len(line), base64.StdEncoding.EncodeToString(line))
    return line, err
}

func (r *GzipReader) GetPos() int {
    return r.pos
}