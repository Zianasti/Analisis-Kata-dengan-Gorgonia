package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gorgonia.org/tensor"
)

func preprocessText(text string) []string {
	// Preprocessing teks: Mengubah teks menjadi huruf kecil dan melakukan tokenisasi kata-kata
	text = strings.ToLower(text)
	words := strings.Split(text, " ")
	return words
}

func loadDataset(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dataset []string
	for scanner.Scan() {
		line := scanner.Text()
		dataset = append(dataset, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dataset, nil
}

func main() {
	// Load dataset dari file data_news.txt
	dataset, err := loadDataset("data_news.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Preprocessing teks
	preprocessedData := make([][]string, len(dataset))
	for i, text := range dataset {
		preprocessedData[i] = preprocessText(text)
	}

	// Membangun kamus kata dari dataset
	vocabulary := make(map[string]int)
	for _, words := range preprocessedData {
		for _, word := range words {
			if _, exists := vocabulary[word]; !exists {
				vocabulary[word] = len(vocabulary)
			}
		}
	}

	// Mengonversi data teks ke dalam bentuk matriks one-hot encoding
	docTermMatrix := tensor.New(tensor.WithShape(len(dataset), len(vocabulary)), tensor.Of(tensor.Float64))
	for i, words := range preprocessedData {
		for _, word := range words {
			wordIndex := vocabulary[word]
			docTermMatrix.SetAt(float64(1), i, wordIndex)
		}
	}

	// Menampilkan hasil matriks document-term
	fmt.Println("Document-Term Matrix:")
	fmt.Println(docTermMatrix)

	// Perform Word Analysis on Document-Term Matrix
	fmt.Println("Word Analysis:")
	wordFrequency := make(map[string]int)
	for _, words := range preprocessedData {
		for _, word := range words {
			wordFrequency[word]++
		}
	}

	// Sortir kata berdasarkan frekuensi lebih dari 10 kali
	var popularWords []string
	for word, frequency := range wordFrequency {
		if frequency > 10 {
			popularWords = append(popularWords, word)
		}
	}

	sort.Strings(popularWords)

	// Print Word Frequencies
	for _, word := range popularWords {
		frequency := wordFrequency[word]
		fmt.Printf("%s: %d\n", word, frequency)
	}

}
