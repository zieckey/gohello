package freader
import (
    "path"
    "os"
    "path/filepath"
    "os/exec"
    "strings"
)

// GetAbsPath gets the absolute path of the giving path p
func GetAbsPath(p string) string {
    if filepath.IsAbs(p) {
        return p
    }

    file, _ := exec.LookPath(os.Args[0])
    exePath, _ := filepath.Abs(file)
    dir := filepath.Dir(exePath)
    fullPath := path.Join(dir, p)
    fullPath, _ = filepath.Abs(fullPath)
    return strings.TrimRight(fullPath, "/\\")
}

// IsDir returns true if given path is a directory,
// or returns false when it's a file or does not exist.
func IsDir(dir string) bool {
    f, e := os.Stat(dir)
    if e != nil {
        return false
    }
    return f.IsDir()
}