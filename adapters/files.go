package adapters

import (
	"fmt"
	"io/ioutil"
)

const (
	filePath = "/files/users/"
)

// FileAdapter has a good shape
type FileAdapter struct {
}

// NewFileAdapter returns the file adapter object
func NewFileAdapter() *FileAdapter {
	return &FileAdapter{}
}

// CreateFile creates a blank file
func (fA *FileAdapter) CreateFile(id int) error {
	// f, err := os.Create(filePath + fmt.Sprint(id))
	// if err != nil {
	// 	return err
	// }

	// defer f.Close()
	// return nil
	return ioutil.WriteFile(filePath+fmt.Sprint(id), []byte("Welcome"), 0644)
}

// Write writes to a file
// func Write(id int, data interface{}) error {

// }
