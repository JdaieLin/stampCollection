import requests

body={
    "user_id":3,
    "coins":100,
    "stamps":[
        {
            "stamp_id":1,
        }
    ]
}
resp=requests.post("http://stamp.chain/api/sweep/sync",json=body)
print("resp:",resp.status_code,resp.text)