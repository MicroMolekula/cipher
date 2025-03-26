package lr

import (
	"flag"
	"fmt"
	"github.com/cipher/rsa"
	"github.com/cipher/verify"
	"log"
	"math/big"
)

var keysFilePath = "./lr/lr4"

func Lr4() {
	genKey := flag.Bool("genKey", false, "Генерирует ключи rsa")
	sign := flag.Bool("signFile", false, "Подписывает файл")
	sigFilePath := flag.String("sigFilePath", "./lr/lr4/file.sig", "Путь для сохранения подписи")
	filePath := flag.String("filePath", "./lr/lr4/file.txt", "Путь к файлу который нужно подписать")
	vrfy := flag.Bool("verify", false, "Проверяет подлинность файла")
	flag.Parse()
	switch {
	case *genKey:
		privateKey, publicKey, err := rsa.GenerateKeys(64)
		if err != nil {
			log.Fatal(err)
		}
		verify.SaveKeys(privateKey, publicKey, keysFilePath)
		fmt.Println("Ключи сгенерированы")
	case *sign:
		hashInt, err := verify.GetHashValueFile(*filePath)
		if err != nil {
			log.Fatal(err)
		}
		privateKey, _, err := verify.GetKeys(keysFilePath)
		if err != nil {
			log.Fatal(err)
		}
		secretKey, err := rsa.HexToKey(privateKey)
		if err != nil {
			log.Fatal(err)
		}
		sig := rsa.Code(hashInt, secretKey)
		if err = verify.SaveSignatureToFile(sig.Text(16), *sigFilePath); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Файл подписан")
		}
	case *vrfy:
		sig, err := verify.GetSignature(*sigFilePath)
		if err != nil {
			log.Fatal(err)
		}
		sigInt, ok := big.NewInt(0).SetString(sig, 16)
		if !ok {
			log.Fatal("Ошибка чтения подписи")
		}
		hashFile, err := verify.GetHashValueFile(*filePath)
		if err != nil {
			log.Fatal(err)
		}
		_, publicKey, err := verify.GetKeys(keysFilePath)
		if err != nil {
			log.Fatal(err)
		}
		pubKey, err := rsa.HexToKey(publicKey)
		if err != nil {
			log.Fatal(err)
		}
		decryptedSigInt := rsa.Decode(sigInt, pubKey)
		if decryptedSigInt.Text(10) == hashFile.Mod(hashFile, pubKey.N).Text(10) {
			fmt.Println("Файл подлинный")
		} else {
			fmt.Println("Файл не подлинный")
		}
	}
}
