package epidemiologicalsimulation

import (
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

//Mundo estura para armazenar rede longa distância
type Mundo struct {
	NomesCidades             []string
	PopulaçãoCidades         []int
	NumeroCidades            int
	TamanhoPopulação         int
	População                []Pessoa
	Cidades                  []Cidade
	Distâncias               [][]float32
	Data                     int
	ProbabilidadeTroca       [][]float32
	ProbabilidadeContagio    []float32
	CidadeInicial            string
	ArquivoNomesCidades      string
	ArquivoPopulaçãoCidades  string
	ArquivoDistanciasCidades string
	TempoSimulação           int
	NumeroVizinhos           int
	Ciclo                    int

	// funcao de probabilidade da contaminaçao
	F func(int) float32
	// funcao de probabilidade da contaminaçao
	FTroca func(float32) float32
}

//init carrega do disco três arquivos com o Nome da Cidade e a População e as Distâncias
func (m *Mundo) init() bool {
	m.carregaNomeCidades()
	m.carregaPopulaçãoCidades()
	m.carregaDistânciasCidades()
	m.initProbabilidadeContagio()
	m.initProbabilidadeTroca()
	//criando as pessoas do mundo
	m.População = make([]Pessoa, m.TamanhoPopulação)
	//distribuindo a populacao mundial nas Cidades
	inicio := 0
	for i := 0; i < m.NumeroCidades; i++ {
		fim := int(m.Cidades[i].TamanhoPopulação) + inicio
		m.Cidades[i].População = m.População[inicio:fim]
		m.Cidades[i].Init()
		inicio = int(fim)
	}
	return true
}

func (m *Mundo) carregaNomeCidades() bool {
	if m.ArquivoNomesCidades == "" {
		log.Print("erro falta passar o nome do arquivoNome")
		return false
	} else {
		arquivo, err := os.Open(m.ArquivoNomesCidades)
		// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
		if err != nil {
			log.Fatalf("Erro ler nomes: %v", err)
			return false
		}
		// Garante que o arquivo sera fechado apos o uso
		defer arquivo.Close()
		// Cria um scanner que lê cada linha do arquivo
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			m.NomesCidades = append(m.NomesCidades, scanner.Text())
		}
	}
	if m.Cidades == nil {
		m.NumeroCidades = len(m.NomesCidades)
		m.Cidades = make([]Cidade, m.NumeroCidades)
	}
	for i, Nome := range m.NomesCidades {
		m.Cidades[i].CodCidade = uint8(i)
		m.Cidades[i].Nome = Nome
		m.Cidades[i].Ciclo = m.Ciclo
	}
	return true
}

func (m *Mundo) carregaPopulaçãoCidades() bool {
	if m.ArquivoPopulaçãoCidades == "" {
		log.Print("erro falta passar o nome do ArquivoPopulaçãoCidades")
		return false
	} else {
		arquivo, err := os.Open(m.ArquivoPopulaçãoCidades)
		// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
		if err != nil {
			log.Fatalf("Erro ler nomes: %v", err)
			return false
		}
		// Garante que o arquivo sera fechado apos o uso
		defer arquivo.Close()
		// Cria um scanner que lê cada linha do arquivo
		var linhas []string
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linhas = append(linhas, scanner.Text())
		}
		m.PopulaçãoCidades = make([]int, len(linhas))
		if m.Cidades == nil {
			m.NumeroCidades = len(linhas)
			m.Cidades = make([]Cidade, len(linhas))
		}
		for i := range linhas {
			pop, err := strconv.ParseInt(linhas[i], 10, 64)
			if err != nil {
				log.Fatalf("Erro: %v", err)
				return false
			}
			m.PopulaçãoCidades[i] = int(pop)
			m.TamanhoPopulação += int(pop)
			m.Cidades[i].TamanhoPopulação = int(pop)
		}
	}
	return true
}

