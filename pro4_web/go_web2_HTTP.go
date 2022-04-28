// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"sync"
// 	"time"
// )

// func testGET1() {
// 	url := "http://apis.juhe.cn/simpleWeather/query?city=北京&key=087d7d10f700d20e27bb753cd806e40b"
// 	// url := "https://www.baidu.com"
// 	r, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()
// 	b, _ := ioutil.ReadAll(r.Body)
// 	fmt.Printf("b: %v\n", string(b))
// }
// func testGET2() {
// 	values := url.Values{}
// 	url, err := url.Parse("http://apis.juhe.cn/simpleWeather/query")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	values.Set("key", "965463c87f88559b922cb7069e7f8c6c")
// 	values.Set("city", "长沙")
// 	url.RawQuery = values.Encode()
// 	urlString := url.String()
// 	r, err := http.Get(urlString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()
// 	b, _ := ioutil.ReadAll(r.Body)
// 	fmt.Printf("b: %v\n", string(b))

// }
// func testParseJson() {
// 	type result struct {
// 		Args   string            `json:"args"`
// 		Header map[string]string `json:"headers"`
// 		Origin string            `json:"origin"`
// 		Url    string            `json:"url"`
// 	}
// 	r, err := http.Get("http://httpbin.org/get")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()
// 	b, _ := ioutil.ReadAll(r.Body)
// 	fmt.Printf("string(b): %v\n", string(b))
// 	var res result
// 	_ = json.Unmarshal(b, &res)
// 	fmt.Printf("res: %v\n", res)
// }
// func testAddHeader() {
// 	type result struct {
// 		Args   string            `json:"args"`
// 		Header map[string]string `json:"headers"`
// 		Origin string            `json:"origin"`
// 		Url    string            `json:"url"`
// 	}
// 	res := result{}
// 	client := &http.Client{}
// 	request, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
// 	request.Header.Add("name", "熊玺皓")
// 	request.Header.Add("age", "20")
// 	response, _ := client.Do(request)
// 	b, _ := ioutil.ReadAll(response.Body)
// 	fmt.Printf("string(b): %v\n", string(b))
// 	json.Unmarshal(b, &res)
// 	fmt.Printf("res: %v\n", res)
// }
// func testPost() {
// 	values := url.Values{}
// 	url := "http://apis.juhe.cn/simpleWeather/query"
// 	values.Set("key", "087d7d10f700d20e27bb753cd806e40b")
// 	values.Set("city", "长沙")
// 	r, err := http.PostForm(url, values)

// 	// values := url.Values{}
// 	// url, err := url.Parse("http://apis.juhe.cn/simpleWeather/query")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// values.Set("key", "087d7d10f700d20e27bb753cd806e40b")
// 	// values.Set("city", "长沙")
// 	// url.RawQuery = values.Encode()
// 	// urlString := url.String()
// 	// r, err := http.PostForm(urlString, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()
// 	b, _ := ioutil.ReadAll(r.Body)
// 	fmt.Printf("b: %v\n", string(b))
// }
// func testPost2() {
// 	url := "http://httpbin.org/post"
// 	values := make(map[string]interface{})
// 	values["name"] = "熊玺皓"
// 	values["age"] = "20"
// 	b, _ := json.Marshal(values)
// 	r, _ := http.Post(url, "application/json", bytes.NewBuffer(b))

// 	// values := url.Values{}
// 	// values.Set("name", "熊玺皓")
// 	// values.Set("age", "20")
// 	// s := values.Encode()
// 	// fmt.Printf("s: %v\n", s)
// 	// r, _ := http.Post(url, "text/html", strings.NewReader(s))

// 	defer r.Body.Close()
// 	b, _ = ioutil.ReadAll(r.Body)
// 	fmt.Printf("string(b): %v\n", string(b))
// }
// func testClient() {
// 	//模拟客户端
// 	values := url.Values{}
// 	values.Set("city", "长沙")
// 	values.Set("key", "087d7d10f700d20e27bb753cd806e40b")
// 	url, _ := url.Parse("http://apis.juhe.cn/simpleWeather/query")
// 	url.RawQuery = values.Encode()
// 	request, _ := http.NewRequest(http.MethodPost, url.String(), nil)
// 	fmt.Printf("request: %v\n", request)
// 	c := http.Client{
// 		Timeout: time.Second * 5,
// 	}
// 	response, _ := c.Do(request)
// 	defer response.Body.Close()
// 	b, _ := ioutil.ReadAll(response.Body)
// 	fmt.Printf("string(b): %v\n", string(b))
// }
// func testHTTPserver() {
// 	f := func(response http.ResponseWriter, request *http.Request) {
// 		io.WriteString(response, "hello world")
// 	}

// 	http.HandleFunc("/hello", f)
// 	err := http.ListenAndServe(":9999", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
// func main() {
// 	testHTTPserver2()
// }

// //使用Handeler实现并发处理:
// type countHandeler struct {
// 	mu sync.Mutex
// 	n  int
// }

// func (h *countHandeler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	h.n++
// 	fmt.Fprintf(w, "count is %v", h.n)
// }
// func testHTTPserver2() {
// 	http.Handle("/count", new(countHandeler))
// 	http.ListenAndServe(":9998", nil)
// }
