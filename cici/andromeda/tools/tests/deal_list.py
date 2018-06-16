import requests

body={
    "stamp_id":8,
    "deal_ids":[2],
}
resp=requests.post("http://localhost:8099/api/deal/list",json=body)
print("resp:",resp.status_code,resp.text)