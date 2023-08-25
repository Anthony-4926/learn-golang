package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file1, _ := os.OpenFile("G:/linux-兄弟连/笔记/4. Linux常用命令.md", os.O_RDONLY, os.ModePerm)
	//file2, _ := os.OpenFile("G:/linux-兄弟连/笔记/4. Linux常用命令2.md", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") && strings.HasSuffix(line, "命令") {
			i := 0
			for line[i] == '#' {
				i++
			}
			//pre := line[:i]
			//<span id="1.1">1.1 命令格式</span>
			id := line[i:]
			id = strings.TrimPrefix(id, " ")
			j := 0
			for id[j] == '#' || id[j] == ' ' || (id[j] >= '0' && id[j] <= '9') || id[j] == '.' {
				j++
			}
			title := id[j:]
			fmt.Printf("[%s](#%s)\n", title, id)
		}
		//file2.WriteString(line + "\n")
	}
	file1.Close()
	//file2.Close()
}
