package epidemiologicalsimulation

import "testing"

func TestContaminação(t *testing.T) {
	var cidade [6]Pessoa

	for i := 0; i < 6; i++ {
		cidade[i] = Pessoa{
			Estado:   1,
			Dia:      0,
			Vizinhos: make([]*Pessoa, 4),
		}
	}

	for i := 0; i < len(cidade[0].Vizinhos); i++ {
		cidade[0].Vizinhos[i] = &cidade[i+1]
	}

	obtido := cidade[0].numeroVizinhosContaminados()
	desejado := uint8(4)
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
