package utility

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const inputLimit = 50

type Interaction struct {
	Reader io.Reader
}

func (i *Interaction) AskUserInput(query string) (string, error) {
	fmt.Println(query + ": ")
	buf := make([]byte, inputLimit)
	l, err := i.Reader.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if l > inputLimit {
		fmt.Println("Error: character limit(50) exceeded")
		return "", errors.New("Error: character limit(50) exceeded")
	}
	return strings.TrimSpace(string(buf[:l])), nil
}
