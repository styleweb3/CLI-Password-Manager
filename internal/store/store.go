package store

import (
	"encoding/json"
	"os"

	"github.com/styleweb3/CLI-pwd-manager/models"
)

const vaultPath = "vault.json"

func Load() ([]models.PasswordEntry, error) {
	if _, err := os.Stat(vaultPath); os.IsNotExist(err) {
		return []models.PasswordEntry{}, nil
	}

	data, err := os.ReadFile(vaultPath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []models.PasswordEntry{}, nil
	}

	var entries []models.PasswordEntry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func Save(entries []models.PasswordEntry) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(vaultPath, data, 0644)
}
