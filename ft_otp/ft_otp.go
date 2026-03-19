package main

import (
	"flag"
	"fmt"
	"ft_otp/internal/encrypt"
	"ft_otp/internal/utils"
	totp "ft_otp/pkg"
	"os"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/terminal"
)

type KeyArgs struct {
	hexKey  string
	keyFile string
}

var keyArgs KeyArgs

func generateQR(secret string) error {
	const (
		label  string = "user"
		issuer string = "ft_otp"
	)
	totpCode := fmt.Sprintf("otpauth://totp/%s/?secret=%s&issuer=%s")
	qrc, err := qrcode.New(totpCode)
	if err != nil {
		return err
	}
	w := terminal.New()
	if err := qrc.Save(w); err != nil {
		return err
	}
	return nil
}

func parseArgs() error {
	const (
		usageHex = "Provide hexadecimal key to store"
		usageKey = "Provide key file to generate otp"
	)
	flag.StringVar(
		&keyArgs.hexKey, "g", "", usageHex,
	)
	flag.StringVar(
		&keyArgs.keyFile, "k", "", usageKey,
	)
	flag.Parse()
	if keyArgs.hexKey == "" && keyArgs.keyFile == "" {
		return utils.ErrArgParse
	}
	return nil
}

func main() {
	if err := parseArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		flag.Usage()
		os.Exit(1)
	}

	if keyArgs.hexKey != "" {
		if err := encrypt.EncryptKey(keyArgs.hexKey); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Key saved succesfully in ft_otp.key")
		if err := generateQR(keyArgs.hexKey); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if keyArgs.keyFile != "" {
		key, err := encrypt.DecryptKey(keyArgs.keyFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		totp.TOTP(key)
	}
}
