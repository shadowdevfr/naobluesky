import asyncio
from twikit import Client as twt
import os
from dotenv import load_dotenv
load_dotenv()

USERNAME = os.environ['USERNAME']
EMAIL = os.environ['EMAIL']
PASSWORD = os.environ['PASSWORD']
BSKY_NAME = os.environ['BSKY_NAME']

# Initialize client
client = twt('en-US')

async def main():
    await client.login(
        auth_info_1=USERNAME ,
        auth_info_2=EMAIL,
        password=PASSWORD
    )
    
    client.save_cookies('cookies.json')

asyncio.run(main())