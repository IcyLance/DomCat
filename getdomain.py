import requests
import time
import os
from dotenv import load_dotenv

load_dotenv()

NS_API_KEY = os.getenv("NS_API_KEY")

def ns_list():
    print("Getting domains from NameSilo")
    page_num = 0
    headers = {"accept":"application/json", "Content-Type":"application/json"}
    domains = []

    while True:
        page_num += 1
        list_auctions_url = f"https://www.namesilo.com/public/api/listAuctions?version=1&type=json&key={NS_API_KEY}&page={page_num}&pageSize=500&buyNow=1"

        r = requests.get(list_auctions_url, headers=headers)

        data = r.json()

        if(data['reply']['body'] == []):
            break
        else:
            for i in data['reply']['body']:
                domains.append(i)
        
        time.sleep(1)

    print(len(domains))
    # for i in domains:
    #     print(i['domain'])
    return domains

ns_list()