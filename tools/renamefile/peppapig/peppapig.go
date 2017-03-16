package main
import (
    "path/filepath"
    "os"
    "fmt"
    "strings"
)

func main() {
    files, err := LookupFiles("F:", "*.avi")
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }

    for _, f := range files {
        orig := f
        f = strings.Replace(f, "Peppa.Pig.", "", -1)
        f = strings.Replace(f, "Peppa.Pig.", "", -1)
        f = strings.Replace(f, "Peppa.Pig.", "", -1)
        f = strings.Replace(f, "Peppa_Pig_", "", -1)
        f = strings.Replace(f, "_", "", -1)
        f = strings.Replace(f, " ", "", -1)
        err = os.Rename(orig, f)
        if err == nil {
            fmt.Printf("rename <%v> to <%v> OK\n", orig, f)
        } else {
            fmt.Printf("rename <%v> to <%v> failed : %v\n", orig, f, err.Error())
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
