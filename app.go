package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
    RemoveCommands bool
    CreateConfig bool
    // Define global session
    client *discordgo.Session
)


// CLI Arguments
func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
    flag.BoolVar(&RemoveCommands, "r", true, "Remove commands after shutdown")
    flag.BoolVar(&CreateConfig, "c", false, "Create config file, if none found")
	flag.Parse()
    
    if Token == "" {
        log.Fatal("No token passed as argument, please pass it as -t [TOKEN_HERE]") //Same as print and then os.exit()
    }

    // Initialize discordgo client
    var err error
    client, err = discordgo.New("Bot " + Token)
    log.Println("Initializing discord bot")
    if err != nil {
        log.Fatalf("Unable to create discord session: %v", err)
    }

    // Add listener, this fires our functions when a command is called
    client.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate){
        if command, found := command_handlers[i.ApplicationCommandData().Name]
        found {
            command(s,i)
        }
    })
}

func main() {
    // Event that fires when we launch the bot
    client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready){
        log.Printf("Succesfully logged in as: %v",client.State.User.Username) 
    })

    // Try to open the session (Websocket connection to discord)
    err := client.Open()
    if err != nil {
       log.Fatalf("Couldn't create websocket: %v", err)
    }

    // Registering commands
    registered_commands := make([]*discordgo.ApplicationCommand, len(commands))
    for i,v := range commands {
        cmd, err := client.ApplicationCommandCreate(client.State.User.ID,"",v)
        if err != nil {
            log.Panicf("Couldn't create command with name %v: %v",v.Name,err)
        }
        registered_commands[i] = cmd
    }

    //Set up jams
    go StartScraper()


 
    defer client.Close()


    //Setup os signal detection to shut down bot
    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, os.Interrupt)
    log.Println("<C-c> to exit")
    <-shutdown

    if RemoveCommands {
        log.Println("Removing commands...")
        for _,v := range registered_commands {
            err := client.ApplicationCommandDelete(client.State.User.ID,"",v.ID)
            if err != nil {
                log.Panicf("Unable to delete command %v: %v",v.Name,err)
            }
        }
    }
    saveConfig()
    
    log.Println("Shutting down...")
}
