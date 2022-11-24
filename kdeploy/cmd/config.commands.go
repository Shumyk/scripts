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
	key, value := args[0], args[1]
	fieldsCollector := tabwriter.NewWriter(os.Stderr, 1, 2, 4, ' ', tabwriter.TabIndent)

	configValue := reflect.ValueOf(&config)
	fieldsNumber := configValue.Elem().NumField()
	for i := 0; i < fieldsNumber; i++ {
		field := configValue.Elem().Type().Field(i)
		if field.Tag.Get("conf") == "no" {
			continue
		}
		if strings.EqualFold(field.Name, key) {
			var valueObj any = value
			if field.Type.Kind() == reflect.Slice {
				valueObj = strings.Split(value, ",")
			}
			SetConfigHandling(field.Name, valueObj)
			return
		}
		_, _ = fmt.Fprintln(fieldsCollector, "\t"+strings.ToLower(field.Name)+"\t:\t"+field.Type.String())
	}
	RedStderr("Non existing property: ", key)
	BoringStderr("Possible configuration properties:")
	_ = fieldsCollector.Flush()
}

func runConfigEdit(cmd *cobra.Command, args []string) {
	fmt.Println("config edit command")
}
