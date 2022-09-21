package create

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/brandtkeller/oscal-builder/internal/types"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

var fileName string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Run: func(cmd *cobra.Command, args []string) {
		// Conduct further error checking here (IE flags/arguments)

		err := startFileCreation(fileName)
		if err != nil {
			fmt.Println("Error string")
		}
	},
}

func Command() *cobra.Command {
	createCmd.Flags().StringVarP(&fileName, "output", "o", "", "output file to create")

	return createCmd
}

func startFileCreation(filename string) (err error) {

	var myComp types.OscalComponentDefinition
	interactiveTraverseComponentDefinition(&myComp)

	return
}

func interactiveTraverseComponentDefinition(doc *types.OscalComponentDefinition) (err error) {
	// reflect.TypeOf(e).Field(i) != "struct"

	fields := reflect.VisibleFields(reflect.TypeOf(struct{ types.OscalComponentDefinition }{}))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
		fmt.Printf("kind: %s\n", reflect.TypeOf(field).Kind())

		if reflect.TypeOf(field).Kind() == reflect.Struct && field.Name == "Metadata" {
			var newMeta types.Metadata
			fmt.Printf("%s\n", newMeta)
			newMeta.Title, _ = editFromStdIn("Title", newMeta.Title)
			clearTerminal()
			fmt.Printf("%s\n", newMeta)
			// metaFields := reflect.VisibleFields(reflect.TypeOf(struct{ types.Metadata }{}))
			// for _, metaf := range metaFields {

			// 	fmt.Printf("Key: %s\tType: %s\n", metaf.Name, metaf.Type)
			// 	fmt.Printf("kind: %s\n", reflect.TypeOf(metaf).Kind())
			// }
		}
	}
	return
}

// func CallClear() {
//     value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
//     if ok { //if we defined a clear func for that platform:
//         value()  //we execute it
//     } else { //unsupported platform
//         panic("Your platform is unsupported! I can't clear terminal screen :(")
//     }
// }

func editFromStdIn(subject string, existingData string) (data string, err error) {
	input, _ := readline.New(subject + ": ")
	defer input.Close()

	datax := existingData
	data2 := []byte(datax)
	input.WriteStdin(data2)

	data, err = input.Readline()

	return
}

func clearTerminal() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
