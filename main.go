package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/paked/configure"
	"github.com/bwmarrin/discordgo"
)

var (
	conf = configure.New()
	confBotToken = conf.String("botToken", "", "The bot token to authenticate to Discord with")
	confTargetServer = conf.String("targetServer", "", "Server ID to read user updates from")
	confAnnounceServer = conf.String("announceServer", "", "Server ID to send user update announcements in")
	confAnnounceChannel = conf.String("announceChannel", "", "Channel ID to send user update announcements to")
	confDebugMode = conf.Bool("debugMode", false, "Debug Mode")
	botToken string = ""
	targetServer string = ""
	announceServer string = ""
	announceChannel string = ""
	debugMode bool = false

	targetServerName string = ""
)

func debugLog(message string) {
	if debugMode {
		fmt.Println(message)
	}
}

func init() {
	conf.Use(configure.NewFlag())
	conf.Use(configure.NewJSONFromFile("config.json"))
}

func main() {
	fmt.Println("Discord-UserJoinAnnounce © JoshuaDoes: 2018.\n")

	fmt.Println("> Loading configuration...")
	conf.Parse()
	botToken = *confBotToken
	targetServer = *confTargetServer
	announceServer = *confAnnounceServer
	announceChannel = *confAnnounceChannel
	debugMode = *confDebugMode
	if (botToken == "" || targetServer == "" || announceServer == "" || announceChannel == "") {
		fmt.Println("> Configuration not properly setup, exiting...")
		return
	} else {
		fmt.Println("> Successfully loaded configuration.")
		debugLog("targetServer: " + targetServer)
		debugLog("announceServer: " + announceServer)
		debugLog("announceChannel: " + announceChannel)
	}

	debugLog("> Creating a new Discord session...")
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("> Error creating Discord session: " + fmt.Sprintf("%v", err))
	}
	
	debugLog("> Registering Ready callback handler...")
	dg.AddHandler(ready)
	debugLog("> Registering GuildMemberAdd callback handler...")
	dg.AddHandler(guildMemberAdd)
	debugLog("> Registering GuildMemberRemove callback handler...")
	dg.AddHandler(guildMemberRemove)
	
	debugLog("> Establishing a connection to Discord...")
	err = dg.Open()
	if err != nil {
		fmt.Println("> Error opening connection to Discord: " + fmt.Sprintf("%v", err))
		return
	}
	
	fmt.Println("> Discord-UserJoinAnnounce has started successfully.")
	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	
	debugLog("> Closing Discord session...")
	dg.Close()
}

func ready(session *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("> Discord is now ready.")
}

func guildMemberAdd(session *discordgo.Session, member *discordgo.GuildMemberAdd) {
	targetServerDetails, err := guildDetails(targetServer, session)
	if err != nil {
		debugLog("> Error getting target server name.")
		targetServerName = "the target server"
	} else {
		targetServerName = targetServerDetails.Name
		debugLog("> Found target server name: " + targetServerName)
	}

	if member.GuildID == targetServer {
		joinedAt := member.JoinedAt
		userID := member.User.ID
		username := member.User.String() //username#0000
		avatarURL := member.User.AvatarURL("2048")
		isVerified := member.User.Verified
		isMFAEnabled := member.User.MFAEnabled
		isBot := member.User.Bot

		joinedAtTimeFormatted := ""
		joinedAtTime, err := time.Parse(time.RFC3339Nano, joinedAt)
		if err != nil {
			debugLog("> Error formatting time: " + joinedAt)
			joinedAtTimeFormatted = joinedAt
		} else {
			joinedAtMonth := joinedAtTime.Month().String()
			joinedAtDay := joinedAtTime.Day()
			joinedAtYear := joinedAtTime.Year()
			joinedAtHour := joinedAtTime.Hour()
			joinedAtMinute := joinedAtTime.Minute()
			joinedAtSecond := joinedAtTime.Second()
			joinedAtTimeFormatted = joinedAtMonth + " " + strconv.Itoa(joinedAtDay) + ", " + strconv.Itoa(joinedAtYear) + " at " + strconv.Itoa(joinedAtHour) + ":" + strconv.Itoa(joinedAtMinute) + ":" + strconv.Itoa(joinedAtSecond)
		}

		debugLog("> Sending typing event...")
		session.ChannelTyping(announceChannel)
		debugLog("> Creating member add embed...")
		embedMemberAdd := NewEmbed().
			SetTitle("New Member").
			SetDescription("A new member joined " + targetServerName + ".").
			AddField("Joined At", joinedAtTimeFormatted).
			AddField("User ID", userID).
			AddField("Username", username).
			AddField("Verified Account", strconv.FormatBool(isVerified)).
			AddField("Multi-Factor Authentication", strconv.FormatBool(isMFAEnabled)).
			AddField("Bot", strconv.FormatBool(isBot)).
			SetImage(avatarURL).
			SetColor(0xff0000).MessageEmbed
		debugLog("> Sending member add embed...")
		_, err = session.ChannelMessageSendEmbed(announceChannel, embedMemberAdd)
		if err != nil {
			debugLog("> Error sending member add embed: " + fmt.Sprintf("%v", err))
		}
	}
}

