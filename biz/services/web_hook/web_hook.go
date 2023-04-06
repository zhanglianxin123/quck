package web_hook

import (
	"github.com/bytedance/gopkg/util/logger"
	git "github.com/go-git/go-git/v5"
	"os/exec"
)

const RELEASED = "released"
const CREATED = "created"

type Hook struct {
	Action  string `json:"action"`
	Release struct {
		Url       string `json:"url"`
		AssetsUrl string `json:"assets_url"`
		UploadUrl string `json:"upload_url"`
		HtmlUrl   string `json:"html_url"`
		Id        int    `json:"id"`
		Author    struct {
			Url     string `json:"url"`
			HtmlUrl string `json:"html_url"`
		} `json:"author"`
		NodeId          string `json:"node_id"`
		TagName         string `json:"tag_name"`
		TargetCommitish string `json:"target_commitish"`
		Name            string `json:"name"`
		Body            string `json:"body"`
	} `json:"release"`
	Repository struct {
		Name    string `json:"name"`
		HtmlUrl string `json:"html_url"`
		Url     string `json:"url"`
	} `json:"repository"`
}

func BuildProject(hook Hook) error {
	if hook.Action == CREATED {
		out, err := exec.Command("rm", "-rf", "./git/"+hook.Repository.Name).CombinedOutput()
		if err != nil {
			return err
		}
		logger.Info(out)
		_, err = git.PlainClone("./git/"+hook.Repository.Name, false, &git.CloneOptions{
			URL:        hook.Repository.HtmlUrl,
			RemoteName: hook.Release.TagName,
		})
		//cloneExec := fmt.Sprintf("git clone -b %s %s", hook.Release.TagName, hook.Repository.HtmlUrl)
		//out, err := exec.Command(cloneExec).CombinedOutput()
		if err != nil {
			return err
		}
		go func() {
			out, err := exec.Command("./build.sh", "./git/"+hook.Repository.Name+"/docker-compose.yaml").CombinedOutput()
			if err != nil {
				logger.Error(err)
			}
			logger.Info(string(out))
		}()
		//logger.Info(fmt.Sprintf("output:%s", response))
	}
	return nil
}
