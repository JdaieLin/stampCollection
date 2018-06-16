import requests

body={
    "user_id":1,
    "deal_id":2,
}
resp=requests.post("http://localhost:8099/api/deal/accept",json=body)
print("resp:",resp.status_code,resp.text)