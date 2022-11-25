package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"reflect"
	. "shumyk/kdeploy/cmd/util"
	"strings"
	"text/tabwriter"
)

func runConfigView(_ *cobra.Command, _ []string) {
	viewBytes, err := yaml.Marshal(config.View())
	ErrorCheck(err, "Couldn't marshal config file")
	fmt.Println(string(viewBytes))
}

func RunConfigSet(_ *cobra.Command, args []string) {
	key, value := args[0], args[1]
	fieldsCollector := tabwriter.NewWriter(os.Stderr, 1, 2, 4, ' ', tabwriter.TabIndent)

	configValue := reflect.ValueOf(&config)
	fieldsNumber := configValue.Elem().NumField()
	for i := 0; i < fieldsNumber; i++ {
		field := configValue.Elem().Type().Field(i)
		// FIXME: this won't work if we have more than 1 value for tag
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

func RunConfigEdit(_ *cobra.Command, _ []string) {
	vim := exec.Command("vim", viper.ConfigFileUsed())
	vim.Stdin, vim.Stdout = os.Stdin, os.Stdout
	err := vim.Run()
	ErrorCheck(err, "Error editing configuration")
}
