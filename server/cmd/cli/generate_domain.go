package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const domainDir = "internal/domain"

func main() {
	fmt.Println("Enter the name of your domain please: ");
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()

	domainName := strings.TrimSpace(scanner.Text())

	if domainName == "" {
		fmt.Println("unable to create domain, invalid input")
		return
	};

	domainPath := filepath.Join(domainDir, domainName);

	if _, err := os.Stat(domainPath); !os.IsNotExist(err) {
		fmt.Println("Error: Domain already exists. Ignoring creation.")
		return
	}

	if err := os.MkdirAll(domainPath, os.ModePerm); err != nil {
		fmt.Println("Error creating domain: ", err)
		return
	}


	files := map[string]string{
		"repository.go": repoTemplate(domainName),
		"service.go" : serviceTemplate(domainName),
	};

	for filename, content := range files {
		filepath := filepath.Join(domainPath, filename);
		if err := os.WriteFile(filepath, []byte(content), 0644); err != nil {
			fmt.Println("Error creating file:", err)
		}
	}

	fmt.Println("✅ Domain", domainName, "created successfully!")
}


func repoTemplate(domain string) string {
	return fmt.Sprintf(`package %s

import (
	"gorm.io/gorm"
)

type %sRepository interface {
}

type %sRepository struct {
	db *gorm.DB
}

func New%sRepository(db *gorm.DB) %sRepository {
	return &%sRepository{db: db}
}
`, domain, parseText(domain, true), parseText(domain, false), cases.Title(language.English).String(domain), parseText(domain, true), parseText(domain, false))
}


func serviceTemplate(domain string) string {
	return fmt.Sprintf(`package %s

type %sService interface {
}

type %sService struct {
	repository %sRepository
}

func New%sService(repo *%sRepository) %sService {
	return &%sService{repository: repo}
}


	`, domain, parseText(domain, true), parseText(domain, false), parseText(domain, true), parseText(domain, true), parseText(domain, true), parseText(domain, true), parseText(domain, false) )
}



func parseText(text string, capitalize bool) string {
	if len(text) == 0 {
		return text
	}

	first := string(text[0])
	rest := text[1:]

	if capitalize {
		return strings.ToUpper(first) + rest
	}
	return strings.ToLower(first) + rest
}
