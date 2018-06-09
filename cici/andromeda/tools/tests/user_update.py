import requests

body={
        "user_id":2,
        "chain" : {
                "Address" : "",
                "Key" : ""
        },
        "name" : "许三多",
        "login_id" : "jack",
        "login_password" : "aabbcc",
        "mail" : "jack@stamp.com",
        "phone" : "13813888139",
        "id_card" : "439509199902024049",
        "coins" : 28,
        "ingots" : 134,
        "books" : [
                {
                        "StampID" : 202,
                        "Status" : 100,
                        "CreateTime" : 1528293631,
                },
                {
                        "StampID" : 203,
                        "Status" : 101,
                        "CreateTime" : 1528293641,
                },
        ],
        "chests" : {
                "list" : [
                        {
                                "Hour" : 0,
                                "Ingot" : 3,
                                "Opened" : True,
                        },
                        {
                                "Hour" : 1,
                                "Ingot" : 2,
                                "Opened" : False,
                        },
                        {
                                "Hour" : 0,
                                "Ingot" : 6,
                                "Opened" : False,
                        },
                ],
        },
}
resp=requests.post("http://localhost:8099/api/user/update",json=body)
print("resp:",resp.status_code,resp.text)