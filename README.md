RSS feed aggregator in Go!

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

It's "Gator", you know, because aggreGATOR 🐊.

CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

There's no server ( other than the database ), it's only intended for local use, but that doesn't mean we can't have multiplayer functionality on a single device

The `JSON` file to keep track of two things:

- Who is currently logged in
- The connection credentials for the PostgreSQL database
- `current_username` is set by the application, so create a config.json with only `db_url`
- THe config file should go in `~/.config/gator.json`

```json
{
  "db_url": "connection_string_goes_here",
  "current_username": "username_goes_here"
}
```
