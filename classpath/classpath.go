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
	return c.userClassPath.toString()
}

// ReadClass 依次从boot ext user类路径搜索class文件
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + classFileSuffix

	byteData, entry, err := c.bootClassPath.readClass(className)
	if err == nil {
		return byteData, entry, err
	}
	byteData, entry, err = c.extClassPath.readClass(className)
	if err == nil {
		return byteData, entry, err
	}
	byteData, entry, err = c.userClassPath.readClass(className)
	if err == nil {
		return byteData, entry, err
	}
	return nil, nil, err
}

// parseBootAndExtClassPath 基于 JRE 目录设置启动和扩展类路径条目。
func (c *Classpath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibDir := filepath.Join(jreDir, jreLibDirName, wildcardClasspathSuffix)
	c.bootClassPath = newWildcardEntry(jreLibDir)
	// jre/lib/ext/*
	jreExtDir := filepath.Join(jreDir, jreLibDirName, jreExtDirName, wildcardClasspathSuffix)
	c.extClassPath = newWildcardEntry(jreExtDir)
}

func (c *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == blank {
		cpOption = defaultUserClasspath
	}
	c.userClassPath = newDirEntry(cpOption)
}

// 优先使用用户输入的-Xjre选项作为jre目录
// 如果没有输入该选项，则在当前目录下寻找jre目录
// 如果找不到 使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != blank && exists(jreOption) {
		return jreOption
	}
	if exists(defaultJreDir) {
		return defaultJreDir
	}
	if javaHome := os.Getenv(javaHomeEnv); javaHome != blank {
		return filepath.Join(javaHome, jreDirName)
	}
	panic(jreDirNotFoundMessage)
}

// 用于判断路径是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}
