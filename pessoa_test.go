package epidemiologicalsimulation

import "testing"

func TestContaminação(t *testing.T) {
	var cidade [6]Pessoa

	for i := 0; i < 6; i++ {
		cidade[i] = Pessoa{
			estado: 1,
			dia:    0,
		}
	}

	for i := 0; i < 5; i++ {
		cidade[0].vizinhos[i] = &cidade[i+1]
	}

	obtido := cidade[0].numeroVizinhosContaminados()
	desejado := uint8(5)
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
