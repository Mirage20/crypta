package io

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadInput(file string) ([]byte, error) {
	if len(file) > 0 {
		return ioutil.ReadFile(file)
	} else {
		return ioutil.ReadAll(os.Stdin)
	}
}

func WriteOutput(data []byte, file string) error {
	if len(file) > 0 {
		return ioutil.WriteFile(file, data, 0644)
	} else {
		fmt.Fprintf(os.Stdout, "%s", data)
		return nil
	}
}
