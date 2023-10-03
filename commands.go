package main

import (
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


func testCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Embeds: jamToEmbed(jamEntries[0]),
        },  
    })
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


func jamToEmbed(jam Jam) []*discordgo.MessageEmbed {
    return []*discordgo.MessageEmbed{
        {
            Title: jam.Name,
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
    }

    command_handlers = map[string]discordgoCommand{
        "test-command": testCommand,
        "status": returnStatus,
    }
)

