import json
from typing import Optional, List, Dict
from fastapi import FastAPI, Form
from pydantic import BaseModel
import markdown2 as md
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

origins = [
    "http://localhost",
    "http://localhost:8080",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
def getEmails(user_type:str,rn:List[int]=None)->List:
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
    elif rn:
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
    return  getEmails(user_type)    

@app.post("/md")
async def renderMD(mdb:str=Form(...)):
    return md.markdown(text=mdb)

@app.post("/send")
async def sendEmail(
    email_to:Optional[str] = Form(None),
    roll_no:Optional[str] = Form(None),
    content:str = Form(...)
    ):
    mdrender=md.markdown(text=content)
    print(roll_no)
    if roll_no is None:
        emails =  getEmails(email_to)
        return {"email":emails,"md":str(mdrender)}
    else:
        roll_nos = [int(x) for x in roll_no.split(",")]
        emails =  getEmails(None,rn=roll_nos)
        return {"email":emails,"md":str(mdrender)}
        # return {"email":"ee@aa","md":"<h2>Error</h2>"} 
