package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

// ZipEntry 表示zip或jar形式的类路径
type ZipEntry struct {
	absoluteZipPath string
	zipReadCloser   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	//  将path转换为绝对路径
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absoluteZipPath: abs, zipReadCloser: nil}
}

func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 如果句柄不存在，则首先获取文件句柄
	if (*z).zipReadCloser == nil {
		err := (*z).openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	// 根据classname查找classFIle
	classFile := (*z).findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found :" + className)
	}
	// 开始获取classFile内容
	data, err := readClassInternal(classFile)
	return data, z, err

}

func (z *ZipEntry) openJar() error {
	readCloser, err := zip.OpenReader(z.absoluteZipPath)
	if err == nil {
		(*z).zipReadCloser = readCloser
	}
	return err
}

func (z *ZipEntry) findClass(className string) *zip.File {
	// 遍历zip文件中每一个文件
	for _, file := range (*z).zipReadCloser.File {
		if file.Name == className {
			return file
		}
	}
	// 未找到则返回null
	return nil
}

func readClassInternal(fileInZip *zip.File) (data []byte, err error) {
	var rc io.ReadCloser
	rc, err = fileInZip.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		var closeErr error = rc.Close()
		if closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()
	data, err = io.ReadAll(rc)
	return
}

func (z *ZipEntry) toString() string {
	return z.absoluteZipPath
}
