package factorymethod

// 实现了 ExportFileApi
type ExportOperate struct{}

func (o *ExportOperate) Create() ExportFileApi {
	return nil
}
func (o *ExportOperate) Export(data string) bool {
	api := o.Create()
	return api.Export(data)
}
