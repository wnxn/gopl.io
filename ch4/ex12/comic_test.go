package ex12

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestType(t *testing.T){
	content := Item{}
	rawStr := `{"month": "4", "num": 571, "link": "", "year": "2009", "news": "", "safe_title": "Can't Sleep", "transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}", "alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.", "img": "https://imgs.xkcd.com/comics/cant_sleep.png", "title": "Can't Sleep", "day": "20"}`
	err := json.Unmarshal([]byte(rawStr), &content)
	if err != nil{
		t.Error(err)
	}
	b,err :=json.MarshalIndent(content,"","  ")
	fmt.Print(string(b))
}

func TestGetContent(t *testing.T) {
	item, err := GetContent(23)
	if err != nil{
		t.Error(err)
	}
	b,err := json.MarshalIndent(item,"","  ")
	fmt.Println(string(b))
}

func TestSearchAsYear(t *testing.T) {
	Download()
	res := SearchAsYear(2006)
	fmt.Printf("Total %d\n", len(res))
	for _, v:=range res{
		fmt.Printf("Title %s, Url %s\n", v.Title, v.Img)
	}
}