package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/point/point/v9/x/incentives/client/cli"
)

var (
	RegisterIncentiveProposalHandler = govclient.NewProposalHandler(cli.NewRegisterIncentiveProposalCmd)
	CancelIncentiveProposalHandler   = govclient.NewProposalHandler(cli.NewCancelIncentiveProposalCmd)
)
