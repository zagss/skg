package test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	path := "path/to/your/file"
	fileName, fileExt := filepath.Split(path)
	fmt.Println("File name:", fileName)
	fmt.Println("File extension:", fileExt)
	parts := strings.SplitN(path, ".", 2)
	fmt.Println("parts:", parts)
}
