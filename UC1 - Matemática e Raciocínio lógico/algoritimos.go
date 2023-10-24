func dataComMaisQueSeteDias(dateString string) bool {
	layout := "2006-01-02"
	today := time.Now()

	date, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Erro ao fazer parsing da data: " + err.Error())
		return false
	}

	difference := today.Sub(date)
	sevenDays := 7 * 24 * time.Hour
	fmt.Println(difference)
	return difference > sevenDays
}

func main() {
	lista := []int{2, 3, 5, 7, 10, 13, 15, 19, 23, 30, 1, 4, 11, 99, 45, 35, 89, 90}
	posicao := buscaBinaria(lista, 23)
	fmt.Println(posicao)
}

func ordenaLista(listaNumeros []int) []int {
	sort.Slice(listaNumeros, func(i, j int) bool {
		return listaNumeros[i] < listaNumeros[j]
	})
	return listaNumeros
}

func buscaBinaria(lista []int, alvo int) (int, int) {
	listaOrdenada := ordenaLista(lista)
	menor := 0
	maior := len(listaOrdenada) - 1
	numeroTentativas := 0

	for menor <= maior {
		numeroTentativas++
		meio := (maior + menor) / 2
		attempt := listaOrdenada[meio]
		if attempt == alvo {
			return meio, numeroTentativas
		}
		if attempt > alvo {
			maior = meio - 1
		}
		if attempt < alvo {
			menor = meio + 1
		}
	}

	return -1, numeroTentativas
}

func maiorParaMenor(lista []int) []int {
	n := len(lista)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if lista[j] < lista[j+1] {
				lista[j], lista[j+1] = lista[j+1], lista[j]
			}
		}
	}

	return lista
}

defer fmt.Println("World")
	fmt.Println("Hello")
	
	
func main() {
	go count1()
	count2()
	count3()
}

func count1() {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func count2() {
	for i := 5; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func count3() {
	for i := 10; i < 15; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

package main
import ("fmt")

func main() {
  var i int

  fmt.Print("Digite um número: ")
  fmt.Scan(&i)
  fmt.Println("Seu número é:", i)
}