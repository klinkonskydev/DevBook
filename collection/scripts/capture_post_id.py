"""Post-response script for collection/posts/create-post.posting.yaml.

Stores the newly created post's id in $POST_ID so the other post requests
(get, update, delete, upvote, downvote) can target it without any manual
copy/pasting.
"""


def on_response(response, posting):
    if response.status_code != 201:
        posting.notify(f"Create post failed ({response.status_code})", severity="error")
        return

    try:
        post = response.json()
        posting.set_variable("POST_ID", post["id"])
        posting.notify(f"Created post #{post['id']} - $POST_ID is ready to use.")
    except Exception as error:
        posting.notify(f"Could not read post id from response: {error}", severity="warning")
