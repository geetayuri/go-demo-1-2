package entities

type Output interface {
	FirstOutput()
	SecondOutput()
}

type Input interface {
	FirstInput() (*[][]int)
}