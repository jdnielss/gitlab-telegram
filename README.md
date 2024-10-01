
# GitLab Merge Request to Telegram

This project retrieves GitLab merge request details and sends them to a Telegram chat. You can run the project using command-line flags to provide necessary configuration.

## Prerequisites

- [GitLab Personal Access Token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)
- [Telegram Bot Token](https://core.telegram.org/bots#botfather)
- [Telegram Chat ID](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)

## Command-Line Flags

The following command-line flags are required to run the project:

- `-t` or `--telegram_token`: Your Telegram bot token.
- `-id` or `--chat_id`: The chat ID where the message will be sent.
- `-gt` or `--gitlab_token`: Your GitLab personal access token.
- `-url` or `--gitlab_base_url`: The base URL of your GitLab instance
- `-pid` or `--project_id`: The ID of your GitLab project.
- `-mid` or `--merge_request_iid`: The IID of the merge request.

### Example Usage

You can run the program using the following command:

```bash
go run main.go   -t="your_telegram_token"   -id="your_chat_id"   -gt="your_gitlab_token"   -url="https://gitlab.com"   -pid="79"   -mid=123
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
