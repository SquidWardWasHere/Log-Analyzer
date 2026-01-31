package main

import (
	"bufio"
	"fmt"
	"log"
	"log-analyzer/internal/analyzer"
	"log-analyzer/internal/report"
	"log-analyzer/internal/tailer"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var cfg analyzer.Config

func init() {
	f, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Config dosyası okunamadı: %v", err)
	}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Fatalf("YAML hatası: %v", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\033[H\033[2J")
		fmt.Println("=====================================")
		fmt.Println("     LOG ANALİZ VE UYARI SİSTEMİ     ")
		fmt.Println("=====================================")
		fmt.Println("1. Dosyaları Tara ve Raporla")
		fmt.Println("2. Canlı İzleme Modu")
		fmt.Println("3. Çıkış ")
		fmt.Println("\nŞeçiminiz: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			fmt.Println("\nAnaliz yapılıyor...")
			results, err := analyzer.AnalyzeFiles(cfg.Files, cfg.Rules)
			if err != nil {
				fmt.Println("Hata: ", err)
			} else {
				fmt.Printf("\nSonuç: %d adet şüpheli kayıt bulundu.\n", len(results))
				if len(results) > 0 {
					if err := report.ExportToCSV(cfg.Settings.OutputPath, results); err == nil {
						fmt.Printf("Rapor kaydedildi: %s\n", cfg.Settings.OutputPath)
					}
				}
			}
			fmt.Println("\nAna menüye dönmek için Enter'a bas...")
			reader.ReadString('\n')
		case "2":
			tailer.StartTailing(cfg.Files, cfg.Rules)

		case "3":
			fmt.Println("Çıkış yapılıyor...")
			return

		default:
			fmt.Println("Geçersiz seçim.")
			time.Sleep(1 * time.Second)
		}
	}
}
