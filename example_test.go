package debug

import (
	"fmt"
	"testing"
	"time"
)

type Student struct {
	Name string
	Age  int
}

// Must implement Stringer interface
func (m Student) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", m.Name, m.Age)
}

func CallFunction() {
	defer CallCost().Log()
	stu := Student{
		Name: "KM",
		Age:  18,
	}
	CheckVar("stu", stu)
	time.Sleep(time.Second)
	SubFunction()
}

func SubFunction() {
	defer CallCost().Log()
	LongString := "This is a long string for test if max length of string will be cut off."
	CheckVar("LongString", LongString)

	time.Sleep(time.Second)

}

// with -tags debug
func TestExample(t *testing.T) {
	defer PrintTimerLogs()
	defer PrintVariables()
	age := 18
	CheckVar("age", age)
	CallFunction()
}
