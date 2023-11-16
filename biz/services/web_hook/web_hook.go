package web_hook

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
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
		SshUrl  string `json:"ssh_url"`
	} `json:"repository"`
}

func BuildProject(hook Hook) error {
	if hook.Action == CREATED {
		out, err := exec.Command("rm", "-rf", "./git/"+hook.Repository.Name).CombinedOutput()
		if err != nil {
			return err
		}
		hlog.Info(out)
		privateKeyFile := "/root/.ssh/id_rsa"
		publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
		if err != nil {
			hlog.Error("generate publickeys failed: %s\n", err.Error())
			return err
		}
		_, err = git.PlainClone("./git/"+hook.Repository.Name, false, &git.CloneOptions{
			Auth:       publicKeys,
			URL:        hook.Repository.SshUrl,
			RemoteName: hook.Release.TagName,
		})
		if err != nil {
			return err
		}
		go func() {
			out, err := exec.Command("./build.sh", "./git/"+hook.Repository.Name+"/docker-compose.yaml").CombinedOutput()
			if err != nil {
				hlog.Error(err)
			}
			hlog.Info(string(out))
		}()
	}
	return nil
}
