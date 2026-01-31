package tailer

import (
	"fmt"
	"strings"
	"time"

	"log-analyzer/internal/analyzer"
	"log-analyzer/internal/utils"

	"github.com/hpcloud/tail"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
)

func StartTailing(files []string, rules []analyzer.Rule) {
	fmt.Println(string(ColorCyan) + "\n--- REAL-TIME MONITORING MODE (Exit: Ctrl+C) ---" + string(ColorReset))

	activeFiles := 0
	for _, f := range files {
		fmt.Printf("%s[STARTED]%s Watching: %s\n", ColorGreen, ColorReset, f)
		go tailFile(f, rules)
		activeFiles++
	}

	if activeFiles == 0 {
		fmt.Println("No files found to monitor.")
		return
	}
	select {}
}

func tailFile(file string, rules []analyzer.Rule) {
	t, err := tail.TailFile(file, tail.Config{Follow: true, ReOpen: true, Poll: true})
	if err != nil {
		fmt.Printf("Error tailing file %s: %v\n", file, err)
		return
	}

	for line := range t.Lines {
		txt := utils.CleanLogLine(line.Text)
		if txt == "" {
			continue
		}

		for _, rule := range rules {
			if strings.Contains(strings.ToLower(txt), strings.ToLower(rule.Keyword)) {
				alert(file, rule, txt)
			}
		}
	}
}

func alert(source string, rule analyzer.Rule, logText string) {
	timestamp := time.Now().Format("15:04:05")
	color := ColorYellow
	if rule.Level == "CRITICAL" {
		color = ColorRed
	}
	fmt.Printf("%s[ALERT] %s | %s | File: %s%s\n", color, timestamp, rule.Name, source, ColorReset)
	fmt.Printf("   └── Content: %s\n", logText)
}
