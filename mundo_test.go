package epidemiologicalsimulation

import (
	"fmt"
	"math"
	"testing"

	"github.com/dcadenas/pagerank"
)

var sergipe = Mundo{
	ArquivoNomesCidades:      "testeNomes.dat",
	ArquivoPopulaçãoCidades:  "testepopulacao.dat",
	ArquivoDistanciasCidades: "testedistancias.dat",
	TempoSimulação:           10,
	NumeroVizinhos:           5,
	CidadeInicial:            "Aracaju",
	Ciclo:                    10,
	P:                        float32(0.2),
	Alpha:                    float32(0.2),
	FTroca: func(x float32) float32 {
		return float32(math.Exp(float64(-x)))
	},
}

func TestCarregaNomesCidades(t *testing.T) {
	obtido := sergipe.carregaNomeCidades()
	desejado := true
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
	if sergipe.NomesCidades == nil {
		t.Errorf("o valor obtido foi %v ", sergipe.NomesCidades)
	}
	if sergipe.Cidades[0].Nome != "Aracaju" {
		t.Errorf("o valor obtido foi %v ", sergipe.Cidades[0].Nome)
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
	sergipe.FTroca = nil
	sergipe.Beta = float32(1)
	obtido := sergipe.Init()
	desejado := true
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}

	obtido2 := sergipe.NumeroCidades
	desejado2 := 2
	if obtido2 != desejado2 {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido2, desejado2)
	}

	obtido1 := sergipe.Cidades[0].TamanhoPopulação
	desejado1 := 10000
	if obtido1 != desejado1 {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido1, desejado1)
	}

	obtido3 := sergipe.Cidades[1].TamanhoPopulação
	desejado3 := 1000
	if obtido1 != desejado1 {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido3, desejado3)
	}

	if sergipe.Cidades[0].Nome != "Aracaju" {
		t.Errorf("o valor obtido foi %v ", sergipe.Cidades[0].Nome)
	}
}

func TestInitProbabilidadeTroca(t *testing.T) {
	obtido0 := sergipe.initProbabilidadeTroca()
	obtido := sergipe.ProbabilidadeTroca[0][0]
	desejado := float32(1)
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
	if obtido0 != true {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestDeslocaPessoas(t *testing.T) {
	sergipe.deslocaPessoas(10)
	obtido := 1
	desejado := 1
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
func TestContamine(t *testing.T) {
	sergipe.Init()

	obtido := sergipe.Cidades[0].Contaminados
	desejado := 1
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", sergipe.Cidades[0].Nome, desejado)
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
func TestUmDia(t *testing.T) {
	sergipe.Init()
	sergipe.UmDia()
	obtido := sergipe.Cidades[0].Contaminados
	desejado := 18
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestUmAno(t *testing.T) {
	sergipe.Init()

	for i := 0; i < 200; i++ {
		sergipe.UmDia()
	}
	obtido := sergipe.Cidades[1].Contaminados
	desejado := 0 //sergipe.Cidades[0].TamanhoPopulação
	if obtido == desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)

	}
}

func TestUmAno2(t *testing.T) {
	sergipe.Init()
	graph := pagerank.New()
	for i := 0; i < 200; i++ {
		sergipe.UmDia()
	}
	for _, k := range sergipe.Links {
		graph.Link(int(k[0]), int(k[1]))
	}
	probabilityFollowingLink := 0.85 // The bigger the number, less probability we have to teleport to some random link
	tolerance := 0.0001              // the smaller the number, the more exact the result will be but more CPU cycles will be needed

	graph.Rank(probabilityFollowingLink, tolerance, func(identifier int, rank float64) {
		fmt.Println("Node", identifier, "rank is", rank)
	})

	obtido := sergipe.Cidades[0].Contaminados
	desejado := sergipe.Cidades[0].TamanhoPopulação
	if obtido > desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
