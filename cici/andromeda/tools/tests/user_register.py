import requests

body={
        "chain" : {
                "Address" : "",
                "Key" : ""
        },
        "name" : "闻一多",
        "login_id" : "balabala",
        "login_password" : "aabbcc",
        "mail" : "wenyiduo@stamp.com",
        "phone" : "13813888135",
        "id_card" : "439509199902024058",
        "coins" : 28,
        "ingots" : 134,
        },
}
resp=requests.post("http://stamp.chain/api/user/register",json=body)
print("resp:",resp.status_code,resp.text)