package entity

type DexFileInfo struct {
	DexHeader    DexHeader
	DexStringIds []DexStringId
	DexTypeIds []DexTypeId
}
