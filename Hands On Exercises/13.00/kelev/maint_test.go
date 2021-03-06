package kelev

import (
	"testing"
	"fmt"
)

func TestYears(t *testing.T)  {
	n := Years(10)
	if n != 70 {
		t.Error("Got", n, "Expected", 70)
	}
}

func TestYearsTwo(t *testing.T)  {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("Got", n, "Expected", 70)
	}
}

func ExampleYears()  {
	fmt.Println(Years(10))
	// Output
	// 70
}

func ExampleYearsTwo()  {
	fmt.Println(YearsTwo(10))
	// Output
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B)  {
	YearsTwo(10)
}
