<template>
<div>
  <h1>MailMon</h1>
<FormulateForm @submit="sndems">
  <FormulateInput
  type="text"
  v-model="sub"
  name="Subject"
  label="Subject of the Email"
  />
  <FormulateInput
  type="textarea"
  v-model="md"
  label="Email Content"
  validation="required|max:200,length"
  validation-name="Email Content"
  error-behavior="live"
  placeholder="Enter content in Markdown Format. Max 200 Characters"
/>
<p v-html="sub"></p>
  <FormulateInput
  v-model="selected"
  :options="{all: 'All Students', cs: 'Students in CS Dept.', it: 'Students in IT', reps: 'Email Class Reps', roll:'Email to specific Roll Numbers'}"
  type="select"
  placeholder="Select an option"
  label="Select whom to send Emails to"
/>
<div v-if="selected === 'roll'">
  <FormulateInput
  v-model="rn"
    type="text"
  />
</div>
<FormulateInput
    type="submit"
    name="Preview MD"
  />
</FormulateForm>
<div v-html="markdown"></div>
<br/>
<div v-for="email in emails" :key="email">
<div>{{email}}</div>
</div>
</div>
</template>

<script>
export default {
  name: 'Index',
  data() {
    return {
      email:'',
      md:'',
      selected:'it',
      emails:'',
      markdown:'',
      rn:'',
      sub:''
    }
  },
  methods:{
    rendermd:function(){
      console.log(this.md)
      var mdForm = new FormData()
      mdForm.append('mdb',this.md)
      this.$http
      .post("http://localhost:8000/md",mdForm,{
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }})
    .then(response=>(this.markdown = response.data))
    .catch(error => console.log(error))
    },
    sndems:function(){
      var mainForm= new FormData()
      mainForm.append("email_to",this.selected)
      mainForm.append("roll_no",this.rn)
      mainForm.append("content",this.md)
      this.$http
      .post("http://localhost:8000/send",mainForm,{
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }})
    .then(response=>(this.markdown = response.data.md,this.emails=response.data.email,console.log(this.emails)))
    .catch(error => console.log(error))
    
    }
  },
}
</script>
<style>
textarea {
  width: 50%;
  height: 200px;
  padding: 12px 20px;
  box-sizing: border-box;
  border: 2px solid #ccc;
  border-radius: 4px;
  background-color: white;
  resize: none;
}
</style>