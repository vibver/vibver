package main

import (
	"bytes"
	"context"
	"log"
	"os"
	"path"

	"github.com/yuin/goldmark"
)

func main() {
	// Output path.
	rootPath := "public"
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	// Create an index page.
	name := path.Join(rootPath, "index.html")
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	fileContent, err := os.ReadFile("vibver.md")
	if err != nil {
		log.Fatalf("failed to read markdown file: %v", err)
	}

	// Convert the markdown to HTML, and pass it to the template.
	var buf bytes.Buffer
	if err := goldmark.Convert(fileContent, &buf); err != nil {
		log.Fatalf("failed to convert markdown to HTML: %v", err)
	}

	// Create an unsafe component containing raw HTML.
	content := Unsafe(buf.String())

	// Write it out.
	err = indexPage("Vibe Versioning Specification","vv🥳.👍.🔥", content).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}
}
