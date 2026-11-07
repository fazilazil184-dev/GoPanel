package views

import (
	"fmt"
	"os"
	"testing"
)

func TestFilePathJoin(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Println("Working dir:", wd)

	ok, err := ExistPath("auth")
	fmt.Println("=> HASIL:", ok, err)
}
