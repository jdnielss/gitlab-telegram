package main

import (
	"flag"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Get GitLab token and project details from environment variables
	telegramToken := flag.String("t", "", "Telegram bot token")
	chatID := flag.String("id", "", "ChatID")
	gitlabToken := flag.String("gt", "", "Gitlab token")
	gitlabBaseURL := flag.String("url", "", "Gitlab base url")
	projectID := flag.String("pid", "", "Gitlab project id")
	mergeRequestIID := flag.Int("mid", 0, "Gitlab merge request id")

	flag.Parse()

	if *telegramToken == "" || *chatID == "" || *gitlabBaseURL == "" || *gitlabToken == "" || *projectID == "" || *mergeRequestIID == 0 {
		fmt.Println("Missing required arguments")
		flag.Usage()
		os.Exit(1)
	}

	git, err := gitlab.NewClient(*gitlabToken, gitlab.WithBaseURL(*gitlabBaseURL))

	if err != nil {
		log.Fatalf("Failed to create client")
	}

	// Fetch merge request details
	mr, _, err := git.MergeRequests.GetMergeRequest(*projectID, *mergeRequestIID, nil)

	if err != nil {
		log.Fatalf("Failed to get merge request")
	}

	message := fmt.Sprintf(
		"Title: %s\nDescription: %s\nState: %s\nTarget Branch: %s\nAuthor: %s (%s)\nWeb URL: %s\nHas Conflicts: %t",
		mr.Title,
		mr.Description,
		mr.State,
		mr.TargetBranch,
		mr.Author.Name,
		mr.Author.Username,
		mr.WebURL,
		mr.HasConflicts,
	)

	// Assignees
	if len(mr.Assignees) > 0 {
		assignees := "\nAssignees:"
		for _, assignee := range mr.Assignees {
			assignees += fmt.Sprintf("\n- %s (%s)", assignee.Name, assignee.Username)
		}
		message += assignees
	}

	// Reviewers
	if len(mr.Reviewers) > 0 {
		reviewers := "\nReviewers:"
		for _, reviewer := range mr.Reviewers {
			reviewers += fmt.Sprintf("\n- %s (%s)", reviewer.Name, reviewer.Username)
		}
		message += reviewers
	}

	// Send the message to Telegram
	sendMessageToTelegram(*telegramToken, *chatID, message)
}

func sendMessageToTelegram(token, chatID, message string) {
	telegramAPI := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)

	// Send the POST request to Telegram API
	resp, err := http.Post(telegramAPI, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Failed to send message to Telegram")
	}
	defer resp.Body.Close() // Proper use of defer to close the response body

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to send message.")
	}

	fmt.Println("Message sent to Telegram successfully!")
}
