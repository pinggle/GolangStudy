package entity

import (
	. "GolangStudy/parseDex/utils"
	"bufio"
	"fmt"
	"strings"
)

const DexHeaderMagicSize uint32 = 8
const DexHeaderSignatureSize uint32 = 20

/// Dex头部信息;

type DexHeader struct {
	Magic         []byte `info:"name" doc:"magic"` //文件标识，一般为"dex\n035\0" => dex + 换行符 + dex版本 + 0
	Checksum      uint32 //adler32校验用来确保Dex文件内容无损
	Signature     []byte /* SHA-1 hash */
	FileSize      uint32 /* length of entire file */
	HeaderSize    uint32 /* offset to start of next section */
	EndianTag     uint32
	LinkSize      uint32
	LinkOff       uint32
	MapOff        uint32
	StringIdsSize uint32
	StringIdsOff  uint32
	TypeIdsSize   uint32
	TypeIdsOff    uint32
	ProtoIdsSize  uint32
	ProtoIdsOff   uint32
	FieldIdsSize  uint32
	FieldIdsOff   uint32
	MethodIdsSize uint32
	MethodIdsOff  uint32
	ClassDefsSize uint32
	ClassDefsOff  uint32
	DataSize      uint32
	DataOff       uint32
}

/// 读取Dex文件头部数据;

func ReadDexHeader(fileReader *bufio.Reader) DexHeader {
	var dexHeader DexHeader

	// 读取Header->Magic
	ReadDataToBuffer("Magic", &dexHeader.Magic, fileReader, DexHeaderMagicSize)
	fmt.Printf("Magic is [%s]\n", strings.Replace(string(dexHeader.Magic), "\n", "\\n", -1))
	ReadDataToInt("Checksum", &dexHeader.Checksum, fileReader, 4)
	ReadDataToBuffer("Signature", &dexHeader.Signature, fileReader, DexHeaderSignatureSize)
	ReadDataToInt("FileSize", &dexHeader.FileSize, fileReader, 4)
	ReadDataToInt("HeaderSize", &dexHeader.HeaderSize, fileReader, 4)
	ReadDataToInt("EndianTag", &dexHeader.EndianTag, fileReader, 4)
	ReadDataToInt("LinkSize", &dexHeader.LinkSize, fileReader, 4)
	ReadDataToInt("LinkOff", &dexHeader.LinkOff, fileReader, 4)
	ReadDataToInt("MapOff", &dexHeader.MapOff, fileReader, 4)
	ReadDataToInt("StringIdsSize", &dexHeader.StringIdsSize, fileReader, 4)
	ReadDataToInt("StringIdsOff", &dexHeader.StringIdsOff, fileReader, 4)
	ReadDataToInt("TypeIdsSize", &dexHeader.TypeIdsSize, fileReader, 4)
	ReadDataToInt("TypeIdsOff", &dexHeader.TypeIdsOff, fileReader, 4)
	ReadDataToInt("ProtoIdsSize", &dexHeader.ProtoIdsSize, fileReader, 4)
	ReadDataToInt("ProtoIdsOff", &dexHeader.ProtoIdsOff, fileReader, 4)
	ReadDataToInt("FieldIdsSize", &dexHeader.FieldIdsSize, fileReader, 4)
	ReadDataToInt("FieldIdsOff", &dexHeader.FieldIdsOff, fileReader, 4)
	ReadDataToInt("MethodIdsSize", &dexHeader.MethodIdsSize, fileReader, 4)
	ReadDataToInt("MethodIdsOff", &dexHeader.MethodIdsOff, fileReader, 4)
	ReadDataToInt("ClassDefsSize", &dexHeader.ClassDefsSize, fileReader, 4)
	ReadDataToInt("ClassDefsOff", &dexHeader.ClassDefsOff, fileReader, 4)
	ReadDataToInt("DataSize", &dexHeader.DataSize, fileReader, 4)
	ReadDataToInt("DataOff", &dexHeader.DataOff, fileReader, 4)
	return dexHeader
}

// http://aospxref.com/android-8.0.0_r36/xref/dalvik/libdex/DexFile.h
///*
// * Direct-mapped "header_item" struct.
// */
//struct DexHeader {
//u1  magic[8];           /* includes version number */
//u4  checksum;           /* adler32 checksum */
//u1  signature[kSHA1DigestLen]; /* SHA-1 hash */
//u4  fileSize;           /* length of entire file */
//u4  headerSize;         /* offset to start of next section */
//u4  endianTag;
//u4  linkSize;
//u4  linkOff;
//u4  mapOff;
//u4  stringIdsSize;
//u4  stringIdsOff;
//u4  typeIdsSize;
//u4  typeIdsOff;
//u4  protoIdsSize;
//u4  protoIdsOff;
//u4  fieldIdsSize;
//u4  fieldIdsOff;
//u4  methodIdsSize;
//u4  methodIdsOff;
//u4  classDefsSize;
//u4  classDefsOff;
//u4  dataSize;
//u4  dataOff;
//};
