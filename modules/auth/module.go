package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/pylons-tech/bdjuno/database"

	"github.com/forbole/juno/v2/modules"
	"github.com/forbole/juno/v2/modules/messages"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/auth module
type Module struct {
	cdc            codec.Marshaler
	db             *database.Db
	messagesParser messages.MessageAddressesParser
}

// NewModule builds a new Module instance
func NewModule(messagesParser messages.MessageAddressesParser, cdc codec.Marshaler, db *database.Db) *Module {
	return &Module{
		messagesParser: messagesParser,
		cdc:            cdc,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}
