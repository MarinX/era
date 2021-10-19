package ui

import "github.com/MarinX/era/rpc"

func (b *GoBridge) Contacts() *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.ContactResponse),
	}

	cnts := new(rpc.ContactResponse)
	err := b.cli.GetContacts(nil, cnts)
	if err != nil {
		resp.Error = err.Error()
	} else {
		sortContactsAsc(cnts.Contacts)
		resp.Data = cnts
	}

	return resp
}

func (b *GoBridge) CreateContact(label, key string) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.ContactResponse),
	}
	err := b.cli.AddContact(&rpc.Contact{Label: label, Key: key}, resp.Data.(*rpc.ContactResponse))
	if err != nil {
		resp.Error = err.Error()
	}

	return resp
}

func (b *GoBridge) RemoveContact(id string) *ResponseBinder {
	resp := &ResponseBinder{}
	err := b.cli.RemoveContact(&rpc.Contact{ID: id}, nil)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) UpdateContact(id, label string) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.ContactResponse),
	}
	err := b.cli.UpdateContact(&rpc.Contact{ID: id, Label: label}, resp.Data.(*rpc.ContactResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}
