package input

import (
	"github.com/nguitarpb/7-solutions/somemodules/entities"
    "encoding/json"
    "os"
    "io"
    "fmt"
)

type Input struct {}

func NewInput() entities.Input {
	return &Input{}
}

func (r *Input) FirstInput() (*[][]int) {
    file, err := os.Open("./hard.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var firstinput [][]int
	json.Unmarshal(data, &firstinput)

    // fmt.Println(firstinput)

	// firstinput := [][]int{
    //     {59},
    //     {73, 41},
    //     {52, 40, 53},
    //     {26, 53, 6, 34},
    // }

	return &firstinput
}