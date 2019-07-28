package factorymethod

import "fmt"

type ExportDBFile struct{}

func (d *ExportDBFile) Export(data string) bool {
	fmt.Println("导出数据", data, "到数据库备份文件")
	return true
}

type ExportDBFileOperate struct {
	ExportOperate
}

func (o *ExportDBFileOperate) Create() ExportFileApi {
	return &ExportDBFile{}
}