func (m *Mundo) carregaDistânciasCidades() bool {
	if m.ArquivoDistanciasCidades == "" {
		log.Print("erro falta passar o nome do arquivoNome")
		return false
	} else {
		arquivo, err := os.Open(m.ArquivoDistanciasCidades)
		// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
		if err != nil {
			log.Fatalf("Erro ler nomes: %v", err)
			return false
		}
		// Garante que o arquivo sera fechado apos o uso
		defer arquivo.Close()
		// Cria um scanner que lê cada linha do arquivo
		var linhas []string
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linhas = append(linhas, scanner.Text())
		}
		m.PopulaçãoCidades = make([]int, len(linhas))
		if m.Cidades == nil {
			m.NumeroCidades = int(math.Sqrt(float64(len(linhas))))
			m.Cidades = make([]Cidade, m.NumeroCidades)
		}
		distância := make([]float32, len(linhas))
		for i := range linhas {
			dist, err := strconv.ParseFloat(linhas[i], 32)
			if err != nil {
				log.Fatalf("Erro: %v", err)
				return false
			}
			distância[i] = float32(dist)
		}
		m.Distâncias = make([][]float32, m.NumeroCidades)
		for i := range m.Distâncias {
			m.Distâncias[i] = make([]float32, m.NumeroCidades)
		}
		k := 0
		for i := 0; i < m.NumeroCidades; i++ {
			for j := 0; j < m.NumeroCidades; j++ {
				m.Distâncias[i][j] = distância[k]
				k++
			}
		}
	}
	return true
}

//initProbabilidadeContagio inicaça a função de porbabilidade de contagio
func (m *Mundo) initProbabilidadeContagio() bool {
	p := make([]float32, m.NumeroVizinhos)
	for i := range p {
		p[i] = m.F(i)
	}
	m.ProbabilidadeContagio = make([]float32, m.NumeroVizinhos)
	copy(m.ProbabilidadeContagio, p)
	return true
}

func (m *Mundo) initProbabilidadeTroca() bool {
	if m.Distâncias == nil {
		m.carregaDistânciasCidades()
	}
	m.ProbabilidadeTroca = make([][]float32, m.NumeroCidades)
	for i := range m.ProbabilidadeTroca {
		m.ProbabilidadeTroca[i] = make([]float32, m.NumeroCidades)
	}
	for i := 0; i < m.NumeroCidades; i++ {
		for j := 0; j < m.NumeroCidades; j++ {
			m.ProbabilidadeTroca[i][j] = m.FTroca(m.Distâncias[i][j])
		}
	}

	return true
}

//deslocaPessoas simula o deslocamento aleatório de pessoas
// não pode ser paralelo por que usa a mesma memoria
func (m *Mundo) deslocaPessoas() {
	if m.Cidades == nil {
		m.init()
	}
	a := &m.População[rand.Intn(m.TamanhoPopulação)]
	b := &m.População[rand.Intn(m.TamanhoPopulação)]
	p := m.ProbabilidadeTroca[a.CodCidade][b.CodCidade]
	if rand.Float32() < p {
		a.Estado, b.Estado = b.Estado, a.Estado
		a.Dia, b.Dia = b.Dia, a.Dia
	}
}

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func lerTexto(caminhoDoArquivo string) ([]string, error) {
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que lê cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return linhas, scanner.Err()
}

//contamine é uma funçao contamina uma Pessoa localizada na Cidade passada como parametro
func (m *Mundo) contamine() {
	for i := range m.Cidades {
		if m.Cidades[i].Nome == m.CidadeInicial {
			y := rand.Intn(int(m.Cidades[i].TamanhoPopulação))
			m.Cidades[i].População[y].Estado = 1
			m.Cidades[i].População[y].Dia = 0
			m.Cidades[i].Contaminados++
		}
	}
}

//umDia  execulta a Mundo de 1 passo de Monte Carlo
func (m *Mundo) umDia() {
	var numCPU = runtime.NumCPU()
	var goroutines int
	m.Data++
	c := make(chan int, numCPU)
	for i := 0; i < m.NumeroCidades; i++ {
		go m.Cidades[i].propaga(&m.Data, &m.ProbabilidadeContagio, c)
		goroutines++
		if goroutines >= numCPU {
			<-c
			goroutines--
		}
	}
	for i := 0; i < goroutines; i++ {
		<-c
	}
}
