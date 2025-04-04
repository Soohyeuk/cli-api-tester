package version

import "runtime/debug"

// Version will be set at build time or use the module version
var Version = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}
	return info.Main.Version
}()
