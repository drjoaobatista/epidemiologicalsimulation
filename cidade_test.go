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
		ciclo:            10,
	}
	aracaju.init()
	L := 100
	L2 := L * L
	if L != aracaju.L {
		t.Errorf("Erro no calculo do L foi %v e o desejado foi %v", aracaju.L, L)
	}
	if L2 != aracaju.tamanhoPopulaçãoQuadrada {
		t.Errorf("Erro no calculo do L2 foi %v e o desejado foi %v", aracaju.tamanhoPopulaçãoQuadrada, L2)
	}
	if float32(10.0/10010.0) != aracaju.erro {
		t.Errorf("Erro no calculo do ERRO foi %v e o desejado foi %v", aracaju.erro, float32(10.0/10010.0))
	}
	// criando a populacao
	aracaju.população = make([]Pessoa, aracaju.tamanhoPopulaçãoQuadrada)
	// configurando as pessoas
	aracaju.setPessoa()
	// configurando os vizinhos numa rede quadrada
	aracaju.vizinhos()
	//contaminando primeira pessoa da cidade
	aracaju.população[0].estado = 1
	aracaju.população[0].dia = 0
	aracaju.contaminados = 1

	//contruindo um vetor de probabilidade de teste
	probabilidade := []float32{0, 0.3, 0.5, 0, 0, 0}

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
