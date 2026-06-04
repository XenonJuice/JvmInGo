package classfile

import "encoding/binary"

type ClassReader struct {
	data  []byte
	index int
}

// u1 = unsigned 1 byte  = 1 字节无符号整数 = 8 bi
func (c *ClassReader) readUint8() uint8 {
	var value uint8
	value = c.data[c.index]
	c.index++
	return value
}

// u2 = unsigned 2 bytes = 2 字节无符号整数 = 16 bit
func (c *ClassReader) readUint16() uint16 {
	var value uint16
	// 这里练习一种手动取出的方式
	//value := uint16(c.readUint8()) << 8
	//value |= uint16(c.readUint8())
	value = binary.BigEndian.Uint16(c.data[c.index:])
	c.index += 2
	return value
}

// u4 = unsigned 4 bytes = 4 字节无符号整数 = 32 bit
func (c *ClassReader) readUint32() uint32 {
	var value uint32
	value = binary.BigEndian.Uint32(c.data[c.index:])
	c.index += 4
	return value
}

func (c *ClassReader) readUint64() uint64 {
	var value uint64
	value = binary.BigEndian.Uint64(c.data[c.index:])
	c.index += 8
	return value
}

// 读取uint16表，表的大小由开头的uint16数据指出
func (c *ClassReader) readUint16s() []uint16 {
	var size uint16
	size = c.readUint16()

	var values []uint16
	values = make([]uint16, size)

	var i int
	for i = 0; i < len(values); i++ {
		values[i] = c.readUint16()
	}
	return values
}

// 用于读取指定数量的字节
func (c *ClassReader) readBytes(length uint32) []byte {
	var bytes []byte
	bytes = c.data[c.index : c.index+int(length)]
	c.index += int(length)
	return bytes
}
