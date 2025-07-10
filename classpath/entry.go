package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找并加载class文件
	readClass(className string) ([]byte, Entry, error)
	toString() string
}

// todo
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return nil
	}
	if strings.HasSuffix(path, "*") {
		return nil
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return nil
	}
	return nil
}
