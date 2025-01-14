import requests
import time
import os
from dotenv import load_dotenv

load_dotenv()

# NS_API_KEY = "39d975214e561cb5ccc860eb4"
NS_API_KEY = os.getenv("NS_API_KEY")


def ns_list():
    print("Getting domains from NameSilo")
    page_num = 0
    headers = {"accept":"application/json", "Content-Type":"application/json"}
    domains = []

    while True:
        page_num += 1
        list_auctions_url = f"https://www.namesilo.com/public/api/listAuctions?version=1&type=json&key={NS_API_KEY}&page={page_num}&pageSize=500"

        r = requests.get(list_auctions_url, headers=headers)

        data = r.json()

        if(data['reply']['body'] == []):
            break
        else:
            # print(page_num)
            for i in data['reply']['body']:
                domains.append(i['domain'])
        
        time.sleep(1)

    print(len(domains))
    return domains

# def gd_list():
#     headers = {f"'Authorization': 'sso-key {GD_API_KEY}:{GD_API_SECRET}'"}

ns_list()