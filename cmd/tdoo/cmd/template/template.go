package template

import (
	_ "embed"
	"fmt"
	_ "text/template"

	"github.com/spf13/cobra"
)

//go:embed template.config.yaml
var tmpl string

var TemplateCmd = &cobra.Command{
	Use:     "get-template",
	Short:   "Get the template yaml file of config",
	Example: "tdoo get-template >> config.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(tmpl)
		return nil
	},
}
