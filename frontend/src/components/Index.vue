<template>
  <div align="center">
    <h1>MailMon</h1>
    <FormulateForm @submit="sndems">
      <FormulateInput
        type="text"
        v-model="sub"
        name="Subject"
        label="Subject of the Email"
        validation="required|max:50,length"
        error-behaviour="live"
        validation-name="Subject"
        placeholder="Enter Subject"
      />
      <FormulateInput
        type="textarea"
        v-model="md"
        label="Email Content"
        validation="required"
        validation-name="Email Content"
        error-behavior="live"
        placeholder="Enter content in Markdown Format."
      />
      <div v-if="md !== ''">
      <FormulateInput
        type="button"
        v-model="premd"
        @click="rendermd"
        label="Preview MarkDown"
      />
      </div>
      <br/>
      <FormulateInput
        v-model="selected"
        type="select"
        placeholder="Select an option"
        label="Select whom to send Emails to"
      />
    <FormulateInput type="submit" name="Send Emails" />
      
    </FormulateForm>
    <div class="render">
    <div v-if="elapsed!==''">Time Took to send Emails: {{elapsed}}</div>
    <div v-if="sub !== ''">Subject: <em>{{sub}}</em></div>
    <div v-if="premd == true" v-html="markdown">
    </div>
    </div>
    <div>{{emresp}}</div>
    <br />
    <!-- <div v-for="email in emails" :key="email">
      <div>{{ email }}</div>
    </div> -->
  </div>
</template>

<script>
export default {
  name: "Index",
  data() {
    return {
      email: "",
      md: "",
      emails: "",
      markdown: "",
      rn: "",
      sub: "",
      premd: false,
      emresp:'',
      elapsed:''
    };
  },
  methods: {
    rendermd: function () {
      console.log(this.md, this.premd);
      var mdForm = new FormData();
      mdForm.append("mdb", this.md);
      this.$http
        .post("http://localhost:8080/md", mdForm, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        })
        .then((response) => (this.markdown = response.data, this.premd = true))
        .catch((error) => console.log(error));
    },
    sndems: function () {
      var mainForm = new FormData();
      mainForm.append("roll_no", this.rn);
      mainForm.append("content", this.md);
      mainForm.append("subject", this.sub);
      this.$http
        .post("http://localhost:8080/send", mainForm, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        })
        .then(
          (response) => (
            (this.markdown = response.data.md),
            (this.emails = response.data.email),
            (this.emresp = response.data.mailresp),
            (this.elapsed = response.data.elapsed),
            console.log(this.emresp)
          )
        )
        .catch((error) => console.log(error));
    },
  },
};
</script>
<style>
textarea {
  width: 75%;
  height: 250px;
  padding: 12px 20px;
  box-sizing: border-box;
  border: 2px solid #ccc;
  border-radius: 4px;
  background-color: white;
  resize: none;
}
.render{
  padding: 12px 20px;
  border:2px solid #ccc;
  border-radius: 4px;
  align-items: right;
}
</style>
