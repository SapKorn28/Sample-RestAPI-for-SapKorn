<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <div class="form">
      <input type="text" v-model="userId" placeholder="ID" />
      <br />
      <input type="text" v-model="passWord" placeholder="Password" />
      <br />
    </div>
    <button @click="getRequest">GetButton</button>
    <button @click="postRequest">PostButton</button>
    <button @click="putRequest">PutButton</button>
    <button @click="deleteRequest">DeleteButton</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "HelloWorld",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
      userId: "",
      passWord: ""
    };
  },
  methods: {
    getRequest: async function() {
      await axios
        .get("http://localhost:7070/api")
        .then(res => {
          console.log("get: ", res);
          this.msg = res.data;
        })
        .catch(error => {
          console.log("error: ", error);
        });
    },
    postRequest: async function() {
      if (this.userId === "" && this.passWord == "") {
        return (this.msg = "Please Input Data!");
      }
      await axios
        .post("http://localhost:7070/api", {
          id: this.userId,
          pass: this.passWord
        })
        .then(res => {
          console.log("post: ", res);
          this.msg = "post method: Hello " + res.data.id + "!";
        })
        .catch(error => {
          console.log("error: ", error);
        });
    },
    putRequest: async function() {
      await axios
        .put("http://localhost:7070/api", {
          id: this.userId,
          pass: this.passWord
        })
        .then(res => {
          console.log("put: ", res);
          this.msg = "put method: OK " + res.data.id + "!";
        })
        .catch(error => {
          console.log("error: ", error);
        });
    },
    deleteRequest: async function() {
      await axios
        .delete("http://localhost:7070/api", {
          id: this.userId,
          pass: this.passWord
        })
        .then(res => {
          console.log("delete!");
          this.msg = "Delete Complete!";
        })
        .catch(error => {
          console.log("error: ", error);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
input {
  margin: 10px 0;
  padding: 10px;
}
</style>
