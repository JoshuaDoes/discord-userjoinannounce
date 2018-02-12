# discord-userjoinannounce

This bot was created with the intention to use it for FiveM communities that prefer using Discord over other chat programs. This allows for the community to have a "babysitter" in their interview or public server to let you know when someone new joins without actually needing to sit in it.

----------------------------------------------------------------

**In these directions:**
- target = The server you want the bot to watch
- announcement = The server you want the bot to announce to

----

# Using the Discord-UserJoinAnnounce

**1. Bot User**
Create a bot user and note down the bot token, then invite the bot to the target
server and the announcement server.

------------------------------------------------------------------

**2. Create the bot**

Creating a bot in discord is stupid easy. First you need to go to [here](https://discordapp.com/developers/applications/me) and click "New Application"
![Application Screen](http://i.imgur.com/FvgfY2Z.png)

Now give your bot a name and a picture, a description isn't necessary.
![New Application Screen](http://i.imgur.com/MOS7yvH.png)

Click "Create Application". On the next page scroll down until you see "Create a bot user" and click that. Also click "Yes, do it!".
![Screen you see after creating a new application then scrolling down a little.](http://i.imgur.com/YAzK5ml.png)
![Yes Do It.](http://i.imgur.com/vkF6Rxo.png)

Now you can get your bot's token by using the "click to reveal" button in the App Bot User section.
![New Bot Page](http://i.imgur.com/xhKMUVU.png)
![Token](http://i.imgur.com/QwCmJJM.png)

There's your token! Now it's time to invite your bot to your server. Don't worry about the bot being up and running for this next step. 
Replace `YOUR_CLIENT_ID_HERE` in this URL ```https://discordapp.com/oauth2/authorize?&client_id=YOUR_CLIENT_ID_HERE&scope=bot&permissions=0``` 
with the Client ID from the page above under App Details, then paste the link into your browser. It will give you a page that looks like this:
![Authorize Bot](http://i.imgur.com/Ggwy0BP.png)

Now select your server in the dropdown, then click "Authorize".
![Authorized](http://i.imgur.com/4cqNcs1.png)

That's it! Now you can start your bot

-------------------------------------------------------------------------


**3. Enable developer mode on your Discord client:**
- Click the cog in the bottom left, near the speaker icon
- On the left side of the new window, click "Appearance"
- Scroll down to the "Advanced" section and turn on "Developer Mode"
- This may require you to restart Discord

**4. Download** 
Get the latest Windows/Linux build of Discord-UserJoinAnnounce from the [releases](https://github.com/JoshuaDoes/discord-userjoinannounce/releases) page.

**5. Configuration:**
- Create a file called `config.json` in the same directory as Discord-UserJoinAnnounce bot.
- Use the below configuration template to set it up.

**Configuration template:**
```JSON
{
	"botToken": "the bot token",
	"targetServer": "the target server ID",
	"announceServer": "the announce server ID",
	"announceChannel": "the announce channel ID",
	"debugMode": false
}
```

**6. Getting your ID's:**
- Do step 3 above
- Right click your target server and select "Copy ID" at the bottom and paste in the config file
- Right click your announcement server and select "Copy ID" at the bottom and paste in the config file
- Right click the text channel you want the bot to make the announcement in and paste in the config file
- Save your config file

**Note**
The bot and config files can exist anywhere on your server, but must be kept together.

**7. Permissions**
- Linux - ```chmod 0755 linux-discord-userjoinannounce-x64``` or ```chmod 0755 linux-discord-userjoinannounce-x86```
- Windows - owner: r,w,x
	  - everyone else: r,x

**8. Starting the bot**
- Windows versions, just run the .exe
- Linux
	- Create a new screen - ```screen -S bot```
	- Change to the bot directory - ```cd /<path to bot folder>/```
	- Start the bot ```./linux-discord-userjoinannounce-x86``` or ```./linux-discord-userjoinannounce-x64```
	- Detach the screen - **Ctrl+a** then **d**
	- To reopen the screen ```screen -r bot```


# How does it work?

After the bot is invited to your target and announce Discord servers, it will
immediately begin listening for new member join events from the target server to
announce them in the announce server's announce channel.

Make sure you bot has any roles and permissions needed to view, post, and attach files
in the announcement channel.

---------------------------------------------------------------------

**NOTE**
This bot isn't complete.  The finished version will have a working announcement of when a user quits the target server

----

# Dependencies

| [configure](https://github.com/paked/configure) |
| [DiscordGo](https://github.com/bwmarrin/discordgo) |

# License
The source code for Discord-UserJoinAnnounce is released under the MIT License. See LICENSE for more details.

## Donations
All donations are appreciated and help pay for the energy drinks that keep me awake long enough to bring you cool new features and projects. Even if it's not much, it helps a lot in the long run!
You can find the donation link here: [Donation Link](https://paypal.me/JoshuaDoes)
