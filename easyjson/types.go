package hconfig

import (
	"fmt"

	"github.com/uber/tchannel-go/thirdparty/github.com/apache/thrift/lib/go/thrift"
	"gopkg.in/yaml.v2"
)

// Config is the Health Controller config.
type Config struct {
	// Pools is the list of pools to monitor.
	Pools []TargetPool `yaml:"pools,omitempty"`

	CfgVersion `yaml:",inline"`
}

// TargetPool is the config for a single pool.
type TargetPool struct {
	// The key for the pool. This needs to match what TC tells Muttley to use.
	Key string `yaml:"key"`

	// The UNS name for the pool.
	UNS string

	// The health check definition.
	Check CheckDefinition `yaml:",omitempty"`
}

// CheckDefinition is the definition of a health check. The health check will
// be applied to all instances.
type CheckDefinition struct {
	// How often to check each instance. In seconds. Defaults to
	// _defaultInterval. Maximum allowed is _maxInterval.
	Interval *uint32 `yaml:"interval,omitempty" json:"interval,omitempty" `

	// The time after which a request is considered failed. In Seconds. Defaults
	// to _defaultTimeout. Maximum allowed is _maxTimeout.
	Timeout *uint32 `yaml:"timeout,omitempty" json:"timeout,omitempty"`

	// If NamedPort is not specified, it is derived from the protocol type
	NamedPort string `yaml:"namedPort,omitempty" json:"namedPort,omitempty"`

	GRPC     *GRPCDefinition     `yaml:"grpc,omitempty" json:"grpc,omitempty"`
	HTTP     *HTTPDefinition     `yaml:"http,omitempty" json:"http,omitempty"`
	TChannel *TChannelDefinition `yaml:"tchannel,omitempty" json:"tchannel,omitempty"`
	TCP      *TCPDefinition      `yaml:"tcp,omitempty" json:"tcp,omitempty"`
}

// HTTPDefinition is the definition of a HTTP health check.
type HTTPDefinition struct {
	// The path to check. Default to be set to
	// "/health?type=traffic&service=<serviceName>"
	Path *string `yaml:"path,omitempty" json:"path,omitempty"`

	// The value to use for the Host: header. Defaults to the IP of the instance.
	Host *string `yaml:"host,omitempty" json:"host,omitempty"`
}

// TChannelDefinition is the definition of a TChannel health check.
type TChannelDefinition struct {
	// The value to use for Tchannel serviceName.
	ServiceName *string `yaml:"serviceName,omitempty" json:"serviceName,omitempty"`
}

// GRPCDefinition is the definition of a gRPC health check.
type GRPCDefinition struct {
	// The value to use for gRPC serviceName.
	ServiceName *string `yaml:"serviceName,omitempty" json:"serviceName,omitempty"`
}

// TCPDefinition is the definition of a TCP health check.
type TCPDefinition struct{}

// CfgVersion holds the config's version information. This data is generated at
// push time by trafficctl.
type CfgVersion struct {
	// Version is the timestamp when this config was created.
	Version int64 `yaml:",omitempty"`
	// SHA of config - the Git commit SHA this config is associated with.
	SHA string `yaml:",omitempty"`
}

// String returns the complete Health Controller config as a string (in yaml format)
func (c Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("Invalid Config: %v", err)
	}
	return string(data)
}

// SortableTargetPools implements sort.Sort interface
// Sort the TargetPools based on their UNS names.
type SortableTargetPools []TargetPool

func (s SortableTargetPools) Len() int           { return len(s) }
func (s SortableTargetPools) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortableTargetPools) Less(i, j int) bool { return s[i].UNS < s[j].UNS }

// DefaultHTTPCheckPath creates http default check path with provided group name
func DefaultHTTPCheckPath(groupName string) *string {
	return thrift.StringPtr(fmt.Sprintf("/health?type=traffic&service=%s", groupName))
}
