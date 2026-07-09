package classfile

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

// MemberInfo 表示字段或方法的基础信息。
// 字段和方法在 class 文件中的结构非常相似，所以可以共用这个类型。
type MemberInfo struct{}

// AttributeInfo 表示 class 文件、字段或方法上的属性信息。
type AttributeInfo interface{}
