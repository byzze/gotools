package filetool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// 将文件夹下所有文件中的的main方法都改为文件名字，因为使用vscode没有发现批量处理的功能，所有自己写了一个满足需求
func BatchUpdateMainToFileName(dirpath string) error {
	_, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		return err
	}
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		logrus.Error(err)
		return err
	}
	for _, uv := range files {
		if uv.IsDir() == true {
			continue
		}
		ext := filepath.Ext(uv.Name())
		fileName := strings.TrimSuffix(uv.Name(), ext)
		filePath := dirpath + "/" + uv.Name()
		dat, err := os.ReadFile(filePath)
		if err != nil {
			logrus.Error(err)
			return err
		}
		newSt := strings.ReplaceAll(string(dat), "func main", "func "+fileName)
		err = os.WriteFile(filePath, []byte(newSt), 0644)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}
