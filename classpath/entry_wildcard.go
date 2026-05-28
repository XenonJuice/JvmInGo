package classpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// newWildcardEntry 解析一个通配类路径，
// 遍历基础目录， 并将所有.jar文件聚合为复合条目。
func newWildcardEntry(path string) CompositeEntry {
	// remove the trailing * from a wildcard classpath to get the base directory
	baseDir := path[:len(path)-1]
	// Entry[] as a return value
	var compositeEntry CompositeEntry
	// the file name
	var name string

	dirs, err := os.ReadDir(baseDir)

	if err != nil {
		fmt.Printf("Error: os.ReadDir() error when reading directory from %q: %v\n", baseDir, err)
		return nil
	}
	// 遍历整个目录，判断每个目录的具体类型信息
	for _, dir := range dirs {

		// 跳过子目录
		if dir.IsDir() {
			continue
		}

		name = dir.Name()
		// 抽取jar文件
		if strings.HasSuffix(name, jarSuffix) ||
			strings.HasSuffix(name, upperJarSuffix) {
			jarPath := filepath.Join(baseDir, name)
			compositeEntry = append(compositeEntry, newZipEntry(jarPath))
		}

	}

	// 这里练习了一下闭包写法，不过我感觉没太必要
	//walkFn := func(filePath string, info os.FileInfo, err error) error {
	//	if err != nil {
	//		return err
	//	}
	//
	//	// 当出现例如扫描lib/目录下jar文件但是扫描出新的目录时，
	//	// 例如 lib/sub/c.jar，则跳过对子目录sub的扫描
	//	if info.IsDir() && filePath != baseDir {
	//		return filepath.SkipDir
	//	}
	//
	//	if strings.HasSuffix(filePath, ".jar") ||
	//		strings.HasSuffix(filePath, ".JAR") {
	//
	//		jarEntry := newZipEntry(filePath)
	//		compositeEntry = append(compositeEntry, jarEntry)
	//	}
	//	return nil
	//}
	//
	//err := filepath.Walk(baseDir, walkFn)
	//if err != nil {
	//	return nil
	//}
	return compositeEntry
}
