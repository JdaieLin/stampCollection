import requests

body={
    "user_id":2,
}
resp=requests.post("http://localhost:8099/api/slot/list",json=body)
print("resp:",resp.status_code,resp.text)