package services

import (
	"fmt"
	"os"
	"strings"
)

//GetGopath returns an error if there is no GOPATH variable set
func GetGopath() (string, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		home := os.Getenv("HOME")
		if home == "" {
			return goPath, fmt.Errorf("Expecting a GOPATH")
		}
		err := os.MkdirAll(home+"/go", 0755)
		if err != nil {
			return "", err
		}
		return home + "/go", nil
	}
	return goPath, nil
	//t.Log("goPath is...",goPath)
}

func CreateProjectFolders(folders ...string) (error){
    // targetRoot := goPath + "/src/github.com/gotstago/qor-projects/" + projectName
    // mask := 0755
		for _, targetFolder := range folders {
            targetFolder := targetFolder
			if err := os.MkdirAll(targetFolder,0755); err != nil {
				return  err
			}
		}
    return nil
}

func RemoveProject(projectFolder string) error{
    return os.RemoveAll(projectFolder)
}

func FormatFileName(typeName string, targetType string) string {
	return fmt.Sprintf("%s_%s.go", strings.ToLower(typeName), strings.ToLower(targetType))
}
