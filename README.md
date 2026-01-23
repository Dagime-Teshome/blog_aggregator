<!-- Improved compatibility of back to top link -->

<a id="readme-top"></a>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#built-with">Built With</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#configuration">Configuration</a></li>
    <li><a href="#commands">Commands</a></li>
    <li><a href="#example-workflow">Example Workflow</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

`gator` is a CLI tool that lets you register users, follow RSS feeds, aggregate posts on a schedule, and discover content.  
Authentication is enforced via middleware, so some commands only run for logged-in users.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

Follow these instructions to get a local copy running.

### Prerequisites

You must have **Go** and **PostgreSQL** installed.

```bash
go version
psql --version
git clone https://github.com/Dagime-Teshome/blog_aggregator.git
cd blog_aggregator
go install github.com/Dagime-Teshome/blog_aggregator@latest
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
## Configuration

Create a config file in your home directory:

```bash
    ~/.gator.config
```

content should have the following

```json
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "username_goes_here"
```

<p align="right">(<a href="#readme-top">back to top</a>)</p> <!-- COMMANDS -->
## Commands

to run commands:

```bash
    gator <command>
```

### `login`

Logs in an existing user. User must already be registered.

**Example:**

```bash
gator login John
```

---

### `register`

Registers and logs in a new user.

- Takes the name of the user as the first argument
- Automatically logs in the user after registration

**Example:**

```bash
gator register John
```

---

### `reset`

Resets the system's tables.

- Deletes all users
- Deletes all feeds
- Deletes all feed follows

**Example:**

```bash
gator reset
```

---

### `users`

Shows a list of users registered in the system.

- Doesn't need any arguments
- Shows the current user with an `*` next to it

**Example:**

```bash
gator users
```

**Output:**

```
John
James
*Rebecca  (currently logged in)
```

---

### `agg`

Aggregates feeds at specified intervals.

- Takes one argument: the interval at which to check feeds for updates
- Requests feeds, parses them, and saves posts in the database
- Accepts time-parsable formats: `1s` (seconds), `1m` (minutes), `1h` (hours), `1ms` (milliseconds)

**Example:**

```bash
gator agg 10m
```

This will aggregate data every 10 minutes.

---

### `addfeed`

Adds a feed to track and aggregate.

- Automatically makes the user who created it follow the feed
- Takes two arguments: feed name and URL

**Example:**

```bash
gator addfeed TechCrunch https://techcrunch.com/feed/
```

---

### `feeds`

Shows the list of feeds in the system.

- Doesn't need any arguments

**Example:**

```bash
gator feeds
```

---

### `follow`

Allows a user to follow a specified feed.

- Takes one argument: the URL of the feed

**Example:**

```bash
gator follow https://techcrunch.com/feed/
```

---

### `following`

Returns a list of feeds the logged-in user is following.

- Uses the logged-in middleware to identify the current user
- Doesn't need any arguments

**Example:**

```bash
gator following
```

---

### `unfollow`

Allows the currently logged-in user to unfollow a feed.

- Takes one argument: the URL of the feed to unfollow

**Example:**

```bash
gator unfollow https://techcrunch.com/feed/
```

---

### `discover`

Returns aggregated posts from feeds the user follows.

- Shows a list of post titles and feed names
- Takes one optional argument to limit the number of posts (defaults to 2)

**Example:**

```bash
gator discover 5
```

This returns 5 posts from feeds that the user follows.

## Example Workflow

Here's a complete workflow demonstrating how to use the Gator RSS aggregator:

```bash
# 1. Register a new user named John
gator register John

# 2. Add a feed to track
gator addfeed TechCrunch https://techcrunch.com/feed/

# 3. Add another feed
gator addfeed TheVerge https://www.theverge.com/rss/index.xml

# 4. Check which feeds are available
gator feeds

# 5. See which feeds you're following
gator following

# 6. Start aggregating feeds every 10 minutes
gator agg 10m

# 7. Discover the latest posts (limit to 5)
gator discover 5

# 8. Follow an existing feed
gator follow https://techcrunch.com/feed/

# 9. Unfollow a feed
gator unfollow https://www.theverge.com/rss/index.xml

# 10. Check all registered users
gator users

# 11. Login as a different user
gator login James

# 12. Reset the entire system (use with caution!)
gator reset
```

### Expected Output Flow

1. **After registration:** User "John" is registered and logged in
2. **After addfeed:** Feed is added and John automatically follows it
3. **After agg:** Feeds are scraped every 10 minutes and posts are saved
4. **After discover:** Shows 5 recent posts from followed feeds with titles and feed names
