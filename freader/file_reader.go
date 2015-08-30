package freader


type FileReader interface {
    ReadFile(file string, pos int) (err error)
    ReadLine() (string, error)
//    GetPos() int
}
