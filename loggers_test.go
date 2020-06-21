package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/jameOne/psm/logger/action"
	"github.com/jameOne/psm/logger/error"
	"github.com/jameOne/psm/logger/input"
	"github.com/jameOne/psm/logger/output"
	"github.com/jameOne/psm/logger/update"
)

func TestLoggers(t *testing.T) {

	var str bytes.Buffer

	act := action.New("test action")
	logFileName := "test.log"

	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%v:", err)
	}
	defer f.Close()

	w := io.MultiWriter(&str, f)
	log.SetOutput(w)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	Action(act)

	fmt.Println(str.String())

	if len(str.String()) == 0 {
		t.Error(
			"log output length equal zero",
		)
	}

	str.Reset()

	err = error.New("test error")
	Error(err)

	fmt.Println(str.String())

	if len(str.String()) == 0 {
		t.Error(
			"log output length equal zero",
		)
	}

	str.Reset()

	inp := input.New("test input")
	Input(inp)

	fmt.Println(str.String())

	if len(str.String()) == 0 {
		t.Error(
			"log output length equal zero",
		)
	}

	str.Reset()

	out := output.New("test output")
	Output(out)

	fmt.Println(str.String())

	if len(str.String()) == 0 {
		t.Error(
			"log output length equal zero",
		)
	}

	str.Reset()

	upd := update.New("test update")
	Update(upd)

	fmt.Println(str.String())

	if len(str.String()) == 0 {
		t.Error(
			"log output length equal zero",
		)
	}

	str.Reset()

	exampleLogFileContents := "A:2020/03/31 16:33:45.130162 loggers.go:59: test action:\nE:2020/03/31 16:33:45.130296 loggers.go:93: test error:\nI:2020/03/31 16:33:45.130308 loggers.go:12: test input:\nO:2020/03/31 16:33:45.130321 loggers.go:16: test output:\nU:2020/03/31 16:33:45.130329 loggers.go:19: test update:\n"

	f, err = os.Open(logFileName)
	if err != nil {
		t.Error(
			"error opening log file",
		)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(
			"error reading log file",
		)
	}

	if len(string(b)) != len(exampleLogFileContents) {
		t.Error(
			"logger written file is not expected length",
		)
	}

	err = os.Remove(logFileName)
	if err != nil {
		t.Error(
			"error removing test.log file",
		)
	}
}
