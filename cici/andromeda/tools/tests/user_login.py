import requests

body={
    "login_id":"balabala",
    "login_password":"aabbcc",
}
resp=requests.post("http://stamp.chain/api/user/login",json=body)
print("resp:",resp.status_code,resp.text)