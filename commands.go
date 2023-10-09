package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
    embedColor int = 15485763
)


// Define bot commands

//Typedef commonly used func signature, so the code is pretty :3
// We cant declare functions with this tho, we still need to repeat the signature for them
type discordgoCommand func(s *discordgo.Session, i *discordgo.InteractionCreate)


func upcomingChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
    config.UpcomingChannelID = i.ChannelID;
    gotChannelID(s,i)
}

func gotChannelID(s *discordgo.Session, i *discordgo.InteractionCreate) {
    err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "Set current channel as the announcement channel",
        },  
    })
    if err != nil {
        log.Printf("User %v unable to run command, reason: %v", i.Member.User.Username, err)
    }
}


func sync(s *discordgo.Session, i *discordgo.InteractionCreate) {
    err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "Syncing",
        },  
    })
    if err != nil {
        log.Printf("User %v unable to run command, reason: %v", i.Member.User.Username, err)
    }
    scrapeItch()
}

func testCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Embeds: jamToEmbed("UOL Game Jam #9"),
        },  
    })
    if err != nil {
        log.Printf("User %v unable to run command, reason: %v", i.Member.User.Username, err)
    }
}

func getStatusEmbed() []*discordgo.MessageEmbed {
    return []*discordgo.MessageEmbed{
        { // Embed body
            Title: "Capy status",
            Color: embedColor,
            Footer: &discordgo.MessageEmbedFooter{IconURL: client.State.User.AvatarURL(""),Text: time.Now().Format("Mon at 15:04")},
            Fields: []*discordgo.MessageEmbedField{
                {
                    Name: "Last sync:",
                    Value: lastSynced.Format("Jan 02, Mon 15:04"),
                    Inline: true,
                },
                {
                    Name: "Sync time:",
                    Value: syncTime.String(),
                    Inline: true,
                },
            },
        }, 
    }
}
func returnStatus(s *discordgo.Session, i *discordgo.InteractionCreate){
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            //Content: "Time since last sync: " + lastSynced.Format("01/02 15:04"),
            Embeds: getStatusEmbed(),
        },
    })
}

func jamToEmbed(jamName string) []*discordgo.MessageEmbed {
    var jam Jam
    val, ok := jamEntries[jamName]
    if ok {
        jam = *val
    } else {
        return []*discordgo.MessageEmbed{{Title: "Jam named " +jamName + " not found."},}
    }
    return []*discordgo.MessageEmbed{
        {
            Title: jamName,
            Color: embedColor,
            URL: jam.JamLink,
            Image: &discordgo.MessageEmbedImage{URL: jam.ImageLink},
            Footer: &discordgo.MessageEmbedFooter{IconURL: jam.ImageLink,Text: time.Now().Format("Mon at 15:04")},
            Fields: []*discordgo.MessageEmbedField{
                {
                    Name: "Duration",
                    Value: jam.Duration,
                    Inline: true,
                },
                {
                    Name: "Starts in",
                    Value: jam.StartsIn.Round(time.Minute * 5).String(),
                    Inline: true,
                },
                {
                    Name: "Participants",
                    Value: jam.Joined,
                    Inline: true,
                },
            },
        },   
    }
}


var (
    commands = []*discordgo.ApplicationCommand{
        {
            Name: "test-command",
            Description: "baller test command",
        },
        {
            Name: "status",
            Description: "Returns bot status",
        },
        {
            Name: "upcoming",
            Description: "Sets the current channel to the upcoming jam announcement channel",
        },
        {
            Name: "sync",
            Description: "yeah",
        },
    }

    command_handlers = map[string]discordgoCommand{
        "test-command": testCommand,
        "status": returnStatus,
        "upcoming": upcomingChannel,
        "sync": sync,
    }
)

