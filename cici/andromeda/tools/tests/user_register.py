import requests

body={
        "chain" : {
                "Address" : "",
                "Key" : ""
        },
        "name" : "闻一多",
        "login_id" : "rose",
        "login_password" : "aabbcc",
        "mail" : "rose@stamp.com",
        "phone" : "13813888133",
        "id_card" : "439509199902024048",
        "coins" : 28,
        "ingots" : 134,
        "books" : [
                {
                        "Stamp_id" : 2,
                        "Status" : 100,
                        "CreateTime" : 1528293631,
                },
                {
                        "Stamp_id" : 3,
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
resp=requests.post("http://localhost:8099/api/user/register",json=body)
print("resp:",resp.status_code,resp.text)