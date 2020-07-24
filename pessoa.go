package epidemiologicalsimulation

import "math/rand"

//Pessoa : tipo usado para representar cada indivíduo do modelo
type Pessoa struct {
	// Estado:susceptível=0, contaminado=1, morto=2 ou imune=3, recuperado=4,
	Estado uint8
	// Examinada: sim =1 não = 0
	Examinada uint8
	//Dia: da contato contado a partir do Dia 0 início da contato apenas 1 no Dia 0
	Dia int
	//Ciclo: tempo de duração da contaminação
	Ciclo uint8
	//Vizinhos : Slices  de ponteiros  para pessoas vizinhas
	Vizinhos []*Pessoa
	//CodCidade é o local onde a Pessoa está
	CodCidade uint8
}

//contato executa um passo markoviano
func (p *Pessoa) contato(data *int, probabilidade *[]float32) int8 {
	var x uint8
	var y int8
	if p.Estado == 0 {
		for i := range p.Vizinhos {
			if p.Vizinhos[i].Estado == 1 {
				x++
			}
		}
		if x > 0 {
			if rand.Float32() < (*probabilidade)[x] {
				p.Estado = 1
				p.Dia = *data
				y = 1
			}
		}
	}
	if p.Estado == 1 {
		if p.Dia-*data >= int(p.Ciclo) {
			y = -1
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
