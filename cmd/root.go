package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/samjtro/go-dsr"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deepseek-reasoner",
	Short: "pure go deepseek-r1 in the cli!",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		client := dsr.NewChatClient()
		start(client)
	},
}

func start(client *dsr.ChatClient) {
	q := read()
	client.AddUserMessage(q)
	res, _ := client.GetNextChatCompletion()
	client.AddMessage(res.Choices[0].Message)
	result := markdown.Render(res.Choices[0].Message.Content, 80, 6)
	fmt.Println()
	fmt.Println(result)
	fmt.Println()
	start(client)
}

func execCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	key, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return key
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
