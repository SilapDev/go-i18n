package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetRootProjectPath() string {

	f, _ := os.Getwd()
	root := fmt.Sprintf("%s/%s", filepath.Dir(f), filepath.Base(f))

	return root

}
