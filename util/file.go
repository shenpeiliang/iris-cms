package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type File struct {
}

//读取json文件
func (f File) ReadJsonFile(filename string) (*File, error) {
	ret := &File{}

	fileReader, err := os.Open(filename)
	if err != nil {
		return ret, err
	}
	defer fileReader.Close()

	fr := io.Reader(fileReader)

	if err = json.NewDecoder(fr).Decode(ret); err != nil {
		return ret, err
	}

	return ret, nil
}

//利用ioutil将file直接读取到[]byte中
func (f File) ReadFileContent(filename string) (string, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fileReader.Close()

	content, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

//一次性读取文件全部内容
func (f File) ReadFileContent1(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

//先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
func (f File) ReadFileContent2(filename string) (string, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fileReader.Close()

	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := fileReader.Read(buf)

		if err != nil && err != io.EOF {
			return "", err
		}

		//说明读取结束
		if n == 0 {
			break
		}

		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	return string(chunk), nil
}
