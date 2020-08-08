<template>
<div>
  <h1>MailMon</h1>
  <label class="label">Email: </label>
  <input type="email" v-model="email"><br/><br/>
  <textarea v-model="message" placeholder="Add Email Content in Markdwon" ></textarea>
  <button v-on:click="rendermd">Preview MD</button><br/>
  <div v-html="markdown"></div>
  <p>{{email}}</p>
  <select v-model="selected">
  <option v-for="option in options" v-bind:key="option.value" >
    {{ option.text }}
  </option>
</select>
<button v-on:click="fetchemails">Click</button>
<span>Selected: {{ selected }}</span>
<div v-for="email in emails" :key="email">
<p>{{email}}</p>
</div>
</div>
</template>

<script>
export default {
  name: 'Index',
  data() {
    return {
      email:'',
      message:'',
      selected:'it',
      options: [
      { text: 'all', value: 'all' },
      { text: 'cs', value: 'cs' },
      { text: 'it', value: 'it' }
    ],
    emails:'',
    markdown:''
    }
  },
  methods:{
    fetchemails:function(){
      console.log(this.selected)
        this.$http
    .get(`http://localhost:8000/email/${this.selected}`)
    .then(response=>(this.emails = response.data))
    .catch(error => console.log(error))
  },
    rendermd:function(){
      console.log(this.message)
      var mdForm = new FormData()
      mdForm.append('mdb',this.message)
      this.$http
      .post("http://localhost:8000/md",mdForm,{
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }})
    .then(response=>(this.markdown = response.data))
    .catch(error => console.log(error))
    }
  },
  mounted(){
    this.$http
    .get(`http://localhost:8000/email/${this.selected}`)
    .then(response=>(this.emails = response.data))
    .catch(error => console.log(error))
  }
}
</script>
<style>
textarea {
  width: 50%;
  height: 150px;
  padding: 12px 20px;
  box-sizing: border-box;
  border: 2px solid #ccc;
  border-radius: 4px;
  background-color: white;
  resize: none;
}
</style>