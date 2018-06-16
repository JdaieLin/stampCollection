import requests

body={
    "user_id":3,
    "stamp_id":77309412371,
}
resp=requests.post("http://stamp.chain/api/stamp/buy",json=body)
print("resp:",resp.status_code,resp.text)