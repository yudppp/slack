## slack

### examples

```go
package main

import (
    "fmt"
    "os"

    "github.com/yudppp/slack"
)

func main() {
    slack.SetWebhookURL(os.Getenv("SLACK_WEBHOOK_URL"))
    slack.SetDefaultUsername("gopher")
    slack.SetDefaultChannel("#general")

    err := slack.NewMessage().SetText("hello every one").SetChannel("#random").Send()
    if err != nil {
        panic(err)
    }

    err = slack.NewMessage().AddAttachment(slack.NewAttachment().SetTitle("hello")).Send()
    if err != nil {
        panic(err)
    }
}

```
