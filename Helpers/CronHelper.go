package Helpers

import (
	"github.com/jasonlvhit/gocron"
	"os"
	. "ys/Log"
)

func CreateCronJob(data map[string]string) {
	gocron.Every(15).Seconds().Do(func() {
		file, err := os.OpenFile("./Data.txt", os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			CreateLogJson("Error", "CreateLogFile", "Error when opening the file while logging the operation to the file.", err.Error())
		}
		for key, val := range data {
			// Save to File
			logDataMessage := `{"key":"` + key + `", "value":" ` + val + `"}`

			_, err = file.WriteString(logDataMessage + "\n")
			if err != nil {
				CreateLogJson("Error", "CreateLogFile", "Error while writing data to the file while logging the operation to the file.", err.Error())
			}
		}
	})
	<-gocron.Start()
}
