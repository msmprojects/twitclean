## TwitClean

This is a dumb little program, largely lifted from [Harold](https://github.com/adamdrake/harold), meant to be scheduled to run somewhere so you don't have to think about it. TwitClean deletes tweets older than a specific number of days from a specific timeline.

Why do this? Because no one needs things they said on Twitter hanging around longer than their relevance half-life. That's probably only about 90 seconds for most tweets, but let's just pretend four weeks is the social media equivalent of prehistory. Twitter is toxic. Use [micro.blog](https://micro.blog) or [Mastodon](https://joinmastodon.org) or email or pretty much anything else instead. You'll be happier and the web will be healthier.

### Prerequisites

- Go 1.11 or higher, since this uses Go modules. Make sure you have `GO111MODULE` set appropriately for your system. You can probably compile this  with earlier versions of Go, but I've not tried that.
- A Twitter application established for your Twitter account at [Twitter Developers](https://developer.twitter.com/). Your application doesn't need to use "Log In With Twitter" but does need read and write permissions, which are the defaults when you create an app. There will be four values assigned to your application under "Keys and Tokens": two consumer API keys and two access tokens. Copy these and keep them safe until you need them later.

### To Build

- Check this repository out locally.
- In the directory, run `go build .` or `go install .`

If you ran `go build`, you should see a binary, `twitclean`, created in the repository directory. Leave it there, put it elsewhere on your path, whatever. If you ran `go install`, the `twitclean` binary has been moved to your Go `bin` directory, which should be in your `$PATH` (which, now that I think about it, is weird when you have Go modules configured, but that's a question for some other day).

### To Run

- Set the following environment variables using the values you generated for  your Twitter app (you kept them safe, right?):

```shell
export TC_TWITTER_CONSUMER_KEY={twitter API key}
export TC_TWITTER_CONSUMER_SECRET={twitter API secret key}
export TC_TWITTER_ACCESS_TOKEN={twitter access token}
export TC_TWITTER_ACCESS_TOKEN_SECRET={twitter secret access token}
```

If you don't set these, `twitclean` will quit with errors.

Then, run the program:

```shell
# deletes anything older than 30 days
twitclean -d 30 
```

The default duration is 28 days:

```shell
# deletes anything older than 28 days
twitclean 
```

Add the program to your favorite job scheduler and enjoy an automated ephemeral social media presence. As of this writing, I run the program in a Github Action, the workflow for which is included in this repository. Please note regarding this, from [GitHub's workflow docs](https://docs.github.com/en/actions/managing-workflow-runs/disabling-and-enabling-a-workflow):

> Warning: To prevent unnecessary workflow runs, scheduled workflows may be disabled automatically. When a public repository is forked, scheduled workflows are disabled by default. In a public repository, scheduled workflows are automatically disabled when no repository activity has occurred in 60 days.

So, if you want this to run for you on GitHub, you'll need to *copy* it in whole into a private repo (sorry, but GitHub will not let you make private copies of public repos) and enable the workflow there.

### WARNING

Running this program is a *destructive* act. 

*You will lose your tweets.*

*You will lose your pinned tweet.*

*You will lose your replies to other people's tweets.*

And to quote [Twitter's help site](https://help.twitter.com/en/using-twitter/missing-tweets): "If you deleted a lot Tweets in a row or ran a mass-deletion program, the Tweet count displayed on your profile may not match the true number of Tweets there are."

### Thanks To

- [Harold](https://github.com/adamdrake/harold)
- [Anaconda](https://github.com/ChimeraCoder/anaconda)
- [Logrus](https://github.com/sirupsen/logrus)

### License

This program is licensed under the terms of the [Blue Oak Model License 1.0.0](https://blueoakcouncil.org/license/1.0.0)
