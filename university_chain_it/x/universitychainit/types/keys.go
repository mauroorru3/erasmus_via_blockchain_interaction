package types

const (
	// ModuleName defines the module name
	ModuleName = "universitychainit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_universitychainit"

	// Version defines the current version the IBC module supports
	Version = "hub-1"

	// PortID is the default port id that module binds to
	PortID = "universitychainit"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("universitychainit-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	StudentInfoKey = "StudentInfo-value-"
)

const (
	TranscriptOfRecordsKey = "TranscriptOfRecords-value-"
)

const (
	PersonalInfoKey = "PersonalInfo-value-"
)

const (
	ResidenceInfoKey = "ResidenceInfo-value-"
)

const (
	ContactInfoKey = "ContactInfo-value-"
)

const (
	AnnualTaxesKey      = "AnnualTaxes-value-"
	AnnualTaxesCountKey = "AnnualTaxes-count-"
)

const (
	TaxesInfoKey = "TaxesInfo-value-"
)

const (
	ErasmusContributionKey = "ErasmusContribution-value-"
)

const (
	ErasmusCareerKey      = "ErasmusCareer-value-"
	ErasmusCareerCountKey = "ErasmusCareer-count-"
)

const (
	ErasmusInfoKey = "ErasmusInfo-value-"
)

const (
	ChainInfoKey = "ChainInfo-value-"
)
