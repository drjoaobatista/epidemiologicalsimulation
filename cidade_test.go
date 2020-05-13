package epidemiologicalsimulation

import (
	"testing"
)

func TestInitCidade(t *testing.T) {
	var aracaju = Cidade{
		nome:             "aracaju",
		codCidade:        1,
		tamanhoPopulação: 10010,
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
	aracaju.população = make([]Pessoa, aracaju.tamanhoPopulaçãoQuadrada)
	aracaju.setPessoa()
	aracaju.vizinhos()
	obtido := 0
	desejado := 0
	if obtido != desejado {
		t.Errorf("o valor obtido foi %v e o desejado foi %v", obtido, desejado)
	}

}
