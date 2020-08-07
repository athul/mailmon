import json
from typing import Optional, List, Dict
from fastapi import FastAPI, Form
from pydantic import BaseModel

app = FastAPI()

async def getEmails(user_type:str,rn:List[int]=None)->List:
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
    return emails

@app.get("/email/{user_type}")
async def getEmail(user_type:str,) -> List:
    return await getEmails(user_type)    



@app.post("/send")
async def sendEmail(
    email_to:Optional[str] = Form(None),
    roll_no:Optional[List] = Form(None),
    content:str = Form(...)
    ):
    print(type(roll_no),len(roll_no))
    if roll_no[0] is "":
        emails = await getEmails(email_to)
        return emails
    else:
        roll_nos = [int(x) for x in roll_no[0].split(",")]
        emails = await getEmails(None,rn=roll_nos)
        return emails
