import requests

body={
    "user_id":3,
}
resp=requests.post("http://stamp.chain/api/chest/list",json=body)
print("resp:",resp.status_code,resp.text)