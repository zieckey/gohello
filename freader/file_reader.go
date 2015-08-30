package freader


type FileReader interface {
    ReadFile(file string, pos int) (err error)
    ReadLine() ([]byte, error)
//    GetPos() int
}
