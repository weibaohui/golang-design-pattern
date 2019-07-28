package factorymethod

// 实现了 ExportFileApi
type ExportOperate struct{}

func (o *ExportOperate) FactoryMethod() ExportFileApi {
	return nil
}
func (o *ExportOperate) Export(data string) bool {
	api := o.FactoryMethod()
	return api.Export(data)
}
