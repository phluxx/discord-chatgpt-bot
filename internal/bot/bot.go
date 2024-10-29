package bot

import (
    "fmt"
    "log"
    "strings"
    "github.com/bwmarrin/discordgo"
    "discord-chatgpt-bot/internal/chatgpt"
    "discord-chatgpt-bot/config"
)

var Bot *discordgo.Session

func Start() error {
    var err error
    Bot, err = discordgo.New("Bot " + config.BotToken)
    if err != nil {
        return fmt.Errorf("error creating a Discord session: %w", err)
    }

    fmt.Println("Discord session created successfully")

    Bot.AddHandler(commandHandler)

    err = Bot.Open()
    if err != nil {
        return fmt.Errorf("error opening connection: %w", err)
    }

    fmt.Println("Bot is now running.")
    return nil
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.Bot {
        return
    }

    if strings.HasPrefix(m.Content, "ChatGPT: ") {
        prompt := strings.TrimSpace(strings.TrimPrefix(m.Content, "ChatGPT: "))
        
        response, err := chatgpt.GetResponse(prompt)
        if err != nil {
            log.Printf("Error getting response from ChatGPT: %v", err)
            s.ChannelMessageSend(m.ChannelID, "Error: Unable to get response from ChatGPT.")
            return
        }

        reply := fmt.Sprintf("<@%s>: %s", m.Author.ID, response)
        s.ChannelMessageSend(m.ChannelID, reply)
    }
}

