//go:build debug

package debug

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

var (
	VariableList = []Variable{}
	PrintBuffer  = strings.Builder{}
	TitleColor   = "\033[1;36m"
	// TODO ColorList
	ColorList = []string{}
	MaxLength = 28

	TopNotationList    = []string{"╭", "─", "┬", "╮\n"}
	MiddleNotationList = []string{"├", "─", "┼", "┤\n"}
	BottomNotationList = []string{"╰", "─", "┴", "╯\n"}
)

// TODO Variable size
type Variable struct {
	Name string
	// TODO use pointer will get better performance ?
	variable any
	Addr     string
}

func CheckVar(name string, a any) {

	_, file, line, _ := runtime.Caller(1)
	rel, err := filepath.Rel(ExecutePath(), file)
	if err != nil {
		panic(err.Error())
		return
	}
	VariableList = append(VariableList, Variable{Name: name, variable: a,
		Addr: fmt.Sprintf("%s:%d", rel, line)})
}

func ExecutePath() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(path)
}

func PrintVariables() {
	lengthList := []int{0, 0, 0, 0}
	for _, v := range VariableList {
		if len(v.Name) > lengthList[0] {
			lengthList[0] = len(v.Name)
		}
		if len(reflect.TypeOf(v.variable).String()) > lengthList[1] {
			lengthList[1] = len(reflect.TypeOf(v.variable).String())
		}
		if len(fmt.Sprint(v.variable)) > lengthList[2] && len(fmt.Sprint(v.variable)) < MaxLength {
			lengthList[2] = len(fmt.Sprint(v.variable))
		}
		if len(v.Addr) > lengthList[3] {
			lengthList[3] = len(v.Addr)
		}
	}

	DrawTop(lengthList)
	DrawTopContent([]string{"Name", "Type", "Value", "Addr"}, lengthList)
	DrawMiddle(lengthList)
	// DrawMiddleContent()
	for _, v := range VariableList {
		str := fmt.Sprint(v.variable)
		if len(str) > lengthList[2] {
			PrintBuffer.WriteString(fmt.Sprintf("│ %-*s │ %-*s │ %-*s │ %-*s │\n",
				lengthList[0],
				v.Name,
				lengthList[1], reflect.TypeOf(v.variable).String(), lengthList[2],
				str[:lengthList[2]-3]+"...", lengthList[3], v.Addr))
		} else {
			PrintBuffer.WriteString(fmt.Sprintf("│ %-*s │ %-*s │ %-*v │ %-*s │\n",
				lengthList[0], v.Name, lengthList[1], reflect.TypeOf(v.variable).String(), lengthList[2], v.variable, lengthList[3], v.Addr))
		}
	}
	DrawBottom(lengthList)
	fmt.Println(PrintBuffer.String())
	PrintBuffer.Reset()
}

func MiddleAlign(ctx string, length int) string {
	return fmt.Sprintf("%*s%s%*s", (length-len(ctx))/2, "", ctx, (length - (length-len(ctx))/2 - len(ctx)), "")
}

// top     ╭ ─ ┬ ╮
// middle  ├ ─ ┼ ┤
// bottom  ╰ ─ ┴ ╯

// TODO Render more color
func Render(ctx string, rgb string) string {
	return rgb + ctx + "\033[0m"
}

// 这样其实还是有意义的我觉得
func DrawTop(lengthList []int) {
	DrawLine(0, lengthList)
}

func DrawMiddle(lengthList []int) {
	DrawLine(1, lengthList)
}

func DrawBottom(lengthList []int) {
	DrawLine(2, lengthList)
}

func DrawLine(__type__ int, lengthList []int) {
	notion := []string{}
	switch __type__ {
	case 0:
		notion = TopNotationList
	case 1:
		notion = MiddleNotationList
	case 2:
		notion = BottomNotationList
	default:
		panic("invalid type")
	}
	PrintBuffer.WriteString(notion[0])
	for i := range lengthList {
		for j := 0; j < lengthList[i]+2; j++ {
			PrintBuffer.WriteString(notion[1])
		}
		if i != len(lengthList)-1 {
			PrintBuffer.WriteString(notion[2])
		}
	}
	PrintBuffer.WriteString(notion[3])
}
