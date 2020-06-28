// example.go
package main

import (
    "log"
    "math/rand"
    "os"
    "time"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/drjoaobatista/epidemiologicalsimulation"
	
)

var nameItems = []string{"1", "2", "3", "4","1", "2", "3", "4" }
var seed = rand.NewSource(time.Now().UnixNano())
var aracaju = Cidade{
	nome:                "aracaju",
	codCidade:           1,
	tamanhoPopulação:    10010,
	numeroVizinhos:      6,
	ciclo:               10,
	numeroTrocaVizinhos: 100,
}

func randInt() []int {
    cnt := len(nameItems)
    r := make([]int, 0)
    for i := 0; i < cnt; i++ {
        r = append(r, int(seed.Int63()) % 50)
    }
    return r
}

func main() {
	aracaju.init()
// criando a populacao
aracaju.população = make([]Pessoa, aracaju.tamanhoPopulação)
// configurando as pessoas
aracaju.setPessoa()
// configurando os vizinhos numa rede quadrada
aracaju.vizinhos()
//contaminando primeira pessoa da cidade
aracaju.população[0].estado = 1
aracaju.população[0].dia = 0
aracaju.contaminados = 1
    // bar := charts.NewBar()
    // bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-bola"}, charts.ToolboxOpts{Show: true})
    // bar.AddXAxis(nameItems).
    //     AddYAxis("bolaA", []int{1,2,3,4}).
    //     AddYAxis("bolaB", randInt())
    // f, err := os.Create("bar.html")
    // if err != nil {
    //     log.Println(err)
    // }
	// bar.Render(f)
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-doida"})
	line.AddXAxis(nameItems).AddYAxis("bolaA", randInt(),  charts.AreaStyleOpts{Opacity: 0.2}, charts.LineOpts{Smooth: true})
	line.AddXAxis(nameItems).AddYAxis("bolaB", randInt())
	f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
	line.Render(f)
}