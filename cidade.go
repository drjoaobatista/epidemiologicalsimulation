package epidemiologicalsimulation

import (
	"log"
	"math/rand"
)

//Cidade rede mundo pequeno unidimesional
type Cidade struct {
	Nome             string
	CodCidade        uint8
	TamanhoPopulação int
	População        []Pessoa
	Contaminados     int
	Susceptivel      int
	MortosImmunes    int
	NumeroVizinhos   int
	TaxaImunidades   float32
	P                float32
	MáximoVizinhos   int
	MínimoVizinhos   int
	MediaVizinhos    float32
	//numero de dias para a doença acabar matando ou imunizando a Pessoa
	Ciclo int
}

//Init calcula os parametros básicos da cidade
func (c *Cidade) Init() {
	c.setPessoa()
	c.SetVizinhos()
	c.estatisticaVizinhos()
}

//SetVizinhos configura os Vizinhos de cada Pessoa criando uma rede Modelo WS (Watts-Strogatz):
func (c *Cidade) SetVizinhos() {
	if len(c.População) > 0 {
		for i := range c.População {
			for j := 1; j <= c.NumeroVizinhos; j++ {
				if rand.Float32() > c.P {
					c.População[i].appendVizinho(&c.População[(i+j)%c.TamanhoPopulação])
				} else {
					for {
						k := (rand.Intn(c.TamanhoPopulação-2*c.NumeroVizinhos) + c.NumeroVizinhos + i) % c.TamanhoPopulação
						if c.População[i].appendVizinho(&c.População[k]) {
							break
						}
					}
				}
			}
		}
		for i := range c.População {
			for j := range c.População[i].Vizinhos {
				c.População[i].Vizinhos[j].appendVizinho(&c.População[i])
			}
		}
	} else {
		log.Print("Erro: é necessário alocar a População ")

	}
}

//estatisticaVizinhos calcula a media o máximo e o mínimo número de vizinhos
func (c *Cidade) estatisticaVizinhos() {
	var nVizinhos, totalVizinhos int
	mínimo := len(c.População[0].Vizinhos)
	máximo := mínimo
	for i := range c.População {
		nVizinhos = len(c.População[i].Vizinhos)
		totalVizinhos += nVizinhos
		if mínimo > nVizinhos {
			mínimo = nVizinhos
		}
		if máximo < nVizinhos {
			máximo = nVizinhos
		}
	}
	c.MáximoVizinhos = máximo
	c.MínimoVizinhos = mínimo
	c.MediaVizinhos = float32(totalVizinhos) / float32(len(c.População))
}

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

//Propaga é um metodo que pormove a propagação da infeção usando o modelo de contato
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
			c.Susceptivel++
		}
	}
}
