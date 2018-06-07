import requests

body={
    "user_id":3,
    "stamp_id":8,
}
resp=requests.post("http://localhost:8099/api/stamp/collect",json=body)
print("resp:",resp.status_code,resp.text)