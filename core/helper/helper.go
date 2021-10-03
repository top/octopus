package helper

import (
	"os"
	"reflect"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func InitLogDir(dir string) error {
	exist, err := PathExists(dir)
	if err != nil {
		return err
	}
	if !exist {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func IsStructureEmpty(x, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}
