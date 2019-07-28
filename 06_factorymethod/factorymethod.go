package factorymethod

import "fmt"

type ExportFileApi interface {
	Export(data string) bool
}
type ExportTxtFile struct{}

func (t *ExportTxtFile) Export(data string) bool {
	fmt.Println("导出数据", data, "到文本文件")
	return true
}

type ExportDBFile struct{}

func (d *ExportDBFile) Export(data string) bool {
	fmt.Println("导出数据", data, "到数据库备份文件")
	return true
}

type ExportOperate struct{}

func (o *ExportOperate) factoryMethod() ExportFileApi {
	return nil
}
func (o *ExportOperate) Export(data string) bool {
	api := o.factoryMethod()
	return api.Export(data)
}

type ExportTxtFileOperate struct {
	ExportOperate
}

func (o *ExportTxtFileOperate) factoryMethod() ExportFileApi {
	return &ExportTxtFile{}
}

type ExportDBFileOperate struct {
	ExportOperate
}

func (o *ExportDBFileOperate) factoryMethod() ExportFileApi {
	return &ExportDBFile{}
}
