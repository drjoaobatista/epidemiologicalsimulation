package epidemiologicalsimulation

import "math/rand"

//pessoa tipo usado para representar cada individuo do modelo
type pessoa struct {
	// estado:susceptivel=0, contaminado=1, morto ou imune=2,
	estado uint8
	//dia: da contaminação contado apartir do dia 0 inicio da contaminação apenas 1 no dia 0
	dia int
	//array de apontadores para pessoas vizinhas
	vizinhos [4]*pessoa
	//codCidade é o local onde a pessoa está
	codCidade uint8
}

//contaminação execulta um passo markroviano
func (p *pessoa) contaminação(data *int, probabilidade *[]float32) uint8 {
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
