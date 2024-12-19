package utils

type VarEnv struct {
	HOSTNAME  string
	FOR_RANGE string
	FOR_PORTS string
	VERBOSE   string
}

// start and end include
type PortRange struct {
	Start int
	End   int
}

type Bag struct {
	PortRange PortRange
	VarEnv    VarEnv
	PortList  []int
	IsRange   bool
	IsList    bool
}
