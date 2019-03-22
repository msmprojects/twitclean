## TwitClean

This is a dumb little program, largely lifted from [Harold](https://github.com/adamdrake/harold), meant to be `cron`ned up somewhere. (In my case, on my laptop, since that machine is generally on all day.) It deletes tweets from a timeline older than a specific number of days.

Why do this? Because no one needs dumb things they said on Twitter lasting longer than their relevance.

### Prerequisites

- Go 1.10 or higher, since this uses Go Modules. Make sure you have `GO111MODULE` set appropriately for your system.
- A Twitter application established for your Twitter account. Your application doesn't need to use "Log In With Twitter" but does need read and write permissions, which are the defaults when you create an app. There will be four values assigned to your application under "Keys and Tokens": two consumer API keys and two access tokens. Copy these and keep them safe until you need them later.

### To Build

- Check this repository out locally.
- In the directory: `go build .`

You should see a binary, `twitclean`, there now. Leave it there, put it elsewhere on your path, whatever.

### To Run

- Set the following environment variables using the values from your Twitter app:

```shell
export TC_TWITTER_CONSUMER_KEY=<twitter API key>
export TC_TWITTER_CONSUMER_SECRET=<twitter API secret key>
export TC_TWITTER_ACCESS_TOKEN=<twitter access token>
export TC_TWITTER_ACCESS_TOKEN_SECRET=<twitter secret access token>
```

Then, run the program.  

```
./twitclean -d 30 # deletes anything older than 30 days
```

The default duration is 28 days:

```
./twitclean # deletes anything older than 30 days
```

Add the program to your local `crontab` and 

### WARNING

*Be warned*, running this program is a *destructive* act. 

*You will lose tweets.*

*You will lose pinned tweets.*

*You will lose replies to other people's tweets.*

### Thanks To

- [Harold](https://github.com/adamdrake/harold)
- [Anaconda](https://github.com/ChimeraCoder/anaconda)
- [Logrus](https://github.com/sirupsen/logrus)

### License

This program is licensed under the terms of the [Blue Oak Model License 1.0.0](https://blueoakcouncil.org/license/1.0.0)
