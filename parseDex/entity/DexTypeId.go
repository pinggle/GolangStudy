package entity

import (
	. "GolangStudy/parseDex/utils"
	"fmt"
	"os"
)

/// 类型信息

type DexTypeId struct {
	descriptorIdx uint32 //DexStringId中的index;
}

/// 读取Dex文件头部数据;

func ReadDexTypeId(typeIdOff uint32, typeIdSize uint32, fileHandler *os.File, dexStringIds []DexStringId) []DexTypeId {
	var dexTypeIds = make([]DexTypeId, typeIdSize)
	var i uint32

	for i = 0; i < typeIdSize; i++ {
		var dexTypeId DexTypeId
		ReadFileDataToInt("DexTypeId.descriptorIdx", &dexTypeId.descriptorIdx, fileHandler, typeIdOff+i*4, 4)
		if dexTypeId.descriptorIdx > 0 && dexTypeId.descriptorIdx < (uint32)(len(dexStringIds)) {
			if dexStringIds[dexTypeId.descriptorIdx].Length > 0 {
				fmt.Printf("DexTypeId.descriptorIdx Print [%d]:[%s]\n", i, string(dexStringIds[dexTypeId.descriptorIdx].Data))
			}
		}
		dexTypeIds[i] = dexTypeId
	}
	return dexTypeIds
}

///*
// * Direct-mapped "type_id_item".
// */
//struct DexTypeId {
//	u4  descriptorIdx;      /* index into stringIds list for type descriptor */
//};
