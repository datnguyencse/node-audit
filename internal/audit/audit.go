package audit

import (
	"fmt"
	"go-node-audit/config"
	"go-node-audit/pkg/rpc"
	"io/ioutil"
	"net/http"
	"time"

	golog "github.com/ipfs/go-log"
)

const EmptyArrayChecksum = "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"

var log = golog.Logger("Audit")

type Audit struct {
	cfg    *config.Config
	failed int
}

func New(cfg *config.Config) *Audit {
	return &Audit{cfg: cfg}
}

func (audit *Audit) Start() error {
	mavis := rpc.NewRPCClient(rpc.JsonRpcUrl(audit.cfg.MavisRpc))
	eternity := rpc.NewRPCClient(rpc.JsonRpcUrl(audit.cfg.EternityRpc))

	log.Infof("Infinity group id: %d, ronin node id: %d", audit.cfg.InfinityGroupId, audit.cfg.RoninNodeGroupId)
	audit.checkErr("Ronin node monitor bot started", audit.cfg.RoninNodeGroupId)
	for {
		mavisBlock, err := mavis.GetLatestBlock()
		if err != nil {
			continue
		}

		eternityBlock, err := eternity.GetLatestBlock()
		if err != nil {
			audit.checkErr("Failed to reach eternity validator node rpc many time.", audit.cfg.RoninNodeGroupId)
			continue
		}

		if eternityBlock.BlockNumber() < mavisBlock.BlockNumber()-audit.cfg.MaxBlockDelay {
			audit.checkErr(fmt.Sprintf("Eternity non-validator node block %d, skymavis block %d, is delayed: %d blocks", eternityBlock.BlockNumber(), mavisBlock.BlockNumber(), mavisBlock.BlockNumber()-eternityBlock.BlockNumber()), audit.cfg.InfinityGroupId)
		}

		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

func (audit *Audit) checkErr(message string, groupID int) {
	log.Infof("Sending message %s to group %d", message, groupID)
	url := fmt.Sprintf("https://api.telegram.org/bot6599839873:AAF0-cjRn_DLtsAwkTYnI-K2vnCKQhUiVoE/sendMessage?chat_id=-%d&text=@here %s", groupID, message)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
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

	fmt.Println(string(body))
}
