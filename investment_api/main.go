package main

func main() {

}

var _ CreditAssigner = (*CreditAssignment)(nil)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignment struct {
}

func (*CreditAssignment) Assig(investment int32) (x300 int32, x500 int32, x700 int32, err error) {
	return
}
