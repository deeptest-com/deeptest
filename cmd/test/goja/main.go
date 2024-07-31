package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 打开TS文件
	f, err := os.Open("math.d.ts")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 创建解析器
	p := astits.NewParser(f)

	for {
		p, err := p.ParseOne()
		if err != nil {
			if err == astits.ErrEOF {
				break
			}
			log.Fatal(err)
		}

		// 处理PAT和PMT
		if pat, ok := p.(*astits.PAT); ok {
			fmt.Printf("Found PAT: %+v\n", pat)
			for _, pmtr := range pat.Programs {
				if pmt, err := astitools.PmtFromRunningFile(f, pmtr.PID); err == nil {
					fmt.Printf("Found PMT for program %d: %+v\n", pmtr.ProgramNumber, pmt)
				} else {
					fmt.Printf("Error while retrieving PMT for program %d: %v\n", pmtr.ProgramNumber, err)
				}
			}
		}
	}
}
