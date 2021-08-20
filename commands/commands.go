package commands

import (
	// "bytes"

	"fmt"
	"io"

	"sshserver/colors"
	"sshserver/utils"
  "sshserver/api"

	"strings"

  "time"
  "log"
)

type cmd struct {
	name        string
	run         func(stream io.Writer, name string, args []string)
	argsInfo    string
}

var (
	cmds = []cmd{
		{"help", HelpCmd, "" },
    {"about", AboutMeCmd, ""},
    {"clear", ClearCmd, ""}}
  
  cmdText = `
    * help - See all the current commands you can run.
    * about - Who is HackermonDev? Learn more with this command.
    * clear - Clears the screen.
`
)

func HelpCmd(stream io.Writer, name string, args []string) {
	utils.AddText(stream, fmt.Sprintf("%shttps://daniel.is-a.dev%s\n\nCommands: %s", colors.Yellow, colors.Reset, cmdText))
}

func AboutMeCmd(stream io.Writer, name string, args []string){
  utils.Type(stream, fmt.Sprintf("Who is HackermonDev??? 🤔🤔🤔\n\n"))

  aboutMe, err := api.GetAboutMeDescription()

  if err != nil{
    log.Println(err)

    utils.Type(stream, fmt.Sprintf("An unexpected error occured. Please contact Hackermon if this error persits."))

    return
  }

  time.Sleep(1 * time.Second)
  utils.Type(stream, fmt.Sprintf(aboutMe))
  utils.AddText(stream, "\n\n")
}

func ClearCmd(stream io.Writer, name string, args []string){
  utils.ClearTerm(stream)
}

func RunCommand(stream io.Writer, text string) {
	cmdName := strings.Split(text, " ")[0]

  if cmdName == ""{
    return
  }

	cmdArgs := strings.Split(text, " ")

	if len(cmdArgs) > 1 {
		cmdArgs = strings.Split(text, "")[1:]
	}

	foundCommandToRun := false

	for _, c := range cmds {
		if c.name == cmdName {
			c.run(stream, cmdName, cmdArgs)

			foundCommandToRun = true
			return
		}
	}

	if foundCommandToRun == false { 
		utils.AddText(stream, fmt.Sprintf("%sThe command \"%s\" was not found on the server.\n", colors.Red, cmdName))
	}
}
