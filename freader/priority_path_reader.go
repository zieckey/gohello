package freader
import (
    "path/filepath"
    "strconv"
)


type PriorityPathReader struct {
    dir string
    priorityLevel int
    readers map[string/*path*/]IPathReader
}

func NewPriorityPathReader(dir string) (ppr *PriorityPathReader, err error) {
    ppr = &PriorityPathReader{}
    ppr.dir = dir
    ppr.priorityLevel = *priorityLevel
    ppr.readers = make(map[string/*path*/]IPathReader)

    if ppr.priorityLevel <= 0 {
        p := filepath.Join(dir, "0")
        ppr.readers[p] = NewPathReader(p)
        ppr.readers[dir] = NewPathReader(dir)
    } else {
        for i := 0; i < ppr.priorityLevel; i++ {
            p := filepath.Join(dir, strconv.Itoa(i))
            ppr.readers[p] = NewPathReader(p)
        }
    }

    return ppr, nil
}