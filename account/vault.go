package account

import (
	"GoPasswords/app/files"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	vault.Save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) DeleteURL(url string) bool {
	changedAccounts := make([]Account, 0)
	for _, account := range vault.Accounts {
		if !strings.Contains(account.Url, url) {
			changedAccounts = append(changedAccounts, account)
		}
	}
	success := (len(vault.Accounts) - len(changedAccounts)) != 0
	vault.Accounts = changedAccounts
	vault.Save()
	return success
}

func (vault *Vault) FindURL(url string) (*[]Account, error) {
	foundAccounts := make([]Account, 0)
	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			foundAccounts = append(foundAccounts, account)
		}
	}
	if len(foundAccounts) == 0 {
		return nil, errors.New("DATA_NOT_FOUND")
	}
	return &foundAccounts, nil
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) Save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
		return
	}
	files.WriteFile(data, "data.json")
}
