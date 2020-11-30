package mains

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/PuerkitoBio/goquery"
	"github.com/kirinlabs/HttpRequest"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	CAMBRIDGE_HOST = "https://dictionary.cambridge.org"
	CAMBRIDGE_URI = "/dictionary/english-chinese-simplified/"
	CACHE_PATH = "cache/"
	REQUEST_TIMEONT = 10
	MAX_ROUTINE_NUM = 10
)


type Example struct {
	Cn string `json:"cn" p:"中文"`
	En string `json:"en" p:"英文"`
}


type Explain struct {
	Phrase string `json:"phrase" p:"短语"`
	Speech string `json:"speech" p:"词性"`
	Tag string `json:"tag" p:"标记"`
	Cn string `json:"cn" p:"中文"`
	En string `json:"en" p:"英文"`
	Example []Example `json:"en" p:"例句"`
}


type Info struct {
	Word       string `json:"word" p:"单词"`
	Speech     string `json:"speech" p:"词性"`
	FullSpeech string `json:"full_speech" p:"单词（全）"`
	UsMp3      string `json:"us_mp3" p:"美式音频"`
	UsPhonetic string `json:"us_phonetic" p:"美式音标"`
	UkMp3      string `json:"uk_mp3" p:"英式读音"`
	UkPhonetic string `json:"uk_phonetic" p:"英式音标"`
	Explain    []Explain `json:"word" p:"释义"`
}


type Excel struct {
	Success []Info
	Fail    []string
}


func putCache(path string, cache string) {
	dirPath := path[0 : strings.LastIndex(path, "/")]
	if e := os.MkdirAll(dirPath, 0777); e != nil {
		fmt.Println(e)
	}

	var content bytes.Buffer
	b := []byte(cache)
	w := zlib.NewWriter(&content)
	w.Write(b)
	w.Close()
	if err := ioutil.WriteFile(path, content.Bytes(), 0777); err != nil {
		fmt.Println(err)
	}
}

func getCache(path string) string {
	cache, _ := ioutil.ReadFile(path)
	var out bytes.Buffer
	r, _ := zlib.NewReader(bytes.NewBuffer(cache))
	io.Copy(&out, r)
	return out.String()
}

func formatExample(example []Example) string {
	str := ""
	for _, item := range example{
		str += item.En + "  " + item.Cn + "\n"
	}
	str = strings.TrimRight(str, "\n")
	return str
}


func requestPage(word string) (string, bool, string) {
	var html string
	toCache := false
	times := 1
	responseTime := ""
	for {
		response, err := HttpRequest.SetHeaders(map[string]string{
			"Accept" : "text/html",
			"User-Agent" : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36",
		}).SetTimeout(REQUEST_TIMEONT).Get(CAMBRIDGE_HOST + CAMBRIDGE_URI + word)

		if err != nil {
			times ++
		} else {
			if response.StatusCode() == 200 {
				html, _ = response.Content()
				responseTime = response.Time()
				toCache = true
				break
			} else if times > 5 {
				html = ""
				break
			} else {
				times ++
			}
		}
	}
	return html, toCache, responseTime
}


func execute(filePath string) Excel {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	var excel Excel
	rows, _ := f.GetRows("Sheet1")

	for line, row := range rows {
		for _, cell := range row {
			fmt.Println("----- 第 "+strconv.Itoa(line+1)+" 行 "+ cell)
			var html string
			toCache := false
			path := CACHE_PATH + strings.Replace(cell, " ", "-", -1) + "/001.bin"

			if _, err := os.Stat(path); err != nil {
				html, toCache, _ = requestPage(cell)
			} else {
				html = getCache(path)
			}

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
			word := dom.Find(".dhw span").Eq(0).Text()

			if word == "" {
				excel.Fail = append(excel.Fail, cell)
			} else {
				var info Info
				info.Word = word
				info.Speech = dom.Find(".dpos").Eq(0).Text()
				info.FullSpeech = dom.Find(".dpos-g").Eq(0).Text()
				info.UkMp3, _ = dom.Find("source").Eq(0).Attr("src")
				info.UsMp3, _ = dom.Find("source").Eq(2).Attr("src")
				info.UkPhonetic = dom.Find(".dpron").Eq(0).Text()
				info.UsPhonetic = dom.Find(".dpron").Eq(1).Text()

				dom.Find(".ddef_block").Each(func(i int, nextDom *goquery.Selection) {
					var explain Explain
					explain.Speech = nextDom.Find(".dgram").Eq(0).Text()
					explain.En = nextDom.Find(".ddef_d").Eq(0).Text()
					explain.Cn = nextDom.Find(".def-body span").Eq(0).Text()
					explain.Phrase = nextDom.Parent().Siblings().Find(".phrase-title").Text()
					explain.Tag = nextDom.Find(".epp-xref").Eq(0).Text()

					nextDom.Find(".dexamp").Each(func(i int, two *goquery.Selection) {
						var example Example
						example.En = two.Find(".deg").Text()
						example.Cn = two.Find(".hdb").Text()
						explain.Example = append(explain.Example, example)
					})
					//fmt.Printf("%+v\n", explain)
					info.Explain = append(info.Explain, explain)
				})
				excel.Success = append(excel.Success, info)

				if toCache {
					putCache(path, html)
				}
			}
		}
	}
	return excel
}


