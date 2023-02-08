package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flowshield/flowshield/fullnode/pkg/contract"
	"github.com/flowshield/flowshield/fullnode/pkg/logger"
	"log"
	"time"

	"github.com/fiatjaf/relayer"
	"github.com/fiatjaf/relayer/storage/postgresql"
	"github.com/kelseyhightower/envconfig"
	"github.com/nbd-wtf/go-nostr"
)

type Relay struct {
	PostgresDatabase string `envconfig:"POSTGRESQL_DATABASE"`

	storage *postgresql.PostgresBackend
}

type FlowShield struct {
	Client   *ethclient.Client
	Instance *contract.Slit
	Auth     *bind.TransactOpts
}
var CS *FlowShield
const (
	FullNode = 1

	Provider = 2
)


func (r *Relay) Name() string {
	return "BasicRelay"
}

func (r *Relay) Storage() relayer.Storage {
	return r.storage
}

func (r *Relay) OnInitialized(*relayer.Server) {}

func (r *Relay) Init() error {
	err := envconfig.Process("", r)
	if err != nil {
		return fmt.Errorf("couldn't process envconfig: %w", err)
	}

	// every hour, delete all very old events
	go func() {
		db := r.Storage().(*postgresql.PostgresBackend)

		for {
			time.Sleep(60 * time.Minute)
			db.DB.Exec(`DELETE FROM event WHERE created_at < $1`, time.Now().AddDate(0, -3, 0).Unix()) // 3 months
		}
	}()

	return nil
}

func (r *Relay) AcceptEvent(evt *nostr.Event) bool {
	// block events that are too large
	jsonb, _ := json.Marshal(evt)
	if len(jsonb) > 10000 {
		return false
	}

	return true
}

func (r *Relay) BeforeSave(evt *nostr.Event) {
	// do nothing
}

func (r *Relay) AfterSave(evt *nostr.Event) {
	// delete all but the 100 most recent ones for each key
	r.Storage().(*postgresql.PostgresBackend).DB.Exec(`DELETE FROM event WHERE pubkey = $1 AND kind = $2 AND created_at < (
      SELECT created_at FROM event WHERE pubkey = $1
      ORDER BY created_at DESC OFFSET 100 LIMIT 1
    )`, evt.PubKey, evt.Kind)
}

func main() {
	//todo 先进行质押注册
	err := InitETH()
	if err != nil {
		log.Fatalf("init eth error: %v", err)
		return
	}

	//原代码
	r := Relay{}
	if err := envconfig.Process("", &r); err != nil {
		log.Fatalf("failed to read from env: %v", err)
		return
	}
	r.storage = &postgresql.PostgresBackend{DatabaseURL: r.PostgresDatabase}
	if err := relayer.Start(&r); err != nil {
		log.Fatalf("server terminated: %v", err)
	}
}

func InitETH() error{
	ctx := context.Background()
	//var cfg *confer.Web3
	//
	//url := cfg.ETH.URL
	//token := cfg.Contract.Token
	//private := cfg.PrivateKey

	url := "https://weechain1.gw106.oneitfarm.com"
	token := "0x056B1B1315304D069D54A4bEAD6eF6E39C7E55fb"
	private := "7259120a1e1f0471d511a14fdb5c619239b267645a356a354e21732a424cc778"

	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	chanID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	contractAdd := common.HexToAddress(token)
	instance, err := contract.NewSlit(contractAdd, client)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(private)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chanID)
	CS = &FlowShield{
		Client:   client,
		Instance: instance,
		Auth:     auth,
	}
	err = CS.stack(ctx)
	if err != nil {
		return err
	}
	return CS.stack(ctx)
}

func (c *FlowShield) stack(ctx context.Context) error {
	logger.Infof("checking if stacked or not...")
	isDeposit, err := c.Instance.IsDeposit(&bind.CallOpts{
		From: c.Auth.From,
	}, FullNode)
	if err != nil {
		return err
	}
	if isDeposit {
		logger.Infof("you have stacked!")
		return nil
	}
	logger.Infof("you have not stacked! trying to stack...")
	// 尝试质押
	tra, err := c.Instance.Stake(c.Auth, FullNode)
	if err != nil {
		return err
	}
	rec, err := bind.WaitMined(ctx, c.Client, tra)
	if err != nil {
		return err
	}
	if rec.Status > 0 {
		logger.Infof("stack succeed !")
		return nil
	}
	return errors.New("sorry,stacked failed")
}