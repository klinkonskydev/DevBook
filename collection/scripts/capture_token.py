"""Post-response script for collection/login/login.posting.yaml.

The login endpoint replies with the raw JWT (no JSON wrapper), so this
decodes it, stores it in $TOKEN, and pulls the "userID" claim out of the
payload into $USER_ID. Every other request in the collection reuses these
two variables, so you only need to run Login once per session.
"""

import base64
import json


def on_response(response, posting):
    token = response.text.strip()

    if response.status_code != 200 or not token:
        posting.notify(
            f"Login failed ({response.status_code}): {response.text}",
            severity="error",
        )
        return

    posting.set_variable("TOKEN", token)

    try:
        payload_b64 = token.split(".")[1]
        payload_b64 += "=" * (-len(payload_b64) % 4)  # restore padding
        payload = json.loads(base64.urlsafe_b64decode(payload_b64))
        posting.set_variable("USER_ID", payload["userID"])
    except Exception as error:
        posting.notify(f"Could not read userID from token: {error}", severity="warning")
        return

    posting.notify("Logged in - $TOKEN and $USER_ID are ready to use.")
