package simplearraysum_test

import (
	"hackerrank_go/simplearraysum"
	"testing"
)

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

func TestSimpleArraySum(t *testing.T) {
	var valorEsperado int32 = 31
	arr := []int32{1, 2, 3, 4, 10, 11}
	valor := simplearraysum.SimpleArraySum(arr)
	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}

}
