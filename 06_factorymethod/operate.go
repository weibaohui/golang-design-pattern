package factorymethod

// 实现了 ExportFileApi
type ExportOperate struct{}

func (o *ExportOperate) factoryMethod() ExportFileApi {
	return nil
}
func (o *ExportOperate) Export(data string) bool {
	api := o.factoryMethod()
	return api.Export(data)
}