func export(excel Excel, fileName string) {
	NewFile := excelize.NewFile()
	line := 1

	NewFile.SetCellValue("Sheet1", "A"+ strconv.Itoa(line), "单词")
	NewFile.SetCellValue("Sheet1", "B"+ strconv.Itoa(line), "词性")
	NewFile.SetCellValue("Sheet1", "C"+ strconv.Itoa(line), "词性（全）")
	NewFile.SetCellValue("Sheet1", "D"+ strconv.Itoa(line), "美式音标")
	NewFile.SetCellValue("Sheet1", "E"+ strconv.Itoa(line), "英式音标")
	NewFile.SetCellValue("Sheet1", "F"+ strconv.Itoa(line), "释义短语")
	NewFile.SetCellValue("Sheet1", "G"+ strconv.Itoa(line), "释义词性")
	NewFile.SetCellValue("Sheet1", "H"+ strconv.Itoa(line), "释义标签")
	NewFile.SetCellValue("Sheet1", "I"+ strconv.Itoa(line), "释义英文")
	NewFile.SetCellValue("Sheet1", "J"+ strconv.Itoa(line), "释义中文")
	NewFile.SetCellValue("Sheet1", "K"+ strconv.Itoa(line), "例子")
	line ++

	for _, word := range excel.Success {
		for index, explain := range word.Explain{
			if index == 0 {
				NewFile.SetCellValue("Sheet1", "A"+ strconv.Itoa(line), word.Word)
				NewFile.SetCellValue("Sheet1", "B"+ strconv.Itoa(line), word.Speech)
				NewFile.SetCellValue("Sheet1", "C"+ strconv.Itoa(line), word.FullSpeech)
				NewFile.SetCellValue("Sheet1", "D"+ strconv.Itoa(line), word.UsPhonetic)
				NewFile.SetCellValue("Sheet1", "E"+ strconv.Itoa(line), word.UkPhonetic)
			}

			NewFile.SetCellValue("Sheet1", "F"+ strconv.Itoa(line), explain.Phrase)
			NewFile.SetCellValue("Sheet1", "G"+ strconv.Itoa(line), explain.Speech)
			NewFile.SetCellValue("Sheet1", "H"+ strconv.Itoa(line), explain.Tag)
			NewFile.SetCellValue("Sheet1", "I"+ strconv.Itoa(line), explain.En)
			NewFile.SetCellValue("Sheet1", "J"+ strconv.Itoa(line), explain.Cn)
			NewFile.SetCellValue("Sheet1", "K"+ strconv.Itoa(line), formatExample(explain.Example))
			line ++
		}
	}

	if excel.Fail != nil {
		line = line + 2
		NewFile.SetCellValue("Sheet1", "A"+ strconv.Itoa(line), "无法爬取的单词")
		line ++
		for _, item := range excel.Fail{
			NewFile.SetCellValue("Sheet1", "A"+ strconv.Itoa(line), item)
			line ++
		}
	}

	if err := NewFile.SaveAs(fileName + "-result.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func routineCacheRun(word string, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	html := ""
	responseTime := ""
	toCache := false
	path := CACHE_PATH + word + "/001.bin"
	if _, err := os.Stat(path); err != nil {
		html, toCache, responseTime = requestPage(word)
		if rand.Intn(10) == 1 {
			fmt.Println("***** 剑桥词典响应时间：" + responseTime)
		}
	}
	if html != "" && toCache {
		putCache(path, html)
	}
	// 完成任务则从ch推出数据
	<- ch
}


func routineCache(fileName string) {
	f, _ := excelize.OpenFile(fileName)
	rows, _ := f.GetRows("Sheet1")
	//受限于带宽设置最大协程数
	ch := make(chan int, MAX_ROUTINE_NUM)
	for _, row := range rows {
		for _, cell := range row {
			// 开启任务协程前往ch塞一个数据
			// 如果ch满了则会处于阻塞，从而达到限制最大协程的功能
			ch <- 1
			go routineCacheRun(cell, ch)
		}
	}
}


func init() {
	if _, e := os.Stat(CACHE_PATH); e != nil {
		if err := os.MkdirAll(CACHE_PATH, 0777); err != nil {
			panic("初始化缓存目录失败")
		}
	}
}


func main() {
	pwd, _ := os.Getwd()
	//获取文件或目录相关信息
	fileList, err := ioutil.ReadDir(pwd)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range fileList {
		fileName := file.Name()
		fileExt := filepath.Ext(fileName)

		if strings.ToLower(fileExt) == ".xlsx" && strings.Index(fileName, "-result.xlsx") == -1 {
			fmt.Println("正在执行：" + file.Name())
			go routineCache(fileName)
			excel := execute(fileName)
			export(excel, strings.Replace(fileName, fileExt, "", -1))
		}
	}
}
