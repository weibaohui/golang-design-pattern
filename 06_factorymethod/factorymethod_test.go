package factorymethod

import "testing"

func TestExport(t *testing.T) {
	dbFileOperate := &ExportDBFileOperate{}
	Export(dbFileOperate.factoryMethod(), "测试数据")

	txtFileOperate := &ExportTxtFileOperate{}
	Export(txtFileOperate.factoryMethod(), "测试数据")

}

func Export(api ExportFileApi, data string) bool {
	return api.Export(data)
}
