package cash

import (
	"github.com/solapi/solapi-go/apirequest"
	"github.com/solapi/solapi-go/types"
)

// Cash struct
type Cash struct {
	Config       map[string]string
	ApiRequester apirequest.ApiRequester
}

// Balance get balance information
func (r *Cash) Balance() (types.Balance, error) {
	request := r.ApiRequester.NewAPIRequest()
	result := types.Balance{}
	params := map[string]string{}
	err := request.GET("cash/v1/balance", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
