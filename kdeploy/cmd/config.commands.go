package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	. "shumyk/kdeploy/cmd/util"
	"strings"
	"text/tabwriter"
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
	key := args[0]
	value := args[1]

	configValue := reflect.ValueOf(&config)

	fmt.Println(config.StatefulSets)
	fmt.Println(len(config.StatefulSets))

	fieldsNumber := configValue.Elem().NumField()
	for i := 0; i < fieldsNumber; i++ {
		field := configValue.Elem().Type().Field(i)
		if field.Tag.Get("conf") != "no" {
			if strings.EqualFold(field.Name, key) {
				var valueObj any = value
				if field.Type.Kind() == reflect.Slice {
					valueObj = strings.Split(value, ",")
				}
				err := SetConfig(field.Name, valueObj)
				ErrorCheck(err, "Could not set config")
				return
			}
		}
	}

	RedStderr("Non existing property: " + key)
	BoringStderr("Possible configuration properties:")

	table := tabwriter.NewWriter(os.Stderr, 1, 2, 4, ' ', 0)
	for i := 0; i < fieldsNumber; i++ {
		field := configValue.Elem().Type().Field(i)
		if field.Tag.Get("conf") != "no" {
			_, _ = fmt.Fprintln(table, "\t"+strings.ToLower(field.Name)+"\t:\t"+field.Type.String())
		}
	}
	_ = table.Flush()
}

func runConfigEdit(cmd *cobra.Command, args []string) {
	fmt.Println("config edit command")
}
