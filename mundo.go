package epidemiologicalsimulation

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

//Mundo estura para armazenar rede longa distância
type mundo struct {
	numeroCidades            int
	tamanhoPopulaçãoQuadrada int
	tamanhoPopulação         int
	população                []pessoa
	cidades                  []cidade
	distâncias               [][]float32
	data                     int
	probabilidadeTroca       [][]float32
	probabilidadeContagio    []float32
	cidadeInicial            string
}

//init carrega do disco três arquivos com o nome da cidade e a população e as distâncias
func (m *mundo) init(s Simulação) {
	var populacaoQuadrada int
	var populaçãoTotal int
	nomes, err := lerTexto(s.arquivoNomes)
	if err != nil {
		log.Fatalf("Erro:", err)
	}
	populaçãoCidade, err := lerTexto(s.arquivoPopulacao)
	if err != nil {
		log.Fatalf("Erro:", err)
	}
	m.cidades = make([]cidade, m.numeroCidades)
	for i, nome := range nomes {
		pop, err := strconv.ParseInt(populaçãoCidade[i], 10, 64)
		if err != nil {
			log.Fatalf("Erro:", err)
		}
		populaçãoTotal += int(pop)
		populacaoQuadrada += m.cidades[i].init(nome, int(pop), uint8(i))
	}
	m.tamanhoPopulaçãoQuadrada = populacaoQuadrada
	m.tamanhoPopulação = populaçãoTotal
	m.população = make([]pessoa, m.tamanhoPopulaçãoQuadrada)
	inicio := 0
	for i := 0; i < m.numeroCidades; i++ {
		fim := m.cidades[i].tamanhoPopulaçãoQuadrada
		m.cidades[i].população = m.população[inicio:fim]
		m.cidades[i].vizinhos()
		m.cidades[i].setPessoa()
		inicio = int(fim)
	}
	m.initProbabilidadeContagio(s.f)
}

//initProbabilidadeContagio inicaça a função de porbabilidade de contagio
func (m *mundo) initProbabilidadeContagio(f func(int) float32) {
	p := make([]float32, 5)
	for i := range p {
		p[i] = f(i)
	}
	copy(m.probabilidadeContagio, p)
}

//deslocaPessoas simula o deslocamento aleatório de pessoas
// não pode ser paralelo por que usa a mesma memoria
func (m *mundo) deslocaPessoas() {
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

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return linhas, scanner.Err()
}

//contamine é uma funçao contamina uma pessoa localizada na cidade passada como parametro
func (m *mundo) contamine() {
	for i := range m.cidades {
		if m.cidades[i].nome == m.cidadeInicial {
			y := rand.Intn(int(m.cidades[i].tamanhoPopulação))
			m.cidades[i].população[y].estado = 1
			m.cidades[i].população[y].dia = 0
		}
	}
}

//umaVolta  execulta a simulação de 1 passo de Monte Carlo
func (m *mundo) umaVolta(data *int) {
	var numCPU = runtime.NumCPU()
	var goroutines int
	c := make(chan int, numCPU)
	for i := 0; i < m.numeroCidades; i++ {
		go m.cidades[i].propaga(data, &m.probabilidadeContagio, c)
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
