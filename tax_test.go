package tax

import (
	"math"
	"testing"
)

// Antes de rodar o teste é necessário iniciar o modulo ( go mod init )
// Em seguida pode ser utilizado o comando abaixo
// go test .
//
// Para ser mais detalhado pode ser utilizado a opção -v
// go test . -v
//
// Texte de cobertura de códgio, faz uma verificação se todas as situações foram testadas
// go test -coverprofile=coverage -v
//
// Para ter um "relatório" sobre qual parte não teve uma cobertura de teste pode ser executado
// go tool cover -html=coverage

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	res := CalculateTax(amount)

	if res != expected {
		t.Errorf("Expected %f but got %f", expected, res)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	lista := []calcTax{
		{0.0, 0.0},
		{500.0, 5.0},
		{999.9, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range lista {
		res := CalculateTax(item.amount)
		if res != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, res)
		}
	}
}

// Para rodar apenas os testes iniciados com BenchmarkCalc por ex execute comando abaixo
// go test -benchmem -count=2 -run=^$ -bench=BenchmarkCalc
// -bench=nome -> executa uma função específica ou todas que começarem com o que foi informado
// -benchmem -> exibe a memória utilizada
// -count=n -> executa o teste o nro de vezes que for informado
//
// Documentação das flags: https://pkg.go.dev/cmd/go#hdr-Testing_flags
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	// Cria uma seed de valores
	seed := []float64{-10.0, -5.0, -1.0, 0.0, 500.0, 1000.0, 1500.0, 19999.99, 20000.0, 25000.0, math.MaxFloat64}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		res := CalculateTax(amount)

		if amount <= 0 {
			if res != 0 {
				t.Errorf("For amount %f, expected 0 but got %f", amount, res)
			}
		} else if amount >= 20000.0 {
			if res != 20 {
				t.Errorf("For amount %f, expected 20 but got %f", amount, res)
			}
		} else if amount >= 1000.0 {
			if res != 10 {
				t.Errorf("For amount %f, expected 10 but got %f", amount, res)
			}
		} else {
			if res != 5 {
				t.Errorf("For amount %f, expected 5 but got %f", amount, res)
			}
		}
	})
}

// func FuzzCalculateTax(f *testing.F) {
// 	// Cria uma seed de valores
// 	seed := []float64{-10.0, -5.0, -1.0, 0.0, 500.0, 100.0, 1500.0}
// 	for _, amount := range seed {
// 		f.Add(amount)
// 	}

// 	f.Fuzz(func(t *testing.T, amount float64) {
// 		res := CalculateTax(amount)
// 		if amount <= 0 && res != 0 {
// 			t.Errorf("Received %f but expected 0", res)
// 		}
// 		if amount >= 20000.0 && res != 20 {
// 			t.Errorf("Received %f but expected 20", res)
// 		}
// 	})
// }

// func FuzzCalculateTax(f *testing.F) {
// 	// Cria uma seed de valores com limites e casos extremos
// 	seed := []float64{-1000.0, -1.0, 0.0, 999.0, 1000.0, 19999.0, 20000.0, 20000.01, 1e6}

// 	for _, amount := range seed {
// 		f.Add(amount)
// 	}

// 	f.Fuzz(func(t *testing.T, amount float64) {
// 		res := CalculateTax(amount)

// 		// Verifica casos negativos e zero
// 		if amount <= 0 && res != 0 {
// 			t.Errorf("Received %f for amount %f, expected 0", res, amount)
// 		}

// 		// Verifica faixa de 1000 a 20000
// 		if amount >= 1000 && amount < 20000 && res != 10.0 {
// 			t.Errorf("Received %f for amount %f, expected 10.0", res, amount)
// 		}

// 		// Verifica faixa acima de 20000
// 		if amount >= 20000 && res != 20.0 {
// 			t.Errorf("Received %f for amount %f, expected 20.0", res, amount)
// 		}
// 	})
// }
