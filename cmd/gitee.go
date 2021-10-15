package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// giteeCmd represents the gitee command
var giteeCmd = &cobra.Command{
	Use:   "gitee",
	Short: "提供 Github 地址，转化为 gitee 的地址",
	Long: "a gitee <github url>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("not found params")
			return
		}

		url := fmt.Sprintf("https://gitee.com/projects/check_project_duplicate?import_url=%s", args[0])
		method := "GET"

		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
		req.Header.Add("Accept", "*/*")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
		req.Header.Add("X-Requested-With", "XMLHttpRequest")
		req.Header.Add("Sec-Fetch-Site", "same-origin")
		req.Header.Add("Sec-Fetch-Mode", "cors")
		req.Header.Add("Sec-Fetch-Dest", "empty")
		req.Header.Add("Referer", "https://gitee.com/projects/import/url")
		req.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Add("Cookie", "gitee-session-n=cWY3U0VYTzVIRjE0RW5vNkp1cWFjU3VmS290ZWNSVFk1ZkcreHhmc1RDRVJEUzg4VGVkVTUya09IaVgzTWFvRE1vaEVwdDgzc1V2dFp0T3c2QXZubFlqQWJLMGNYRnE1bVFDMEM1SU1CTlRRNXBCTWVIUG1GNGQyZ1RlRmZMYzJ3a2xQcUxjY0xSYUJmY3hOR1g4Zi9NdkxER0ZFT2R0d3pqeDVEZ3FhWUlaYUp0elllWFp3am45RjQzeGc5VXplSnAzTVVKLzN6REpGZ1AyalNTZGE4TDE0QUkzb1B2NXJHUkIyNUFabkhtdDl1VnFCdmhVSUNiWjE2QnhQVFRBKzZrdDJERUhJK0g5aE5QdHppdHRpemZMblBxTWlvQ0RrYlQyaVRBeHRKMXdxSG9hRUNtd1R0QjR5NkNrN0ZYWU4rSkh6QTREc0ZRUGdvdHhoMDB3KzZyTWhWanQ0bitudVZGZTh5SVRyOHpjPS0tUEQ5WXZjZUdHU3N1bDgwR25Hckp6UT09--5b1c34cc708091c8e11107f507788e6b526ba1a6; gitee_user=true; oschina_new_user=false; user_locale=zh-CN")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		compile := regexp.MustCompile("<a href=\"(.*)\" ")
		submatch := compile.FindStringSubmatch(string(body))
		if len(submatch) == 2 {
			fmt.Println(fmt.Sprintf("%s.git", submatch[1]))
		} else {
			log.Printf("没找到仓库")
		}
	},
}

func init() {
	rootCmd.AddCommand(giteeCmd)
}
