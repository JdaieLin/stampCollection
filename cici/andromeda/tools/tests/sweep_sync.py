import requests

body={
    "stamps":[
        {
            "stamp_id":1,
        }
    ]
}
resp=requests.post("http://localhost:8099/api/sweep/sync",json=body)
print("resp:",resp.status_code,resp.text)