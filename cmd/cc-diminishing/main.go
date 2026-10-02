// Command cc-diminishing 跑控制递减样例。
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"example.com/control"
)

// Run 执行一次命令行调用，返回退出码。
func Run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("cc-diminishing", flag.ContinueOnError)
	flags.SetOutput(stderr)
	sample := flags.String("sample", "curve", "curve / category / window / scan")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	switch *sample {
	case "curve":
		tracker := control.New(10000, 3000)
		first := tracker.Apply("u1", "stun", 1000, 0)
		second := tracker.Apply("u1", "stun", 1000, 1)
		third := tracker.Apply("u1", "stun", 1000, 2)
		fourth := tracker.Apply("u1", "stun", 1000, 3)
		fmt.Fprintf(stdout, "first=%d second=%d third=%d fourth=%d\n", first, second, third, fourth)
	case "category":
		tracker := control.New(10000, 3000)
		tracker.Apply("u1", "stun", 1000, 0)
		root := tracker.Apply("u1", "root", 1000, 1)
		fmt.Fprintf(stdout, "root=%d\n", root)
	case "window":
		tracker := control.New(10000, 3000)
		tracker.Apply("u1", "stun", 1000, 0)
		second := tracker.Apply("u1", "stun", 1000, 10001)
		fmt.Fprintf(stdout, "second=%d\n", second)
	case "scan":
		tracker := control.New(10000, 3000)
		for index := 0; index < 3000; index++ {
			tracker.Apply(fmt.Sprintf("u-%d", index), "stun", 1000, 0)
		}
		for index := 0; index < 3000; index++ {
			tracker.Apply("u-0", "stun", 1000, int64(index))
		}
		fmt.Fprintf(stdout, "scanned=%d\n", tracker.Scanned())
	default:
		fmt.Fprintln(stderr, "需要 --sample curve|category|window|scan")
		return 2
	}
	return 0
}

func main() {
	os.Exit(Run(os.Args[1:], os.Stdout, os.Stderr))
}
