package rpc

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Key struct {
	ID        string    `json:"id"`
	Key       string    `json:"key"`
	Private   string    `json:"private,omitempty"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
}

type KeyResponse struct {
	Key  *Key   `json:"key,omitempty"`
	Keys []*Key `json:"keys,omitempty"`
}

func (a *App) GenerateKey(req *Key, res *KeyResponse) error {
	if len(req.Label) == 0 {
		req.Label = "My Key"
	}
	prv, pub, err := a.service.CreateKey()
	if err != nil {
		return err
	}
	req.ID = uuid.NewString()
	req.Key = pub
	req.Private = prv
	req.CreatedAt = time.Now().UTC()
	err = a.store.Create("keys", req.ID, req)
	if err != nil {
		return err
	}
	// clear private key from response
	req.Private = ""
	res.Key = req
	return nil
}

func (a *App) GetKeys(req *string, res *KeyResponse) error {
	err := a.store.ReadAll("keys", &res.Keys)
	if err != nil {
		return err
	}
	// clear private key
	for i := range res.Keys {
		res.Keys[i].Private = ""
	}
	return nil
}

func (a *App) UpdateKey(req *Key, res *KeyResponse) error {
	if len(req.ID) == 0 {
		return errors.New("missing id")
	}
	model := new(Key)
	err := a.store.Read("keys", req.ID, model)
	if err != nil {
		return err
	}
	if len(req.Label) > 0 {
		model.Label = req.Label
	}
	err = a.store.Update("keys", model.ID, model)
	if err != nil {
		return err
	}

	// clear private key
	model.Private = ""
	res.Key = model
	return nil
}

func (a *App) RemoveKey(req *Key, res *KeyResponse) error {
	if len(req.ID) == 0 {
		return errors.New("missing id")
	}
	model := new(Key)
	err := a.store.Read("keys", req.ID, model)
	if err != nil {
		return err
	}
	err = a.store.Delete("keys", req.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ShowPrivateKey(req *Key, res *KeyResponse) error {
	if len(req.ID) == 0 {
		return errors.New("missing id")
	}
	model := new(Key)
	err := a.store.Read("keys", req.ID, model)
	if err != nil {
		return err
	}
	res.Key = model
	return nil
}

func (a *App) ImportKeys(req *string, res *string) error {
	return errors.New("//TODO")
}
