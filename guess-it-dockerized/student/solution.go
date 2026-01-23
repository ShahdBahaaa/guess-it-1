package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// دالة للقيمة المطلقة
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// slice لتخزين الأرقام اللي دخلها المستخدم
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()

		// تحويل الرقم من string إلى int
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("ERROR: Invalid input")
			continue
		}

		numbers = append(numbers, num)

		// لو الرقم الأول → مش ممكن نحدد range
		if len(numbers) == 1 {
			continue
		}

		// لو الرقم الثاني → الفارق بين الرقمين
		if len(numbers) == 2 {
			diff := abs(numbers[1] - numbers[0])
			lower := numbers[1] - diff
			upper := numbers[1] + diff
			fmt.Println(lower, upper)
			continue
		}

		// لو الرقم >= 3 → ناخد آخر فرقين ونحسب المتوسط
		last := numbers[len(numbers)-1]
		prev := numbers[len(numbers)-2]
		prevPrev := numbers[len(numbers)-3]

		diff1 := abs(last - prev)
		diff2 := abs(prev - prevPrev)
		avgDiff := (diff1 + diff2) / 2

		lower := last - avgDiff
		upper := last + avgDiff
		fmt.Println(lower, upper)
	}
}
