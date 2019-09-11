package main

import (
	"flag"
	"fmt"
	"github.com/skip2/go-qrcode"

	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	newFlag *flag.FlagSet
	help    bool
	txFile  string
	qrSize  int
)

func init() {
	newFlag = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	newFlag.BoolVar(&help, "h", false, "display help")
	newFlag.IntVar(&qrSize, "s", 256, "QR Code size")
	newFlag.StringVar(&txFile, "f", "tx.txt", "unsigned tx file")
	newFlag.Usage = usage
}

func usage() {
	_, _ = fmt.Println("Options:")
	newFlag.PrintDefaults()
}

func main() {
	_ = newFlag.Parse(os.Args[1:])
	if help {
		newFlag.Usage()
		return
	}
	tx, err := ioutil.ReadFile(txFile)
	if err != nil {
		panic(err)
	}
	//remove '/n'
	err = qrcode.WriteFile(string(tx[:len(tx)-1]), qrcode.Low, qrSize, "tx.png")
	if err != nil {
		panic(err)
	}
}

func makeUnsignedSendTxSigQr() {
	logSep := "-------------------------------"
	//step 1: build tx
	fromAddr := "coinex1hhh9afg5n2jjuephtjj4j8yyc2uqqw3unszrjc"
	toAddr := "coinex1rd3tgkzd8q8akaug53hnqwhr378xfeljchmzls"
	chainId := "coinexdex-test1"
	amount := 1000000
	token := "cet"
	gas := 6000
	fees := 200000
	feeMoney := "cet"
	outFilePath := "tx.txt"
	txStr := fmt.Sprintf("cetcli tx send %s %d%s "+
		"--from=%s "+
		"--chain-id=%s "+
		"--gas=%d "+
		"--fees=%d%s "+
		"--generate-unsigned-tx",
		toAddr, amount, token, fromAddr, chainId, gas, fees, feeMoney)

	fmt.Println(txStr)
	var cmd *exec.Cmd

	args := strings.Split(txStr, " ")
	cmd = exec.Command(args[0], args[1:]...)
	fd, err := os.OpenFile(outFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	cmd.Stdout = fd
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fd.Close()
	tx, err := ioutil.ReadFile(outFilePath)
	if err != nil {
		panic(err)
	}
	//remove '/n'
	fmt.Println(string(tx[:len(tx)-1]))
	fmt.Println(logSep)
	err = qrcode.WriteFile(string(tx[:len(tx)-1]), qrcode.Low, 256, "qr.png")
	if err != nil {
		panic(err)
	}

}

//not use now
func makeUnsignedSendTxSig() {
	logSep := "-------------------------------"
	//step 1: build tx
	fromAddr := "coinex1hhh9afg5n2jjuephtjj4j8yyc2uqqw3unszrjc"
	toAddr := "coinex1rd3tgkzd8q8akaug53hnqwhr378xfeljchmzls"
	chainId := "coinexdex-test1"
	amount := 1000000
	token := "cet"
	gas := 6000
	fees := 200000
	feeMoney := "cet"
	outFilePath := "tx.txt"
	txStr := fmt.Sprintf("cetcli tx send %s %d%s "+
		"--from=%s "+
		"--chain-id=%s "+
		"--gas=%d "+
		"--fees=%d%s "+
		"--generate-only",
		toAddr, amount, token, fromAddr, chainId, gas, fees, feeMoney)

	var cmd *exec.Cmd

	args := strings.Split(txStr, " ")
	cmd = exec.Command(args[0], args[1:]...)
	fd, err := os.OpenFile(outFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	cmd.Stdout = fd
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fd.Close()
	tx, err := ioutil.ReadFile(outFilePath)
	if err != nil {
		panic(err)
	}
	//remove '/n'
	fmt.Println(string(tx[:len(tx)-1]))
	fmt.Println(logSep)

	//step 2: get account seq and num
	queryAccountStr := fmt.Sprintf("cetcli query account %s --trust-node", fromAddr)
	fmt.Println(queryAccountStr)
	args = strings.Split(queryAccountStr, " ")
	cmd = exec.Command(args[0], args[1:]...)
	fd, err = os.OpenFile(outFilePath, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	cmd.Stdout = fd
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fd.Close()
	account, err := ioutil.ReadFile(outFilePath)
	if err != nil {
		panic(err)
	}
	//remove '/n'
	accountStr := string(account[:len(account)-1])
	fmt.Println(accountStr)
	index := strings.Index(accountStr, "accountnumber: ")
	if index == -1 || (index+len("accountnumber: ") == len(accountStr)) {
		panic("get accountnumber failed")
	}
	accountNumStr := strings.Split(accountStr[index+len("accountnumber: "):], "\n")[0]
	fmt.Println(accountNumStr)
	fmt.Println(logSep)

	accountNum, err := strconv.ParseUint(accountNumStr, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(accountNum)
	fmt.Println(logSep)

	seqSearchStr := accountStr[index:]
	index = strings.Index(seqSearchStr, "sequence: ")
	if index == -1 || (index+len("sequence: ") == len(seqSearchStr)) {
		panic("get accountnumber failed")
	}
	sequenceNumStr := strings.Split(seqSearchStr[index+len("sequence: "):], "\n")[0]
	sequenceNum, err := strconv.ParseUint(sequenceNumStr, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(sequenceNum)
	fmt.Println(logSep)

	//step 3: build unsigned tx
}
