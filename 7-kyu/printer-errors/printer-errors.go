package kata

import "fmt"

func PrinterError(s string) string {
	erros := 0
	for _, char := range s {
		posicao := char - 'a'
		if posicao > 12 {
			erros++
		}
	}
	return fmt.Sprintf("%d/%d", erros, len(s))
}
