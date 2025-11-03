package ds

import (
	"fmt"
	"log"
	"strings"
)

func toString[T any](data []T) string {
	var b strings.Builder
	b.WriteString("[")
	for i := 0; i < len(data)-1; i++ {
		str := fmt.Sprintf("%v, ", data[i])
		_, err := b.WriteString(str)
		if err != nil {
			log.Fatalf("could not write in strings builder: %v", err)
		}
	}
	str := fmt.Sprintf("%v]%s", data[len(data)-1], fmt.Sprintln())
	_, err := b.WriteString(str)
	if err != nil {
		log.Fatalf("could not write in strings builder: %v", err)
	}
	return b.String()
}
