package rpc

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID        string    `json:"id"`
	Label     string    `json:"label"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}

type ContactResponse struct {
	Contact  *Contact   `json:"contact,omitempty"`
	Contacts []*Contact `json:"contacts,omitempty"`
}

func (a *App) GetContacts(req *string, res *ContactResponse) error {
	err := a.store.ReadAll("contacts", &res.Contacts)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetContact(req *Contact, res *ContactResponse) error {
	if len(req.ID) == 0 {
		return errors.New("missing id")
	}
	model := new(Contact)
	err := a.store.Read("contacts", req.ID, model)
	if err != nil {
		return err
	}
	res.Contact = model
	return nil
}

func (a *App) RemoveContact(req *Contact, res *ContactResponse) error {
	if len(req.ID) == 0 {
		return errors.New("missing id")
	}
	model := new(Contact)
	err := a.store.Read("contacts", req.ID, model)
	if err != nil {
		return err
	}
	return a.store.Delete("contacts", req.ID)
}

func (a *App) UpdateContact(req *Contact, res *ContactResponse) error {
	if len(req.ID) == 0 {
		return errors.New("missing id")
	}
	model := new(Contact)
	err := a.store.Read("contacts", req.ID, model)
	if err != nil {
		return err
	}
	if len(req.Label) > 0 {
		model.Label = req.Label
	}
	if len(req.Key) > 0 {
		model.Key = req.Key
	}
	err = a.store.Update("contacts", model.ID, model)
	if err != nil {
		return err
	}
	res.Contact = model
	return nil
}

func (a *App) AddContact(req *Contact, res *ContactResponse) error {
	if !a.service.IsRecipient(req.Key) {
		return errors.New("not a valid key")
	}
	if len(req.Label) == 0 {
		req.Label = "Default Label"
	}
	req.CreatedAt = time.Now().UTC()
	req.ID = uuid.NewString()
	err := a.store.Create("contacts", req.ID, req)
	if err != nil {
		return err
	}
	res.Contact = req
	return nil
}

func (a *App) ImportContacts(req *string, res *string) error {
	return errors.New("//TODO")
}
