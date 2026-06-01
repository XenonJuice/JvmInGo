package classpath

import "os"

// 路径列表分隔符
const pathListSeparator = string(os.PathListSeparator)

const (
	blank                   = ""
	wildcardClasspathSuffix = "*"
	jarSuffix               = ".jar"
	upperJarSuffix          = ".JAR"
	zipSuffix               = ".zip"
	upperZipSuffix          = ".ZIP"
	classFileSuffix         = ".class"
	classNotFoundMessage    = "class not found :"
	defaultUserClasspath    = "."
	defaultJreDir           = "./jre"
	javaHomeEnv             = "JAVA_HOME"
	jreDirName              = "jre"
	jreLibDirName           = "lib"
	jreExtDirName           = "ext"
	jreDirNotFoundMessage   = "Can not find jre directory, please input -Xjre option or set JAVA_HOME environment variable"
)
