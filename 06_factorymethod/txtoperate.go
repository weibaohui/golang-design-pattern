package factorymethod

import "fmt"

type ExportTxtFile struct{}

func (t *ExportTxtFile) Export(data string) bool {
	fmt.Println("导出数据", data, "到文本文件")
	return true
}

// 组合了ExportOperate
type ExportTxtFileOperate struct {
	ExportOperate
}

func (o *ExportTxtFileOperate) Create() ExportFileApi {
	return &ExportTxtFile{}
}
