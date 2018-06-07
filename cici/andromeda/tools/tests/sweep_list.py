import requests

body={
    "max":10,
}
resp=requests.post("http://localhost:8099/api/sweep/list",json=body)
print("resp:",resp.status_code,resp.text)