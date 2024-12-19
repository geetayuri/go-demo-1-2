package output

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nguitarpb/7-solutions/somemodules/entities"
)

type outputUse struct {
	InputRepo entities.Input
}

// Constructor
func NewUsecase(inputRepo entities.Input) entities.Output {
	return &outputUse{
		InputRepo: inputRepo,
	}
}

func (u *outputUse) FirstOutput() {
	data := *(u.InputRepo.FirstInput())

	sum := data[0][0] // เริ่มจาก row 0 column 0
	index := 0        // index = 0

	// เริ่มจากแถวที่ 2
	for i := 1; i < len(data); i++ {
		// Get แถวปัจจุบัน
		row := data[i]

		// เลือกค่าสูงสุดจากแต่ละคู่ ถ้าตัวขวามากกว่าตัวซ้าย index จะ +1
		// fmt.Println(index)
		if index < len(row)-1 && row[index+1] > row[index] {
			index++ // index +1
		}

		// นำค่าที่เลือกมาบวก
		sum += row[index]
	}

	fmt.Println(sum)
}

func (u *outputUse) SecondOutput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string: ")
	encoded, _ := reader.ReadString('\n')
	encoded = strings.TrimSpace(encoded)

	var list []int
	list = append(list, 0)

	for _, char := range encoded {
		if char == '=' {
			list = append(list, list[len(list)-1])
		} else if char == 'L' {
			list = append(list, list[len(list)-1]-1)
		} else if char == 'R' {
			list = append(list, list[len(list)-1]+1)
		}
	}

	minValue := list[0]
	for _, v := range list {
		if v < minValue {
			minValue = v
		}
	}

	var preList []int
	for _, v := range list {
		preList = append(preList, v-minValue)
	}

	minimizedList := minimizeListSum(preList)
	// fmt.Println(minimizedList)

	var result string
	for _, num := range minimizedList {
		result += fmt.Sprintf("%d", num)
	}

	fmt.Println(result)
}

func minimizeListSum(originalList []int) []int {
	n := len(originalList)
	newList := make([]int, n)

	// fmt.Println(originalList)
	// fmt.Println(newList)

	newList[0] = 0

	for i := 1; i < n; i++ {
		if originalList[i-1] < originalList[i] {
			newList[i] = newList[i-1] + 1
		} else if originalList[i-1] > originalList[i] {
			j := i
			k := 0
			for j != 0 && newList[j-1] <= k && originalList[j-1] >= originalList[j] {
				newList[j-1] = newList[j-1] + 1
				j--
				k++
			}
			newList[i] = 0
		} else {
			newList[i] = newList[i-1]
		}
		// fmt.Println(newList)
	}

	return newList
}
