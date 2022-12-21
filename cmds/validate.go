package cmds

import (
	"encoding/xml"
	"github.com/antchfx/xmlquery"
	"github.com/nyudlts/go-aspace"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(ValidateCmd)
}

var ValidateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {},
}

func validateFile(path string) error {
	filename := filepath.Base(path)
	log.Printf("[INFO] validating %s", filename)
	//check that the file has an .xml extension
	if !xmlPtn.MatchString(filename) {
		log.Printf("[WARNING] file %s does not end in .xml skipping", filename)
		return nil
	}

	//get the bytes of a file
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[ERROR] could not read %s", filename)
		return nil
	}

	//check that ead is well-formed
	if err = xml.Unmarshal(fileBytes, new(interface{})); err != nil {
		log.Printf("[ERROR] %s is not well-formed", filename)
		return nil
	}

	//validate against schema
	if err = aspace.ValidateEAD(fileBytes); err != nil {
		log.Printf("[ERROR] %s is not valid to EAD 2002 schema", filename)
		return nil
	}

	//validate urls
	if err = urlValidation(fileBytes); err != nil {
		log.Printf("[ERROR] %s contains invalid links", filename)
	}

	log.Printf("[INFO] %s is valid", filename)
	return nil
}

func urlValidation(fa []byte) error {
	doc, err := xmlquery.Parse(strings.NewReader(string(fa)))
	if err != nil {
		return err
	}

	daos := xmlquery.Find(doc, "//dao")
	if len(daos) > 0 {
		for _, dao := range daos {
			href := dao.SelectAttr("xlink:href")
			_, err := url.ParseRequestURI(href)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
