package epidemiologicalsimulation

//padrão guardião
import "fmt"

//Simulação parematros de entrada
type Simulação struct {
	arquivoNomes      string
	arquivoPopulacao  string
	arquivoDistancias string
	diasSimulado      int
	f                 func(int) float32
}

//Simular é a rotina para iniciar a simulação
func (s Simulação) Simular() string {
	if s.verificarEntradas() {
		fmt.Println("iniciado simulação")
		var sergipe mundo
		sergipe.init(s)
		sergipe.contamine()
		for i := 0; i < s.diasSimulado; i++ {
			sergipe.umaVolta(&i)
			sergipe.deslocaPessoas()
		}
		return "ok"
	}
	return "erro"
}

func (s Simulação) verificarEntradas() bool {
	if s.arquivoDistancias == "" {
		fmt.Println("nome do arquivoDistancias")
		return false
	}
	if s.arquivoNomes == "" {
		fmt.Println("nome do arquivoNome")
		return false
	}
	if s.arquivoPopulacao == "" {
		fmt.Println("nome do arquivoNome")
		return false
	}
	return true
}
