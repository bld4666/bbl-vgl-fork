package netparams

import (
	"fmt"
	"math/big"
	"time"

	"github.com/babylonlabs-io/vigilante/types"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// newHashFromStr converts the passed big-endian hex string into a
// chainhash.Hash.  It only differs from the one available in chainhash in that
// it panics on an error since it will only (and must only) be called with
// hard-coded, and therefore known good, hashes.
func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		// Ordinarily I don't like panics in library code since it
		// can take applications down without them having a chance to
		// recover which is extremely annoying, however an exception is
		// being made in this case because the only way this can panic
		// is if there is an error in the hard-coded hashes.  Thus it
		// will only ever potentially panic on init and therefore is
		// 100% predictable.
		panic(err)
	}
	return hash
}

var (
	testNet3GenesisHash = newHashFromStr("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
	genesisMerkleRoot   = newHashFromStr("4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b")

	genesisCoinbaseTx = wire.MsgTx{
		Version: 1,
		TxIn: []*wire.TxIn{
			{
				PreviousOutPoint: wire.OutPoint{
					Hash:  chainhash.Hash{},
					Index: 0xffffffff,
				},
				SignatureScript: []byte{
					0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
					0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
					0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
					0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
					0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
					0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
					0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
					0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
					0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
					0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
				},
				Sequence: 0xffffffff,
			},
		},
		TxOut: []*wire.TxOut{
			{
				Value: 0x12a05f200,
				PkScript: []byte{
					0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
					0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
					0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
					0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
					0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
					0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
					0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
					0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
					0x1d, 0x5f, 0xac, /* |._.| */
				},
			},
		},
		LockTime: 0,
	}

	testNet3GenesisBlock = wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:    1,
			PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
			MerkleRoot: *genesisMerkleRoot,       // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
			Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
			Bits:       0x1d00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
			Nonce:      0x18aea41a,               // 414098458
		},
		Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
	}
	bigOne           = big.NewInt(1)
	testNet3PowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)
)

func GetBTCParams(net string) (*chaincfg.Params, error) {
	switch net {
	case types.BtcMainnet.String():
		return &chaincfg.MainNetParams, nil
	// case types.BtcTestnet.String():
	// 	return &chaincfg.TestNet3Params, nil
	case types.BtcSimnet.String():
		return &chaincfg.SimNetParams, nil
	case types.BtcRegtest.String():
		return &chaincfg.RegressionNetParams, nil
	case types.BtcTestnet.String():
		res := &chaincfg.Params{
			Name:        "fractaltest",
			Net:         0xdab5bffa,
			DefaultPort: "18333",
			DNSSeeds: []chaincfg.DNSSeed{
				{"dnsseed.fractalbitcoin.io", true},
			},

			// Chain parameters
			GenesisBlock:             &testNet3GenesisBlock,
			GenesisHash:              testNet3GenesisHash,
			PowLimit:                 testNet3PowLimit,
			PowLimitBits:             0x1d00ffff,
			BIP0034Height:            1, // 0000000023b3a96d3484e5abb3755c413e7d41500f8e2a5c3f0dd01299cd8ef8
			BIP0065Height:            1, // 00000000007f6655f22f98e72ed80d8b06dc761d5da09df0fa1dc4be4f861eb6
			BIP0066Height:            1, // 000000002104c8c45e99a8853285a3b592602a3ccde2b832481da85e9e4ba182
			CoinbaseMaturity:         100,
			SubsidyReductionInterval: 210000,
			TargetTimespan:           time.Hour * 24,   // 14 days
			TargetTimePerBlock:       time.Minute * 10, // 10 minutes
			RetargetAdjustmentFactor: 4,                // 25% less, 400% more
			ReduceMinDifficulty:      false,
			MinDiffReductionTime:     time.Minute * 20, // TargetTimePerBlock * 2
			GenerateSupported:        false,

			// Checkpoints ordered from oldest to newest.
			Checkpoints: []chaincfg.Checkpoint{},

			// Consensus rule change deployments.
			//
			// The miner confirmation window is defined as:
			//   target proof of work timespan / target proof of work spacing
			RuleChangeActivationThreshold: 1512, // 75% of MinerConfirmationWindow
			MinerConfirmationWindow:       2016,
			Deployments: [chaincfg.DefinedDeployments]chaincfg.ConsensusDeployment{
				chaincfg.DeploymentTestDummy: {
					BitNumber: 28,
					DeploymentStarter: chaincfg.NewMedianTimeDeploymentStarter(
						time.Unix(1199145601, 0), // January 1, 2008 UTC
					),
					DeploymentEnder: chaincfg.NewMedianTimeDeploymentEnder(
						time.Unix(1230767999, 0), // December 31, 2008 UTC
					),
				},
				chaincfg.DeploymentTestDummyMinActivation: {
					BitNumber:                 22,
					CustomActivationThreshold: 1815,    // Only needs 90% hash rate.
					MinActivationHeight:       10_0000, // Can only activate after height 10k.
					DeploymentStarter: chaincfg.NewMedianTimeDeploymentStarter(
						time.Time{}, // Always available for vote
					),
					DeploymentEnder: chaincfg.NewMedianTimeDeploymentEnder(
						time.Time{}, // Never expires
					),
				},
				chaincfg.DeploymentCSV: {
					BitNumber: 0,
					DeploymentStarter: chaincfg.NewMedianTimeDeploymentStarter(
						time.Unix(1456790400, 0), // March 1st, 2016
					),
					DeploymentEnder: chaincfg.NewMedianTimeDeploymentEnder(
						time.Unix(1493596800, 0), // May 1st, 2017
					),
				},
				chaincfg.DeploymentSegwit: {
					BitNumber: 1,
					DeploymentStarter: chaincfg.NewMedianTimeDeploymentStarter(
						time.Unix(1462060800, 0), // May 1, 2016 UTC
					),
					DeploymentEnder: chaincfg.NewMedianTimeDeploymentEnder(
						time.Unix(1493596800, 0), // May 1, 2017 UTC.
					),
				},
				chaincfg.DeploymentTaproot: {
					BitNumber: 2,
					DeploymentStarter: chaincfg.NewMedianTimeDeploymentStarter(
						time.Unix(1619222400, 0), // April 24th, 2021 UTC.
					),
					DeploymentEnder: chaincfg.NewMedianTimeDeploymentEnder(
						time.Unix(1628640000, 0), // August 11th, 2021 UTC
					),
					CustomActivationThreshold: 1512, // 75%
				},
			},

			// Mempool parameters
			RelayNonStdTxs: true,

			// Human-readable part for Bech32 encoded segwit addresses, as defined in
			// BIP 173.
			Bech32HRPSegwit: "tb", // always tb for test net

			// Address encoding magics
			PubKeyHashAddrID:        0x6f, // starts with m or n
			ScriptHashAddrID:        0xc4, // starts with 2
			WitnessPubKeyHashAddrID: 0x03, // starts with QW
			WitnessScriptHashAddrID: 0x28, // starts with T7n
			PrivateKeyID:            0xef, // starts with 9 (uncompressed) or c (compressed)

			// BIP32 hierarchical deterministic extended key magics
			HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
			HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

			// BIP44 coin type used in the hierarchical deterministic path for
			// address generation.
			HDCoinType: 1,
		}
		return res, nil
	}
	return nil, fmt.Errorf(
		"BTC network with name %s does not exist. should be one of {%s, %s, %s, %s, %s}",
		net,
		types.BtcMainnet.String(),
		types.BtcTestnet.String(),
		types.BtcSimnet.String(),
		types.BtcRegtest.String(),
		types.BtcSignet.String(),
	)
}
