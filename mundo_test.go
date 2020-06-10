package epidemiologicalsimulation

import (
	"log"
	"math"
	"testing"
)

var sergipe = Mundo{
	arquivoNomesCidades:      "testeNomes.dat",
	arquivoPopulaçãoCidades:  "testepopulacao.dat",
	arquivoDistanciasCidades: "testedistancias.dat",
	TempoSimulação:           10,
	numeroVizinhos:           5,
	f: func(n int) float32 {
		return float32(1 - math.Pow((1-0.1), float64(n)))
	},
}

func TestLerTexto(t *testing.T) {
	obtido, err := lerTexto("testeNomes.dat")
	if err != nil {
		log.Fatalf("Erro ler nomes: %v", err)
	}
	//desejado := []string{"Aracaju", "Lagarto"}
	if obtido == nil {
		t.Errorf("o valor obtido foi %v ", obtido)
	}

	obtido2 := len(obtido)
	desejado := 2
	if obtido2 != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido2, desejado)
	}

}

func TestCarregaNomesCidades(t *testing.T) {
	obtido := sergipe.carregaNomeCidades()
	desejado := true
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
	if sergipe.nomesCidades == nil {
		t.Errorf("o valor obtido foi %v ", sergipe.nomesCidades)
	}

}

func TestCarregapopulaçãoCidades(t *testing.T) {
	obtido := sergipe.carregaPopulaçãoCidades()
	desejado := true
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}

}

func TestCarregaDistânciasCidades(t *testing.T) {
	obtido := sergipe.carregaDistânciasCidades()
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

	obtido2 := sergipe.numeroCidades
	desejado2 := 2
	if obtido2 != desejado2 {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido2, desejado2)
	}

	obtido1 := sergipe.cidades[0].tamanhoPopulação
	desejado1 := 10000
	if obtido1 != desejado1 {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido1, desejado1)
	}

	obtido3 := sergipe.cidades[1].tamanhoPopulação
	desejado3 := 1000
	if obtido1 != desejado1 {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido3, desejado3)
	}
}

func TestInitProbabilidadeContagio(t *testing.T) {
	sergipe.initProbabilidadeContagio()
	obtido := sergipe.probabilidadeContagio[0]
	desejado := float32(0.0)
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
