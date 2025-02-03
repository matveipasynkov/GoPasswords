package account

import (
	"GoPasswords/app/encrypter"
	"GoPasswords/app/output"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func (vault *VaultWithDb) AddAccount(acc Account) {
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

func (vault *VaultWithDb) DeleteURL(url string) bool {
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

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) (*[]Account, error) {
	foundAccounts := make([]Account, 0)
	for _, account := range vault.Accounts {
		if checker(account, str) {
			foundAccounts = append(foundAccounts, account)
		}
	}
	if len(foundAccounts) == 0 {
		return nil, errors.New("DATA_NOT_FOUND")
	}
	return &foundAccounts, nil
}

func (vault *VaultWithDb) FindLogin(login string) (*[]Account, error) {
	foundAccounts := make([]Account, 0)
	for _, account := range vault.Accounts {
		if strings.Contains(account.Login, login) {
			foundAccounts = append(foundAccounts, account)
		}
	}
	if len(foundAccounts) == 0 {
		return nil, errors.New("DATA_NOT_FOUND")
	}
	return &foundAccounts, nil
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.vault")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) Save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError(err.Error())
		return
	}
	encData := vault.enc.Encrypt(data)
	vault.db.Write(encData)
}
