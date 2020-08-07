import json
from typing import Optional, List
from fastapi import FastAPI, Form
from pydantic import BaseModel

app = FastAPI()

# @app.get("/emails")
# async def getAll(cs:Optional[bool]=False,it:Optional[bool]=False,reps:Optional[bool]=False):
#     if cs:
#         with open('data/emails-cs.json') as f:
#             data = json.load(f)
#     elif it:
#         with open('data/emails-it.json') as f:
#             data = json.load(f)
#     else:
#         with open('data/emails-all.json') as f:
#             data = json.load(f)
#     return data

@app.get("/email/{user_type}")
async def getEmail(user_type:str,rn:Optional[List[int]]) -> List:
    if user_type == "cs":
        with open('data/emails-cs.json') as f:
            data = json.load(f)
        emails = [data[i]['Email'] for i,_ in enumerate(data)]
    elif user_type == "it":
        with open('data/emails-it.json') as f:
            data = json.load(f)
        emails = [data[i]['Email'] for i,_ in enumerate(data)]
    elif user_type == "reps":
        with open('data/emails-rep.json') as f:
            data = json.load(f)
        emails = [data[i]['Email'] for i,_ in enumerate(data)]
    elif rn :
        rolls = []
        emails = []
        for item in rn:
            rolls.append(item)
        with open('data/emails-cs.json') as f:
            jfile = json.load(f)
        for i,_ in enumerate(jfile):
            for rn in rolls:
                if jfile[i]['rn'] == rn:
                    emails.append(jfile[i]['Email'])
    else:
        with open('data/emails-all.json') as f:
            data = json.load(f)
        emails = [data[i]['Email'] for i,_ in enumerate(data)]
    # print(rn,type(rn))
    return {"emails":emails}



@app.post("/send")
async def sendEmail(
    email_type:str = Form("all"),
    roll_no:Optional[List] = Form(None),
    ):
    roll_nos = [int(x) for x in roll_no[0].split(",")]
    return email_type,roll_nos