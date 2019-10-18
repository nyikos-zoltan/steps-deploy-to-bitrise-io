// Package converters contains the interface that is required to be a package a test result converter.
// It must be possible to set files from outside(for example if someone wants to use
// a pre-filtered files list), need to return Junit4 xml test result, and needs to have a
// Detect method to see if the converter can run with the files included in the test result dictionary.
// (So a converter can run only if the dir has a TestSummaries.plist file for example)
package converters

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-steplib/steps-deploy-to-bitrise-io/test/checkstyle"
	"github.com/bitrise-steplib/steps-deploy-to-bitrise-io/test/converters/junitxml"
	"github.com/bitrise-steplib/steps-deploy-to-bitrise-io/test/converters/xcresult"
	"github.com/bitrise-steplib/steps-deploy-to-bitrise-io/test/converters/xcresult3"
)

// Intf is the required interface a converter need to match
type Intf interface {
	XML() (interface{}, error)
	Detect([]string) bool
}

type CheckstyleConverter struct {
	files []string
}

func (c *CheckstyleConverter) XML() (interface{}, error) {
	var xmlContent checkstyle.XML

	for _, file := range c.files {
		data, err := fileutil.ReadBytesFromFile(file)
		if err != nil {
			return nil, err
		}

		var fileData checkstyle.XML
		checkstyleError := xml.Unmarshal(data, &fileData)
		if checkstyleError != nil {
			return nil, checkstyleError
		}

		xmlContent.Files = append(xmlContent.Files, fileData.Files...)
	}

	return xmlContent, nil
}

func (c *CheckstyleConverter) Detect(files []string) bool {
	fmt.Println("Detect called!")
	for _, file := range files {
		if strings.HasSuffix(file, "checkstyle.xml") {
			c.files = append(c.files, file)
		} else {
			fmt.Printf("%s was skipped!\n", file)
		}
	}
	return len(c.files) > 0
}

var converters = []Intf{
	&CheckstyleConverter{},
	&junitxml.Converter{},
	&xcresult.Converter{},
	&xcresult3.Converter{},
}

// List lists all supported converters
func List() []Intf {
	return converters
}
