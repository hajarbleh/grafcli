package utility

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ListFiles lists all files (excluding directories) in given dirPth. Returns the
// path of the files.
func ListFiles(dirPth string) ([]string, error) {
	infos, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, info := range infos {
		if !info.IsDir() {
			ret = append(ret, filepath.Join(dirPth, info.Name()))
		}
	}
	return ret, nil
}

// SanitizeFilename lower case & replace non-alphanumeric characters with
// underscores from given string.
func SanitizeFilename(name string) string {
	name = regexp.MustCompile(`\W+`).ReplaceAllString(name, "_")
	return strings.ToLower(name)
}

// WriteFile writes data into filePth. Will also create the directories first if
// required.
func WriteFile(data []byte, filePth string) error {
	if err := os.MkdirAll(filepath.Dir(filePth), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(filePth, data, 0644)
}
