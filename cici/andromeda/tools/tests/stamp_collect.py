import requests

body={
    "user_id":3,
    "stamp_id":197568495852,
}
resp=requests.post("http://stamp.chain/api/stamp/collect",json=body)
print("resp:",resp.status_code,resp.text)