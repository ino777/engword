package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"regexp"
	"encoding/json"
)
const (
	url string = "https://toiguru.jp/toeic-vocabulary-list"
)
func main() {
	//
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("http err:", err)
	}
	defer res.Body.Close()
	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	html := string(byteArray)
	
	// 単語取得
	rep := regexp.MustCompile(`<td>(.*?)</td>`)
	result := rep.FindAllString(html, -1)
	
	// En-Jp Map
	ejMap  := map[string]string{}
	
	var en, jp string
	var tdIndex int
	for i:=0; i<len(result); i++ {
	    ej := strings.Split(result[i], "<br>")
	    if len(ej) == 1 {
	        continue
	    }
	   
	    en = ej[0][4:]
	   
	    tdIndex = strings.Index(ej[1], "</td>")
	    if tdIndex == -1 {
	        jp = ej[2]
	    } else {
	        jp = ej[1][:tdIndex]
	    }
	   
	    ejMap[en] = jp
	}
	
	// Jsonに変換
	bytes, err := json.Marshal(ejMap)
	if err != nil {
	    log.Fatal("Json marshal err:", err)
	}
	
	fmt.Println(string(bytes))
}
