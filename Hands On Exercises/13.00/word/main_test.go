package word

import "testing"

func TestUseCount(t *testing.T) {
	m := UseCount("ehod shtaim shelosh arba hamesh hamesh hamesh shesh sheva shmonay tesha eser")
	for k, v := range m {
		switch k {
		case ehod:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case shtaim:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case shelosh:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case arba:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case hamesh:
			if v!= 3 {
				t.Error("Received", n, "Expected", 3)
			}
		case shesh:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case sheva:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case shmonay:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case tesha:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		case eser:
			if v!= 1 {
				t.Error("Received", n, "Expected", 1)
			}
		}
	}
}

func ExampleCount()  {
	fmt.Println(Count("ehod shtaim shelosh arba hamesh shesh sheva shmonay tesha eser"))
	// Output:
	// 10
}

func BenchmarkCount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
	Count(poem.SongOfMyself)
	}
}

func BenchmarkUseCount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
	UseCount(poem.SongOfMyself)
	}
}
