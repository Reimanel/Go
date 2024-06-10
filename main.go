package main

import "fmt"

func main() {
	// Definir os pontos dos times manualmente
	timeA := 3
	timeB := 15

	// Calcular combinações possíveis para cada time
	combinacoesTimeA := calcularCombinacoes(timeA)
	combinacoesTimeB := calcularCombinacoes(timeB)

	// Imprimir as combinações possíveis
	fmt.Println("Combinações possíveis:")
	fmt.Println("Time A (", timeA, "pontos):", combinacoesTimeA)
	fmt.Println("Time B (", timeB, "pontos):", combinacoesTimeB)
}

// Função para calcular as combinações possíveis para uma determinada pontuação
func calcularCombinacoes(pontos int) int {
	combinacoes := 0

	// Usar loops para contar combinações válidas
	for touchdowns := 0; touchdowns <= pontos/6; touchdowns++ {
		for extras := 0; extras <= touchdowns*2; extras++ { // Extras podem ser 0, 1 ou 2 por touchdown
			for fieldGoals := 0; fieldGoals <= pontos/3; fieldGoals++ {
				// Calcular a pontuação total
				total := touchdowns*6 + extras + fieldGoals*3
				// Verificar se a pontuação total é igual aos pontos necessários
				if total == pontos {
					combinacoes++
				}
			}
		}
	}

	return combinacoes
}
