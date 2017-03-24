package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/juju/loggo"
)

var log = loggo.GetLogger("app")
var ptr, helpPtr = 0, 0

func loadFile(args []string) string {
	if len(args) < 2 {
		log.Errorf("You must give me a filename.")
		os.Exit(1)
	}
	fileName := args[1]
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("File %s doesn't read", fileName)
	}

	str := fmt.Sprintf("%s", fileBytes)

	return str
}

func main() {
	log.Infof("Loading brainfuck JIT compiler...")
	fileData := loadFile(os.Args)
	log.Infof("Data loaded.")
	log.Debugf("print: %s", fileData)

	tape := []uint8{0}
	var index = 0
	for ; index < len(fileData); index++ {
		operation := fileData[index]
		switch operation {
		case '>':
			ptr++
			if len(tape) <= ptr {
				tape = append(tape, 0)
			}
		case '<':
			if ptr > 0 {
				ptr--
			}
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			fmt.Print(string(tape[ptr]))
		case '[':
			if tape[ptr] == 0 {
				for depth := 1; depth > 0; {
					index++
					newInstruction := fileData[index]
					if newInstruction == '[' {
						depth++
					} else if newInstruction == ']' {
						depth--
					}
				}
			}
		case ']':
			for depth := 1; depth > 0; {
				index--
				newInstruction := fileData[index]
				if newInstruction == '[' {
					depth--
				} else if newInstruction == ']' {
					depth++
				}
			}
			index--
		}
	}
}
