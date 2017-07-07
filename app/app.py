import os
from cent import Client

url = os.environ['CENTRIFUGO_URL']
secret_key = os.environ['CENTRIFUGO_SECRET']

# initialize client instance.
client = Client(url, secret_key, timeout=1)

# publish data into channel
channel = "all"
data = {"input": "test"}
client.publish(channel, data)
