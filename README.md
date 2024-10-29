# Discord ChatGPT Bot

This project is a Discord bot written in Golang that responds to user prompts with ChatGPT-generated responses. The bot listens for messages in a Discord server that start with a specific prefix (`ChatGPT:`) and sends the prompt to the OpenAI API for a response. It uses the `discordgo` library for interacting with Discord and the OpenAI API for text generation.

## Features

- Listens for messages with the prefix `ChatGPT:`.
- Sends prompts to OpenAI's ChatGPT API.
- Replies in Discord channels, mentioning the user who requested the prompt.

## Project Structure

```
discord-chatgpt-bot/
├── config/
│   └── config.go           # Holds configurations for bot token and API key
├── cmd/
│   └── main.go             # Main entry point for the bot
├── internal/
│   ├── bot/
│   │   ├── bot.go          # Bot initialization and message handling
│   └── chatgpt/
│       └── client.go       # ChatGPT client for interacting with the OpenAI API
└── go.mod
```

## Setup and Configuration

### Prerequisites

- [Go](https://golang.org/) (version 1.16+)
- A Discord bot token (see [Discord Developer Portal](https://discord.com/developers/applications))
- An OpenAI API key (see [OpenAI API](https://platform.openai.com/overview))

### Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/yourusername/discord-chatgpt-bot.git
    cd discord-chatgpt-bot
    ```

2. **Initialize the Go module**:

    ```bash
    go mod tidy
    ```

3. **Set Environment Variables**:

   Set your Discord bot token and OpenAI API key in your environment.

    ```bash
    export DISCORD_BOT_TOKEN="your_discord_token"
    export CHATGPT_API_KEY="your_openai_api_key"
    ```

4. **Run the bot**:

    ```bash
    go run cmd/main.go
    ```

### Usage

1. Invite your bot to a Discord server using an invite link from the [Discord Developer Portal](https://discord.com/developers/applications).
2. In the server, type a message in the format:

    ```
    ChatGPT: your question or prompt here
    ```

3. The bot will respond with an answer generated by ChatGPT.

### Troubleshooting

- **Bot Not Responding**: Verify the bot has permission to read and send messages in the channel.
- **Invalid API Key**: Double-check your OpenAI API key and ensure it is set correctly in your environment.
- **404/400 Errors**: Confirm that your `client.go` configuration and the OpenAI API endpoint are correct.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to open issues or submit pull requests if you'd like to improve the project!
