package internal

import (
	"fmt"
	"math/rand"
	"os"
)

// This is what we would be responsible for writing to a file
func SaveData1(path string, data []byte) error {
	//cool part about go is that you can have multiple return types here specifically an error and a file
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return fp.Sync()
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()
	if _, err = fp.Write(data); err != nil { // 1. save to the temporary file
		return err
	}
	if err = fp.Sync(); err != nil { //2. fsync
		return err
	}
	err = os.Rename(tmp, path) // 3. replace the temp file with
	return err
}
