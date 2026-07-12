package classfile

// MemberInfo 表示 class 文件中的字段信息或方法信息。
// 字段表和方法表的结构相同，所以共用这个结构体。
type MemberInfo struct {
	// cp 指向当前 class 文件的常量池，用来根据索引查找字段名、方法名、描述符等信息。
	cp ConstantPool

	// accessFlags 是字段或方法的访问标志，例如 public、private、static、final 等。
	accessFlags uint16

	// nameIndex 是字段名或方法名在常量池中的索引。
	nameIndex uint16

	// descriptorIndex 是字段描述符或方法描述符在常量池中的索引。
	descriptorIndex uint16

	// attributes 保存字段或方法上的属性信息，例如 Code、ConstantValue、Exceptions 等。
	attributes []AttributeInfo
}

// 读取字段表或者方法表。
//
// class 文件中字段表和方法表的格式类似：
//
//	fields_count  u2
//	fields        field_info[fields_count]
//
//	methods_count u2
//	methods       method_info[methods_count]
//
// 所以这里先读取一个 uint16，得到字段或方法的数量，
// 然后按照这个数量继续读取每一个 MemberInfo。
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段或者方法。
//
// field_info 和 method_info 的结构相同：
//
//	access_flags     u2
//	name_index       u2
//	descriptor_index u2
//	attributes_count u2
//	attributes       attribute_info[attributes_count]
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

// 从常量池查找字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// 从常量池查找字段或者方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
