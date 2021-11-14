package Handlers

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	. "ys/Helpers"
	. "ys/Log"
	. "ys/Models"
)

func StartCroneJob() {
	CreateCronJob(cache.GetAll())
}

// WriteDataFromFileToCache : Uygulama tekrar başlayınca dosyada kayıtlı olan dataları cache içine yükleyecek.
func WriteDataFromFileToCache() {
	if _, err := os.Stat("Data.txt"); err == nil {
		f, err := os.OpenFile("Data.txt", os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("open file error: %v", err)
			return
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("read file line error: %v", err)
				return
			}
			var rm DataModel
			_ = line // GET the line string
			err = json.Unmarshal([]byte(line), &rm)
			if err != nil {
				CreateLogJson("Error", "WriteDataFromFileToCache", "Json UnMarshal", err.Error())
			}
			cache.Set(rm.Key, rm.Value)
		}
	}
}
