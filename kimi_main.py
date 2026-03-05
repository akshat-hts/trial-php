
import requests, base64, time

invoke_url = "https://integrate.api.nvidia.com/v1/chat/completions"
stream = False


headers = {
  "Authorization": "Bearer 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",
  "Accept": "text/event-stream" if stream else "application/json", 
  "Content-Type": "application/json"
}

payload = {
  "model": "moonshotai/kimi-k2.5",
  "messages": [{"role":"user","content":"what is machine learning in one line"}],
  "max_tokens": 50,
  "temperature": 1.00,
  "top_p": 1.00,
  "stream": False,
  
}


print("loading..."); 
start = time.time(); 
response = requests.post(invoke_url, headers=headers, json=payload, timeout=60)
elapsed_time = time.time() - start; 
print("Response Loaded in: ", elapsed_time)

result = response.json(); 
print("Status Code:", response.status_code)
print("Raw Response:", response.text)

if response.status_code == 200:
    result = response.json()
    print(result)