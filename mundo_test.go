package epidemiologicalsimulation

import "testing"

func TestInitMundo(t *testing.T) {
	obtido := 1
	desejado := 1
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
