// Command gen generates Go model definitions and JSON schema properties from models.yml.
//
// Usage:
//
//	go run ./cmd/gen/ -models > internal/model/generated.go
//	go run ./cmd/gen/ -schemas > internal/model/schema_generated.go
package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	genModels := flag.Bool("models", false, "Generate model definitions")
	genSchemas := flag.Bool("schemas", false, "Generate JSON schema properties")
	modelsPath := flag.String("path", "", "Path to models.yml (default: auto-detect)")
	flag.Parse()

	if !*genModels && !*genSchemas {
		log.Fatal("specify -models or -schemas")
	}

	path := *modelsPath
	if path == "" {
		// Try common locations.
		candidates := []string{
			"../../lib/openslides-go/meta/models.yml",
			"../lib/openslides-go/meta/models.yml",
			"lib/openslides-go/meta/models.yml",
		}
		for _, c := range candidates {
			if _, err := os.Stat(c); err == nil {
				path = c
				break
			}
		}
		if path == "" {
			log.Fatal("could not find models.yml, use -path flag")
		}
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("open models.yml: %v", err)
	}
	defer f.Close()

	parsed, err := parseModels(f)
	if err != nil {
		log.Fatalf("parse models.yml: %v", err)
	}

	if *genModels {
		if err := writeModelDefs(os.Stdout, parsed); err != nil {
			log.Fatalf("write model defs: %v", err)
		}
	}

	if *genSchemas {
		if err := writeSchemaDefs(os.Stdout, parsed); err != nil {
			log.Fatalf("write schema defs: %v", err)
		}
	}
}
