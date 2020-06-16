package epidemiologicalsimulation

//TODO misturar as interações criar n vizinhos

import (
	"fmt"
)

//Cidade1D rede mundo pequeno unidimesional
type Cidade1D struct {
	nome                     string
	codCidade                uint8
	tamanhoPopulação         int
	tamanhoPopulaçãoQuadrada int
	L                        int
	erro                     float32
	população                []Pessoa
	contaminados             int
	susceptivel              int
	mortosImmunes            int
	numeroVizinhos           int
	//numero de dias para a doença acabar matando ou imunizando a Pessoa
	ciclo int
}

//init calcula os parametros básicos da cidade
func (c *Cidade1D) init() int {
	c.L = c.tamanhoPopulação
	c.erro = 0
	c.tamanhoPopulaçãoQuadrada = c.tamanhoPopulação
	c.susceptivel = c.tamanhoPopulaçãoQuadrada
	return c.tamanhoPopulação
}

//vizinhos configura os vizinhos de cada Pessoa
// precisa ser chamado depois que a população for alocada
// configurada para uma rede quadrada precisa configurar para variar as interações no futuro

func (c *Cidade1D) vizinhos() {
	//s é o sucessor a é o antecessor
	if len(c.população) > 0 {
		for i := range c.população {
			for j := 1; j <= c.numeroVizinhos/2; j++ {
				if i-j < 0 {
					c.população[i].vizinhos[j] = &c.população[c.tamanhoPopulação+(i-j)]
				} else {
					c.população[i].vizinhos[j] = &c.população[i-j]
				}
				if i+j >= c.tamanhoPopulação {
					c.população[i].vizinhos[j] = &c.população[-c.tamanhoPopulação+(i+j)]
				} else {
					c.população[i].vizinhos[j] = &c.população[i-j]
				}
			}
		}
	} else {
		fmt.Println("Erro: é necessário alocar a populacao ")
	}
}

//propaga testa todas as pessoas da Cidade1D  propagando a doença de acordo com a
//prbabilidade empirica de contato
// essa é a rotina paralela
// x numero de infectados no cilclo
// y numero de mortos no ciclo
func (c *Cidade1D) propaga(data *int, probabilidade *[]float32, x chan int) {
	var dx, dy int
	for i := range c.população {
		if c.população[i].estado == 0 {
			dx += int(c.população[i].contato(data, probabilidade))
		} else {
			if c.população[i].estado == 1 && (*data-c.população[i].dia) > c.ciclo {
				c.população[i].estado = 2
				dx--
				dy++
			}
		}
	}
	c.contaminados += dx
	c.mortosImmunes += dy
	x <- 0
}

func (c *Cidade1D) setPessoa() {
	for i := range c.população {
		c.população[i].codCidade = c.codCidade
		c.população[i].estado = 0
	}
}
