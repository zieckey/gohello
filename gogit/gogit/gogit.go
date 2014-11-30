package gogit

import (
	"os"
	"path"
	"runtime"
	"strings"
)

var PathSeperator = ":"

func init() {
	if runtime.GOOS == "windows" {
		PathSeperator = ";"
	}
}


// return $GOPATH/src
func getSrcPath() string {
	gopath := os.Getenv("GOPATH")
	paths := strings.Split(gopath, PathSeperator)
	gosrcpath := gopath
	if len(paths) > 1 {
		gosrcpath = paths[len(paths)-1]
	}

	return path.Join(gosrcpath, "src")
}

func getAuthor(url string) string {
	prefix := "github.com/"
	pos := strings.Index(url, prefix)
	if pos < 0 {
		return ""
	}

	postfix := url[pos+len(prefix):]
	pos = strings.IndexByte(postfix, '/')
	if pos < 0 {
		return ""
	}

	return postfix[:pos]
}


// If project = https://github.com/zieckey/goini
// and GOPATH=d:\go;e:\goworkspace\go
// this function will return "e:\goworkspace\go\src\github\zieckey"
func GetProjectParentPath(url string) string {
	gosrcpath := getSrcPath()
	author := getAuthor(url)
	if len(author) == 0 {
		println("The input argument format ERROR. CANNOT found author")
		return ""
	}

	s := path.Join(gosrcpath, "github.com", author)
	return strings.Replace(s, "\\", "/", -1)
}
