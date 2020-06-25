package epidemiologicalsimulation

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

//Mundo estura para armazenar rede longa distância
type Mundo struct {
	nomesCidades             []string
	populaçãoCidades         []int
	numeroCidades            int
	tamanhoPopulação         int
	população                []Pessoa
	cidades                  []Cidade
	distâncias               [][]float32
	data                     int
	probabilidadeTroca       [][]float32
	probabilidadeContagio    []float32
	cidadeInicial            string
	arquivoNomesCidades      string
	arquivoPopulaçãoCidades  string
	arquivoDistanciasCidades string
	tempoSimulação           int
	numeroVizinhos           int
	ciclo                    int

	// funcao de probabilidade da contaminaçao
	f func(int) float32
	// funcao de probabilidade da contaminaçao
	fTroca func(float32) float32
}

//init carrega do disco três arquivos com o nome da Cidade e a população e as distâncias
func (m *Mundo) init() bool {
	m.carregaNomeCidades()
	m.carregaPopulaçãoCidades()
	m.carregaDistânciasCidades()
	m.initProbabilidadeContagio()
	m.initProbabilidadeTroca()
	for i := 0; i < m.numeroCidades; i++ {
		m.tamanhoPopulação += m.cidades[i].init()
	}
	//criando as pessoas do mundo
	m.população = make([]Pessoa, m.tamanhoPopulação)

	//distribuindo a populacao mundial nas cidades
	inicio := 0
	for i := 0; i < m.numeroCidades; i++ {
		fim := int(m.cidades[i].tamanhoPopulação) + inicio
		m.cidades[i].população = m.população[inicio:fim]
		m.cidades[i].vizinhos()
		m.cidades[i].setPessoa()
		inicio = int(fim)
	}
	return true
}

func (m *Mundo) carregaNomeCidades() bool {
	if m.arquivoNomesCidades == "" {
		fmt.Println("nome do arquivoNome")
		return false
	} else {
		arquivo, err := os.Open(m.arquivoNomesCidades)
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
			m.nomesCidades = append(m.nomesCidades, scanner.Text())
		}
	}
	if m.cidades == nil {
		m.numeroCidades = len(m.nomesCidades)
		m.cidades = make([]Cidade, m.numeroCidades)
	}
	for i, nome := range m.nomesCidades {
		m.cidades[i].codCidade = uint8(i)
		m.cidades[i].nome = nome
		m.cidades[i].ciclo = m.ciclo
	}
	return true
}

func (m *Mundo) carregaPopulaçãoCidades() bool {
	if m.arquivoPopulaçãoCidades == "" {
		fmt.Println("nome do arquivoPopulaçãoCidades")
		return false
	} else {
		arquivo, err := os.Open(m.arquivoPopulaçãoCidades)
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
		m.populaçãoCidades = make([]int, len(linhas))
		if m.cidades == nil {
			m.numeroCidades = len(linhas)
			m.cidades = make([]Cidade, len(linhas))
		}
		for i := range linhas {
			pop, err := strconv.ParseInt(linhas[i], 10, 64)
			if err != nil {
				log.Fatalf("Erro: %v", err)
				return false
			}
			m.populaçãoCidades[i] = int(pop)
			m.tamanhoPopulação += int(pop)
			m.cidades[i].tamanhoPopulação = int(pop)
		}
	}
	return true
}

func (m *Mundo) carregaDistânciasCidades() bool {
	if m.arquivoDistanciasCidades == "" {
		fmt.Println("nome do arquivoNome")
		return false
	} else {
		arquivo, err := os.Open(m.arquivoDistanciasCidades)
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
		m.populaçãoCidades = make([]int, len(linhas))
		if m.cidades == nil {
			m.numeroCidades = int(math.Sqrt(float64(len(linhas))))
			m.cidades = make([]Cidade, m.numeroCidades)
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
		m.distâncias = make([][]float32, m.numeroCidades)
		for i := range m.distâncias {
			m.distâncias[i] = make([]float32, m.numeroCidades)
		}
		k := 0
		for i := 0; i < m.numeroCidades; i++ {
			for j := 0; j < m.numeroCidades; j++ {
				m.distâncias[i][j] = distância[k]
				k++
			}
		}
	}
	return true
}

//initProbabilidadeContagio inicaça a função de porbabilidade de contagio
func (m *Mundo) initProbabilidadeContagio() bool {
	p := make([]float32, m.numeroVizinhos)
	for i := range p {
		p[i] = m.f(i)
	}
	m.probabilidadeContagio = make([]float32, m.numeroVizinhos)
	copy(m.probabilidadeContagio, p)
	return true
}

func (m *Mundo) initProbabilidadeTroca() bool {
	if m.distâncias == nil {
		m.carregaDistânciasCidades()
	}
	m.probabilidadeTroca = make([][]float32, m.numeroCidades)
	for i := range m.probabilidadeTroca {
		m.probabilidadeTroca[i] = make([]float32, m.numeroCidades)
	}
	for i := 0; i < m.numeroCidades; i++ {
		for j := 0; j < m.numeroCidades; j++ {
			m.probabilidadeTroca[i][j] = m.fTroca(m.distâncias[i][j])
		}
	}

	return true
}

//deslocaPessoas simula o deslocamento aleatório de pessoas
// não pode ser paralelo por que usa a mesma memoria
func (m *Mundo) deslocaPessoas() {
	if m.cidades == nil {
		m.init()
	}
	a := &m.população[rand.Intn(m.tamanhoPopulação)]
	b := &m.população[rand.Intn(m.tamanhoPopulação)]
	p := m.probabilidadeTroca[a.codCidade][b.codCidade]
	if rand.Float32() < p {
		a.estado, b.estado = b.estado, a.estado
		a.dia, b.dia = b.dia, a.dia
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
	for i := range m.cidades {
		if m.cidades[i].nome == m.cidadeInicial {
			y := rand.Intn(int(m.cidades[i].tamanhoPopulação))
			m.cidades[i].população[y].estado = 1
			m.cidades[i].população[y].dia = 0
			m.cidades[i].contaminados++
		}
	}
}

//umDia  execulta a Mundo de 1 passo de Monte Carlo
func (m *Mundo) umDia() {
	var numCPU = runtime.NumCPU()
	var goroutines int
	m.data++
	c := make(chan int, numCPU)
	for i := 0; i < m.numeroCidades; i++ {
		go m.cidades[i].propaga(&m.data, &m.probabilidadeContagio, c)
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
