import requests

body={
    "user_id":3,
    "stamp_ids":[8],
    "price":12.3,
}
resp=requests.post("http://localhost:8099/api/deal/create",json=body)
print("resp:",resp.status_code,resp.text)