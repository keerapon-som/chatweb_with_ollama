import requests
import json

def TestGenerate():
    url = "http://localhost:8080/api/generate"
    headers = {
        "Content-Type": "application/json"
    }

    data = {
        "model": "llama3.2:1b",
        "prompt": "ฉันมาละ"
    }
    response = requests.post(url, headers=headers, json=data, stream=True)
    resporder = ""
    for line in response.iter_lines():
        if line:
            decoded_line = line.decode('utf-8')
            jsondata = json.loads(decoded_line)
            resporder = resporder + jsondata['response']

    print(resporder)

def TestChatGenerate():
    url = "http://localhost:8080/api/chat"
    headers = {
        "Content-Type": "application/json"
    }
    
    data = {
        "model": "llama3.2:1b",
        "messages":[
            {
                "role": "user",
                "content": "1 + 1 เท่ากับเท่าไหร่"
            },
            {
                "role": "assistant",
                "content": "5"
            },
            {
                "role": "user",
                "content": "คุณตอบผิดหรือป่าว เมื่อกี้คุณตอบว่าอะไรนะ แล้วคำตอบจริงๆมันคืออะไร ?"
            },
        ],
    }
    response = requests.post(url, headers=headers, json=data, stream=True)
    resporder = ""
    for line in response.iter_lines():
        if line:
            decoded_line = line.decode('utf-8')
            jsondata = json.loads(decoded_line)
            resporder = resporder + jsondata['message']["content"]

    print(resporder)

TestChatGenerate()