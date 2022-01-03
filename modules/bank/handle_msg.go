package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/pylons-tech/juno/modules/messages"
	"github.com/rs/zerolog/log"

	"github.com/pylons-tech/bdjuno/database"
	bankutils "github.com/pylons-tech/bdjuno/modules/bank/utils"
	"github.com/pylons-tech/bdjuno/modules/utils"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	juno "github.com/pylons-tech/juno/types"
)

// HandleMsg handles any message updating the involved addresses balances
func HandleMsg(
	tx *juno.Tx, msg sdk.Msg, getAddresses messages.MessageAddressesParser, bankClient banktypes.QueryClient,
	cdc codec.Codec, db *database.Db,
) error {
	addresses, err := getAddresses(cdc, msg)
	if err != nil {
		log.Error().Str("module", "bank").Str("operation", "refresh balances").
			Err(err).Msgf("error while refreshing balances after message of type %s", proto.MessageName(msg))
	}

	return bankutils.RefreshBalances(tx.Height, utils.FilterNonAccountAddresses(addresses), bankClient, db)
}
