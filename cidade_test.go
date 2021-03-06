package epidemiologicalsimulation

import (
	"math"
	"runtime"
	"testing"
)

func TestCidade(t *testing.T) {
	var aracaju = Cidade{
		Nome:             "aracaju",
		CodCidade:        1,
		TamanhoPopulação: 1000000,
		NumeroVizinhos:   3,
		Ciclo:            10,
		P:                0.1,
		F: func(n int) float32 {
			return float32(1 - math.Pow((1-0.5), float64(n)))
		},
	}
	// criando a população
	aracaju.População = make([]Pessoa, aracaju.TamanhoPopulação)
	aracaju.Init()
	//contaminando primeira pessoa da cidade
	aracaju.População[0].Estado = 1
	aracaju.População[0].Dia = 0
	aracaju.Contaminados = 1

	//usando goroutines
	var numCPU = runtime.NumCPU()
	canal := make(chan int, numCPU)
	var goroutines int
	for data := 0; data < 100; data++ {
		go aracaju.propaga(&data, canal)
		goroutines++
		if goroutines >= numCPU {
			<-canal
			goroutines--
		}
	}
	for i := 0; i < goroutines; i++ {
		<-canal
	}
	obtido := aracaju.Contaminados
	desejado := aracaju.TamanhoPopulação
	if obtido > desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestCidade1(t *testing.T) {
	var aracaju = Cidade{
		Nome:             "aracaju",
		CodCidade:        1,
		TamanhoPopulação: 1000000,
		NumeroVizinhos:   3,
		Ciclo:            10,
		P:                0.1,
		F: func(n int) float32 {
			return float32(1 - math.Pow((1-0.5), float64(n)))
		},
	}
	// criando a população
	aracaju.População = make([]Pessoa, aracaju.TamanhoPopulação)
	aracaju.Init()
	//contaminando primeira pessoa da cidade
	aracaju.População[0].Estado = 1
	aracaju.População[0].Dia = 0
	aracaju.Contaminados = 1
	//usando goroutines
	for data := 0; data < 100; data++ {
		aracaju.Propaga(&data)
	}
	obtido := aracaju.Contaminados
	desejado := aracaju.TamanhoPopulação
	if obtido > desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestCidade2(t *testing.T) {
	//testando a inicialização
	var aracaju = Cidade{
		Nome:             "aracaju",
		CodCidade:        1,
		TamanhoPopulação: 1000,
		NumeroVizinhos:   3,
		Ciclo:            10,
		P:                float32(0.4),
		F: func(n int) float32 {
			return float32(1 - math.Pow((1-0.5), float64(n)))
		},
	}
	// criando a população
	aracaju.População = make([]Pessoa, aracaju.TamanhoPopulação)
	aracaju.Init()
	//contaminando primeira pessoa da cidade
	aracaju.População[0].Estado = 1
	aracaju.População[0].Dia = 0
	aracaju.Contaminados = 1
	obtido := aracaju.MediaVizinhos
	desejado := float32(6)
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}

func TestCidade3(t *testing.T) {
	//testando a inicialização
	var aracaju = Cidade{
		Nome:             "aracaju",
		CodCidade:        1,
		TamanhoPopulação: 1000,
		NumeroVizinhos:   3,
		Ciclo:            10,
		P:                float32(0.4),
		F: func(n int) float32 {
			return float32(1 - math.Pow((1-0.5), float64(n)))
		},
	}
	// criando a população
	aracaju.População = make([]Pessoa, aracaju.TamanhoPopulação)
	aracaju.Init()
	aracaju.Init()
	aracaju.Init()
	//contaminando primeira pessoa da cidade
	aracaju.População[0].Estado = 1
	aracaju.População[0].Dia = 0
	aracaju.Contaminados = 1
	obtido := aracaju.MáximoVizinhos
	desejado := 11
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}
}
