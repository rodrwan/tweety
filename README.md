# Tweety

Tweety is a simple command line tool to manage Twitter accounts. This tools allows you to
post new tweets, retweet, and track some words via Twitter stream.

# Configuration

To use this tool you should set the following variables in your environment.
```
TWITTER_CONSUMER_KEY=...
TWITTER_CONSUMER_SECRET=...
TWITTER_ACCESS_TOKEN=...
TWITTER_ACCESS_TOKEN_SECRET=...
```

To get those values you need to create a new app in [Twitter](https://apps.twitter.com/) and generate a new access token.
### Post a new Tweet

```sh
$ tweety post "message"
```

### Sniff a word or words (separates by commas)

```sh
$ tweety sniff "#golang,#reactjs"
```

### Retweet (in progress)

soon :P
