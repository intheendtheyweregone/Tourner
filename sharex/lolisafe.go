package sharex

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

// HandleUpload Upload the file from the local server to the destination origin.
func HandleUpload(f multipart.File, h *multipart.FileHeader, s string) ([]byte, error) {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, err := writer.CreateFormFile("files[]", h.Filename)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(f)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	part.Write(b)
	writer.Close()

	req, _ := http.NewRequest("POST", s, buf)

	req.Header.Add("authority", s)
	// * Will need to change dynamically for ChibiSafe
	req.Header.Add("accept", "accept: application/vnd.chibisafe.json")
	req.Header.Add("accept-language", "en-US,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("origin", s)
	req.Header.Add("referer", s)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, error := http.DefaultClient.Do(req)

	if error != nil {
		fmt.Println("[Tourner] Something went wrong with the request")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("\n [Tourner] successfully got response from origin: %v", s)
	return body, nil
}
