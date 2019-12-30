package saver

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Saver interface {
	Save(url string, reader io.Reader) error
}

type saver struct {}

func NewSaver() *saver {
	return &saver{}
}

func (s *saver) Save (url string, reader io.Reader) error {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	name := strings.TrimPrefix(url, "http://")
	filename := fmt.Sprintf("store/%s.html", name)
	err = ioutil.WriteFile(filename, bytes,  0644)
	if err != nil {
		return err
	}
	return nil
}