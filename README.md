# mydiscordbot

From: https://discordapp.com/developers/applications/me
Adding Bots to Guilds
A URL can be generated that redirects authenticated users to the add-bot flow, by using the following format (this utilizes the OAuth2 authentication flow, without a callback URL):
```
https://discordapp.com/api/oauth2/authorize?client_id=157730590492196864&scope=bot&permissions=0
```
client_id is your bot application's ID and permissions is an integer following the permissions format.

## Commands:
- !help - send the user a list of commands.
