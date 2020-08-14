import json
from typing import Optional, List, Dict
from fastapi import FastAPI, Form
from pydantic import BaseModel
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail 
import markdown2 as md
import logging
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



async def sendEmails(emails:List,sub:str,body:str):
    resp_code = []
    for email in emails:
        mail = Mail(
            from_email="athul8720@gmail.com",
            subject=sub,
            to_emails=email,
            html_content=body
            )
        logging.error(email)
        try:
            sg = SendGridAPIClient()
            resp = sg.send(mail)
            resp_code.append([email,str(resp.status_code)])
        except Exception as e:
            resp_code.append([email,None])
            logging.error(f"Could not send email to {to} due {str(e)}")
    return resp_code
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
async def getEmailfromreq(user_type:str,) -> List:
    return await getEmails(user_type)    

@app.post("/md")
async def renderMD(mdb:str=Form(...)):
    return md.markdown(text=mdb)

@app.post("/send")
async def sendEmailfromreq(
    subject:str= Form(...),
    email_to:Optional[str] = Form(None),
    roll_no:Optional[str] = Form(None),
    content:str = Form(...)
    ):
    mdrender=md.markdown(text=content)
    print(roll_no)
    if roll_no is None:
        emails = await getEmails(email_to)
        mailresp = await sendEmails(emails,subject,str(mdrender))
        return {"email":emails,"subject":subject,"md":str(mdrender),"mailresp":mailresp}
    else:
        roll_nos = [int(x) for x in roll_no.split(",")]
        emails = await getEmails(None,rn=roll_nos)
        mailresp = await sendEmails(emails,subject,str(mdrender))
    return {"email":emails,"subject":subject,"md":str(mdrender),"mailresp":mailresp}
        # return {"email":"ee@aa","md":"<h2>Error</h2>"} 
