package epidemiologicalsimulation

import "testing"

var sergipe = Mundo{
	arquivoNomesCidades:      "testeNomes.dat",
	arquivoPopulaçãoCidades:  "testepopulacao.dat",
	arquivoDistanciasCidades: "testedistancias.dat",
	TempoSimulação:           10,
}

func TestCarregaArquivos(t *testing.T) {
	obtido := sergipe.carregaArquivos()
	desejado := true
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestInitMundo(t *testing.T) {
	obtido := sergipe.init()
	desejado := true
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestMundo(t *testing.T) {
	obtido := 1
	desejado := 1
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
