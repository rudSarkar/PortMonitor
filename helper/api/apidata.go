package api

import (
	"bufio"
	"context"
	"fmt"
	config "github.com/rudSarkar/PortMonitor/helper/database"
	"github.com/rudSarkar/PortMonitor/helper/options"
	pstruct "github.com/rudSarkar/PortMonitor/helper/pstructure"
	"github.com/rudSarkar/PortMonitor/helper/slackwebhook"
	"log"
	"net/http"
	"os"
	"regexp"
)

func GetOutput() {
	options := options.ParseOptions()
	regexVal := regexp.MustCompile(`^[0-9]+\/tcp(\s*)(open|filtered)(\s*)[\w'-]+`)

	// get env URI
	URI := os.Getenv("MONGODB_URI")
	ApiKey := os.Getenv("HACKERTARGET_API")
	HookUrl := os.Getenv("SLACK_WEBHOOK")
	PortCollection := config.IsConnected(URI).Database("nmap-output").Collection("ports")

	if options.File != "" {
		fileName, err := os.Open(options.File)
		if err != nil {
			log.Fatalln(err)
		}
		defer fileName.Close()

		scanner := bufio.NewScanner(fileName)

		for scanner.Scan() {
			req, err := http.NewRequest("GET", "https://api.hackertarget.com/nmap/", nil)
			if err != nil {
				log.Fatalln(err)
			}
			q := req.URL.Query()
			q.Add("q", scanner.Text())
			q.Add("apikey", ApiKey)
			req.URL.RawQuery = q.Encode()

			ApiUri, _ := http.Get(req.URL.String())
			if err != nil {
				log.Fatalln(err)
			}
			defer ApiUri.Body.Close()

			readBody := bufio.NewScanner(ApiUri.Body)

			for readBody.Scan() {
				if regexVal.MatchString(readBody.Text()) {
					singlePort := pstruct.Ports{Domain: scanner.Text(), Port: readBody.Text()}

					var result pstruct.Ports
					find := PortCollection.FindOne(context.TODO(), singlePort).Decode(&result)

					if find != nil {
						_, err := PortCollection.InsertOne(context.TODO(), singlePort)
						if err != nil {
							log.Fatalln(err)
						}

						slackwebhook.SlackSendMessage(HookUrl, fmt.Sprintf(":rotating_light:  `%s` port found opened of `%s`", singlePort.Port, singlePort.Domain))
					}
				}
			}
		}
	}
}
