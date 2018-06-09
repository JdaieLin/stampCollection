import requests

body={
    "login_id":"balabala",
    "login_password":"aabbcc",
}
resp=requests.post("http://localhost:8099/api/user/login",json=body)
print("resp:",resp.status_code,resp.text)