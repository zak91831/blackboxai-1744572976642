package myplugin

import (
    "deja_vu/pkg/plugins"
)

// Register registers the plugin with the Deja Vu system
func Register() {
    plugins.RegisterPlugin(NewPlugin())
}
