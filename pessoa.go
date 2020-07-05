package epidemiologicalsimulation

import "math/rand"

//Pessoa tipo usado para representar cada individuo do modelo
type Pessoa struct {
	// Estado:susceptivel=0, contaminado=1, morto=2 ou imune=3, recuperado=4,
	Estado uint8
	// atestado: sim =1 nao = 0
	Examinada uint8
	//Dia: da contato contado apartir do Dia 0 inicio da contato apenas 1 no Dia 0
	Dia int
	//array de apontadores para pessoas vizinhas
	Vizinhos []*Pessoa
	//CodCidade é o local onde a Pessoa está
	CodCidade uint8
}

//contato execulta um passo markroviano
func (p *Pessoa) contato(data *int, probabilidade *[]float32) uint8 {
	var x, y uint8
	for i := range p.Vizinhos {
		if p.Vizinhos[i].Estado == 1 {
			x++
		}
	}
	if x > 0 {
		if rand.Float32() < (*probabilidade)[x] {
			p.Estado = 1
			p.Dia = *data
			y++
		}
	}
	return y
}

func (p *Pessoa) numeroVizinhosContaminados() uint8 {
	var x uint8
	for i := range p.Vizinhos {
		if p.Vizinhos[i] != nil {
			if p.Vizinhos[i].Estado == 1 {
				x++
			}
		}
	}
	return x
}

func (p *Pessoa) appendVizinho(novoViz *Pessoa) bool {
	for viz := range p.Vizinhos {
		if (p.Vizinhos[viz]) == novoViz {
			return false
		}
	}
	p.Vizinhos = append(p.Vizinhos, novoViz)
	return true
}
