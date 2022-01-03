package utils

import (
	"context"

	"github.com/pylons-tech/juno/client"

	"github.com/pylons-tech/bdjuno/types"

	"github.com/cosmos/cosmos-sdk/types/query"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func GetSigningInfos(height int64, slashingClient slashingtypes.QueryClient) ([]types.ValidatorSigningInfo, error) {
	var signingInfos []slashingtypes.ValidatorSigningInfo

	header := client.GetHeightRequestHeader(height)

	var nextKey []byte
	var stop = false
	for !stop {
		res, err := slashingClient.SigningInfos(
			context.Background(),
			&slashingtypes.QuerySigningInfosRequest{
				Pagination: &query.PageRequest{
					Key:   nextKey,
					Limit: 1000, // Query 1000 signing infos at a time
				},
			},
			header,
		)
		if err != nil {
			return nil, err
		}

		nextKey = res.Pagination.NextKey
		stop = len(res.Pagination.NextKey) == 0
		signingInfos = append(signingInfos, res.Info...)
	}

	infos := make([]types.ValidatorSigningInfo, len(signingInfos))
	for index, info := range signingInfos {
		infos[index] = types.NewValidatorSigningInfo(
			info.Address,
			info.StartHeight,
			info.IndexOffset,
			info.JailedUntil,
			info.Tombstoned,
			info.MissedBlocksCounter,
			height,
		)
	}
	return infos, nil
}
