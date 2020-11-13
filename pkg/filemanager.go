package pkg

import (
	"bufio"
	"log"
	"os"
)

func SaveToFile(gameData []string) {
	file, err := os.OpenFile("wishlist.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	dataWriter := bufio.NewWriter(file)

	for _, data := range gameData {
		_, err := dataWriter.WriteString(data + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}

	dataWriter.Flush()
	file.Close()
}