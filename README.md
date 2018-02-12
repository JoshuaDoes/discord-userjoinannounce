# discord-userjoinannounce

### A Discord bot created to announce user join/leave messages from a different server.

----

## Using the official up-to-date version of Discord-UserJoinAnnounce

1. Download the latest Windows/Linux build of Discord-UserJoinAnnounce from the [releases](https://github.com/JoshuaDoes/discord-userjoinannounce/releases) page.
2. Create a bot user and note down the bot token, then invite the bot to the target
server and the announcement server.
3. Enable developer mode on your Discord client and note down the server ID of the
server you'd like to target, the server ID of the server you'd like to announce to,
and the channel ID of the channel you'd like to announce to.
4. Create a file called `config.json` in the same directory as Discord-UserJoinAnnounce.
Use the below configuration template to set it up.
5. Run Discord-UserJoinAnnounce and enjoy!

**Configuration template:**
```JSON
{
	"botToken": "the bot token",
	"targetServer": "the target server ID",
	"announceServer": "the announce server ID",
	"announceChannel": "the announce channel ID",
	"debugMode": true
}
```

## How does it work?

After the bot is invited to your target and announce Discord servers, it will
immediately begin listening for member join/leave events from the target server to
announce them in the announce server's announce channel.

----

## Dependencies

| [configure](https://github.com/paked/configure) |
| [DiscordGo](https://github.com/bwmarrin/discordgo) |

## License
The source code for Discord-UserJoinAnnounce is released under the MIT License. See LICENSE for more details.

## Donations
All donations are appreciated and help pay for the energy drinks that keep me awake long enough to bring you cool new features and projects. Even if it's not much, it helps a lot in the long run!
You can find the donation link here: [Donation Link](https://paypal.me/JoshuaDoes)