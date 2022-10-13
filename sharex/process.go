package sharex

import (
	"fmt"
	"mime/multipart"
)

func ProcessUpload(file multipart.File, handler *multipart.FileHeader, site string, siteType string) ([]byte, error) {
	body := make(chan string)
	errs := make(chan error, 1)

	switch siteType {
	case "LoliSafe":
		go func() {
			_, err := HandleUpload(file, handler, site)
			if err != nil {
				fmt.Printf("\n[Tourner] %v", err)
			}
		}()
		msg := <-body
		msg2 := <-errs

		return []byte(msg), msg2

	default:
		fmt.Println("?")
	}

	return nil, nil
}
