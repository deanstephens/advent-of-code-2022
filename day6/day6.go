package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type buffer struct {
	bufferString  string
	startLocation int
}

func Part1() int {
	buffers := read()
	fmt.Println(buffers)

	b := strings.Split(buffers[0].bufferString, "")

	return getBufferStart(b, 4)
}

func Part2() int {
	buffers := read()
	fmt.Println(buffers)

	b := strings.Split(buffers[0].bufferString, "")

	return getBufferStart(b, 14)
}

func getBufferStart(bufferString []string, codeLength int) int {
	for i := codeLength; i < len(bufferString); i++ {
		if checkIsSliceUnique(bufferString[i-codeLength : i]) {
			return i
		}
	}

	return -1
}

func checkIsSliceUnique(s []string) bool {
	var m = make(map[string]bool)

	for _, char := range s {
		m[char] = true
	}

	return len(m) == len(s)
}

func read() []buffer {
	var buffers []buffer

	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	for scanner.Scan() {
		text := scanner.Text()
		buffers = append(buffers, buffer{
			bufferString:  text,
			startLocation: 0,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return buffers
}
