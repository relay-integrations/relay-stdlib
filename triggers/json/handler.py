from relay_sdk import Interface, WebhookServer
from quart import Quart, request
import logging
import json

relay = Interface()
app = Quart('json')

logging.getLogger().setLevel(logging.INFO)

@app.route('/', methods=['POST'])
async def handler():
    r = await request.get_json(force=True)
    logging.info("Received the following webhook payload: \n%s", json.dumps(r, indent=4))
    relay.events.emit(r)
    return {}, 202, {}

if __name__ == '__main__':
    WebhookServer(app).serve_forever()
