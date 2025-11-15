package main

import (
//	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	 "github.com/weiji6/hacker-support/httptool"
	"github.com/weiji6/hacker-support/encrypt"
)


func main(){
	url1 := "https://gtainmuxi.muxixyz.com/api/v1/organization/code"
	reqForpassport, err := http.NewRequest("GET", url1, nil)
	checkError(err)
	reqForpassport.Header.Add("Code","MUXI_task_001")

	myClient := new(http.Client)
	myClient.Timeout = 10 * time.Second
	
	res, err := myClient.Do(reqForpassport)
	checkError(err)
	res.Body.Close()

	passport := res.Header.Values("Passport")[0]
	
	fmt.Println("passport:")
	fmt.Println(passport)
	fmt.Println()
	

	resp, err := myClient.Do(reqForpassport)
	checkError(err)
	res.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(body))

	url2 := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key"
	reqForKey, err := http.NewRequest("GET", url2, nil)
	checkError(err)


	reqForKey.Header.Add("Passport", passport)
	reqForKey.Header.Add("Code","MUXI_task_001")

	respOfKey, err := myClient.Do(reqForKey)
	checkError(err)
	defer respOfKey.Body.Close()

	body2, err := io.ReadAll(respOfKey.Body)
	checkError(err)
	fmt.Println(string(body2))

	keyAndcode, err := Base64URLDecode("c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=")
	checkError(err)
	fmt.Println(string(keyAndcode))

	key := "MuxiStudio203304"
	code := "for {go func(){time.Sleep(1*time.Hour)}()}"

	secretCode, err := encrypt.AESEncryptOutInBase64([]byte(code),[]byte(key))
	checkError(err)
	fmt.Println(string(secretCode))


	reqForAtk, err := httptool.NewRequest("PUT", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", string(secretCode), 0)
	checkError(err)

	reqForAtk.AddHeader("Passport",passport)
	

	respOfAtk, err := reqForAtk.SendRequest()
	checkError(err)
	
	respOfAtk.ShowBody()

	reqForIris, err := httptool.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", "", 0)
	checkError(err)
	reqForIris.AddHeader("Passport",passport)
	respOfIris, err := reqForIris.SendRequest()
	checkError(err)
	respOfIris.ShowBody()

	reqForSample, err := httptool.NewRequest("GET","http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample","", 0)
	checkError(err)
	reqForSample.AddHeader("Passport",passport)
	respOfSample, err := reqForSample.SendRequest()
	checkError(err)
	respOfSample.Save("D:/杂项文件/test.jpg")

	reqForRec, err := httptool.NewRequest(httptool.POSTMETHOD, "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", "D:/杂项文件/test.jpg", httptool.FILE)
	checkError(err)
	reqForRec.AddHeader("Passport",passport)

	respOfRec, err := reqForRec.SendRequest()
	checkError(err)
	respOfRec.ShowBody()
/*
	for k, v := range res.Header {
		fmt.Println(k,":",v)
	}
	for k, v := range respOfKey.Header {
		fmt.Println(k,":",v)
	}
	respOfAtk.ShowHeader()
	respOfIris.ShowHeader()
	respOfSample.ShowHeader()
	respOfRec.ShowHeader()*/
	reqFinal, err := httptool.NewRequest(httptool.POSTMETHOD, "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", "D:/MyGo/MUXI-Learning/week3(till 25.11.16)/Heap/main.go", httptool.FILE)
	reqFinal.AddHeader("Passport",passport)
	checkError(err)
	respFinal, err  := reqFinal.SendRequest()
	checkError(err)
	respFinal.ShowBody()

}

func checkError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func Base64URLDecode(s string)([]byte, error){
	padCount := (4 - len(s) % 4) % 4
	s += strings.Repeat("=", padCount)

	return base64.URLEncoding.DecodeString(s)
}

func JWTEncoding(jwt string)(map[string]interface{}, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3{
		return nil, fmt.Errorf("无效的 JWT 格式")
	}

	payloadBytes, err := Base64URLDecode(parts[1])
	if err != nil {
	return nil, fmt.Errorf("payload 解析失败：%v", err)
	}

	var payload map[string]interface{}

	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("JSON 解析失败:%v", err)
	}

	return  payload, nil
}