func guildMemberRemove(session *discordgo.Session, member *discordgo.GuildMemberRemove) {
	targetServerDetails, err := guildDetails(targetServer, session)
	if err != nil {
		debugLog("> Error getting target server name.")
		targetServerName = "the target server"
	} else {
		targetServerName = targetServerDetails.Name
		debugLog("> Found target server name: " + targetServerName)
	}

	if member.GuildID == targetServer {
		joinedAt := member.JoinedAt
		userID := member.User.ID
		username := member.User.String() //username#0000
		avatarURL := member.User.AvatarURL("2048")
		isVerified := member.User.Verified
		isMFAEnabled := member.User.MFAEnabled
		isBot := member.User.Bot

		joinedAtTimeFormatted := ""
		joinedAtTime, err := time.Parse(time.RFC3339Nano	, joinedAt)
		if err != nil {
			debugLog("> Error formatting time: " + joinedAt)
			joinedAtTimeFormatted = joinedAt
		} else {
			joinedAtMonth := joinedAtTime.Month().String()
			joinedAtDay := joinedAtTime.Day()
			joinedAtYear := joinedAtTime.Year()
			joinedAtHour := joinedAtTime.Hour()
			joinedAtMinute := joinedAtTime.Minute()
			joinedAtSecond := joinedAtTime.Second()
			joinedAtTimeFormatted = joinedAtMonth + " " + strconv.Itoa(joinedAtDay) + ", " + strconv.Itoa(joinedAtYear) + " at " + strconv.Itoa(joinedAtHour) + ":" + strconv.Itoa(joinedAtMinute) + ":" + strconv.Itoa(joinedAtSecond)
		}
		
		debugLog("> Sending typing event...")
		session.ChannelTyping(announceChannel)
		debugLog("> Creating member remove embed...")
		embedMemberRemove := NewEmbed().
			SetTitle("Lost Member").
			SetDescription("A member left " + targetServerName + ".").
			AddField("Joined At", joinedAtTimeFormatted).
			AddField("User ID", userID).
			AddField("Username", username).
			AddField("Verified Account", strconv.FormatBool(isVerified)).
			AddField("Multi-Factor Authentication", strconv.FormatBool(isMFAEnabled)).
			AddField("Bot", strconv.FormatBool(isBot)).
			SetImage(avatarURL).
			SetColor(0xff0000).MessageEmbed
		debugLog("> Sending member remove embed...")
		_, err = session.ChannelMessageSendEmbed(announceChannel, embedMemberRemove)
		if err != nil {
			debugLog("> Error sending member remove embed: " + fmt.Sprintf("%v", err))
		}
	}
}

func guildDetails(guildID string, s *discordgo.Session) (*discordgo.Guild, error) {
	guildDetails, err := s.State.Guild(guildID)
	if err != nil {
		return nil, err
	}
	return guildDetails, nil
}