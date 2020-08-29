package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const _timeout int = 20000
const _readWriteTimeout int = 60000
const _ignoreSSLCheck bool = true
const _disableWebProxy bool = false

func DoGet(url string, textParams map[string]interface{}, headerParams map[string]string) (result []byte, err error) {
	var request *http.Request
	var response *http.Response
	var body []byte
	if len(textParams) > 0 {
		url = BuildRequestUrl(url, textParams)
	}
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	if len(headerParams) > 0 {
		request = GetWebRequest(request, headerParams)
	}
	var client = &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	return body, err
}

func DoPost(url string, textParams map[string]interface{}, headerParams map[string]string, otherParams map[string]interface{}) (result []byte, err error) {
	var request *http.Request
	var response *http.Response
	var body []byte
	var paramStr string
	var client = &http.Client{}
	if len(textParams) > 0 {
		paramStr, err = BuildPostRequestUrl(textParams)
	}
	if len(headerParams) > 0 {
		GetWebRequest(request, headerParams)
	}
	if len(otherParams) > 0 {
		url = BuildRequestUrl(url, otherParams)
	}
	fmt.Println(url)
	request, err = http.NewRequest("POST", url, strings.NewReader(paramStr))
	if err != nil {
		return result, err
	}
	response, err = client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	return body, err
}

func DoFilePost(url string, textParams map[string]interface{}, paramName, path string) (result []byte, err error) {
	var request *http.Request
	var response *http.Response
	var body []byte
	var client = &http.Client{}
	var file *os.File
	var part io.Writer
	file, err = os.Open(path)
	if err != nil {
		return body, err
	}
	defer file.Close()
	if len(textParams) > 0 {
		url = BuildRequestUrl(url, textParams)
	}
	_body := &bytes.Buffer{}
	writer := multipart.NewWriter(_body)
	part, err = writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return body, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return body, err
	}
	//for key, val := range textParams {
	//	_ = writer.WriteField(key, val.(string))
	//}
	err = writer.Close()
	if err != nil {
		return body, err
	}

	request, err = http.NewRequest("POST", url, _body)
	if err != nil {
		return body, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	response, err = client.Do(request)
	if err != nil {
		return body, err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return body, err
	}
	return body, err
}

func BuildPostRequestUrl(params map[string]interface{}) (result string, err error) {
	var marshal []byte
	marshal, err = json.Marshal(params)
	if err != nil {
		return result, err
	}
	return string(marshal), err
}

func BuildRequestUrl(url string, params map[string]interface{}) string {
	if len(params) == 0 {
		return url
	}
	return url + "?" + BuildQuery(params)
}

func BuildQuery(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	query := ""
	hasParam := false
	for k, v := range params {
		if k != "" && v != "" {
			if hasParam {
				query += "&"
			}
			query = query + fmt.Sprintf("%s=%s", k, UrlEncode(v.(string)))
			hasParam = true
		}
	}
	return query
}

func UrlEncode(v string) string {
	return url.QueryEscape(v)
	//return v
}

func GetWebRequest(req *http.Request, headerParams map[string]string) *http.Request {
	if len(headerParams) == 0 {
		return req
	}
	for k, v := range headerParams {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return req
}
