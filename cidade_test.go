package epidemiologicalsimulation

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestCidade(t *testing.T) {
	//testando a inicializacao
	var aracaju = Cidade{
		nome:             "aracaju",
		codCidade:        1,
		tamanhoPopulação: 10010,
		numeroVizinhos:   6,
		ciclo:            10,
	}
	aracaju.init()
	// criando a populacao
	aracaju.população = make([]Pessoa, aracaju.tamanhoPopulação)
	// configurando as pessoas
	aracaju.setPessoa()
	// configurando os vizinhos numa rede quadrada
	aracaju.vizinhos()
	//contaminando primeira pessoa da cidade
	aracaju.população[0].estado = 1
	aracaju.população[0].dia = 0
	aracaju.contaminados = 1

	//contruindo um vetor de probabilidade de teste
	probabilidade := []float32{0, 0.3, 0.5, 0, 0, 0, 0, 0, 05}

	//usando goruntimes
	var numCPU = runtime.NumCPU()
	canal := make(chan int, numCPU)
	var goroutines int
	var data int
	data = 1
	for data := 0; data < 10; data++ {
		go aracaju.propaga(&data, &probabilidade, canal)
		goroutines++
		if goroutines >= numCPU {
			<-canal
			goroutines--
		}
	}
	for i := 0; i < goroutines; i++ {
		<-canal
	}

	obtido := aracaju.contaminados
	desejado := 4
	if obtido < desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
		t.Error(aracaju.população[1].vizinhos[0].vizinhos[0].contato(&data, &probabilidade))
		t.Error(rand.Float32())
	}

}

func TestCidade2(t *testing.T) {
	//testando a inicializacao
	var aracaju = Cidade{
		nome:             "aracaju",
		codCidade:        1,
		tamanhoPopulação: 100,
		numeroVizinhos:   6,
		ciclo:            10,
	}
	aracaju.init()
	// criando a populacao
	aracaju.população = make([]Pessoa, aracaju.tamanhoPopulação)
	// configurando as pessoas
	aracaju.setPessoa()
	// configurando os vizinhos numa rede quadrada
	aracaju.vizinhos()
	//contaminando primeira pessoa da cidade
	aracaju.população[0].estado = 1
	aracaju.população[0].dia = 0
	aracaju.contaminados = 1

	//contruindo um vetor de probabilidade de teste
	probabilidade := []float32{0, 0.3, 0.5, 0, 0, 0, 0, 0, 05}

	//usando goruntimes
	var numCPU = runtime.NumCPU()
	canal := make(chan int, numCPU)
	var goroutines int
	var data int
	data = 1
	for data := 0; data < 10; data++ {
		go aracaju.propaga(&data, &probabilidade, canal)
		goroutines++
		if goroutines >= numCPU {
			<-canal
			goroutines--
		}
	}
	for i := 0; i < goroutines; i++ {
		<-canal
	}

	obtido := aracaju.contaminados
	desejado := 4
	if obtido < desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
		t.Error(aracaju.população[1].vizinhos[0].vizinhos[0].contato(&data, &probabilidade))
		t.Error(rand.Float32())
	}

}
