package main

import (
	. "ys/Handlers"
	. "ys/cmd"
)

func main() {
	WriteDataFromFileToCache() // Uygulama tekrar başlayınca dosyada kayıtlı olan dataları cache içine yükleyecek.
	go StartCroneJob()         // Cron-job olarak ayarlanmış, belirli bir sürede çalışmasını istediğimiz fonksiyonu çalıştıracak.
	Run()                      // handler ve api startını verecek.
}
