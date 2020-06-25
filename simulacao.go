package epidemiologicalsimulation

//padrão guardião
import (
	"fmt"
	"log"
	"os"
)

//Simulação parematros de entrada
type Simulação struct {
	arquivoNomes      string
	arquivoPopulacao  string
	arquivoDistancias string
	diasSimulado      int
	// funcao de probabilidade da contaminaçao
	f func(int) float32
}

//Simular é a rotina para iniciar a simulação
func (s Simulação) Simular() string {
	if s.verificarEntradas() {
		fmt.Println("iniciado simulação")
		var sergipe Mundo
		sergipe.init()
		sergipe.contamine()
		for i := 0; i < s.diasSimulado; i++ {
			sergipe.umDia()
			sergipe.deslocaPessoas()
		}
		return "ok"
	}
	return "erro"
}

//#TODO: trocar o print por loge aprender a usar o log
func (s Simulação) verificarEntradas() bool {

	if s.arquivoDistancias == "" {
		fmt.Println("nome do arquivoDistancias ausente")
		return false
	} else {
		_, err := os.Open(s.arquivoDistancias)
		if err != nil {
			log.Fatal(err)
			return false
		}

	}
	if s.arquivoNomes == "" {
		fmt.Println("nome do arquivoNome")
		return false
	} else {
		_, err := os.Open(s.arquivoDistancias)
		if err != nil {
			log.Fatal(err)
			return false
		}

	}
	if s.arquivoPopulacao == "" {
		fmt.Println("nome do arquivoNome")
		return false
	} else {
		_, err := os.Open(s.arquivoDistancias)
		if err != nil {
			log.Fatal(err)
			return false
		}

	}
	return true
}
