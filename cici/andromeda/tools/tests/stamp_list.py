import requests

body={
    "user_id":3,
    "status":[101,100],
}
resp=requests.post("http://localhost:8099/api/stamp/list",json=body)
print("resp:",resp.status_code,resp.text)