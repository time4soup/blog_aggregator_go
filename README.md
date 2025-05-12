# blog_aggregator_go

This program is a blog aggregator to collect posts from RSS feeds
and be able to browse these feeds using different accounts.
Posts are requested from website RSS feeds urls, stored in an SQL database
and linked to the users who subscribe.

Requirements:
This program requires all of these programs and files to function.
* Postgres
* Go
* a config file name ".gatorconfig.json" in the root directory containing: 
{
    "db_url":"postgres://zakirk:@localhost:5432/gator",
    "current_user_name":
}
* install the gator CLI with command 'go install github.com/time4soup/blog_aggregator_go'

Using gator:
The general flow of using gator is first, registering users and adding feeds,
then aggregating feeds in the background to be able to browse them.

Here are the commands to be able to use gator effectively:
* 'gator register <name>' : registers user <name> as a valid user in the database
* 'gator users' : lists all valid users that have been registered
* 'gator login <name>' : logs in to user <name> as long as <name> is a valid user that has been registered
* 'gator addfeed <name> <url>' : Adds a new feed with name <name> connected to the RSS url <url>. Follows this feed for the current user
* 'gator feeds' : lists all added feeds
* 'gator follow <url>' : follows the feed to the url <url> for current user. feed for this url must already be added
* 'gator unfollow <url>' : unfollows the feed to the url <url> for the current user. feed for this url must already be added and followed
* 'gator following' : lists all feeds the current user is following
* 'gator agg <time>' : collects posts from all added feeds. searches for new posts every <time> interval
* 'gator browse <amount>' : prints most recent <amount> posts collected by agg command. only prints posts by feeds that the current user follows
* 'gator reset' : WARNING! USE CAREFULLY! deletes all information for gator including, registered users, added feeds, followed feeds, and posts collected