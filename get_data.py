import sys
import requests
from bs4 import BeautifulSoup
import pandas as pd
import json

Charting_Link = "https://chartink.com/screener/"
Charting_url = 'https://chartink.com/screener/process'

def GetDataFromChartink(payload):
    payload = {'scan_clause': payload}

    with requests.Session() as s:
        # Send GET request to get the csrf-token
        r = s.get(Charting_Link)
        soup = BeautifulSoup(r.text, "html.parser")
        csrf = soup.select_one("[name='csrf-token']")['content']
        s.headers['x-csrf-token'] = csrf

        # Send POST request with payload
        r = s.post(Charting_url, data=payload)

        df = pd.DataFrame()

        # Process the data and populate the DataFrame
        for item in r.json()['data']:
            if len(item) > 0:
                df = pd.concat([df, pd.DataFrame.from_dict(item, orient='index').T], ignore_index=True)

        # Convert the DataFrame to a JSON string and return it
        return df.to_json(orient='records')

# Main function to handle the script execution
if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Error: Please provide the scan_clause.")
        sys.exit(1)

    # Get the scan_clause from command-line arguments
    scan_clause = sys.argv[1]
    result = GetDataFromChartink(scan_clause)
    print(result)  # Output the JSON string