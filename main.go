package main

import (
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"

	"sshserver/colors"
	"sshserver/commands"
	"sshserver/utils"

	"golang.org/x/term"

	_ "image/png"
	"time"

	"github.com/qeesung/image2ascii/convert"
)

func CloseServerHandler(s ssh.Session) {
	s.Exit(0)

	log.Println(fmt.Sprintf("connection from %s closed.", s.RemoteAddr()))
}

func StartServer(port string) {
	log.Println("started ssh server.")

	ssh.Handle(func(s ssh.Session) {
		log.Println(fmt.Sprintf("new connection from %s", s.RemoteAddr()))

		utils.ClearTerm(s)

		// Attempt my logo (assets/hackermon.png) in shell using ascii
		pty, _, _ := s.Pty()

		convertOptions := convert.DefaultOptions
		convertOptions.FixedWidth = pty.Window.Width
		convertOptions.FixedHeight = pty.Window.Height

		// Create the image converter
		converter := convert.NewImageConverter()

		logoAscii := converter.ImageFile2ASCIIString("assets/hackermon.png", &convertOptions)

		utils.AddText(s, "Loading...\n")
		utils.AddText(s, logoAscii)
		utils.AddText(s, "Loading...\n")

		time.Sleep(1 * time.Second)
		utils.ClearTerm(s)

		utils.Type(s, fmt.Sprintf("%sWelcome to %sdaniel.is-a.dev%s! Type %shelp%s to get started.%s\n\n", colors.Green, colors.Yellow, colors.Green, colors.Yellow, colors.Green, colors.Reset))

		term := term.NewTerminal(s, fmt.Sprintf("%s[%s@daniel.is-a.dev]%s$ ", colors.Green, s.User(), colors.Reset))

		for {
			command, err := term.ReadLine()

			if err != nil {
				CloseServerHandler(s)
				break
			}

			commands.RunCommand(term, command, s)
			log.Println(fmt.Sprintf("%s ran command \"%s\"", s.RemoteAddr(), command))

		}
	})

	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}

func main() {
	port := "22"

	log.Println("starting ssh server")
	StartServer(port)
}
