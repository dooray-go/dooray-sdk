> 해당 기능은 비공식이며 , 커뮤니티 기여로 작성하는 도구입니다. NHN Dooray 서비스에서 제공하는 기능이 아님을 밝혀 둡니다.

# Dooray! API in Go

This is the Dooray Client Library for Go.

This library is based on the [Dooray! API Documentation](https://helpdesk.dooray.com/share/pages/9wWo-xwiR66BO5LGshgVTg/2939987647631384419).

## Installing
### go get
```
$ go get -u github.com/dooray-go/dooray-sdk
```

## Features

| Category | Feature | Method | Description |
|----------|---------|--------|-------------|
| **[Socket Mode](socketmode/README.md)** | Agent | `NewAgent` | Real-time event handling via WebSocket |
| **Messenger** | Webhook | `PostWebhook` | Send messages via webhook |
| | Direct Send | `DirectSend` | Send a 1:1 message (`POST .../direct-send`) |
| | Get Channels | `GetChannels` | List joined channels |
| | Create Channel | `CreateChannel` | Create a direct or private channel |
| | Join Channel | `JoinChannelMembers` | Add members to a channel |
| | Leave Channel | `LeaveChannelMembers` | Remove members from a channel |
| | Send Message | `SendMessage` | Send a channel message (`POST .../logs`) |
| | Update Channel Log | `UpdateChannelLog` | Edit a channel message |
| | Delete Channel Log | `DeleteChannelLog` | Delete a channel message |
| | Reply Channel Log | `ReplyChannelLog` | Reply to a channel message |
| | Create And Send Thread | `CreateAndSendThread` | Create a thread and send |
| | Create Thread From Log | `CreateAndSendThreadFromLog` | Create a thread from an existing log |
| **Project** | Get Projects | `GetProjects` | Retrieve list of projects |
| | Get Posts | `GetPosts` | Retrieve posts from a project |
| | Get Posts (Options) | `GetPostsWithOptions` | Retrieve posts with full query parameters (paging, filters, date, sort) |
| | Get Post | `GetPost` | Retrieve a single post (task) by ID |
| | Create Post | `CreatePost` | Create a new post in a project |
| **Calendar** | Get Calendars | `GetCalendars` | Retrieve list of calendars |
| | Get Events | `GetEvents` | Retrieve events from calendars |
| | Create Event | `CreateEvent` | Create a new calendar event |
| | Update Event | `UpdateEvent` | Update an existing calendar event |
| | Delete Event | `DeleteEvent` | Delete a calendar event |
| **Wiki** | Get Wikis | `GetWikis` | List accessible wikis |
| | Get Page | `GetPage` | Get a wiki page by page ID |
| | Get Wiki Page | `GetWikiPage` | Get a page in a wiki |
| | Get Pages | `GetPages` | List sibling pages (one depth) |
| | Create Page | `CreatePage` | Create a wiki page |
| | Update Page | `UpdatePage` | Update page title and body |
| | Update Page Title | `UpdatePageTitle` | Update page title only |
| | Update Page Content | `UpdatePageContent` | Update page body only |
| | Update Page Referrers | `UpdatePageReferrers` | Replace page referrers |
| | Move Page | `MovePage` | Move or reorder a page |
| | Delete Page | `DeletePage` | Delete a wiki page |
| | Create Comment | `CreateComment` | Add a page comment |
| | Get Comments | `GetComments` | List page comments |
| | Get Comment | `GetComment` | Get a comment |
| | Update Comment | `UpdateComment` | Update a comment |
| | Delete Comment | `DeleteComment` | Delete a comment |
| | Get Shared Links | `GetSharedLinks` | List page shared links |
| | Upload Page File | `UploadPageFile` | Upload an attachment or inline image |
| | Download Page File | `DownloadPageFile` | Download a page file |
| | Delete Page File | `DeletePageFile` | Delete a page file |
| | Upload Wiki File | `UploadWikiFile` | Upload a wiki file |
| | Download Attach File | `DownloadAttachFile` | Download by attachFileId |

## Messenger WebHook Example
```go
package main

import (
    "context"
    "github.com/dooray-go/dooray-sdk"
    "log"
    "time"
)

func main() {
    ctx1 := context.Background()
    subCtx1, _ := context.WithTimeout(ctx1, 3*time.Second)
    doorayErr := dooray.PostWebhookContext(subCtx1, "[Your WebHook URL]", &dooray.WebhookMessage{
        BotName: "dooray-go",
        Text:    "Hello",
    })
    
    if doorayErr != nil {
        log.Printf("dial error: %s", doorayErr.Error())
    }
}
```

## OpenApi Examples

### Get Posts from a Project
```go
package main

import (
    "fmt"
    "log"

    "github.com/dooray-go/dooray-sdk/openapi/project"
)

func main() {
    projectClient := project.NewDefaultProject()
    projectID := "your-project-id"

    // Simple: filter by member IDs and workflow classes
    response, err := projectClient.GetPosts("your-dooray-api-key", projectID, "member-id-1,member-id-2", "registered,working")
    if err != nil {
        log.Fatalf("Failed to get posts: %s", err)
    }

    fmt.Printf("Total posts: %d\n", response.TotalCount)
    for _, post := range response.Result {
        fmt.Printf("Post #%d: %s\n", post.Number, post.Subject)
    }
}
```

### Get Posts with Options
```go
package main

import (
    "fmt"
    "log"

    "github.com/dooray-go/dooray-sdk/openapi/project"
)

func main() {
    projectClient := project.NewDefaultProject()
    projectID := "your-project-id"

    size := 10
    page := 0
    toMemberSize := 1

    response, err := projectClient.GetPostsWithOptions("your-dooray-api-key", projectID, project.GetPostsOptions{
        // Paging
        Page: &page,
        Size: &size,

        // Filters
        PostWorkflowClasses: "registered,working",
        ToMemberIds:         "member-id",
        ToMemberSize:        &toMemberSize,  // 1: toMemberIds[0]이 혼자 담당인 업무
        TagIds:              "tag-id-1,tag-id-2",
        MilestoneIds:        "milestone-id",

        // Date filters (today, thisweek, prev-{N}d, next-{N}d, or ISO8601 range)
        CreatedAt: "prev-7d",
        DueAt:     "next-30d",

        // Sort (prefix with - for descending)
        Order: "-createdAt",
    })
    if err != nil {
        log.Fatalf("Failed to get posts: %s", err)
    }

    fmt.Printf("Total posts: %d\n", response.TotalCount)
    for _, post := range response.Result {
        fmt.Printf("Post #%d: %s (status: %s)\n", post.Number, post.Subject, post.WorkflowClass)
    }
}
```

### Create a Post in a Project
```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/dooray-go/dooray-sdk/openapi/project"
    model "github.com/dooray-go/dooray-sdk/openapi/model/project"
    "github.com/dooray-go/dooray-sdk/utils"
)

func main() {
    // Create a project client
    projectClient := project.NewDefaultProject()

    // Set due date (24 hours from now)
    dueDate := utils.JsonTime(time.Now().Add(24 * time.Hour))

    // Create a post request
    postRequest := model.PostRequest{
        Subject: "New Task",
        Body: model.PostBody{
            MimeType: "text/html",
            Content:  "<p>This is a new task created via API</p>",
        },
        Users: &model.PostUsers{
            To: []model.PostRecipient{
                {
                    Type: "member",
                    Member: &model.PostMember{
                        OrganizationMemberID: "member-id",
                    },
                },
            },
        },
        DueDate:  &dueDate,
        Priority: "normal", // urgent | high | normal | low
        TagIDs:   []string{"tag-id-1", "tag-id-2"},
    }

    // Create the post
    projectID := "your-project-id"
    response, err := projectClient.CreatePost("your-dooray-api-key", projectID, postRequest)
    if err != nil {
        log.Fatalf("Failed to create post: %s", err)
    }

    fmt.Printf("Post created successfully! ID: %s\n", response.Result.ID)
}
```

### Get a Project Post
```go
package main

import (
    "fmt"
    "log"

    "github.com/dooray-go/dooray-sdk/openapi/project"
)

func main() {
    projectClient := project.NewDefaultProject()

    response, err := projectClient.GetPost("your-dooray-api-key", "project-id", "post-id")
    if err != nil {
        log.Fatalf("Failed to get post: %s", err)
    }

    fmt.Printf("Post: %s\n%s\n", response.Result.Subject, response.Result.Body.Content)
}
```

### Get Wiki Pages
```go
package main

import (
    "fmt"
    "log"

    "github.com/dooray-go/dooray-sdk/openapi/wiki"
)

func main() {
    wikiClient := wiki.NewDefaultWiki()

    wikis, err := wikiClient.GetWikis("your-dooray-api-key")
    if err != nil {
        log.Fatalf("Failed to list wikis: %s", err)
    }
    fmt.Printf("wikis: %d\n", wikis.TotalCount)

    page, err := wikiClient.GetPage("your-dooray-api-key", "page-id")
    if err != nil {
        log.Fatalf("Failed to get page: %s", err)
    }
    fmt.Printf("%s\n%s\n", page.Result.Subject, page.Result.Body.Content)
}
```

### Update a Calendar Event
```go
package main

import (
    "fmt"
    "log"

    "github.com/dooray-go/dooray-sdk/openapi/calendar"
    model "github.com/dooray-go/dooray-sdk/openapi/model/calendar"
)

func main() {
    calendarClient := calendar.NewDefaultCalendar()

    // Only non-empty fields are updated; omit fields you want to leave unchanged.
    subject := "Updated meeting"
    location := "Meeting Room B"
    response, err := calendarClient.UpdateEvent("your-dooray-api-key", "calendar-id", "event-id", model.UpdateEventRequest{
        Subject:  &subject,
        Location: &location,
    })
    if err != nil {
        log.Fatalf("Failed to update event: %s", err)
    }

    fmt.Printf("Event updated: %v\n", response.Header.IsSuccessful)
}
```

### Delete a Calendar Event
```go
package main

import (
    "fmt"
    "log"

    "github.com/dooray-go/dooray-sdk/openapi/calendar"
    model "github.com/dooray-go/dooray-sdk/openapi/model/calendar"
)

func main() {
    calendarClient := calendar.NewDefaultCalendar()

    // Single event: deleteType defaults to "this"
    response, err := calendarClient.DeleteEvent("your-dooray-api-key", "calendar-id", "event-id", model.DeleteEventRequest{})
    if err != nil {
        log.Fatalf("Failed to delete event: %s", err)
    }
    fmt.Printf("Event deleted: %v\n", response.Header.IsSuccessful)

    // Recurring occurrence: use the id from GetEvents (may include a timestamp suffix)
    // this / wholeFromThis / whole
    _, err = calendarClient.DeleteEvent("your-dooray-api-key", "calendar-id", "event-id-20240228T013000Z", model.DeleteEventRequest{
        DeleteType: model.DeleteTypeWholeFromThis,
    })
    if err != nil {
        log.Fatalf("Failed to delete recurring event: %s", err)
    }
}
```

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.
