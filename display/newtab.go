package display

import (
	"io/ioutil"

	"github.com/makeworld-the-better-one/amfora/config"
)

//nolint
var defaultNewTabContent = `# New Tab

You've opened a new tab. Use the bar at the bottom to browse around. You can start typing in it by pressing the space key.

Press the ? key at any time to bring up the help, and see other keybindings. Most are what you expect.

Happy browsing!

=> //gemini.circumlunar.space What is Gemini?

## Directories

=> //medusae.space medusae.space
=> //geminispace.info geminispace.info search engine
=> //discovery.geminiprotocol.com Gemini Discovery

## Internal Pages

=> about:bookmarks Bookmarks
=> about:subscriptions Subscriptions
=> about:about All internal pages

## Learn more about Amfora!

=> https://github.com/makeworld-the-better-one/amfora Amfora homepage
=> https://github.com/makeworld-the-better-one/amfora/wiki Amfora Wiki [GitHub]
=> gemini://makeworld.space/amfora-wiki/ Amfora Wiki [On Gemini!]
`

// Read the new tab content from a file if it exists or fallback to a default page.
func getNewTabContent() string {
	data, err := ioutil.ReadFile(config.NewTabPath)
	if err == nil {
		return string(data)
	}
	return defaultNewTabContent
}
