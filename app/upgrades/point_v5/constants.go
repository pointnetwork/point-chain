package point_v5

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v0.0.5"
	// MainnetUpgradeHeight defines the Point mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 18450331 //18440331 + 2500*4
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{
		"darwin/arm64":"https://github.com/pointnetwork/point-chain/releases/download/v0.0.5/point-chain_0.0.5_Darwin_arm64.tar.gz",
		"darwin/x86_64":"https://github.com/pointnetwork/point-chain/releases/download/v0.0.5/point-chain_0.0.5_Darwin_x86_64.tar.gz",
		"linux/arm64":"https://github.com/pointnetwork/point-chain/releases/download/v0.0.5/point-chain_0.0.5_Linux_arm64.tar.gz",
		"linux/x86_64":"https://github.com/pointnetwork/point-chain/releases/download/v0.0.5/point-chain_0.0.5_Linux_x86_64.tar.gz",
		"windows/x86_64":"https://github.com/pointnetwork/point-chain/releases/download/v0.0.5/point-chain_0.0.5_Windows_x86_64.zip"
	}}'`
)
