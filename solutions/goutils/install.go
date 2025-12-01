package goutils

import (
	"fmt"
	"log"
	"os"
	"io"
	"net/http"
	"github.com/joho/godotenv"
)

func ImportInput(day uint) string {
	err := godotenv.Load() // 👈 load .env file
    if err != nil {
    	log.Fatal(err)
    }
    
	API_COOKIE := os.Getenv("AOC_API_COOKIE")
	cookie_session := fmt.Sprintf("session=%s", API_COOKIE)

	getUrl := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, _ := http.NewRequest("GET", getUrl, nil)

	req.Header.Set("Cookie", cookie_session)
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Failed to get http response!")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read http body!")
	}

	return string(body)
}



func main() {
	var day uint = 1
	fmt.Println(ImportInput(day))
}
