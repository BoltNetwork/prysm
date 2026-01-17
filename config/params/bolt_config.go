package params

import (
	"math"
	"time"

	"github.com/OffchainLabs/prysm/v7/consensus-types/primitives"
	"github.com/OffchainLabs/prysm/v7/encoding/bytesutil"
)

// BoltName is the name of the Bolt network config
const BoltName = "bolt-testnet"

// UseBoltNetworkConfig uses the Bolt beacon chain specific network config.
func UseBoltNetworkConfig() {
	OverrideBeaconNetworkConfig(boltNetworkConfig)
}

// BoltConfig returns the configuration for Bolt network
func BoltConfig() *BeaconChainConfig {
	if boltBeaconConfig.ForkVersionSchedule == nil {
		boltBeaconConfig.InitializeForkSchedule()
	}
	return boltBeaconConfig.Copy()
}

var boltNetworkConfig = &NetworkConfig{
	AttSubnetKey:               "attnets",
	SyncCommsSubnetKey:         "syncnets",
	CustodyGroupCountKey:       "cgc",
	MinimumPeersInSubnetSearch: 20,
	ContractDeploymentBlock:    0,
	BootstrapNodes: []string{
		// Bolt Network Testnet bootnode
		"enr:-Nm4QMu6W0tgXGw5IXK3xmj3fo6rTo3vOm89oyGWuG3e5Tg9ILzmCsGQHq51-MRzDEKt2FAjtD-w1z2mRtI239F1pPOGAZu5viXkh2F0dG5ldHOI__________-DY2djgYCEZXRoMpBBpfhXcAAAAP__________gmlkgnY0gmlwhNUgGaqDbmZkhAAAAACEcXVpY4IyyIlzZWNwMjU2azGhAhcPxXt6za3A_8WQuTdsFiuIA_CrEJ7lSBn7UwT1NkGwiHN5bmNuZXRzD4N0Y3CCMsiDdWRwgi7g",
	},
}

