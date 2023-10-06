package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func promptChoice(message string, choices []string) string {
	fmt.Println(message)
	fmt.Printf("Choices: %s\n", strings.Join(choices, ", "))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	return choice
}

func main() {
	fmt.Println("Welcome to the Interactive Story!")
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("\nHello, %s! You find yourself at a crossroad in the forest.\n", name)
	choice := promptChoice("You have two paths to choose from: 'left' or 'right'.", []string{"left", "right"})

	switch choice {
	case "left":
		fmt.Println("\nYou chose the left path and found an old cabin.")
		choice2 := promptChoice("Inside the cabin, you discover a hidden treasure map. Will you follow the map ('follow') or continue your journey ('continue')?", []string{"follow", "continue"})
		switch choice2 {
		case "follow":
			fmt.Println("\nYou decide to follow the map's clues...")
			// Add more story content and choices here...
		case "continue":
			fmt.Println("\nYou leave the cabin and continue your journey.")
			// Add more story content and choices here...
		default:
			fmt.Println("\nInvalid choice. You're paralyzed by indecision.")
		}
	case "right":
		fmt.Println("\nYou chose the right path and encountered a mystical waterfall.")
		choice3 := promptChoice("You notice a hidden cave behind the waterfall. Do you want to explore the cave ('explore') or avoid it ('avoid')?", []string{"explore", "avoid"})
		switch choice3 {
		case "explore":
			fmt.Println("\nYou enter the cave and find a chest filled with ancient artifacts.")
			// Add more story content and choices here...
		case "avoid":
			fmt.Println("\nYou decide not to risk it and continue your journey.")
			// Add more story content and choices here...
		default:
			fmt.Println("\nInvalid choice. Fear holds you back.")
		}
	default:
		fmt.Println("\nInvalid choice. You're paralyzed by indecision.")
	}

	// End of the story
	fmt.Printf("\nThanks for playing, %s! Goodbye.\n", name)
}
