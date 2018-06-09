import requests

body={
    "user_id":2,
    "chests":[
        {
            "hour":22,
            "opened":True,
        },
        {
            "hour":9,
            "opened":True,
        }
    ]
}
resp=requests.post("http://localhost:8099/api/chest/sync",json=body)
print("resp:",resp.status_code,resp.text)