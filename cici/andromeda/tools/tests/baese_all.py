import requests

body={
}
resp=requests.post("http://localhost:8099/api/base/all",json=body)
print("resp:",resp.status_code,resp.text)