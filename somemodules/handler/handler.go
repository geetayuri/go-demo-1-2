package handlers

import (
	"github.com/nguitarpb/7-solutions/somemodules/entities"
	"github.com/nguitarpb/7-solutions/somemodules/output"
	"github.com/nguitarpb/7-solutions/somemodules/input"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Handler struct {
	output entities.Output
}

func NewHandler() *Handler {
	repo := input.Input{}
	uc := output.NewUsecase(&repo)

	return &Handler{output: uc}
}

func (h *Handler) StartApplication() {
	fmt.Println("Handler: Starting application...")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Exam Answer you need (1 or 2): ")
	encoded, _ := reader.ReadString('\n')
	encoded = strings.TrimSpace(encoded)
	if(encoded == "1"){
		h.output.FirstOutput()
	}else if(encoded == "2"){
		h.output.SecondOutput()
	}else{
		fmt.Print("Wrong input.")
	}

	// data := h.output.FirstOutput()
	// fmt.Println(*data)
}