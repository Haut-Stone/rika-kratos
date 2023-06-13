package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func getDBData(grade int, question, imageUrl string) (id int, err error) {

	// 构造uuid
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("构造 uuid 失败: %s", err)
		return
	}
	str := strings.NewReplacer("{{IMAGE_URL}}", imageUrl, "{{OCR_TXT}}", question).Replace(Template)

	params := map[string]interface{}{
		"subject_id": cast.ToString(2),
		"grade_id":   cast.ToString(grade),
		"origin":     cast.ToString(1002),
		"source_id":  cast.ToString(1007),
		"origin_id":  u2.String(),
		"question":   str,
		"book_id":    cast.ToString(99),
	}

	var u = url.Values{}
	for k, v := range params {
		u.Add(k, fmt.Sprintf("%v", v))
	}
	fmt.Println(u)
	req, _ := http.NewRequest("POST", DbOnlineUrl, strings.NewReader(u.Encode()))

	// 设置 header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("caller", "recommend-th-0510")
	resp, err := (&http.Client{}).Do(req) //nolint:bodyclose
	if err != nil {
		fmt.Println("获取数据库失败", err)
		return 0, err
	}
	var dbResp DbResp
	respByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respByte, &dbResp)
	if err != nil {
		fmt.Println("解析DB返回数据出错", err)
		return 0, err
	}
	fmt.Println(dbResp)

	if dbResp.ErrorCode != 0 {
		fmt.Println("访问DB接口出错", err)
		return 0, err
	}

	return cast.ToInt(dbResp.Data.Id), nil
}

func getAIData(questionId int) (id int, err error) {
	reqData := "question_id=" + cast.ToString(questionId)
	req, _ := http.NewRequest("POST", AnswerOnlineUrl, strings.NewReader(string(reqData)))

	appKey := "pyb02011bc4c7fcd"
	appSecret := "7f8d3651d6f6a293a7b96260f373a23c"
	timeStamp := cast.ToString(time.Now().Unix())
	md5Str := appKey + timeStamp + appSecret
	data := []byte(md5Str)
	md5New := md5.New()
	md5New.Write(data)
	sign := hex.EncodeToString(md5New.Sum(nil))

	// 设置 header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Tal-DeviceId", "recommend-th-0510")
	req.Header.Set("X-Tal-Timestamp", timeStamp)
	req.Header.Set("X-Tal-AccountId", "recommend-th-0510")
	req.Header.Set("X-Tal-Sign", sign)
	req.Header.Set("X-Tal-AppKey", appKey)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("获取AI推荐数据失败", err)
		return 0, err
	}
	var aiResp AIResp
	respByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respByte, &aiResp)
	if err != nil {
		fmt.Println("解析AI返回数据出错", err)
		return 0, err
	}

	if cast.ToInt(aiResp.ErrCode) != 0 {
		fmt.Println("访问AI接口出错出错", aiResp.ErrCode, aiResp.ErrMsg)
		return 0, err
	}

	fmt.Println("AI推荐题目信息:", aiResp.Data)

	if len(aiResp.Data.MediumRecommend) > 0 {
		return cast.ToInt(aiResp.Data.MediumRecommend[0].QuestionId), nil
	}
	return 0, nil
}

func getGpt4MathData(prompt, question, asr string) (gptStr string, err error) {

	str := strings.NewReplacer("{{asr}}", asr, "{{question}}", question).Replace(prompt)

	// 构造请求
	info := Gpt4Req{
		Message:         str,
		Model:           "gpt-4",
		PresencePenalty: 0,
		Temperature:     0,
	}
	reqData, _ := json.Marshal(info)
	fmt.Println("gpt4请求信息：--> ", string(reqData))
	req, _ := http.NewRequest("POST", Gpt4Url, strings.NewReader(string(reqData)))

	resp, err := (&http.Client{}).Do(req) //nolint:bodyclose
	if err != nil {
		fmt.Println("获取gpt4失败", err)
		return "", err
	}
	var gpt4Resp Gpt4Resp
	respByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respByte, &gpt4Resp)
	if err != nil {
		fmt.Println("解析gpt4返回数据出错", err)
		return "", err
	}
	if gpt4Resp.ErrorCode != 0 {
		fmt.Println("访问gpt4接口出错", err, gpt4Resp)
		return "", err
	}

	return gpt4Resp.Data.Result, nil
}

func FileToRows(fileName, tableName string) (*excelize.File, [][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	rows, err := f.GetRows(tableName)
	if err != nil {
		return nil, nil, err
	}
	return f, rows, nil
}

func timeCost(start time.Time) {
	tc := time.Since(start).Seconds()
	fmt.Printf("time cost = %v\n", tc)
}

func GetText4QuestionData(question string) (id int, err error) {

	params := map[string]interface{}{
		"words": question,
	}

	var u = url.Values{}
	for k, v := range params {
		u.Add(k, fmt.Sprintf("%v", v))
	}
	fmt.Println(u)
	req, _ := http.NewRequest("POST", Text4QuestionOnlineUrl, strings.NewReader(u.Encode()))

	// 设置 header online
	u2, _ := uuid.NewV4()
	var (
		timestamp = strconv.Itoa(int(time.Now().Unix()))
		appKey    = "101135b0bc4c7fcddecb3dbbe7692198"
		appSecret = "51607f8df33a23cd6f6a293a7b962367"
		nonce     = fmt.Sprintf("%d-%s", time.Now().UnixMicro(), u2)
	)
	req.Header.Set("X-Qz-AccountId", "123456789")
	req.Header.Set("X-Qz-DeviceId", "script_123456")

	// 设置 header 测试
	//u2, _ := uuid.NewV4()
	//var (
	//	timestamp = strconv.Itoa(int(time.Now().Unix()))
	//	appKey    = "a8054bc4bfc4babbec4d7a42deee4880"
	//	appSecret = "62c17c869d2d45d1dcfd0c28223fc4ed"
	//	nonce     = fmt.Sprintf("%d-%s", time.Now().UnixMicro(), u2)
	//)

	// 通用
	md5Str := "X-Qz-Timestamp=" + timestamp + "&X-Qz-Nonce=" + nonce + "&X-Qz-AppKey=" + appKey + "&app_secret=" + appSecret
	data := []byte(md5Str)
	md5New := md5.New()
	md5New.Write(data)
	sign := hex.EncodeToString(md5New.Sum(nil))
	req.Header.Set("X-Qz-Nonce", nonce)
	req.Header.Set("X-Qz-Timestamp", timestamp)
	req.Header.Set("X-Qz-AppKey", appKey)
	req.Header.Set("X-Qz-Sign", sign)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("caller", "search-xxj-0613")
	resp, err := (&http.Client{}).Do(req) //nolint:bodyclose
	if err != nil {
		fmt.Println("获取数据库失败", err)
		return 0, err
	}
	var dbResp Text4QuestionResp
	respByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respByte, &dbResp)
	if err != nil {
		fmt.Println("解析Text4Question返回数据出错", err)
		return 0, err
	}
	fmt.Println(dbResp)

	if dbResp.ErrorCode != 0 {
		fmt.Println("访问Text4QuestionResp接口出错", err)
		return 0, errors.New(dbResp.ErrorCode, "", dbResp.ErrorMsg)
	}

	if len(dbResp.Data.QuestionArr) == 0 {
		return -1, nil
	}

	fmt.Println("traceId: ", dbResp.TraceId)

	return cast.ToInt(dbResp.Data.QuestionArr[0].QuestionId), nil
}
