package entity

import (
	. "GolangStudy/parseDex/utils"
	"bufio"
	"fmt"
	"os"
)

/// 字符串信息

type DexStringId struct {
	StringDataOff uint32

	//=================
	Length uint32 // 字符串长度;
	Data   []byte // 字符串内容;
}

/// 读取Dex文件头部数据;

func ReadDexStringId(stringIdSize uint32, fileReader *bufio.Reader, fileHandler *os.File) []DexStringId {
	var dexStringIds = make([]DexStringId, stringIdSize)
	var i uint32

	for i = 0; i < stringIdSize; i++ {
		var dexStringId DexStringId
		ReadDataToInt("DexStringId.StringDataOff", &dexStringId.StringDataOff, fileReader, 4)
		ReadFileDataToInt("DexStringId.Length", &dexStringId.Length, fileHandler, dexStringId.StringDataOff, 1)
		ReadFileDataToBuffer("DexStringId.Data", &dexStringId.Data, fileHandler, dexStringId.StringDataOff+1, dexStringId.Length)
		if dexStringId.Length > 0 {
			fmt.Printf("DexStringId String Print [%d]:[%s]\n", i, string(dexStringId.Data))
		}
		dexStringIds[i] = dexStringId
	}
	//ReadDataToInt("Checksum", &dexHeader.Checksum, fileReader, 4)
	return dexStringIds
}

///*
// * Direct-mapped "string_id_item".
// */
//struct DexStringId {
//	u4 stringDataOff;      /* file offset to string_data_item */
//};
