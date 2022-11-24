package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"reflect"
	. "shumyk/kdeploy/cmd/util"
	"strings"
)

func runConfig(cmd *cobra.Command, args []string) {
	fmt.Println("config command")
}

func runConfigView(_ *cobra.Command, _ []string) {
	viewBytes, err := yaml.Marshal(config.View())
	ErrorCheck(err, "Couldn't marshal config file")
	fmt.Println(string(viewBytes))
}

func runConfigSet(cmd *cobra.Command, args []string) {
	key := cases.Title(language.English).String(args[0])
	value := args[1]

	configValue := reflect.ValueOf(config)
	field := configValue.FieldByName(key)
	if field.IsValid() {
		err := SetConfig(key, value)
		ErrorCheck(err, "Could not set config")
	} else {
		RedStderr("Non existing property: " + key)
		BoringStderr("Possible configuration properties:")
		for i := 0; i < configValue.NumField(); i++ {
			tag := configValue.Type().Field(i).Tag
			if tag.Get("conf") != "no" {
				fieldName := configValue.Type().Field(i).Name
				fmt.Println("	-", strings.ToLower(fieldName))
			}
		}
	}
}

func runConfigEdit(cmd *cobra.Command, args []string) {
	fmt.Println("config edit command")
}
