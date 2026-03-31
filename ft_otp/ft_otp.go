package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"ft_otp/internal/encrypt"
	"ft_otp/internal/utils"
	"ft_otp/pkg/totp"
	"os"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
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
	encodedSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	totpCode := fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s", label, encodedSecret, issuer)
	qrc, err := qrcode.New(totpCode)
	if err != nil {
		return err
	}
	w, err := standard.New("./otp-code.jpeg")
	if err != nil {
		return err
	}
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
		fmt.Println("Do you want to generate a QR code? (y/n)")
		for {
			var input string
			_, err := fmt.Scanln(&input)
			switch {
			case err != nil:
				return
			case input == "y" || input == "yes":
				if err := generateQR(keyArgs.hexKey); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				fmt.Println("Code succesfully generated in ./otp-code.jpeg")
				return
			case input == "no" || input == "n" || input == "":
				return
			}
		}
	}
	if keyArgs.keyFile != "" {
		key, err := encrypt.DecryptKey(keyArgs.keyFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		code := totp.TOTP(key)
		fmt.Printf("Code: %06s\n", code)
	}
}
