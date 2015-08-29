package freader_test

import (
    . "github.com/zieckey/gohello/freader"
    "testing"
    "github.com/bmizerany/assert"
    "runtime"
    "path/filepath"
)

func TestGetAbsPath(t *testing.T) {
    var dir string
    if runtime.GOOS == "windows" {
        dir = "c:/"
        assert.Equal(t, filepath.IsAbs(dir), true)
    } else {
        dir = "/etc"
        assert.Equal(t, filepath.IsAbs(dir), true)
    }
    assert.Equal(t, GetAbsPath(dir), dir)
}