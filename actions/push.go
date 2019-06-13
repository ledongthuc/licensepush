package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LicensePushCmd(cmd *cobra.Command, args []string) {
	sourcePath := "./"
	if len(args) > 0 {
		sourcePath = args[0]
	}
	licenseContent, err := getLicenseContent(sourcePath)
	if err != nil {
		fmt.Printf("ERROR: can't read license's content: %s \n", err)
		return
	}
	licensePush(sourcePath, licenseContent)
}

func getLicenseContent(sourcePath string) (string, error) {
	if viper.IsSet("license") {
		if licenseContent := viper.GetString("license"); licenseContent != "" {
			return licenseContent, nil
		}
	}

	localLicense := "./LICENSE"
	raw, err := ioutil.ReadFile(localLicense)
	if err != nil {
		return "", errors.Wrap(err, "Read from ./LICENSE")
	}
	return string(raw), nil
}

func licensePush(sourcePath, content string) error {
	patterns := GetCommentPatterns()
	err := filepath.Walk(sourcePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil // Next
			}

			if supportedConfig, ok := patterns[filepath.Ext(path)]; ok {
				err := WriteFile(path, supportedConfig, content)
				if err != nil {
					return err
				}
			}
			return nil
		})
	return err
}

func WriteFile(filePath string, config Config, content string) error {
	replacement := config.GetReplacement(content)
	sourceCodeB, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	sourceCode := string(sourceCodeB)
	if config.AddTop && !strings.HasPrefix(sourceCode, replacement) {
		sourceCode = fmt.Sprintf("%s\n%s", replacement, sourceCode)
	}
	if config.AddBottom && !strings.HasSuffix(sourceCode, replacement) {
		sourceCode = fmt.Sprintf("%s\n%s", sourceCode, replacement)
	}

	return ioutil.WriteFile(filePath, []byte(sourceCode), 0644)
}
