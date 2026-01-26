package tailer

import (
	"fmt"
	"log-analyzer/internal/analyzer"
	"log-analyzer/internal/utils"
	"strings"
	"time"

	"github.com/hpcloud/tail"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

func StartTailing(files []string, rules []analyzer.Rule) {
	fmt.Println(string(ColorCyan) + "\n--- GERÇEK ZAMANLI İZLEME MODU (Çıkış: Ctrl+C) ---" + string(ColorReset))
	activeFiles := 0
	for _, f := range files {
		go tailFile(f, rules)
		activeFiles++
	}
	if activeFiles == 0 {
		fmt.Println("İzlenecek dosya bulunamadı. ")
		return
	}
	select {}
}

func tailFile(file string, rules []analyzer.Rule) {
	t, err := tail.TailFile(file, tail.Config{Follow: true, Poll: true})
	if err != nil {
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
	fmt.Printf("%s[UYARI] %s | %s | Dosya: %s%s\n", color, timestamp, rule.Name, source, ColorReset)
	fmt.Printf(" İçerik: %s\n", logText)
}
