package main

import (
	"fmt"
	"log"
	"os"
	// "path/filepath"

	"visrosa/sgr" // Replace with your actual module path
)

func main() {
	// Method 1: Load from JSON file
	jsonFile := "ansi_codes.json"
	if err := loadRegistryFromFile(jsonFile); err != nil {
		log.Printf("Could not load from file: %v", err)

		// Method 2: Generate JSON using your extractor
		if err := generateAndLoadRegistry(); err != nil {
			log.Fatalf("Could not generate registry: %v", err)
		}
	}

	// Now you can use the codes normally
	fmt.Println("Testing SGR codes loaded from JSON:")

	// Basic usage
	fmt.Print(sgr.Reset.Apply())
	fmt.Print(sgr.Bold.Apply())
	fmt.Print("Bold Text")
	fmt.Print(sgr.Reset.Apply())
	fmt.Println()

	// Color usage
	fmt.Print(sgr.FgRed.Apply())
	fmt.Print("Red Text")
	fmt.Print(sgr.Reset.Apply())
	fmt.Println()

	// 256 color usage
	fmt.Print(sgr.Fg.Color(196).Apply())
	fmt.Print("Bright Red (256 color)")
	fmt.Print(sgr.Reset.Apply())
	fmt.Println()

	// RGB color usage
	fmt.Print(sgr.Fg.RGB(255, 128, 0).Apply())
	fmt.Print("Orange (RGB)")
	fmt.Print(sgr.Reset.Apply())
	fmt.Println()

	// Cursor operations
	fmt.Print(sgr.Cursor.Save.Apply())
	fmt.Print("Saved position")
	fmt.Print(sgr.Cursor.Forward.Apply()) // This would need parameters
	fmt.Print(sgr.Cursor.Restore.Apply())
	fmt.Println()

	// Registry introspection
	listAvailableCodes()
}

func loadRegistryFromFile(filename string) error {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	// Read JSON file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Initialize registry
	if err := sgr.InitFromJSON(jsonData); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	fmt.Printf("Successfully loaded registry from %s\n", filename)
	return nil
}

func generateAndLoadRegistry() error {
	// You could run your extractor here or embed the JSON
	// For now, let's assume you have the JSON as a string constant

	// This would be the output from your sgr_extractor.go
	jsonData := `{
		"version": "1.0.0",
		"generated": "2024-01-01T00:00:00Z",
		"code_types": {
			"sgr": 0,
			"cursor": 1,
			"screen": 2,
			"control": 3,
			"osc": 4
		},
		"param_types": {
			"none": 0,
			"numeric": 1,
			"string": 2,
			"rgb": 3,
			"position": 4,
			"meta": 5
		},
		"codes": {
			"reset": {
				"id": "reset",
				"name": "Reset",
				"code": "0",
				"type": "sgr",
				"param_type": "none",
				"description": "Reset all attributes to default",
				"examples": [
					{
						"description": "Reset all formatting",
						"output": "\\x1b[0m"
					}
				]
			}
		}
	}`

	// Initialize registry from embedded JSON
	if err := sgr.InitFromJSON([]byte(jsonData)); err != nil {
		return fmt.Errorf("failed to initialize registry: %v", err)
	}

	fmt.Println("Successfully generated and loaded registry")
	return nil
}

func listAvailableCodes() {
	fmt.Println("\nAvailable codes in registry:")

	// Get a few sample codes to demonstrate
	codes := []string{"reset", "bold", "fg_red", "cursor_up", "erase_display"}

	for _, id := range codes {
		if code, exists := sgr.GetCode(id); exists {
			fmt.Printf("- %s: %s (%s)\n", id, code.Name, code.Description)
		}
	}
}

// Alternative: Embed JSON directly in your sgr package
func embedJSONInPackage() {
	// You could add this to your sgr package:

	const embeddedJSON = `{...}` // Your full JSON here

	// In init() function:
	// if err := InitFromJSON([]byte(embeddedJSON)); err != nil {
	//     panic(fmt.Sprintf("Failed to load embedded registry: %v", err))
	// }
}
