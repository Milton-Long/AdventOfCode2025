package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
//	fmt.Println("Smile Bomb!")
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer input.Close()

	reader := bufio.NewScanner(input)
	for reader.Scan(){
		line := reader.Text()
		fmt.Println(line)
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

}
