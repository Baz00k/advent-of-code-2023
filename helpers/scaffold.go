package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func Scaffold(day int) {
	dayStr := fmt.Sprintf("%02d", day)

	// Create the directory for the day
	os.Mkdir("day"+dayStr, 0755)

	// Define the necessary files and their corresponding template files
	files := map[string]string{
		"a.go":      "templates/a.go.tmpl",
		"a_test.go": "templates/a_test.go.tmpl",
		"b.go":      "templates/b.go.tmpl",
		"b_test.go": "templates/b_test.go.tmpl",
		"input.txt": "templates/input.txt.tmpl",
		"main.go":   "templates/main.go.tmpl",
		"test.txt":  "templates/test.txt.tmpl",
	}

	for file, templateFile := range files {
		// Parse the template file
		tmpl, err := template.ParseFiles(templateFile)
		if err != nil {
			log.Fatalf("Failed to parse template file: %v", err)
		}

		// Create the new file
		newFile, err := os.Create("day" + dayStr + "/" + file)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}

		// Execute the template with the day as the data
		err = tmpl.Execute(newFile, map[string]string{"Day": dayStr})
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
		}

		newFile.Close()
	}

	// ... existing code ...

	// Add the new day to the solve.go file in the root directory
	solveFile, _ := os.ReadFile("solve.go")
	solveContent := string(solveFile)

	// Add the import for the new day
	importLine := fmt.Sprintf("\t\"github.com/Baz00k/advent-of-code-2023/day%02d\"\n)", day)
	solveContent = strings.Replace(solveContent, ")", importLine, 1)

	// Add the new day to the solveFuncs map
	solveFuncLine := fmt.Sprintf("\t%d: day%02d.Solve,\n}", day, day)
	solveContent = strings.Replace(solveContent, "}", solveFuncLine, 1)

	os.WriteFile("solve.go", []byte(solveContent), 0644)
}
