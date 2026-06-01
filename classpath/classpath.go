package classpath

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	var cp *Classpath = new(Classpath)
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (c *Classpath) String() string {
	return ""
}

func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}

func (c *Classpath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibDir := filepath.Join(jreDir, "lib", wildcardClasspathSuffix)
	c.bootClassPath = newWildcardEntry(jreLibDir)
	// jre/lib/ext/*
	jreExtDir := filepath.Join(jreDir, "lib", "ext", wildcardClasspathSuffix)
	c.extClassPath = newWildcardEntry(jreExtDir)
}

func (c *Classpath) parseUserClassPath(cpOption string) {
}

// 优先使用用户输入的-Xjre选项作为jre目录
// 如果没有输入该选项，则在当前目录下寻找jre目录
// 如果找不到 使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != blank && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != blank {
		return filepath.Join(javaHome, "jre")
	}
	panic("Can not find jre directory, please input -Xjre option or set JAVA_HOME environment variable")
}

// 用于判断路径是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}
