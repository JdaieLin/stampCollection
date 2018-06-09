import requests

body={
    "user_id":3,
    "stamp_ids":[1,2],
}
resp=requests.post("http://localhost:8099/api/stamp/get",json=body)
print("resp:",resp.status_code,resp.text)