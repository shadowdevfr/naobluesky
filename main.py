import asyncio
from twikit import Client as twt
from blueskysocial import Client, Post, Image
import hashlib
import schedule
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
    
    file = open("lasttweet.txt","r")

    user = await client.get_user_by_screen_name("Naolib_Direct")
    tweets = await client.get_user_tweets(user.id, 'Tweets')    

    if "RT @" in tweets[0].text:
        file.close()
        exit()
        
    md5 = hashlib.md5(tweets[0].text.encode('utf-8')).hexdigest()    
    
    if file.read() == md5:
        print("on tweet pas c'est le meme")
        file.close()
        exit()
    print(md5)
    
    bsky = Client()
    bsky.authenticate(BSKY_NAME, PASSWORD)
    post = Post(tweets[0].text)
    bsky.post(post)
    
    file.close()
    file = open("lasttweet.txt","w")
    file.write(md5)
    file.close()
    
def run():
    asyncio.run(main())
    
run()
schedule.every(5).minutes.do(run)