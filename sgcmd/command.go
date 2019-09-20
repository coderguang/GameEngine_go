package sgcmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
	"github.com/coderguang/GameEngine_go/sgthread"
)

func init() {
	cmdMap = make(map[string]CmdData)
	cmdKeys = []string{}
	RegistCmd("q", "[\"q\"]:exit system", exitCmd)
}

type CmdData struct {
	Cmd     string
	Help    string
	FuncPtr func([]string)
}

var cmdMap map[string]CmdData
var cmdKeys []string

func exitCmd(cmdstr []string) {
	sglog.Info("exit the system by cmd")
	sgserver.StopAllServer()
	sgthread.SleepBySecond(3)
	os.Exit(0)
}

func RegistCmd(cmd string, help string, funcPtr func(cmdstr []string)) bool {

	if _, ok := cmdMap[cmd]; ok {
		sglog.Error("cmd ", cmd, " already be regist")
		return false
	}

	data := CmdData{}
	data.Cmd = cmd
	data.Help = help
	data.FuncPtr = funcPtr

	cmdMap[cmd] = data
	cmdKeys = append(cmdKeys, cmd)
	sort.Strings(cmdKeys)
	return true
}

func runCmd(cmd string, cmdstr []string) bool {
	data, ok := cmdMap[cmd]
	if !ok {
		sglog.Error("cmd ", cmd, " not regist")
		return false
	}
	data.FuncPtr(cmdstr)
	return true
}

func printHelp() {
	sglog.Debug("you should input a cmd like this format (string json array):")
	sglog.Debug("[\"cmdstring\",\"param1\",\"param2\"]")

	for _, v := range cmdKeys {
		cmd, ok := cmdMap[v]
		if ok {
			sglog.Debug(cmd.Help)
		}
	}
}

func StartCmdWaitInput() bool {
	sglog.Info("\n====please input your cmd====")
	var str string
	fmt.Scanln(&str)
	sglog.Debug("input data is ", str, " ===========>")

	cmdInputData := []string{}
	err := json.Unmarshal([]byte(str), &cmdInputData)
	if err != nil {
		//sglog.Error("input not a valid json array(string)")
		return false
	}
	if len(cmdInputData) <= 0 {
		return false
	}
	if !runCmd(cmdInputData[0], cmdInputData) {
		return false
	}
	sglog.Info("cmd run complete============\n\n")
	return true
}

func StartCmdWaitInputLoop() {
	for {
		if !StartCmdWaitInput() {
			printHelp()
		}
	}
}
