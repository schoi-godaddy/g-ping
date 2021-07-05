package cmd

import (
	"fmt"
	"g-ping/utils"
	"math"
	"strings"

	"github.com/fatih/color"
	"github.com/go-ping/ping"
	"github.com/gosuri/uilive"
	"github.com/spf13/cobra"
)

var (
	Target string
	q      []int

	maxWidth  = 40
	maxHeight = 30

	yelloCutover = 30
	redCutover   = 60

	blue   = color.New(color.FgBlue).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

func init() {
	q = make([]int, 0)
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&Target, "target", "t", "", "Target to ping (required)")
	startCmd.MarkFlagRequired("target")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "(Short) Start g-ping",
	Long:  "(Long) Starts g-ping command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		pinger, err := ping.NewPinger(Target)
		pinger.SetPrivileged(true)

		if err != nil {
			return err
		}
		writer := uilive.New()
		writer.Start()

		pinger.OnRecv = onReceive(writer)

		if err := pinger.Run(); err != nil {
			return err
		}

		return nil
	},
}

func onReceive(writer *uilive.Writer) func(pkt *ping.Packet) {
	return func(pkt *ping.Packet) {
		q = append(q, int(pkt.Rtt.Milliseconds()))

		if len(q) == maxWidth {
			q = q[1:]
		}

		var dy int
		board := make([][]string, 0)
		min, max, total, avg := utils.GetStats(q)
		if float64(max) > avg*1.5 {
			dy = int(math.Round(float64(max) / float64(maxHeight)))
		} else {
			dy = int(math.Round(avg / (float64(maxHeight) / 2.0)))
		}

		for _, n := range q {
			temp := make([]string, 0)
			for i := 0; i < maxHeight; i++ {
				c := " "
				if i*dy < n {
					c = "#"
				}

				if i*dy >= redCutover {
					temp = append([]string{red(c)}, temp...)
				} else if i*dy >= yelloCutover {
					temp = append([]string{yellow(c)}, temp...)
				} else {
					temp = append([]string{green(c)}, temp...)
				}
			}
			board = append(board, temp)
		}

		if len(board) == maxWidth {
			board = board[1:]
		}

		newBoard, _ := utils.Zip(board...)

		var s string

		s += fmt.Sprintf("%s<%dms %s<%dms %s>%dms\n", green("Green"), yelloCutover, yellow("Yellow"), redCutover, red("Red"), redCutover)
		s += fmt.Sprintf("%d bytes from %s(%s): icmp_seq=%d time=%s\n",
			pkt.Nbytes, Target, yellow(pkt.IPAddr), pkt.Seq, blue(pkt.Rtt))
		s += fmt.Sprintf("min=%d max=%d total=%d avg=%.2f dy=%d\n\n", min, max, total, avg, dy)

		for _, b := range newBoard {
			s += strings.Join(b, " ")
			s += "\n"
		}

		fmt.Fprintf(writer, "%s\n", s)
	}
}
