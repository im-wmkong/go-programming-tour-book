package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
)

var (
	calculateTime string
	duration      string
)

var (
	timeCmd = &cobra.Command{
		Use:   "time",
		Short: "时间格式处理",
		Long:  "时间格式处理",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	nowTimeCmd = &cobra.Command{
		Use:   "now",
		Short: "获取当前时间",
		Long:  "获取当前时间",
		Run:   nowTimeCmdRun,
	}
	calculateTimeCmd = &cobra.Command{
		Use:   "calc",
		Short: "计算所需时间",
		Long:  "计算所需时间",
		Run:   calculateTimeCmdRun,
	}
)

func nowTimeCmdRun(cmd *cobra.Command, args []string) {
	nowTime := timer.GetNowTime()
	log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
}

func calculateTimeCmdRun(cmd *cobra.Command, args []string) {
	var currentTimer time.Time
	var layout = "2006-01-02 15:04:05"
	if calculateTime == "" {
		currentTimer = timer.GetNowTime()
	} else {
		var err error
		if !strings.Contains(calculateTime, " ") {
			layout = "2006-01-02"
		}
		currentTimer, err = time.Parse(layout, calculateTime)
		if err != nil {
			t, _ := strconv.Atoi(calculateTime)
			currentTimer = time.Unix(int64(t), 0)
		}
	}
	t, err := timer.GetCalculateTime(currentTimer, duration)
	if err != nil {
		log.Fatalf("timer.GetCalculateTime err: %v", err)
	}
	log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算时间，有效单位为时间戳或以格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
