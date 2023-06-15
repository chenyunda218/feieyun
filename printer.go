package feieyun

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Printer struct {
	Endpoint string
	Ukey     string
	Sn       string
	User     string
}

type ResponseBody struct {
	Ret                string  `json:"ret"`
	Msg                string  `json:"msg"`
	Data               *string `json:"data"`
	ServerExecutedTime int     `json:"serverExecutedTime"`
}

const PrintApiName = "Open_printMsg"

func (p Printer) print(content string) error {
	timestamp := fmt.Sprint(time.Now().Unix())
	sig := fmt.Sprintf("%s%s%s", p.User, p.Ukey, timestamp)
	h := sha1.New()
	io.WriteString(h, sig)
	v := make(map[string][]string)
	v["user"] = []string{p.User}
	v["stime"] = []string{timestamp}
	v["sig"] = []string{fmt.Sprintf("%x", h.Sum(nil))}
	v["apiname"] = []string{PrintApiName}
	v["sn"] = []string{p.Sn}
	v["content"] = []string{content}
	qs := url.Values(v)
	resp, err := http.PostForm(p.Endpoint+"/Api/Open/", qs)
	defer resp.Body.Close()
	var b ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&b)
	fmt.Println(b)
	return err
}
