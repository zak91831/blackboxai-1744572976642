package plugins

// Plugin defines the interface that all plugins must implement
type Plugin interface {
    // Name returns the name of the plugin
    Name() string

    // Version returns the version of the plugin
    Version() string

    // Initialize initializes the plugin with the provided configuration
    Initialize(config map[string]interface{}) error

    // Start starts the plugin
    Start() error

    // Stop stops the plugin
    Stop() error

    // Hooks returns a map of hook functions that the plugin provides
    Hooks() map[string]interface{}
}

var registeredPlugins []Plugin

// RegisterPlugin registers a new plugin with the system
func RegisterPlugin(p Plugin) {
    registeredPlugins = append(registeredPlugins, p)
}

// GetPlugins returns all registered plugins
func GetPlugins() []Plugin {
    return registeredPlugins
}
