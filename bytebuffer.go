package bytebuffer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

const (
	//Big ...
	Big = "big"
	//Little ...
	Little = "little"
)

//Buffer ...
type Buffer struct {
	packetBuffer bytes.Buffer
	Endian       string
}

//NewBuffer ...
func NewBuffer(types string) *Buffer {
	return &Buffer{
		Endian: types,
	}
}

//PutShort ...
func (obj *Buffer) PutShort(value int) {

	if obj.Endian != Big && obj.Endian != Little {
		fmt.Println("Invalid endianness, must be big or little")
		return
	}

	var buff []byte = make([]byte, 2)

	if obj.Endian == Big {
		binary.BigEndian.PutUint16(buff, uint16(value))
	} else if obj.Endian == Little {
		binary.LittleEndian.PutUint16(buff, uint16(value))
	}

	obj.packetBuffer.Write(buff)
}

//GetShort ...
func (obj *Buffer) GetShort() uint16 {

	var tempBuff = obj.packetBuffer.Bytes()
	var shortValue []byte = tempBuff[:2]
	var restValue []byte = tempBuff[2:]
	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer

	if obj.Endian == Big {
		return binary.BigEndian.Uint16(shortValue)
	}

	return binary.LittleEndian.Uint16(shortValue)
}

//PutInt ...
func (obj *Buffer) PutInt(value int) {

	if obj.Endian != Big && obj.Endian != Little {
		fmt.Println("Invalid endianness, must be big or little")
		return
	}

	var buff []byte = make([]byte, 4)

	if obj.Endian == Big {
		binary.BigEndian.PutUint32(buff, uint32(value))
	} else if obj.Endian == Little {
		binary.LittleEndian.PutUint32(buff, uint32(value))
	}

	obj.packetBuffer.Write(buff)
}

//GetInt ...
func (obj *Buffer) GetInt() uint32 {

	var tempBuff = obj.packetBuffer.Bytes()
	var intValue = tempBuff[:4]
	var restValue = tempBuff[4:]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer

	if obj.Endian == Big {
		return binary.BigEndian.Uint32(intValue)
	}

	return binary.LittleEndian.Uint32(intValue)
}

//PutLong ...
func (obj *Buffer) PutLong(value int) {

	if obj.Endian != Big && obj.Endian != Little {
		fmt.Println("Invalid endianness, must be big or little")
		return
	}

	var buff []byte = make([]byte, 8)

	if obj.Endian == Big {
		binary.BigEndian.PutUint64(buff, uint64(value))
	} else if obj.Endian == Little {
		binary.LittleEndian.PutUint64(buff, uint64(value))
	}
	obj.packetBuffer.Write(buff)
}

//GetLong ...
func (obj *Buffer) GetLong() uint64 {

	var tempBuff = obj.packetBuffer.Bytes()
	var longValue = tempBuff[:8]
	var restValue = tempBuff[8:]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer

	if obj.Endian == Big {
		return binary.BigEndian.Uint64(longValue)
	}

	return binary.LittleEndian.Uint64(longValue)

}

//PutFloat ...
func (obj *Buffer) PutFloat(value float32) {

	if obj.Endian != Big && obj.Endian != Little {
		fmt.Println("Invalid endianness, must be big or little")
		return
	}

	var bits uint32 = math.Float32bits(value)
	var buff []byte = make([]byte, 4)

	if obj.Endian == Big {
		binary.BigEndian.PutUint32(buff, bits)
	} else if obj.Endian == Little {
		binary.LittleEndian.PutUint32(buff, bits)
	}

	obj.packetBuffer.Write(buff)
}

//GetFloat ...
func (obj *Buffer) GetFloat() float32 {

	var tempBuff = obj.packetBuffer.Bytes()
	var floatValue = tempBuff[:4]
	var restValue = tempBuff[4:]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer

	if obj.Endian == Big {
		return math.Float32frombits(binary.BigEndian.Uint32(floatValue))
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(floatValue))
}

//PutDouble ...
func (obj *Buffer) PutDouble(value float64) {

	if obj.Endian != Big && obj.Endian != Little {
		fmt.Println("Invalid endianness, must be big or little")
		return
	}

	var bits uint64 = math.Float64bits(value)
	var buff []byte = make([]byte, 8)

	if obj.Endian == Big {
		binary.BigEndian.PutUint64(buff, bits)
	} else if obj.Endian == Little {
		binary.LittleEndian.PutUint64(buff, bits)
	}

	obj.packetBuffer.Write(buff)
}

//GetDouble ...
func (obj *Buffer) GetDouble() float64 {

	var tempBuff = obj.packetBuffer.Bytes()
	var doubleValue = tempBuff[:8]
	var restValue = tempBuff[8:]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer

	if obj.Endian == Big {
		return math.Float64frombits(binary.BigEndian.Uint64(doubleValue))
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(doubleValue))
}

//Put ...
func (obj *Buffer) Put(value []byte) {
	obj.packetBuffer.Write(value)
}

//PutByte ...
func (obj *Buffer) PutByte(value byte) {
	var tempByte []byte
	tempByte = append(tempByte, value)
	obj.packetBuffer.Write(tempByte)
}

//Get ...
func (obj *Buffer) Get(size int) []byte {

	var tempBuff = obj.packetBuffer.Bytes()
	var value = tempBuff[:size]
	var restValue = tempBuff[size:]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer
	return value
}

//GetByte ...
func (obj *Buffer) GetByte() []byte {

	var tempBuff = obj.packetBuffer.Bytes()
	var value = tempBuff[:1]
	var restValue = tempBuff[1:]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(restValue)
	obj.packetBuffer = byteBuffer
	return value
}

//Array ...
func (obj *Buffer) Array() []byte {
	return obj.packetBuffer.Bytes()
}

//Size ...
func (obj *Buffer) Size() int {
	return len(obj.packetBuffer.Bytes())
}

//Flip ...
func (obj *Buffer) Flip() {

	var bytesArr = obj.packetBuffer.Bytes()
	for i, j := 0, len(bytesArr)-1; i < j; i, j = i+1, j-1 {
		bytesArr[i], bytesArr[j] = bytesArr[j], bytesArr[i]
	}

	var byteBuffer bytes.Buffer
	byteBuffer.Write(bytesArr)
	obj.packetBuffer = byteBuffer
}

//Clear ...
func (obj *Buffer) Clear() {
	var byteBuffer bytes.Buffer
	obj.packetBuffer = byteBuffer
}

//Slice ...
func (obj *Buffer) Slice(start int, end int) error {

	var bytesArr = obj.packetBuffer.Bytes()

	if len(bytesArr) < (start + end) {
		return errors.New("Buffer does not contain that much of limit")
	}

	bytesArr = bytesArr[start:end]
	var byteBuffer bytes.Buffer
	byteBuffer.Write(bytesArr)
	obj.packetBuffer = byteBuffer
	return nil
}
