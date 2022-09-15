package matrix

import (
	"fmt"
)

type Matrix struct {
	Data [][]int
	El   []int
	M    int
	N    int
}

func (m *Matrix) GetRow(i int) []int {
	if i >= m.M {
		msg := fmt.Sprintf("row index [%d] out of range [%d] \n", i, m.M)
		panic(msg)
	}

	return m.Data[i]
}

func (m *Matrix) GetCol(j int) []int {
	if j >= m.N {
		msg := fmt.Sprintf("col index [%d] out of range [%d] \n", j, m.N)
		panic(msg)
	}

	col := make([]int, m.M)
	for i := 0; i < m.M; i++ {
		col[i] = m.Data[i][j]
	}

	return col
}

func (m *Matrix) AppendRow(x []int) {

}

func (m *Matrix) Print() {
	fmt.Println("------ normal print ------")
	for i := 0; i < m.M; i++ {
		row := m.GetRow(i)
		for j := 0; j < m.N; j++ {
			fmt.Printf("%d\t", row[j])
		}
		fmt.Println()
	}
	fmt.Println("------ normal print ------")
}

func (m *Matrix) PrintRow(i int) {
	fmt.Println("------ row print ------")
	row := m.GetRow(i)
	for j := 0; j < m.N; j++ {
		fmt.Printf("%d\t", row[j])
	}
	fmt.Println()
	fmt.Println("------ row print ------")
}

func (m *Matrix) PrintCol(j int) {
	fmt.Println("------ col print ------")
	col := m.GetCol(j)
	for i := 0; i < m.M; i++ {
		fmt.Printf("%d\n", col[i])
	}
	fmt.Println("------ col print ------")
}

func (m *Matrix) SnakeLikePrint() {
	// TODO: print with start element
	// start from(0, 0)

}

func (m *Matrix) CircleLikePrint() {
	// TODO: print with start element
	// start from(0, 0)

	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	// right -> down -> left -> up
	for count := 0; count <= m.M+m.N+1; count++ {
		switch count % 4 {
		case 0:
			// right
			row := m.GetRow(0)
			fmt.Println(row)
			m.Data = m.Data[1:m.M]
			m.M--
		case 1:
			// down
			col := m.GetCol(m.N - 1)
			fmt.Println(col)
			for i := 0; i < m.M; i++ {
				m.Data[i] = m.Data[i][:m.N-1]
			}
			m.N--
		case 2:
			// left
			row := m.GetRow(m.M - 1)
			reverse(row)
			fmt.Println(row)
			m.Data = m.Data[:m.M-1]
			m.M--
		case 3:
			// up
			col := m.GetCol(0)
			reverse(col)
			fmt.Println(col)
			for i := 0; i < m.M; i++ {
				m.Data[i] = m.Data[i][1:]
			}
			m.N--
		default:
			fmt.Println("end")
		}
	}
}

func New(row_count int, col_count int, el ...int) *Matrix {
	if row_count*col_count != len(el) {
		msg := fmt.Sprintf("row_count * col_count != len(el) [%d*%d != %d] \n", row_count, col_count, len(el))
		panic(msg)
	}

	m := &Matrix{
		M:  row_count,
		N:  col_count,
		El: el,
	}

	for len(el) != 0 {
		m.Data = append(m.Data, el[:col_count])
		el = el[col_count:]
	}

	return m
}
