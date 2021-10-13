package service

import "testing"

func TestKeys(t *testing.T) {
	service := New()

	priv := "AGE-SECRET-KEY-1WNXJUVVRTMCKUJU6NQN9ALGCLFX623G0PKWZRDHHL0A0AXUKM5MSJWTW2Y"
	pub := "age1vl5csek7565t2pugr930n5a9ma26wrer0egc4vsfutg9ehl0pcfqu9x25m"

	identity := service.IsIdentity(priv)
	if !identity {
		t.Errorf("%s should be our identity, but it says it's not!\n", priv)
		return
	}

	identity = service.IsIdentity("bla")
	if identity {
		t.Error("bla should fail as identity but it passes!")
		return
	}

	recp := service.IsRecipient(pub)
	if !recp {
		t.Errorf("%s should be a recipient but it's not!\n", pub)
		return
	}

	recp = service.IsIdentity("bla")
	if recp {
		t.Error("bla should fail as recipient but is passes!")
		return
	}
}
