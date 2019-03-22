## TwitClean

This is a dumb little program, largely lifted from [Harold](https://github.com/adamdrake/harold), meant to be `cron`ned up somewhere. (In my case, on my laptop, since that machine is generally on all day.) It deletes tweets from a specific timeline older than a specific number of days.

Why do this? Because no one needs things they said on Twitter hanging around longer than their relevance half-life. That's probably only about 90 seconds for most tweets, but let's just pretend four weeks is the social media equivalent of prehistory. Twitter is toxic. Use [micro.blog](https://micro.blog) or [Mastodon](https://joinmastodon.org) or email or pretty much anything else instead. You'll be happier.

### Prerequisites

- Go 1.11 or higher, since this uses Go Modules. Make sure you have `GO111MODULE` set appropriately for your system. You can probably make this compile with earlier versions of Go, but I've not tested that.
- A Twitter application established for your Twitter account at [Twitter Developers](https://developer.twitter.com/). Your application doesn't need to use "Log In With Twitter" but does need read and write permissions, which are the defaults when you create an app. There will be four values assigned to your application under "Keys and Tokens": two consumer API keys and two access tokens. Copy these and keep them safe until you need them later.

### To Build

- Check this repository out locally.
- In the directory, run `go build .` or `go install .`

If you ran `go build`, you should see a binary, `twitclean`, created in the repository directory. Leave it there, put it elsewhere on your path, whatever. If you ran `go install`, the `twitclean` binary has been moved to your Go `bin` directory, which should be in your `$PATH` (which, now that I think about it, is weird when you have Go modules configured, but that's a question for some other day).

### To Run

- Set the following environment variables using the values you generated for  your Twitter app (you kept them safe, right?):

```shell
export TC_TWITTER_CONSUMER_KEY=<twitter API key>
export TC_TWITTER_CONSUMER_SECRET=<twitter API secret key>
export TC_TWITTER_ACCESS_TOKEN=<twitter access token>
export TC_TWITTER_ACCESS_TOKEN_SECRET=<twitter secret access token>
```

If you don't set these, `twitclean` will quit with errors.

Then, run the program:

```
# deletes anything older than 30 days
twitclean -d 30 
```

The default duration is 28 days:

```
# deletes anything older than 28 days
twitclean 
```

Add the program to your local `crontab` and enjoy an automated ephemeral social media presence.

```crontab
# run twitclean -d 30 at 3pm every day
00 15 * * * twitclean -d 30 2>&1 >/dev/null | terminal-notifier
```

In the `crontab` example above, I'm using [terminal-notifier](https://github.com/julienXX/terminal-notifier) to post the execution results to my Mac's notifications. The `twitclean` executable prints all of its output to `stderr` so a little output redirection is necessary, hence the `2>&1 >/dev/null |` bits.

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
