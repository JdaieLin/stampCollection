import requests

body={
    "user_id":2,
    "serial_id":1,
    "serial_order":5,
}
resp=requests.post("http://localhost:8099/api/slot/unlock",json=body)
print("resp:",resp.status_code,resp.text)