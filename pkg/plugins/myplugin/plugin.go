package myplugin

import (
    "fmt"
    "deja_vu/pkg/plugins"
)

// Plugin implements a new feature
type Plugin struct {
    config map[string]interface{}
}

// NewPlugin creates a new plugin
func NewPlugin() plugins.Plugin {
    return &Plugin{}
}

// Name returns the name of the plugin
func (p *Plugin) Name() string {
    return "myplugin"
}

// Version returns the version of the plugin
func (p *Plugin) Version() string {
    return "1.0.0"
}

// Initialize initializes the plugin with the provided configuration
func (p *Plugin) Initialize(config map[string]interface{}) error {
    p.config = config
    return nil
}

// Start starts the plugin
func (p *Plugin) Start() error {
    fmt.Println("[MyPlugin] Plugin started")
    return nil
}

// Stop stops the plugin
func (p *Plugin) Stop() error {
    fmt.Println("[MyPlugin] Plugin stopped")
    return nil
}

// Hooks returns a map of hook functions that the plugin provides
func (p *Plugin) Hooks() map[string]interface{} {
    return map[string]interface{}{
        "my_hook": p.myHookFunction,
    }
}

// myHookFunction is a hook function
func (p *Plugin) myHookFunction(args ...interface{}) {
    fmt.Println("[MyPlugin] Hook function called")
}
