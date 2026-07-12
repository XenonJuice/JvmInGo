package classfile

import "fmt"

type ClassFile struct {
	// magic 是 class 文件的魔数，固定为 0xCAFEBABE。
	// 这个值通常只在读取时校验，不需要保存在 ClassFile 结构体里。
	// magic uint32

	// minorVersion 表示 class 文件的次版本号。
	minorVersion uint16

	// majorVersion 表示 class 文件的主版本号，用来判断 class 文件对应的 Java 版本。
	majorVersion uint16

	// constantPool 是常量池，保存类名、字段名、方法名、字符串字面量等常量信息。
	constantPool ConstantPool

	// accessFlags 是类的访问标志，例如 public、final、interface、abstract 等。
	accessFlags uint16

	// thisClass 是当前类在常量池中的索引。
	thisClass uint16

	// superClass 是父类在常量池中的索引，java.lang.Object 的父类索引为 0。
	superClass uint16

	// interfaces 保存当前类实现的接口在常量池中的索引列表。
	interfaces []uint16

	// fields 保存类或接口中声明的字段信息。
	fields []*MemberInfo

	// methods 保存类或接口中声明的方法信息。
	methods []*MemberInfo

	// attributes 保存 class 文件级别的属性信息。
	attributes []AttributeInfo
}

// ConstantPool 表示 class 文件中的常量池。
// 这里只先定义类型，具体结构会在实现常量池解析时补充。
type ConstantPool []ConstantInfo

// ConstantInfo 表示常量池中的一个常量项。
type ConstantInfo interface{}

// AttributeInfo 表示 class 文件、字段或方法上的属性信息。
type AttributeInfo interface{}

// 将[]byte解析成ClassFile结构体
func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("panic recovered: %v", r)
			}
		}
	}()
	classReader := &ClassReader{data: classData}
	classFile = &ClassFile{}
	classFile.read(classReader)
	return
}

func (classFile *ClassFile) read(reader *ClassReader) {
	classFile.readAndCheckMagic(reader)
	classFile.readAndCheckVersion(reader)
	classFile.constantPool = readConstantPool(reader)
	classFile.accessFlags = reader.readUint16()
	classFile.thisClass = reader.readUint16()
	classFile.superClass = reader.readUint16()
	classFile.interfaces = reader.readUint16s()
	classFile.fields = readMembers(reader, classFile.constantPool)
	classFile.methods = readMembers(reader, classFile.constantPool)
	classFile.attributes = readAttributes(reader, classFile.constantPool)
}

// 检查魔数
func (classFile *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		// todo 由于目前暂时无法抛出异常，就先panic
		panic("Lava.lang.ClassFormatError: magic number in class file is incorrect")
	}
}

// 检查版本号 Java8 支持45.0-52.0的class文件
func (classFile *ClassFile) readAndCheckVersion(reader *ClassReader) {
	classFile.minorVersion = reader.readUint16()
	classFile.majorVersion = reader.readUint16()
	switch classFile.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if classFile.minorVersion == 0 {
			return
		}
		// todo 由于目前暂时无法抛出异常，就先panic
		panic("Lava.lang.ClassFormatError: unsupported major version")

	}
}

// getter
func (classFile *ClassFile) MajorVersion() uint16 {
	return classFile.majorVersion
}

func (classFile *ClassFile) MinorVersion() uint16 {
	return classFile.minorVersion
}

func (classFile *ClassFile) ConstantPool() ConstantPool {
	return classFile.constantPool
}

func (classFile *ClassFile) AccessFlags() uint16 {
	return classFile.accessFlags
}

func (classFile *ClassFile) ThisClass() uint16 {
	return classFile.thisClass
}

func (classFile *ClassFile) ClassName() string {
	return classFile.constantPool.getClassName(classFile.thisClass)
}

func (classFile *ClassFile) SuperClass() uint16 {
	return classFile.superClass
}

// 从常量池中查找接口名
func (classFile *ClassFile) SuperClassName() string {
	if classFile.superClass > 0 {
		return classFile.constantPool.getClassName(classFile.superClass)
	}
	return ""
}

// 从常量池中查找接口名
func (classFile *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(classFile.interfaces))
	for i, constantPoolIndex := range classFile.interfaces {
		interfaceNames[i] = classFile.constantPool.getClassName(constantPoolIndex)
	}
	return interfaceNames
}

func (classFile *ClassFile) Interfaces() []uint16 {
	return classFile.interfaces
}

func (classFile *ClassFile) Fields() []*MemberInfo {
	return classFile.fields
}

func (classFile *ClassFile) Methods() []*MemberInfo {
	return classFile.methods
}

func (classFile *ClassFile) Attributes() []AttributeInfo {
	return classFile.attributes
}
