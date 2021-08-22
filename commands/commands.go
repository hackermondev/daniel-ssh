package commands

import (
	"bytes"
	"strconv"

	"fmt"
	"io"

	"github.com/gliderlabs/ssh"

	"sshserver/api"
	"sshserver/colors"
	"sshserver/utils"

	"strings"

	"github.com/tj/go-spin"

	"log"
	"time"
)

type cmd struct {
	name     string
	run      func(stream io.Writer, name string, args []string, session ssh.Session)
	argsInfo string
}

var (
	cmds = []cmd{
		{"help", HelpCmd, ""},
		{"about", AboutMeCmd, ""},
		{"clear", ClearCmd, ""},
		{"blogs", BlogsCmd, ""},
		{"echo", EchoCmd, "<text>"},
		{"exit", ExitCmd, ""}}

	cmdText = `
    * help - See all the current commands you can run.
    * about - Who is HackermonDev? Learn more with this command.
    * clear - Clears the screen.
    * blogs - Read my blogs from SSH.
    * echo <text> - The server will echo the text you specified.
    * exit - Exit the SSH session.
`
)

func HelpCmd(stream io.Writer, name string, args []string, session ssh.Session) {
	utils.AddText(stream, fmt.Sprintf("%shttps://daniel.is-a.dev%s\n\nCommands: %s", colors.Yellow, colors.Reset, cmdText))
}

func AboutMeCmd(stream io.Writer, name string, args []string, session ssh.Session) {
	utils.Type(stream, fmt.Sprintf("Who is HackermonDev??? 🤔🤔🤔\n\n"))

	aboutMe, err := api.GetAboutMeDescription()

	if err != nil {
		log.Println(err)

		utils.Type(stream, fmt.Sprintf("An unexpected error occured. Please contact Hackermon if this error persits."))

		return
	}

	time.Sleep(1 * time.Second)
	utils.Type(stream, fmt.Sprintf(aboutMe))
	utils.AddText(stream, "\n\n")
}

func ClearCmd(stream io.Writer, name string, args []string, session ssh.Session) {
	utils.ClearTerm(stream)
}

func ExitCmd(stream io.Writer, name string, args []string, session ssh.Session) {
	utils.AddText(stream, "Bye bye 👋\n")
	session.Close()
}

func BlogsCmd(stream io.Writer, name string, args []string, session ssh.Session) {
	isLoading := true

	utils.AddText(stream, "\n")
	go func() {
		// Cool terminal loading spinner

		s := spin.New()

		for {
			if isLoading == false {
				break
			}

			text := fmt.Sprintf("\r\033[36mLoading blogs\033[m %s ", s.Next())

			utils.AddText(stream, text)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	blogs, err := api.GetBlogs()

	isLoading = false

	// Clear terminal line
	utils.AddText(stream, "\033[2K")

	if err != nil {
		log.Println(err)

		utils.Type(stream, fmt.Sprintf("%s An unexpected error occured. Please contact Hackermon if this error persits. %s", colors.Red, colors.Reset))

		return
	}

	commandType := ""

	if len(args) > 1 {
		commandType = args[1]
	}

	if commandType == "view" {
		if len(args) < 2 {
			utils.Type(stream, fmt.Sprintf("%s You need a specify a blog ID to lookup. %s", colors.Red, colors.Reset))
			return
		}

		blogID, _ := strconv.Atoi(args[2])

		var blog api.Blog

		for i := 0; i < len(blogs); i++ {
			b := blogs[i]

			if b.Id == blogID {
				blog = b
				break
			}
		}

		// detect empty variable (which means no blog was found)
		if (blog == api.Blog{}) {
			utils.Type(stream, fmt.Sprintf("%s The blog ID you specified was not found. %s", colors.Red, colors.Reset))
		}

		text := utils.RenderMarkdownTerminal(blog.Data)

		if err != nil {
			log.Println(err)

			utils.Type(stream, fmt.Sprintf("%s An unexpected error occured. Please contact Hackermon if this error persits. %s", colors.Red, colors.Reset))

			return
		}

		utils.AddText(stream, fmt.Sprintf(`%s %s %s 

--------------------------
%s %s
--------------------------
%s
    
%s On terminal, the markdown and pictures may not be rendered correctly. To see the full blog post go to https://daniel.is-a.dev/blog/%s %s
`, colors.Green, blog.Title, colors.Reset, colors.Gray, text, colors.Reset, colors.Red, blog.Slug, colors.Reset))
		return
	}

	var text bytes.Buffer

	text.WriteString("\n--------------------------\n")
	for i := 0; i < len(blogs); i++ {
		blog := blogs[i]

		text.WriteString(fmt.Sprintf("%s %s %s \n%s (%s) %s \n%s\n%s", colors.Gray, blog.PublishedAt, colors.Green, colors.Green, strconv.Itoa(blog.Id), blog.Title, blog.Teaser, colors.Reset))

		text.WriteString("--------------------------\n")
	}

	text.WriteString(fmt.Sprintf("\n\n %s Use %sblogs view <id>%s command to view the full blog%s\n", colors.Cyan, colors.Yellow, colors.Cyan, colors.Reset))

	utils.AddText(stream, text.String())
}

func EchoCmd(stream io.Writer, name string, args []string, session ssh.Session) {
	if len(args) < 1 {
		utils.AddText(stream, "\n")
		return
	}

	text := strings.Join(args[1:], " ")

	utils.AddText(stream, fmt.Sprintf("%s\n", text))
}

func RunCommand(stream io.Writer, text string, session ssh.Session) {
	cmdName := strings.Split(text, " ")[0]

	if cmdName == "" {
		return
	}

	cmdArgs := strings.Split(text, " ")

	foundCommandToRun := false

	for _, c := range cmds {
		if c.name == cmdName {
			c.run(stream, cmdName, cmdArgs, session)

			foundCommandToRun = true
			break
		}
	}

	if foundCommandToRun == false {
		utils.AddText(stream, fmt.Sprintf("%sThe command \"%s\" was not found on the server.\n", colors.Red, cmdName))
	}
}
