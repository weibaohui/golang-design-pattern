package abstractfactory

import "testing"

func TestRun(t *testing.T) {
	engineer := &ComputerEngineer{}
	schema1 := &Schema2{}
	engineer.MakeComputer(schema1)
}
