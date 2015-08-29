package freader


type IPathReader interface {
    ReadLine() (string, error)
}

type PathReader struct {
    dir string
}

func NewPathReader(dir string) *PathReader {
    r := &PathReader{}
    r.dir = dir
    return r
}

func (r *PathReader) ReadLine() (line string, err error)  {

    return line, err
}