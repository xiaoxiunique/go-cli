
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
	"os/exec"
	"strings"
)

var isEn bool

// tCmd represents the t command
var tCmd = &cobra.Command{
	Use:   "t",
	Short: "command transition",
	Long: "a t word",
	Run: func(cmd *cobra.Command, args []string) {

		credential := common.NewCredential(
			"AKIDud6JFuJRIsBDzUislNinb7ML2TdW8TWI",
			"Xprc9UkqhJemgKglZ72g39yXppZ3GBqE",
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "tmt.tencentcloudapi.com"
		client, _ := tmt.NewClient(credential, "ap-chengdu", cpf)

		request := tmt.NewTextTranslateRequest()

		var source = "zh"
		var target = "en"
		var word = strings.Join(args, " ")

		isEn, _:= cmd.Flags().GetBool("isEn")
		if !isEn {
			source = "en"
			target = "zh"
		}
		request.SourceText = common.StringPtr(word)
		request.Source = common.StringPtr(source)
		request.Target = common.StringPtr(target)
		request.ProjectId = common.Int64Ptr(0)

		response, err := client.TextTranslate(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}
		res := &AutoGenerated{}
		json.Unmarshal([]byte(response.ToJsonString()), res)
		fmt.Printf(word + " -> ")
		color.Green(res.Response.TargetText)
		bash := fmt.Sprintf("echo %s | pbcopy",res.Response.TargetText)
		_, err = exec.Command("bash", "-c", bash).Output()
		if err == nil {
			fmt.Println("~")
		}
	},
}
type AutoGenerated struct {
	Response struct {
		TargetText string `json:"TargetText"`
		Source     string `json:"Source"`
		Target     string `json:"Target"`
		RequestID  string `json:"RequestId"`
	} `json:"Response"`
}

func init() {
	rootCmd.AddCommand(tCmd)
	tCmd.Flags().BoolVarP(&isEn, "isEn", "z", false, "????????? en -> ch")
}
