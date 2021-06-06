package smodels

import "github.com/kwanifi/numiscan-api/dmodels"

type ProposalVote struct {
	Title       string `json:"title"`
	IsValidator bool   `json:"is_validator"`
	dmodels.ProposalVote
}
