package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/huttotw/pkpass-go"
)

func main() {
	r := gin.Default()
	r.GET("", CreatePkPass)

	r.Run(":8080")
}

func CreatePkPass(c *gin.Context) {
	p12FileName := "certs/flexable.p12"
	p12FilePassword := "password"
	keyFile := "certs/signerKey.key"
	certFile := "certs/signerCert.pem"
	wwdrFile := "certs/wwdr.pem"
	passRawFilePath := "pass.raw" //in this folder we need to store all type of image and pass.json
	passFileNameWithPath := "pass.raw/pass.json"

	pkpassFileName := "Coupon.pkpass"

	if err := CreateP12File(certFile, keyFile, wwdrFile, p12FileName, p12FilePassword); err != nil {
		log.Println("CreatePkPass CreateP12File Error: ", err.Error())
	}

	if err := CreatePassJson(passFileNameWithPath); err != nil {
		log.Println("CreatePkPass CreatePassJson: ", err.Error())
	}

	cert, err := os.Open(p12FileName)
	if err != nil {
		log.Println("CreatePkPass cert Open Error: ", err.Error())
	}
	defer cert.Close()

	r, err := pkpass.New(passRawFilePath, p12FilePassword, cert)
	if err != nil {
		log.Println("CreatePkPass pkpass.NewError: ", err.Error())
		panic(err)
	}

	f, err := os.Create(pkpassFileName)
	if err != nil {
		log.Println("CreatePkPass create .pkpass file Error: ", err.Error())
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		log.Println("CreatePkPass copy file Error: ", err.Error())
	}

	c.Header("Content-Type", "application/vnd.apple.pkpass")
	c.FileAttachment(pkpassFileName, pkpassFileName)
}
func CreatePassJson(passFileNameWithPath string) error {
	passData := Pass{
		
	}
	passJSON, err := json.Marshal(passData)

	if err != nil {
		return err
	}

	err = os.WriteFile(passFileNameWithPath, passJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}
func CreateP12File(certFile, keyFile, wwdrFile, generateP12FileName, generateFilePassword string) error {
	cmd := exec.Command(
		"openssl",
		"pkcs12",
		"-export",
		"-out", generateP12FileName,
		"-inkey", keyFile,
		"-in", certFile,
		"-certfile", wwdrFile,
		"-password", "pass:"+generateFilePassword,
	)

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute command: %s\n", err.Error())
		return err
	}
	return nil
}
