package logger

import (
	"fmt"
	"io/ioutil"
	"time"
)

func Log(err error) {
	bytes, _ := ioutil.ReadFile("./data/log.txt")

	newRow := fmt.Sprintf("%v %s",
		time.Now().Format(time.RFC1123),
		err.Error(),
	)

	prev := string(bytes)
	output := fmt.Sprintf("%s", newRow)

	if len(prev) > 0 {
		output = fmt.Sprintf("%s\n%s", prev, newRow)
	}

	ioutil.WriteFile("./data/log.txt", []byte(output), 0644)
}
