package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan nilai A dan B (Gunakan spasi untuk memisahkan kedua nilai tersebut): ")
    input, _ := reader.ReadString('\n')
	
	inputArray := strings.Split(input, " ")
	aText := strings.TrimSpace(inputArray[0])
	bText := strings.TrimSpace(inputArray[1])
	a,_ := strconv.Atoi(aText) 
	b,_ := strconv.Atoi(bText)
	result := a + b

	fmt.Printf("Hasil penjumlahan A dan B adalah %d \n", result)

}
