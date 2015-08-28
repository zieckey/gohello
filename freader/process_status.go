package freader
import (
    "os"
    "bufio"
    "io"
    "time"
    "fmt"
    "strings"
    "sort"
    "strconv"
)

type FileProcessingTime struct {
    Start time.Time
    End time.Time
}

type ProcessStatus struct {
    processedFiles map[string]FileProcessingTime    // The processed files and the time when starting to process and end

    // The content format of status file :
    //  It is a text file. Every line represents a processed file.
    //  The line has 3 part
    //      1. start processing date time
    //      2. end of processing date time
    //      3. the name of the file
    // For example: 2015/08/28-20:42:12.1231 2015/08/28-20:43:23.3123 /home/s/data/log/xxx.log
    statusFile string       // The path of the status file which used to store the status information of all processed files
    statusFileFp *os.File   // The file pointer to the status file

}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
    _, err := os.Stat(path)
    return err == nil || os.IsExist(err)
}

func NewProcessStatus(statusFile string) (ps *ProcessStatus, err error) {
    ps = &ProcessStatus{}
    ps.statusFile = statusFile
    ps.processedFiles = make(map[string]FileProcessingTime)

    if IsExist(statusFile) {
        ps.statusFileFp, err = os.OpenFile(statusFile, os.O_RDWR | os.O_APPEND, 0755)
        if err != nil {
            fmt.Printf("open status file <%v> failed : %v\n", err.Error())
            return nil, err
        }
        if err = ps.parse(); err != nil {
            return nil, err
        }
    } else {
        ps.statusFileFp, err = os.OpenFile(statusFile, os.O_CREATE | os.O_RDWR, 0755)
        if err != nil {
            fmt.Printf("open status file <%v> failed : %v\n", err.Error())
            return nil, err
        }
    }

    return ps, nil
}

func (ps *ProcessStatus) GetProcessedFiles() map[string]FileProcessingTime {
    return ps.processedFiles
}

func (ps *ProcessStatus) parse() error {
    r := bufio.NewReader(ps.statusFileFp)
    for {
        line, err := r.ReadString('\n')
        if len(line) > 0 {
            line = strings.TrimSpace(line)
            var start,end,path string
            fmt.Sscanf(line, "%s %s %s", &start, &end, &path)
            var t FileProcessingTime
            t.Start, err = time.Parse("2006/01/02-15:04:05.9999", start)
            if err != nil {
                return fmt.Errorf("ERROR line <%v> %v", line, err.Error())
            }
            t.End, err = time.Parse("2006/01/02-15:04:05.9999", end)
            if err != nil {
                return fmt.Errorf("ERROR line <%v> %v", line, err.Error())
            }
            if len(path) == 0 {
                return fmt.Errorf("ERROR line <%v>, path empty", line)
            }
            ps.processedFiles[path] = t
        }
        if err == io.EOF {
            break
        }
    }
    return nil
}

func (ps *ProcessStatus) OnFileProcessingFinished(path string, startProcessing time.Time) {
    var t FileProcessingTime
    t.Start = startProcessing
    t.End = time.Now()
    ps.processedFiles[path] = t

    w := bufio.NewWriter(ps.statusFileFp)
    w.WriteString(t.Start.Format("2006/01/02-15:04:05.9999 "))
    w.WriteString(t.End.Format("2006/01/02-15:04:05.9999 "))
    w.WriteString(path)
    w.WriteString("\n")
    w.Flush()
}

func (ps *ProcessStatus) OnFileDeleted(path string) {
    delete(ps.processedFiles, path)
}

func (ps *ProcessStatus) Close()  {
    ps.statusFileFp.Close()
    tmp := ps.statusFile + ".tmp." + strconv.FormatInt(time.Now().UnixNano(), 10)
    if err := ps.saveAllTo(tmp); err != nil {
        panic(err.Error())
    } // flush all data to files
}

type StringArray []string

func (ss StringArray) Len() int {
    return len(ss)
}
func (ss StringArray) Less(i, j int) bool {
    return ss[i] < ss[j]
}

func (ss StringArray) Swap(i, j int) {
    ss[i], ss[j] = ss[j], ss[i]
}


func (ps *ProcessStatus) saveAllTo(tmpFilePath string) error {
    fp, err := os.OpenFile(tmpFilePath, os.O_CREATE | os.O_RDWR, 0755)
    if err != nil {
        return err
    }
    defer fp.Close()
    var files StringArray
    for k, _ := range ps.processedFiles {
        files = append(files, k)
    }
    sort.Sort(files)

    w := bufio.NewWriter(fp)
    for _, f := range files {
        if t, ok := ps.processedFiles[f]; ok {
            w.WriteString(t.Start.Format("2006/01/02-15:04:05.9999 "))
            w.WriteString(t.End.Format("2006/01/02-15:04:05.9999 "))
            w.WriteString(f)
            w.WriteString("\n")
        }
    }
    w.Flush()

    bak := ps.statusFile + ".bak." + strconv.FormatInt(time.Now().UnixNano(), 10)
    err = os.Rename(ps.statusFile, bak)
    if err != nil {
        return fmt.Errorf("os.Rename <%v> to <%v> failed : %v", ps.statusFile, bak, err.Error())
    }
    os.Rename(tmpFilePath, ps.statusFile)
    if err != nil {
        return fmt.Errorf("os.Rename <%v> to <%v> failed : %v", tmpFilePath, ps.statusFile, err.Error())
    }

    return nil
}
