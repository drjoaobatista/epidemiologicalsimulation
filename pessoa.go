package epidemiologicalsimulation

import "math/rand"

//Pessoa tipo usado para representar cada individuo do modelo
type Pessoa struct {
	// estado:susceptivel=0, contaminado=1, morto=2 ou imune=3, recuperado=4,
	estado uint8
	// atestado: sim =1 nao = 0
	examinada uint8
	//dia: da contato contado apartir do dia 0 inicio da contato apenas 1 no dia 0
	dia int
	//array de apontadores para pessoas vizinhas
	vizinhos [4]*Pessoa //#TODO: deixar número de vizinhos variaveis
	//codCidade é o local onde a Pessoa está
	codCidade uint8
}

//contato execulta um passo markroviano
func (p *Pessoa) contato(data *int, probabilidade *[]float32) uint8 {
	var x uint8
	for i := range p.vizinhos {
		if p.vizinhos[i].estado == 1 {
			x++
		}
	}
	if x > 0 {
		if rand.Float32() < (*probabilidade)[x] {
			p.estado = 1
			p.dia = *data
			return 1
		}
	}
	return 0
}

func (p *Pessoa) numeroVizinhosContaminados() uint8 {
	var x uint8
	for i := range p.vizinhos {
		if p.vizinhos[i] != nil {
			if p.vizinhos[i].estado == 1 {
				x++
			}
		}
	}
	return x
}
