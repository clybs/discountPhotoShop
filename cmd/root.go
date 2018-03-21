package cmd

import (
	"fmt"
	"os"
	"strconv"

	"bufio"
	"strings"

	"github.com/spf13/cobra"

	"github.com/clybs/discountPhotoShop/artist"
	"github.com/clybs/discountPhotoShop/draw"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "discountPhotoShop",
	Short: "Drawing on a budget",
	Long:  `This application will draw something. That's a promise.'`,
	Run:   start,
}

var at artist.Artist
var dr draw.Draw

func Draw(drawStepsCommands []string, drawStepsParams, canvas [][]string) ([][]string, bool) {
	for i, params := range drawStepsParams {
		// Check if the param values are valid
		paramsValid := true

		for j, param := range params {
			// Check if this is the fill command
			isFillCommand := strings.ToUpper(drawStepsCommands[i]) == "B"
			isCharacterParam := j == 2

			if isFillCommand && isCharacterParam {
				continue
			}

			_, err := strconv.Atoi(param)
			if err != nil {
				paramsValid = false
				break
			}
		}

		// Process only valid commands
		if paramsValid {
			switch strings.ToUpper(drawStepsCommands[i]) {
			case "C":
				canvas = at.CreateBlankCanvas(toInt(params[0]), toInt(params[1]))
			case "L":
				canvas = at.CreateLine(toInt(params[0]), toInt(params[1]), toInt(params[2]), toInt(params[3]), canvas)
			case "R":
				canvas = at.CreateRectangle(toInt(params[0]), toInt(params[1]), toInt(params[2]), toInt(params[3]), canvas)
			case "B":
				canvas = at.CreateFill(params[2], toInt(params[0]), toInt(params[1]), canvas)
			}
		}
	}

	// Get last command
	lastCommand := drawStepsCommands[len(drawStepsCommands)-1]
	quit := strings.ToUpper(lastCommand) == "Q"

	return canvas, quit
}

// Execute adds all child commands to the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)
	validCommands := []string{"C", "L", "R", "B", "Q"}
	drawStepsCommands := make([]string, 0)
	drawStepsParams := make([][]string, 0)
	canvas := make([][]string, 0)
	quit := false

	// Scan user input
	for scanner.Scan() {
		input := scanner.Text()
		commandParam := getCommandParams(input)

		// Check validity of params and commands
		paramCounter := 0
		paramSet := make([]string, 0)
		for i, v := range commandParam {
			// Check if paramCounter was set
			if paramCounter > 0 {
				// Update draw steps for params
				paramSet, paramCounter, drawStepsParams = updateDrawStepsParams(v, paramSet, paramCounter, drawStepsParams)
			} else if paramCounter != -1 && isStringInSlice(strings.ToUpper(v), validCommands) {
				// Update draw steps for commands
				paramCounter, commandParam, drawStepsCommands = updateDrawStepsCommands(v, i, paramCounter, commandParam, drawStepsCommands)
			}
		}

		// Start drawing
		canvas, quit = Draw(drawStepsCommands, drawStepsParams, canvas)

		// Display drawing
		dr.Display(canvas)

		// Check if quit mode
		if quit {
			os.Exit(0)
		}

		// Reset all steps
		drawStepsCommands = make([]string, 0)
		drawStepsParams = make([][]string, 0)
	}

	// Gotcha errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}

func updateDrawStepsCommands(v string, i, paramCounter int, commandParam, drawStepsCommands []string) (int, []string, []string) {
	// Get allowed param count
	allowedParamCount := expectedParamLength(v)

	// Check if allowedParamCount is within range
	if len(commandParam) >= i+1+allowedParamCount {
		// Add to draw steps
		drawStepsCommands = append(drawStepsCommands, v)

		// Get allowed parameter length
		paramCounter = expectedParamLength(v)
	}
	return paramCounter, commandParam, drawStepsCommands
}

func updateDrawStepsParams(v string, paramSet []string, paramCounter int, drawStepsParams [][]string) ([]string, int, [][]string) {
	paramSet = append(paramSet, v)
	paramCounter -= 1

	// Check if set is complete
	if paramCounter == 0 {
		drawStepsParams = append(drawStepsParams, paramSet)

		// Reset the paramSet
		paramSet = make([]string, 0)
	}

	return paramSet, paramCounter, drawStepsParams
}
