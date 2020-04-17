package evt_link

import "github.com/fsn-dev/cryptoCoins/internal/evt/evtapi/client"

type GetTrxIdForLinkIdRequest struct {
	LinkId string `json:"link_id"`
}

type GetTrxIdForLinkIdResult struct {
	TrxId    string `json:"trx_id"`
	BlockNum string `json:"block_num"`
}

func (it *Instance) GetTrxIdForLinkId(linkId string) (*GetTrxIdForLinkIdResult, *client.ApiError) {
	response := &GetTrxIdForLinkIdResult{}

	err := it.client.Post(it.Path("get_trx_id_for_link_id"), &GetTrxIdForLinkIdRequest{linkId}, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
