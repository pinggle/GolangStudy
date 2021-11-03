package utils

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

const DebugMode bool = false

/// 从指定reader中读取指定长度的数据,以小端方式读出来,存储到指定buffer;

func ReadDataToInt(tag string, outPtr *uint32, fileReader *bufio.Reader, size uint32) {
	byteBuffer := make([]byte, size)
	fileReader.Read(byteBuffer)
	*outPtr = binary.LittleEndian.Uint32(byteBuffer)
	if DebugMode {
		fmt.Printf("%s is [%d] => 0x[%X]\n", tag, *outPtr, *outPtr)
	}
}

func ReadFileDataToInt(tag string, outPtr *uint32, fileHandler *os.File, readOffset uint32, size uint32) {

	if 0 == readOffset {
		return
	}

	byteBuffer := make([]byte, size)
	fileHandler.Seek(int64(readOffset), 0)
	fileHandler.Read(byteBuffer)
	if 1 == size {
		// 读取1byte;
		*outPtr = uint32(byteBuffer[0])
	} else if 4 == size {
		*outPtr = binary.LittleEndian.Uint32(byteBuffer)
	} else {
		fmt.Errorf("unsupport size :[%d]", size)
	}
	if DebugMode {
		fmt.Printf("%s is [%d] => 0x[%X]\n", tag, *outPtr, *outPtr)
	}
}

func ReadFileDataToBuffer(tag string, outPtr *[]byte, fileHandler *os.File, readOffset uint32, size uint32) {

	if 0 == readOffset || 0 == size {
		return
	}

	byteBuffer := make([]byte, size)
	fileHandler.Seek(int64(readOffset), 0)
	fileHandler.Read(byteBuffer)
	*outPtr = byteBuffer
	if DebugMode {
		fmt.Printf("%s [%d] is [%d] => 0x[%X]\n", readOffset, tag, *outPtr, *outPtr)
	}
}

func ReadDataToBuffer(tag string, outPtr *[]byte, fileReader *bufio.Reader, size uint32) {
	byteBuffer := make([]byte, size)
	fileReader.Read(byteBuffer)
	*outPtr = byteBuffer
	if DebugMode {
		fmt.Printf("%s is [%d] => 0x[%X]\n", tag, *outPtr, *outPtr)
	}
}
