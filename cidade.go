package epidemiologicalsimulation

//TODO misturar as interações criar n Vizinhos

import (
	"fmt"
	"math/rand"
)

//Cidade rede mundo pequeno unidimesional
type Cidade struct {
	Nome                string
	CodCidade           uint8
	TamanhoPopulação    int
	População           []Pessoa
	Contaminados        int
	Susceptivel         int
	MortosImmunes       int
	NumeroVizinhos      int
	NumeroTrocaVizinhos int
	TaxaImunidades      float32
	//numero de dias para a doença acabar matando ou imunizando a Pessoa
	Ciclo int
}

//Init calcula os parametros básicos da cidade
func (c *Cidade) Init() int {
	c.setPessoa()
	c.SetVizinhos()
	c.Susceptivel = c.TamanhoPopulação
	return c.TamanhoPopulação
}

//Vizinhos configura os Vizinhos de cada Pessoa
// precisa ser chamado depois que a População for alocada
// configurada para uma rede quadrada precisa configurar para variar as interações no futuro

func (c *Cidade) SetVizinhos() {
	//s é o sucessor a é o antecessor

	if len(c.População) > 0 {
		for i := range c.População {
			for j := 1; j <= c.NumeroVizinhos/2; j++ {
				if i-j < 0 {
					c.População[i].Vizinhos[j-1] = &c.População[c.TamanhoPopulação+(i-j)]
				} else {
					c.População[i].Vizinhos[j-1] = &c.População[i-j]
				}
				if i+j >= c.TamanhoPopulação {
					c.População[i].Vizinhos[j-1+c.NumeroVizinhos/2] = &c.População[-c.TamanhoPopulação+(i+j)]
				} else {
					c.População[i].Vizinhos[j-1+c.NumeroVizinhos/2] = &c.População[i+j]
				}
			}
		}
		for i := 0; i < c.NumeroTrocaVizinhos; i++ {
			k := rand.Intn(c.TamanhoPopulação)
			l := rand.Intn(c.TamanhoPopulação)
			m := rand.Intn(c.NumeroVizinhos)
			n := rand.Intn(c.NumeroVizinhos)
			c.População[k].Vizinhos[m], c.População[l].Vizinhos[n] = c.População[l].Vizinhos[n], c.População[k].Vizinhos[m]
		}

	} else {
		fmt.Println("Erro: é necessário alocar a População ")
	}
}

//propaga testa todas as pessoas da Cidade  propagando a doença de acordo com a
//prbabilidade empirica de contato
// essa é a rotina paralela
// x numero de infectados no cilclo
// y numero de mortos no Ciclo
func (c *Cidade) propaga(data *int, probabilidade *[]float32, x chan int) {
	var dx, dy int
	for i := range c.População {
		if c.População[i].Estado == 0 {
			dx += int(c.População[i].contato(data, probabilidade))
		} else {
			if c.População[i].Estado == 1 && (*data-c.População[i].Dia) > c.Ciclo {
				c.População[i].Estado = 2
				dx--
				dy++
			}
		}
	}
	c.Contaminados += dx
	c.MortosImmunes += dy
	x <- 0
}

func (c *Cidade) Propaga(data *int, probabilidade *[]float32) {
	var dx, dy int
	for i := range c.População {
		if c.População[i].Estado == 0 {
			dx += int(c.População[i].contato(data, probabilidade))
		} else {
			if c.População[i].Estado == 1 && (*data-c.População[i].Dia) > c.Ciclo {
				c.População[i].Estado = 0
				dx--
				dy++
			}
		}
	}
	c.Contaminados += dx
	c.MortosImmunes += dy
}

func (c *Cidade) setPessoa() {
	for i := range c.População {
		c.População[i].CodCidade = c.CodCidade
		if rand.Float32() < c.TaxaImunidades {
			c.População[i].Estado = 2
		} else {
			c.População[i].Estado = 0
		}

		c.População[i].Vizinhos = make([]*Pessoa, c.NumeroVizinhos)
	}
}
