# approveapi-go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/approveapi/approveapi-go)

Go API bindings for the [ApproveAPI HTTP API](https://approveapi.com).

ApproveAPI is a simple API to request a user's real-time approval on anything via email, sms + mobile push.

Features (almost complete):
- [x] Send Prompt
  - [x] web redirect actions (i.e. magic links)
  - [x] custom approve/reject buttons
  - [x] metadata
  - [x] long polling
- [x] Retrieve Prompt
- [x] Check Prompt status 
- [x] Futures support
- [ ] Webhook callbacks

## Install

Install the dependencies:
```
go get "github.com/approveapi/approveapi-go"
```

## Import:
```golang
import "github.com/approveapi/approveapi-go"
```

## Getting Started

To get started, we create a client:

```go
client := approveapi.CreateClient("sk_test_yourapikeyhere")
```

Now we can make API calls. For example, let's send an approval prompt to confirm a financial transaction.

```go
longPoll := true
prompt, _, err := client.ApproveApi.CreatePrompt(
    approveapi.CreatePromptRequest {
      User: "alice@approveapi.com",
      Body: `A transfer of $1337.45 from acct 0294 to acct 1045
          has been initiated. Do you want to authorize this transfer?`,
      ApproveText: "Authorize",
      RejectText: "Reject",
      LongPoll: &longPoll,
    },
)
if err != nil {
    fmt.Printf("%#v\n", err)
    return
}
if prompt.Answer != nil {
    if prompt.Answer.Result {
        fmt.Println("Request approved")
    } else {
        fmt.Println("Request rejected")
    }
} else {
    fmt.Println("No response yet")
}

```

## Documentation

Full documentation is available here: [approveapi.com/docs](https://www.approveapi.com/docs/?go).


