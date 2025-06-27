package main

import (
	// Modules of packwiz
	"github.com/AvalonServers/packwiz/cmd"
	_ "github.com/AvalonServers/packwiz/curseforge"
	_ "github.com/AvalonServers/packwiz/github"
	_ "github.com/AvalonServers/packwiz/migrate"
	_ "github.com/AvalonServers/packwiz/modrinth"
	_ "github.com/AvalonServers/packwiz/settings"
	_ "github.com/AvalonServers/packwiz/url"
	_ "github.com/AvalonServers/packwiz/utils"
)

func main() {
	cmd.Execute()
}
