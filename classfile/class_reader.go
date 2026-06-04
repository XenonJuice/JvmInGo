package classfile

type ClassReader struct {
	data []byte
}

// u1 = unsigned 1 byte  = 1 字节无符号整数 = 8 bi
func (c *ClassReader) readUint8() uint8 {

}

// u2 = unsigned 2 bytes = 2 字节无符号整数 = 16 bit
func (c *ClassReader) readUint16() uint16 {

}

// u4 = unsigned 4 bytes = 4 字节无符号整数 = 32 bit
func (c *ClassReader) readUint32() uint32 {

}

func (c *ClassReader) readUint64() uint64 {

}

func (c *ClassReader) readUint16s() []uint16 {

}

func (c *ClassReader) readBytes(length uint32) []byte {

}
