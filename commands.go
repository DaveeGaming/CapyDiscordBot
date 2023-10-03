package main

import (
	"github.com/bwmarrin/discordgo"
)

// Define bot commands

//Typedef commonly used func signature, so the code is pretty :3
// We cant declare functions with this tho, we still need to repeat the signature for them
type discordgoCommand func(s *discordgo.Session, i *discordgo.InteractionCreate)


func testCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "Oh wow, our first slash command!",
        },  
    })
}

func getStatusEmbed() []*discordgo.MessageEmbed {
    return []*discordgo.MessageEmbed{
        { // Embed body
            Title: "Capy status",
            Color: 15485763,
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

