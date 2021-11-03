package service

import (
	. "GolangStudy/parseDex/entity"
	"bufio"
	"fmt"
	"os"
)

// ParseDex
// 解析Dex参考:http://aospxref.com/android-8.0.0_r36/xref/dalvik/libdex/DexFile.h
///**
func ParseDex(inputFile string) {

	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		fmt.Println("文件不存在:", inputFile)
		return
	}
	fmt.Println("文件大小:", fileInfo.Size())

	fileHandler, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	// 声明退出函数前关闭文件
	defer fileHandler.Close()

	var dexFileInfo DexFileInfo
	fileReader := bufio.NewReader(fileHandler)
	dexFileInfo.DexHeader = ReadDexHeader(fileReader)
	dexFileInfo.DexStringIds = ReadDexStringId(dexFileInfo.DexHeader.StringIdsSize, fileReader, fileHandler)
	dexFileInfo.DexTypeIds = ReadDexTypeId(dexFileInfo.DexHeader.TypeIdsOff, dexFileInfo.DexHeader.TypeIdsSize, fileHandler, dexFileInfo.DexStringIds)
	//fileHandler.Seek(10,0)
	//fileHandler.Read(2)

}

func init() {
	fmt.Println("ParseDex, init() ...")
}
