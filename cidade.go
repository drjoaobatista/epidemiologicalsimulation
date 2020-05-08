package epidemiologicalsimulation

import "math"

//Cidade rede mundo pequeno
type Cidade struct {
	nome                     string
	codCidade                uint8
	tamanhoPopulação         int
	tamanhoPopulaçãoQuadrada int
	L                        int
	erro                     float32
	população                []Pessoa
	contamiados              int
	susceptivel              int
	mortosImmunes            int
	//numero de dias para a doença acabar matando ou imunizando a Pessoa
	ciclo int
}

//init cria uma matriz retangular de pessoas
func (c *Cidade) init(nome string, população int, cod uint8) int {
	c.codCidade = cod
	c.nome = nome
	c.tamanhoPopulação = população
	c.L = int(math.Round(math.Sqrt(float64(população))))
	c.erro = (float32(c.L*c.L) - float32(c.tamanhoPopulação)) / float32(c.tamanhoPopulação)

	return int(c.L * c.L)
}

//vizinhos configura os viznhos de cada Pessoa
// precisa ser chamado depois que a populaçao for alocada
// configuarada para uma rede quadrada precisa configurara para variar as interações no futuro
func (c *Cidade) vizinhos() {
	//s é o sucessor a é o antecessor
	var a, s []int
	for i := 0; i < c.L; i++ {
		s[i] = i + 1
		a[i] = i - 1
	}
	s[c.L-1] = 0
	a[0] = c.L - 1

	for i := range c.população {
		lin, col := i/c.L, i%c.L
		v0 := lin*c.L + a[col]
		v1 := lin*c.L + s[col]
		v2 := a[lin]*c.L + col
		v3 := s[lin]*c.L + col
		c.população[i].vizinhos[0] = &c.população[v0]
		c.população[i].vizinhos[1] = &c.população[v1]
		c.população[i].vizinhos[2] = &c.população[v2]
		c.população[i].vizinhos[3] = &c.população[v3]
	}

}

//propaga testa todas as pessoas da Cidade  propagando a doença de acordo com a
//prbabilidade empirica de contaminação
// essa é a rotina paralela
// x numero de infectados no cilclo
// y numero de mortos no ciclo
func (c *Cidade) propaga(data *int, probabilidade *[]float32, x chan int) {
	var dx, dy int
	for i := range c.população {
		if c.população[i].estado == 0 {
			c.população[i].contaminação(data, probabilidade)
			dx++
		} else {
			if c.população[i].estado == 1 && (*data-c.população[i].dia) > c.ciclo {
				c.população[i].estado = 2
				dx--
				dy++
			}
		}
	}
	c.contamiados += dx
	c.mortosImmunes += dy
	x <- 0
}

func (c *Cidade) setPessoa() {
	for i := range c.população {
		c.população[i].codCidade = c.codCidade
	}
}
