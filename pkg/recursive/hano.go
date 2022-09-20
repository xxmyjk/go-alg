package recursive

import "fmt"

type Tower struct {
	Data []int
}

// mv from t1 to t2
func MoveTo(t1, t2 *Tower) {
	// if len(t1.Data) == 0 || len(t2.Data) == 0 || t1.Data[0] > t2.Data[0] {
	// 	// this should never happen.
	// 	// fmt.Println("not move")
	// 	return
	// }

	t2.Data = append([]int{t1.Data[0]}, t2.Data...)
	t1.Data = t1.Data[1:]
}

func NewTower(length int) *Tower {
	t := &Tower{}

	for i := 1; i <= length; i++ {
		t.Data = append(t.Data, i)
	}

	return t
}

type H struct {
	Count   int
	A, B, C *Tower
}

func (h *H) Print(A, B, C *Tower) {
	// print
	h.Count++
	fmt.Printf("------%d------\n", h.Count)

	fmt.Print("A: ")
	for i, num := range A.Data {
		fmt.Print(i, " [", num, "] ")
	}
	fmt.Println()

	fmt.Print("B: ")
	for i, num := range B.Data {
		fmt.Print(i, " [", num, "] ")
	}
	fmt.Println()

	fmt.Print("C: ")
	for i, num := range C.Data {
		fmt.Print(i, " [", num, "] ")
	}
	fmt.Println()

	fmt.Printf("------%d------\n", h.Count)
}

func (h *H) Hano(A, B, C *Tower, n int) {
	// A --> C
	if n == 1 {
		MoveTo(A, C)
		h.Print(A, B, C)
		return
	}

	// A --> B
	h.Hano(A, C, B, n-1)
	MoveTo(A, C)
	h.Hano(B, A, C, n-1)
}
