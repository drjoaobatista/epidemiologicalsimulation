package epidemiologicalsimulation

import "testing"

var covid = Simulação{
	arquivoNomes:      "testeNomes.dat",
	arquivoPopulacao:  "testepopulacao.dat",
	arquivoDistancias: "testedistancias.dat",
	diasSimulado:      10,
	// funcao de probabilidade da contaminaçao
	//f func(int) float32
}

func TestVerificarEntradas(t *testing.T) {
	obtido := covid.verificarEntradas()
	desejado := true
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
