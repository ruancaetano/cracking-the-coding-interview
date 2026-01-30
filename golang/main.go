package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/chapter1"
	"github.com/ruancaetano/cracking-the-coding-interview/golang/chapter2"
	"github.com/ruancaetano/cracking-the-coding-interview/golang/chapter3"
	"github.com/ruancaetano/cracking-the-coding-interview/golang/chapter4"
)

type exercise struct {
	name string
	run  func()
}

type chapter struct {
	name      string
	exercises []exercise
}

var chapters = []chapter{
	{
		name: "Chapter 1",
		exercises: []exercise{
			{name: "Exercise 1 - Is Unique", run: chapter1.RunChapter1Exercise1},
			{name: "Exercise 2 - Check Permutation", run: chapter1.RunChapter1Exercise2},
			{name: "Exercise 3 - URLify", run: chapter1.RunChapter1Exercise3},
			{name: "Exercise 4 - Palindrome Permutation", run: chapter1.RunChapter1Exercise4},
			{name: "Exercise 5 - One Away", run: chapter1.RunChapter1Exercise5},
			{name: "Exercise 6 - String Compression", run: chapter1.RunChapter1Exercise6},
			{name: "Exercise 7 - Rotate Matrix", run: chapter1.RunChapter1Exercise7},
			{name: "Exercise 8 - Zero Matrix", run: chapter1.RunChapter1Exercise8},
			{name: "Exercise 9 - String Rotation", run: chapter1.RunChapter1Exercise9},
		},
	},
	{
		name: "Chapter 2",
		exercises: []exercise{
			{name: "Exercise 1 - Remove Dups", run: chapter2.RunChapter2Exercise1},
			{name: "Exercise 2 - Return Kth to Last", run: chapter2.RunChapter2Exercise2},
			{name: "Exercise 3 - Delete Middle Node", run: chapter2.RunChapter2Exercise3},
			{name: "Exercise 4 - Partition", run: chapter2.RunChapter2Exercise4},
			{name: "Exercise 5 - Sum Lists", run: chapter2.RunChapter2Exercise5},
			{name: "Exercise 6 - Palindrome", run: chapter2.RunChapter2Exercise6},
			{name: "Exercise 7 - Intersection", run: chapter2.RunChapter2Exercise7},
			{name: "Exercise 8 - Loop Detection", run: chapter2.RunChapter2Exercise8},
		},
	},
	{
		name: "Chapter 3",
		exercises: []exercise{
			{name: "Exercise 1 - Three in One", run: chapter3.RunChapter3Exercise1},
			{name: "Exercise 2 - Stack Min", run: chapter3.RunChapter3Exercise2},
			{name: "Exercise 3 - Stack of Plates", run: chapter3.RunChapter3Exercise3},
			{name: "Exercise 4 - Queue via Stacks", run: chapter3.RunChapter3Exercise4},
			{name: "Exercise 5 - Sort Stack", run: chapter3.RunChapter3Exercise5},
			{name: "Exercise 6 - Animal Shelter", run: chapter3.RunChapter3Exercise6},
		},
	},
	{
		name: "Chapter 4",
		exercises: []exercise{
			{name: "Exercise 1 - Route Between Nodes", run: chapter4.RunChapter4Exercise1},
			{name: "Exercise 2 - Minimal Tree", run: chapter4.RunChapter4Exercise2},
			{name: "Exercise 3 - List of Depths", run: chapter4.RunChapter4Exercise3},
			{name: "Exercise 4 - Check Balanced", run: chapter4.RunChapter4Exercise4},
			{name: "Exercise 5 - Validate BST", run: chapter4.RunChapter4Exercise5},
			{name: "Exercise 6 - Successor", run: chapter4.RunChapter4Exercise6},
			{name: "Exercise 7 - Build Order", run: chapter4.RunChapter4Exercise7},
			{name: "Exercise 8 - First Common Ancestor", run: chapter4.RunChapter4Exercise8},
			{name: "Exercise 9 - BST Sequences", run: chapter4.RunChapter4Exercise9},
			{name: "Exercise 10 - Check Subtree", run: chapter4.RunChapter4Exercise10},
			{name: "Exercise 11 - Random Node", run: chapter4.RunChapter4Exercise11},
			{name: "Exercise 12 - Paths with Sum", run: chapter4.RunChapter4Exercise12},
		},
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Cracking the Coding Interview - Go Exercises ===")
		fmt.Println("Select a chapter:")
		for i, ch := range chapters {
			fmt.Printf("%d) %s\n", i+1, ch.name)
		}
		fmt.Println("0) Exit")
		fmt.Print("Enter choice: ")

		choice, err := readInt(reader)
		if err != nil || choice < 0 || choice > len(chapters) {
			fmt.Println("Invalid choice. Please enter a number from the list.")
			fmt.Println()
			continue
		}

		if choice == 0 {
			fmt.Println("Goodbye!")
			return
		}

		runChapterMenu(reader, chapters[choice-1])
	}
}

func runChapterMenu(reader *bufio.Reader, ch chapter) {
	for {
		fmt.Printf("\n=== %s ===\n", ch.name)
		fmt.Println("Select an exercise to run:")
		for i, ex := range ch.exercises {
			fmt.Printf("%d) %s\n", i+1, ex.name)
		}
		fmt.Println("0) Back to chapters")
		fmt.Print("Enter choice: ")

		choice, err := readInt(reader)
		if err != nil || choice < 0 || choice > len(ch.exercises) {
			fmt.Println("Invalid choice. Please enter a number from the list.")
			fmt.Println()
			continue
		}

		if choice == 0 {
			fmt.Println()
			return
		}

		selected := ch.exercises[choice-1]
		fmt.Printf("\n--- Running %s - %s ---\n\n", ch.name, selected.name)
		selected.run()
		fmt.Println("\n--- Finished ---")
		fmt.Println("Press Enter to continue...")
		_, _ = reader.ReadString('\n')
	}
}

func readInt(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}
