package RGR

import (
	"arch-lab4/parser"
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func BenchmarkParser(b *testing.B) {

	file, err := os.Create("./test.txt")
	if err != nil {
		log.Fatal(err)
	}

	linesToWrite := []string{"print Hello!", "delete hello l"}

	for k := 0; k < 18; k++ {
		linesToWrite = append(linesToWrite, linesToWrite...)

		for _, line := range linesToWrite {
			file.WriteString(line + "\n")
		}

		b.Run(fmt.Sprintf("Len=%d", len(linesToWrite)), func(b *testing.B) {
			b.StopTimer()
			if input, err := os.Open("./test.txt"); err == nil {
				defer input.Close()
				scanner := bufio.NewScanner(input)
				for scanner.Scan() {
					commandLine := scanner.Text()
					b.StartTimer()
					parser.Parse(commandLine)
				}
			}
		})

	}
}
