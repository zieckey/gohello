package main
import (
    "path/filepath"
    "os"
    "fmt"
)

func main() {
    files, err := LookupFiles("E:/1/lianjiachengjiaofangyuan/2015-11-07.ori", "*.html")
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    //TODO:ERROR The process cannot access the file because it is being used by another process
    for _, f := range files {
        f1 := filepath.Join(f, ".20151107.html")
        err = os.Rename(f, f1)
        if err == nil {
            fmt.Printf("rename <%v> to <%v> OK\n", f, f1)
        } else {
            fmt.Printf("rename <%v> to <%v> failed : %v\n", f, f1, err.Error())
        }

    }
}


func LookupFiles(dir string, pattern string) ([]string, error) {
    var files []string = make([]string, 0, 5)

    err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }

        if f.IsDir() {
            return nil
        }

        if ok, err := filepath.Match(pattern, f.Name()); err != nil {
            return err
        } else if ok {
            files = append(files, path)
        }
        return nil
    })

    return files, err
}
