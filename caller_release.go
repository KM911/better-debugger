//go:build release

package debug

type Caller struct{}

func CallCost() *Caller { return nil }
func (t *Caller) Log() {
}
func PrintTimerLogs() {
}
