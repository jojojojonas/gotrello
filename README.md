# gotrello

With this library you should be able to interact with the Trello API in the easiest way possible. I'm just gradually working the features into this library, so it doesn't have a full feature set by any means. If you want to help me, feel free to extend the code!

## Install

To import the library into your project you have to execute the following command.

```console
go get github.com/jojojojonas/gotrello
```

## Card

### CreatecCard

To create a new card you need to provide the following data. Please enter the date in the following format: 2006-01-02.

```go
card, err := gotrello.CreateCart("key", "token", "list", "name", "description", "date")
if err != nil {
	fmt.Println(err)
}
```

### Create card member

To assign a member to a card the following function must be called.

```go
member, err := gotrello.CreateCartMember("key", "token", "card", "member")
if err != nil {
fmt.Println(err)
}
```