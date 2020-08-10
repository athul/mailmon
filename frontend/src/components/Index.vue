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
        validation="required|max:200,length"
        validation-name="Email Content"
        error-behavior="live"
        placeholder="Enter content in Markdown Format. Max 200 Characters"
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
        :options="{
          all: 'All Students',
          cs: 'Students in CS Dept.',
          it: 'Students in IT',
          reps: 'Email Class Reps',
          roll: 'Email to Roll Numbers',
        }"
        type="select"
        placeholder="Select an option"
        label="Select whom to send Emails to"
      />
      <div v-if="selected === 'roll'">
        <FormulateInput
          v-model="rn"
          type="text"
          validation="required"
          validation-name="Roll Number"
          error-behavior="live"
          placeholder ="Enter Roll Numbers of Students like 1,24,48,etc..."
        />
      </div>
      
    <FormulateInput type="submit" name="Send Emails" />
      
    </FormulateForm>
    <div class="render">
    <div v-if="sub !== ''">Subject: <em>{{sub}}</em></div>
    <div v-if="premd == true" v-html="markdown">
    </div>
    </div>
    <br />
    <div v-for="email in emails" :key="email">
      <div>{{ email }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Index",
  data() {
    return {
      email: "",
      md: "",
      selected: "",
      emails: "",
      markdown: "",
      rn: "",
      sub: "",
      premd: false,
    };
  },
  methods: {
    rendermd: function () {
      console.log(this.md, this.premd);
      var mdForm = new FormData();
      mdForm.append("mdb", this.md);
      this.$http
        .post("http://localhost:8000/md", mdForm, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        })
        .then((response) => (this.markdown = response.data, this.premd = true))
        .catch((error) => console.log(error));
    },
    sndems: function () {
      var mainForm = new FormData();
      mainForm.append("email_to", this.selected);
      mainForm.append("roll_no", this.rn);
      mainForm.append("content", this.md);
      mainForm.append("subject", this.sub);
      this.$http
        .post("http://localhost:8000/send", mainForm, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        })
        .then(
          (response) => (
            (this.markdown = response.data.md),
            (this.emails = response.data.email),
            console.log(response.data.mailresp)
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
