//go:build debug

package debug

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Caller struct {
	start    time.Time
	duration time.Duration
	name     string
	caller   string
	Addr     string
}

var (
	CallerList = []*Caller{}
	MainStart  = time.Now()
)

func CallCost() *Caller {
	pc1, _, _, _ := runtime.Caller(1)
	pc2, file, line, _ := runtime.Caller(2)
	rel, err := filepath.Rel(ExecutePath(), file)
	if err != nil {
		panic(err.Error())
	}
	caller := &Caller{start: time.Now(), name: runtime.FuncForPC(pc1).Name(), caller: runtime.FuncForPC(pc2).Name(), Addr: fmt.Sprintf("%s:%d", rel, line)}
	CallerList = append(CallerList, caller)
	return caller
}

func (t *Caller) Log() {
	t.duration = time.Now().Sub(t.start)
}

func PrintTimerLogs() {
	titleList := []string{"Name", "Caller", "Cost", "Rate", "Addr"}
	lengthList := []int{0, 0, 0, 0, 0}
	for _, v := range CallerList {
		if len(v.name) > lengthList[0] {
			lengthList[0] = len(v.name)
		}
		if len(v.caller) > lengthList[1] {
			lengthList[1] = len(v.caller)
		}
		if len(fmt.Sprint(v.duration)) > lengthList[2] {
			lengthList[2] = len(fmt.Sprint(v.duration))
		}
		if len(fmt.Sprintf("%.4f",
			float64(v.duration.Nanoseconds()*1000/time.Now().Sub(MainStart).Nanoseconds())/1000)) > lengthList[3] {
			lengthList[3] = len(fmt.Sprintf("%.4f",
				float64(v.duration.Nanoseconds()*1000/time.Now().Sub(MainStart).Nanoseconds())/1000))
		}
		if len(v.Addr) > lengthList[4] {
			lengthList[4] = len(v.Addr)
		}
	}
	DarwCallerTable(titleList, lengthList)
}

func DarwCallerTable(titleList []string, lengthList []int) {
	DrawTop(lengthList)
	DrawTopContent(titleList, lengthList)
	DrawMiddle(lengthList)
	for _, v := range CallerList {
		PrintBuffer.WriteString(fmt.Sprintf("│ %s │ %s │ %s │ %s │ %s │\n",
			MiddleAlign(fmt.Sprint(v.name), lengthList[0]),
			MiddleAlign(fmt.Sprint(v.caller), lengthList[1]),
			MiddleAlign(fmt.Sprint(v.duration), lengthList[2]),
			MiddleAlign(fmt.Sprintf("%.4f", float64(v.duration.Nanoseconds()*1000/time.Now().Sub(MainStart).Nanoseconds())/1000), lengthList[3]),
			MiddleAlign(fmt.Sprint(v.Addr), lengthList[4])))
	}
	DrawBottom(lengthList)
	fmt.Println(PrintBuffer.String())
}

func DrawTopContent(titleList []string, lengthList []int) {
	for i := range titleList {
		PrintBuffer.WriteString("│")
		PrintBuffer.WriteString(Render(MiddleAlign(titleList[i],
			lengthList[i]+2), TitleColor))
	}
	PrintBuffer.WriteString("│\n")
}
