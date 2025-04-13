package main

import (
    "fmt"
    "os"
    "deja_vu/pkg/plugins/myplugin"
    "github.com/urfave/cli/v2"
)

func main() {
    // Register the new plugin
    myplugin.Register()

    // Initialize CLI application
    app := initializeApp()
    
    // Run the application
    if err := app.Run(os.Args); err != nil {
        fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
        os.Exit(1)
    }
}

func initializeApp() *cli.App {
    app := &cli.App{
        Name:  "deja_vu",
        Usage: "A vulnerability scanner with time-travel capabilities",
        Commands: []*cli.Command{
            {
                Name:    "scan",
                Aliases: []string{"s"},
                Usage:   "Start a new scan",
                Action: func(c *cli.Context) error {
                    fmt.Println("Starting new scan...")
                    return nil
                },
            },
            {
                Name:    "plugin",
                Aliases: []string{"p"},
                Usage:   "Plugin management commands",
                Subcommands: []*cli.Command{
                    {
                        Name:  "list",
                        Usage: "List all installed plugins",
                        Action: func(c *cli.Context) error {
                            fmt.Println("Installed plugins:")
                            fmt.Printf("- %s (v%s)\n", myplugin.NewPlugin().Name(), myplugin.NewPlugin().Version())
                            return nil
                        },
                    },
                    {
                        Name:  "start",
                        Usage: "Start a specific plugin",
                        Action: func(c *cli.Context) error {
                            plugin := myplugin.NewPlugin()
                            return plugin.Start()
                        },
                    },
                    {
                        Name:  "stop",
                        Usage: "Stop a specific plugin",
                        Action: func(c *cli.Context) error {
                            plugin := myplugin.NewPlugin()
                            return plugin.Stop()
                        },
                    },
                },
            },
        },
    }
    return app
}
