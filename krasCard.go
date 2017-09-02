package krasCard

import (
	"krasCard/common"
	"krasCard/transportCard"
)

func GetTransportCard(id string) (*transportCard.TransportCard, error) {
	csrf, err := common.GetCSRF()
	if err != nil {
		return nil, err
	}
	resp, err := transportCard.GetInfoPage(csrf, id)
	if err != nil {
		return nil, err
	}
	info, err := transportCard.ParseInfo(resp)
	if err != nil {
		return nil, err
	}

	return info, nil
}
