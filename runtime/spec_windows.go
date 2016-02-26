package runtime

// Temporary Windows version of the spec in lieu of opencontainers/specs having
// Windows support currently.
import "github.com/docker/containerd/specs"

type platformSpec specs.WindowsSpec
type processSpec specs.Process
