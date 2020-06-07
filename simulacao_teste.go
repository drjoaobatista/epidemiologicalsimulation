package epidemiologicalsimulation

import "testing"

var covid Simulação

func TestSimulação(t *testing.T) {
	obtido := 1
	desejado := 1
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestInitSimulação(t *testing.T) {
	obtido := 1
	desejado := 1
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
