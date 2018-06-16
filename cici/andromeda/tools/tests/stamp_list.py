import requests

body={
    "user_id":3,
    "status":[101],
}
resp=requests.post("http://stamp.chain/api/stamp/list",json=body)
print("resp:",resp.status_code,resp.text)