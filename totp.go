package main

import (
	"github.com/pquerna/otp/totp"
	"time"
)

func generateOTP(key string) string {
	res, _ := totp.GenerateCode(key, time.Now())
	return res
}

/*
func totpDisplay(key *otp.Key, data []byte) {
	fmt.Printf("Issuer:			 %s\n", key.Issuer())
	fmt.Printf("Account Name: %s\n", key.AccountName())
	fmt.Printf("Secret:			 %s\n", key.Secret())
	fmt.Println("Writing PNG to qr-code.png....")
	_ = ioutil.WriteFile("qr-code.png", data, 0644)
	fmt.Println("")
	fmt.Println("Please add your TOTP to your OTP Application now!")
	fmt.Println("")
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

func totpMain() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:			"Example.com",
		AccountName: "alice@example.com",
	})
	if err != nil {
		panic(err)
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	_ = png.Encode(&buf, img)

	// display the QR code to the user.
	totpDisplay(key, buf.Bytes())

	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	valid := totp.Validate(passcode, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passocde!")
		os.Exit(1)
	}
]
 */
