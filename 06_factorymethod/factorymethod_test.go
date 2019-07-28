package factorymethod

import "testing"

func TestExport(t *testing.T) {
	dbFileOperate := &ExportDBFileOperate{}
	Export(dbFileOperate.Create(), "测试数据")

	txtFileOperate := &ExportTxtFileOperate{}
	Export(txtFileOperate.Create(), "测试数据")

}

func Export(api ExportFileApi, data string) bool {

	return api.Export(data)
}
