package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		logEntry := fmt.Sprintf("%s - [%s] \"%s %s\" %v\n",
			r.RemoteAddr,
			startTime.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.RequestURI,
			time.Since(startTime))

		logToFile(logEntry)
	})
}

func logToFile(message string) {
	filePath := "fichiers/logfile.txt"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erreur lors de l'ouverture du fichier journal :", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(message); err != nil {
		log.Println("Erreur lors de l'écriture dans le fichier journal :", err)
	}
}
