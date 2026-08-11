# API collection (Posting)

This is a [Posting](https://posting.sh) collection for testing the devbook API locally.

## 1. Start the API

1. Make sure a MySQL instance matching `api/.env` is running (`DB_HOST`,
   `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`).
2. Load the schema and seed data:
   ```sh
   mysql -h 127.0.0.1 -P 3406 -u golang -p < api/sql/sql.sql
   mysql -h 127.0.0.1 -P 3406 -u golang -p devbook < api/sql/data.sql
   ```
3. From `api/`, run `go run main.go`. It listens on `http://localhost:5000`
   by default (`API_PORT` in `api/.env`), matching `$BASE_URL` below.

## 2. Run the collection

```sh
posting -c collection
```

`collection/posting.env` is loaded automatically and already points at the
local API and a seeded demo user.

## 3. Try it out

1. Run **login/Login** first. Every seeded user (see `api/sql/data.sql`)
   shares the password `123456`; the collection defaults to
   `ana.silva@example.com`. A script stores the returned JWT and user id in
   `$TOKEN` / `$USER_ID`, so you don't need to copy/paste tokens into
   other requests.
2. Run any request under `users/` or `posts/` - they all reuse `$TOKEN`
   (and `$USER_ID` where the endpoint is scoped to "your own" account).
3. Run **posts/Create Post** before the other `posts/` requests that need
   an existing post (Get/Update/Delete/Upvote/Downvote) - it stores the
   new post's id in `$POST_ID` for them to reuse.

Requests marked "Destructive" in their description (delete user, delete
post) mutate seed data - re-run `api/sql/data.sql` afterwards to reset it.
