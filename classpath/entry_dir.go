package classpath

import (
	"os"
	"path/filepath"
)

// DirEntry DirEntry，表示目录形式的类路径
type DirEntry struct {
	absoluteDirPath string
}

func newDirEntry(path string) *DirEntry {
	// 将path转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 返回一个DirEntry结构体的指针
	return &DirEntry{absoluteDirPath: absDir}
}

// readClass 在目录中找到，读取和返回类名称指定的类文件的内容。
// 它将类数据作为字节切片返回，如果无法读取文件，则返回错误。
func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absoluteDirPath, className)
	data, err := os.ReadFile(fileName)
	return data, d, err
}

// toString 返回绝对目录路径作为字符串表示的 DirEntry.
func (d *DirEntry) toString() string {
	return d.absoluteDirPath
}
