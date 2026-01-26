package analyzer

import (
	"bufio"
	"fmt"
	"log-analyzer/internal/utils"
	"os"
	"strings"
	"time"
)

type Config struct {
	Settings struct {
		OutputPath string `yaml:"output_path"`
	} `yaml:"settings"`
	Files []string `yaml:"files"`
	Rules []Rule   `yaml:"rules"`
}

type Rule struct {
	Name    string `yaml:"name"`
	Keyword string `yaml:"keyword"`
	Level   string `yaml:"level"`
}

type LogResult struct {
	Timestamp string
	Level     string
	Message   string
	Source    string
}

func AnalyzeFiles(files []string, rules []Rule) ([]LogResult, error) {
	var results []LogResult

	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("BİLGI: Dosya bulunamadı veya erişilemiyor (%s) - Atlanıyor.\n", filePath)
			continue
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := utils.CleanLogLine(scanner.Text())
			if line == "" {
				continue
			}

			for _, rule := range rules {
				if strings.Contains(strings.ToLower(line), strings.ToLower(rule.Keyword)) {
					results = append(results, LogResult{
						Timestamp: time.Now().Format("2006-01-02 15:04:05"),
						Level:     rule.Level,
						Message:   fmt.Sprintf("[%s] tespit edildi: %s", rule.Name, line),
						Source:    filePath,
					})
				}
			}
		}
		file.Close()
	}
	return results, nil
}
