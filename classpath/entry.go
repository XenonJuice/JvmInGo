package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

const (
	blank                   = ""
	wildcardClasspathSuffix = "*"
	jarSuffix               = ".jar"
	upperJarSuffix          = ".JAR"
	zipSuffix               = ".zip"
	upperZipSuffix          = ".ZIP"
	classNotFoundMessage    = "class not found :"
)

type Entry interface {
	// 寻找并加载class文件
	readClass(className string) ([]byte, Entry, error)
	toString() string
}

// todo
func newEntry(path string) Entry {
	// 如果path中包含路径列表分隔符，说明它是由多个classpath组成的路径列表。
	// example on Unix/macOS: target/classes:lib/gson-2.10.1.jar:lib/slf4j-api-2.0.13.jar
	// example on Windows: target\classes;lib\gson-2.10.1.jar;lib\slf4j-api-2.0.13.jar
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	// 如果path以*结尾，说明它是通配符classpath，用来加载某个目录下的所有jar包。
	// example: lib/*
	if strings.HasSuffix(path, wildcardClasspathSuffix) {
		return newWildcardEntry(path)
	}

	// 如果path以.jar或.zip结尾，说明它是压缩包形式的classpath。
	// example: lib/gson-2.10.1.jar
	if strings.HasSuffix(path, jarSuffix) ||
		strings.HasSuffix(path, upperJarSuffix) ||
		strings.HasSuffix(path, zipSuffix) ||
		strings.HasSuffix(path, upperZipSuffix) {
		return newZipEntry(path)
	}

	// 其它情况按普通目录classpath处理。
	// example: target/classes
	return newDirEntry(path)
}
