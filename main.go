package main

import (
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
  "sshserver/colors"
  "sshserver/utils"
  "sshserver/commands"

  "golang.org/x/term"
)


func CloseServerHandler(s ssh.Session){
  utils.Type(s, "\nGoodbye.\n")

  log.Println(fmt.Sprintf("connection from %s closed.", s.RemoteAddr()))
}

func StartServer(port string) {
  log.Println("started ssh server.")

	ssh.Handle(func(s ssh.Session) {
    log.Println(fmt.Sprintf("new connection from %s", s.RemoteAddr()))

    utils.ClearTerm(s)

		utils.Type(s, fmt.Sprintf("%sWelcome to %sdaniel.is-a.dev%s! Type %shelp%s to get started.%s\n\n", colors.Green, colors.Yellow, colors.Green, colors.Yellow, colors.Green, colors.Reset))
    
    term := term.NewTerminal(s, fmt.Sprintf("%s[%s@daniel.is-a.dev]%s$ ", colors.Green, s.User(), colors.Reset))
    
    for {
      command, err := term.ReadLine()

      if err != nil{
        CloseServerHandler(s)
        break
      }

      commands.RunCommand(s, command)
      log.Println(fmt.Sprintf("%s ran command %s %s", s.RemoteAddr(), command, colors.Reset))

    }
	})

	log.Fatal(ssh.ListenAndServe(":" + port, nil))
}

func main() {
	port := "22"

  StartServer(port)
}
