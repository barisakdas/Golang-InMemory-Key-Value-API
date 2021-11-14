package Log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ServerRequestLog : Handler üzerine istek geldiği zaman bu isteği loga çevirip kaydedecek.
func ServerRequestLog(handlerType, endpoint string) {
	logData :=
		`{
"Type" : "Info",
"Endpoint" : "` + endpoint + `",
"Handler" : "` + handlerType + `",
"Description" : "A request was sent to the endpoint.",
"CreatedDate" : "` + strings.Split(time.Now().String(), "+")[0] + `"
}`

	WriteLogToFile(logData, "RequestLogs.txt")
}

// CreateLogJson : Çalışan fonksiyonlar üzerindeki işlemlerin logunu tutacak olan fonksiyon
func CreateLogJson(logType, logFunction, logName, logErrorMessage string) {
	logData :=
		`{
"Type" : "` + logType + `",
"Function" : "` + logFunction + `",
"Name" : "` + logName + `",
"Description" : "` + logErrorMessage + `",
"CreatedDate" : "` + strings.Split(time.Now().String(), "+")[0] + `"
}`

	fmt.Println(logData)
	WriteLogToFile(logData, "RunTimeLogs.txt")
}

// WriteLogToFile : Elde edilen logları ilgili dosyalara yazacak olan fonksiyon
func WriteLogToFile(logDataMessage, logFileName string) {
	file, err := os.OpenFile("./"+logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		CreateLogJson("Error", "CreateLogFile", "Error when opening the file while logging the operation to the file.", err.Error())
	}

	defer file.Close()

	_, err = file.WriteString(logDataMessage + ",\n")
	if err != nil {
		CreateLogJson("Error", "CreateLogFile", "Error while writing data to the file while logging the operation to the file.", err.Error())
	}
}
