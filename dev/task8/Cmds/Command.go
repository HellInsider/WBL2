package Cmds

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
	"strconv"
)

type ICommand interface {
	Exec(args []string) error
}

type Command struct {
	name         string
	argsRequired int
}

func (c *Command) neArgsErr(n int) {
	fmt.Println("Error in ", c.name, "\nNot enough arguments. Got ", n, ",expected", c.argsRequired)
}

type CD struct {
	cmd Command
}
type PWD struct {
	cmd Command
}
type ECHO struct {
	cmd Command
}
type KILL struct {
	cmd Command
}
type PS struct {
	cmd Command
}
type QUIT struct {
	cmd Command
}

func NewCD() *CD {
	return &CD{Command{"-cd", 1}}
}
func NewPWD() *PWD {
	return &PWD{Command{"-pwd", 0}}
}
func NewECHO() *ECHO {
	return &ECHO{Command{"-echo", -1}}
}
func NewKILL() *KILL {
	return &KILL{Command{"-kill", 1}}
}
func NewPS() *PS {
	return &PS{Command{"-ps", 1}}
}
func NewQUIT() *QUIT {
	return &QUIT{Command{"/quit", 1}}
}

func (c *CD) Exec(args []string) error {
	if len(args) < c.cmd.argsRequired {
		c.cmd.neArgsErr(len(args))
	} else {
		err := os.Chdir(args[0])
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *PWD) Exec(args []string) error {

	path, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = fmt.Println(path)

	return err
}

func (c *ECHO) Exec(args []string) error {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
	return nil
}

func (c *KILL) Exec(args []string) error {
	var err error
	var pid int

	if len(args) < c.cmd.argsRequired {
		c.cmd.neArgsErr(len(args))
	} else {
		pid, err = strconv.Atoi(args[0])
		if err == nil {
			if proc, ok := os.FindProcess(pid); ok != nil {
				fmt.Println("Err:", pid, "not found")
				err = ok
			} else if ok = proc.Kill(); ok != nil {
				fmt.Println("kill: ", ok.Error())
				err = ok
			}
			fmt.Println("Proc with pid: ", pid, " killed")
		}
	}
	return err
}

func (c *PS) Exec(args []string) error {
	procList, err := ps.Processes()
	if err != nil {
		return err
	}

	for proc := range procList {
		var process ps.Process
		process = procList[proc]
		_, err = fmt.Println(process.Executable(), " ", process.Pid(), " ", process.PPid())
	}
	return err
}

func (c *QUIT) Exec(args []string) error {
	fmt.Println("Exit...")
	os.Exit(0)
	return nil
}
