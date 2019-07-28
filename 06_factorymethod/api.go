package factorymethod

type ExportFileApi interface {
	Export(data string) bool
}