var boltBeaconConfig = &BeaconChainConfig{
	// Constants (Non-configurable)
	FarFutureEpoch:           math.MaxUint64,
	FarFutureSlot:            math.MaxUint64,
	BaseRewardsPerEpoch:      4,
	DepositContractTreeDepth: 32,
	GenesisDelay:             0,

	// Misc constant.
	TargetCommitteeSize:            128,
	MaxValidatorsPerCommittee:      2048,
	MaxCommitteesPerSlot:           64,
	MinPerEpochChurnLimit:          4,
	ChurnLimitQuotient:             64,
	ShuffleRoundCount:              90,
	MinGenesisActiveValidatorCount: 64,
	MinGenesisTime:                 1768331470,
	TargetAggregatorsPerCommittee:  16,
	HysteresisQuotient:             4,
	HysteresisDownwardMultiplier:   1,
	HysteresisUpwardMultiplier:     5,

	// Gwei value constants.
	MinDepositAmount:          1 * 1e9,
	MaxEffectiveBalance:       32 * 1e9,
	EjectionBalance:           16 * 1e9,
	EffectiveBalanceIncrement: 1 * 1e9,

	// Initial value constants.
	BLSWithdrawalPrefixByte:         byte(0),
	ETH1AddressWithdrawalPrefixByte: byte(1),
	CompoundingWithdrawalPrefixByte: byte(2),
	BuilderWithdrawalPrefixByte:     byte(3),
	BuilderIndexSelfBuild:           primitives.BuilderIndex(math.MaxUint64),
	ZeroHash:                        [32]byte{},

	// Time parameter constants - 5 SECOND SLOTS
	MinAttestationInclusionDelay:     1,
	SecondsPerSlot:                   5,
	SlotDurationMilliseconds:         5000,
	SlotsPerEpoch:                    32,
	SqrRootSlotsPerEpoch:             5,
	MinSeedLookahead:                 1,
	MaxSeedLookahead:                 4,
	EpochsPerEth1VotingPeriod:        64,
	SlotsPerHistoricalRoot:           8192,
	MinValidatorWithdrawabilityDelay: 256,
	ShardCommitteePeriod:             256,
	MinEpochsToInactivityPenalty:     4,
	Eth1FollowDistance:               2048,

	// Fork choice algorithm constants.
	ProposerScoreBoost:              40,
	ReorgHeadWeightThreshold:        20,
	ReorgParentWeightThreshold:      160,
	ReorgMaxEpochsSinceFinalization: 2,
	IntervalsPerSlot:                3,

	// Time-based protocol parameters.
	ProposerReorgCutoffBPS: primitives.BP(1667),
	AttestationDueBPS:      primitives.BP(3333),
	AggregrateDueBPS:       primitives.BP(6667),
	SyncMessageDueBPS:      primitives.BP(3333),
	ContributionDueBPS:     primitives.BP(6667),

	// Bolt Chain ID
	DepositChainID:         1337,
	DepositNetworkID:       1337,
	DepositContractAddress: "0x00000000219ab540356cBB839Cbe05303d7705Fa",

	// Validator params.
	RandomSubnetsPerValidator:         1 << 0,
	EpochsPerRandomSubnetSubscription: 1 << 8,

	SecondsPerETH1Block: 5,

	// State list length constants.
	EpochsPerHistoricalVector: 65536,
	EpochsPerSlashingsVector:  8192,
	HistoricalRootsLimit:      16777216,
	ValidatorRegistryLimit:    1099511627776,

	// Reward and penalty quotients constants.
	BaseRewardFactor:               64,
	WhistleBlowerRewardQuotient:    512,
	ProposerRewardQuotient:         8,
	InactivityPenaltyQuotient:      67108864,
	MinSlashingPenaltyQuotient:     128,
	ProportionalSlashingMultiplier: 1,

	// Max operations per block constants.
	MaxProposerSlashings:             16,
	MaxAttesterSlashings:             2,
	MaxAttesterSlashingsElectra:      1,
	MaxAttestations:                  128,
	MaxAttestationsElectra:           8,
	MaxDeposits:                      16,
	MaxVoluntaryExits:                16,
	MaxWithdrawalsPerPayload:         16,
	MaxBlsToExecutionChanges:         16,
	MaxValidatorsPerWithdrawalsSweep: 16384,

	// BLS domain values.
	DomainBeaconProposer:              bytesutil.Uint32ToBytes4(0x00000000),
	DomainBeaconAttester:              bytesutil.Uint32ToBytes4(0x01000000),
	DomainRandao:                      bytesutil.Uint32ToBytes4(0x02000000),
	DomainDeposit:                     bytesutil.Uint32ToBytes4(0x03000000),
	DomainVoluntaryExit:               bytesutil.Uint32ToBytes4(0x04000000),
	DomainSelectionProof:              bytesutil.Uint32ToBytes4(0x05000000),
	DomainAggregateAndProof:           bytesutil.Uint32ToBytes4(0x06000000),
	DomainSyncCommittee:               bytesutil.Uint32ToBytes4(0x07000000),
	DomainSyncCommitteeSelectionProof: bytesutil.Uint32ToBytes4(0x08000000),
	DomainContributionAndProof:        bytesutil.Uint32ToBytes4(0x09000000),
	DomainApplicationMask:             bytesutil.Uint32ToBytes4(0x00000001),
	DomainApplicationBuilder:          bytesutil.Uint32ToBytes4(0x00000001),
	DomainBLSToExecutionChange:        bytesutil.Uint32ToBytes4(0x0A000000),
	DomainBeaconBuilder:               bytesutil.Uint32ToBytes4(0x0B000000),

	// Prysm constants.
	GweiPerEth:                 1000000000,
	BLSSecretKeyLength:         32,
	BLSPubkeyLength:            48,
	DefaultBufferSize:          10000,
	WithdrawalPrivkeyFileName:  "/shardwithdrawalkey",
	ValidatorPrivkeyFileName:   "/validatorprivatekey",
	RPCSyncCheck:               1,
	EmptySignature:             [96]byte{},
	DefaultPageSize:            250,
	MaxPeersToSync:             15,
	SlotsPerArchivedPoint:      2048,
	GenesisCountdownInterval:   time.Minute,
	ConfigName:                 BoltName,
	PresetBase:                 "mainnet",
	BeaconStateFieldCount:          21,
	BeaconStateAltairFieldCount:    24,
	BeaconStateBellatrixFieldCount: 25,
	BeaconStateCapellaFieldCount:   28,
	BeaconStateDenebFieldCount:     28,
	BeaconStateElectraFieldCount:   37,
	BeaconStateFuluFieldCount:      38,

	// Slasher related values.
	WeakSubjectivityPeriod:          54000,
	PruneSlasherStoragePeriod:       10,
	SlashingProtectionPruningEpochs: 512,

	// Weak subjectivity values.
	SafetyDecay: 10,

	// ALL FORKS AT EPOCH 0
	GenesisEpoch:         0,
	GenesisForkVersion:   []byte{0x10, 0x00, 0x00, 0x00},
	AltairForkVersion:    []byte{0x20, 0x00, 0x00, 0x00},
	AltairForkEpoch:      0,
	BellatrixForkVersion: []byte{0x30, 0x00, 0x00, 0x00},
	BellatrixForkEpoch:   0,
	CapellaForkVersion:   []byte{0x40, 0x00, 0x00, 0x00},
	CapellaForkEpoch:     0,
	DenebForkVersion:     []byte{0x50, 0x00, 0x00, 0x00},
	DenebForkEpoch:       0,
	ElectraForkVersion:   []byte{0x60, 0x00, 0x00, 0x00},
	ElectraForkEpoch:     0,
	FuluForkVersion:      []byte{0x70, 0x00, 0x00, 0x00},
	FuluForkEpoch:        0,

	// Participation flag indices.
	TimelySourceFlagIndex: 0,
	TimelyTargetFlagIndex: 1,
	TimelyHeadFlagIndex:   2,

	// Incentivization weight values.
	TimelySourceWeight: 14,
	TimelyTargetWeight: 26,
	TimelyHeadWeight:   14,
	SyncRewardWeight:   2,
	ProposerWeight:     8,
	WeightDenominator:  64,

	// Validator related values.
	TargetAggregatorsPerSyncSubcommittee: 16,
	SyncCommitteeSubnetCount:             4,

	// Misc values.
	SyncCommitteeSize:            512,
	InactivityScoreBias:          4,
	InactivityScoreRecoveryRate:  16,
	EpochsPerSyncCommitteePeriod: 256,

	// Updated penalty values.
	InactivityPenaltyQuotientAltair:         50331648,
	MinSlashingPenaltyQuotientAltair:        64,
	ProportionalSlashingMultiplierAltair:    2,
	MinSlashingPenaltyQuotientBellatrix:     32,
	ProportionalSlashingMultiplierBellatrix: 3,
	InactivityPenaltyQuotientBellatrix:      16777216,

	// Light client
	MinSyncCommitteeParticipants: 1,
	MaxRequestLightClientUpdates: 128,

	// Bellatrix
	TerminalBlockHashActivationEpoch: 18446744073709551615,
	TerminalBlockHash:                [32]byte{},
	TerminalTotalDifficulty:          "0",
	MaxBytesPerTransaction:           1073741824,
	MaxTransactionsPerPayload:        1048576,
	BytesPerLogsBloom:                256,
	MaxExtraDataBytes:                32,
	EthBurnAddressHex:                "0x0000000000000000000000000000000000000000",
	DefaultBuilderGasLimit:           uint64(60000000),

	// Mevboost circuit breaker
	MaxBuilderConsecutiveMissedSlots: 3,
	MaxBuilderEpochMissedSlots:       5,
	ExecutionEngineTimeoutValue:      8,

	// Subnet value
	BlobsidecarSubnetCount:        6,
	BlobsidecarSubnetCountElectra: 9,

	MaxPerEpochActivationChurnLimit:  128,
	MinEpochsForBlobsSidecarsRequest: 4096,
	MaxRequestBlobSidecars:           768,
	MaxRequestBlocksDeneb:            128,
	FieldElementsPerBlob:             4096,
	MaxBlobCommitmentsPerBlock:       4096,
	KzgCommitmentInclusionProofDepth: 17,
	DeprecatedMaxBlobsPerBlock:       6,

	// Electra values
	MinPerEpochChurnLimitElectra:          128_000_000_000,
	MaxPerEpochActivationExitChurnLimit:   256_000_000_000,
	MaxEffectiveBalanceElectra:            2048_000_000_000,
	MinSlashingPenaltyQuotientElectra:     4096,
	WhistleBlowerRewardQuotientElectra:    4096,
	PendingDepositsLimit:                  134_217_728,
	PendingPartialWithdrawalsLimit:        134_217_728,
	PendingConsolidationsLimit:            262_144,
	MinActivationBalance:                  32_000_000_000,
	MaxConsolidationsRequestsPerPayload:   2,
	MaxPendingPartialsPerWithdrawalsSweep: 8,
	MaxPendingDepositsPerEpoch:            16,
	FullExitRequestAmount:                 0,
	MaxWithdrawalRequestsPerPayload:       16,
	MaxDepositRequestsPerPayload:          8192,
	UnsetDepositRequestsStartIndex:        math.MaxUint64,
	DeprecatedMaxBlobsPerBlockElectra:     9,
	DeprecatedTargetBlobsPerBlockElectra:  6,
	MaxRequestBlobSidecarsElectra:         1152,

	// Fulu values
	MaxRequestDataColumnSidecars:          16384,
	DataColumnSidecarSubnetCount:          128,
	SamplesPerSlot:                        8,
	NumberOfCustodyGroups:                 128,
	CustodyRequirement:                    4,
	MinEpochsForDataColumnSidecarsRequest: 4096,
	ValidatorCustodyRequirement:           8,
	BalancePerAdditionalCustodyGroup:      32_000_000_000,

	// Networking parameters.
	MaxPayloadSize:                  10 * 1 << 20,
	AttestationSubnetCount:          64,
	AttestationPropagationSlotRange: 32,
	MaxRequestBlocks:                1 << 10,
	TtfbTimeout:                     5,
	RespTimeout:                     10,
	MaximumGossipClockDisparity:     500,
	MessageDomainInvalidSnappy:      [4]byte{00, 00, 00, 00},
	MessageDomainValidSnappy:        [4]byte{01, 00, 00, 00},
	MinEpochsForBlockRequests:       33024,
	EpochsPerSubnetSubscription:     256,
	AttestationSubnetExtraBits:      0,
	AttestationSubnetPrefixBits:     6,
	SubnetsPerNode:                  2,
	NodeIdBits:                      256,

	BlobSchedule: []BlobScheduleEntry{},
}