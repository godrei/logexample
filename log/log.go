package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var out io.Writer

// Command ...
func Command(s string) {
	if _, err := write(s, map[string]string{
		"type": "cmd",
		"time": fmt.Sprintf("%s", time.Now().Format("2006:01:02-15:04:05")),
	}); err != nil {
		panic(err)
	}
}

// Info ...
func Info(s string) {
	if _, err := write(s, map[string]string{
		"type": "inf",
		"time": fmt.Sprintf("%s", time.Now().Format("2006:01:02-15:04:05")),
	}); err != nil {
		panic(err)
	}
}

func write(str string, tags map[string]string) (int, error) {
	l := map[string]string{}
	for k, v := range tags {
		l["@"+k] = v
	}
	l["msg"] = str

	b, err := json.Marshal(l)
	if err != nil {
		return 0, err
	}

	return out.Write(append(b, []byte("\n")...))
}

// SetOut ...
func SetOut(w io.Writer) {
	out = w
}

func init() {
	SetOut(os.Stdout)
}
