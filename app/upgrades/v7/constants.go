package v7

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v7.0.0"
	// MainnetUpgradeHeight defines the Point mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 2_476_000
	// TODO: TestnetUpgradeHeight defines the Point testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 2_176_500
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/pointnetwork/point-chain/releases/download/v7.0.0/point_7.0.0_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/pointnetwork/point-chain/releases/download/v7.0.0/point_7.0.0_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/pointnetwork/point-chain/releases/download/v7.0.0/point_7.0.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/pointnetwork/point-chain/releases/download/v7.0.0/point_7.0.0_Linux_amd64.tar.gz","windows/x86_64":"https://github.com/pointnetwork/point-chain/releases/download/v7.0.0/point_7.0.0_Windows_x86_64.zip"}}'`

	// FaucetAddressFrom is the inaccessible secp address of the Testnet Faucet
	FaucetAddressFrom = "point1z4ya98ga2xnffn2mhjym7tzlsm49ec23890sze"
	// FaucetAddressTo is the new eth_secp address of the Testnet Faucet
	FaucetAddressTo = "point1ujm4z5v9zkdqm70xnptr027gqu90f7lxjr0fch"

	// ContributorAddrFrom is the lost address of an early contributor
	ContributorAddrFrom = "point1659xwt0hnu5humgek7scefhnpcm2w6hyvy4fsq"
	// ContributorAddrTo is the new address of an early contributor
	ContributorAddrTo = "point1pktlmqrz448cuazl98tqmsj4kjwpqpmaa0cjcf"
)
