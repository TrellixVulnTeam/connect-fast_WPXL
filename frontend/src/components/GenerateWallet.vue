<script>
import { GenerateAccount, GetAccount } from '../../wailsjs/go/main/App';

export default {
  name: "Generate Wallet ",
  data(){
      return {
          password1: "", 
          password2: "",
          error: "", 
          clickable: true
      }
  },
  watch:{
    password1(newPassword){
      if(newPassword.length < 6){
          this.error = "Your first password is too short!"
      }
      if(newPassword.length >= 6){
        this.error = null;
      }
    }, 
    password2(new2){
     
      
      if(new2.length < 6){
        this.error = "Password 2 is too short"
      }
      if(new2.length > 5){
        this.error = "Password matched"
        
        if(this.password1 == new2){
          this.error = "you can go now."
          this.clickable = true
        }
      }
    }
  },
  methods:{
     generate(){
        if(this.password1 != this.password2){
          this.error = "password don't match"
          return
        }
        
        GenerateAccount(this.password1).then(
          result => console.log(result)
        )
        
        //TODO: IMPLEMENT THE BACKEND HERE. 
        
     }, 
     retrieve(){
        GetAccount("koenkoen").then(
          restult => console.log(result)
        )
     } 
   }
};
</script>

<template>
  <main>
    <div>
       <div id="result" class="result">
         <button @click="retrieve">Read lockfile</button>
         Creation of identity withing our operating system. It is important to carefully read all the steps to understand how the process is done.
          Watch out! You need to complete the whole process before exiting the program. This will take around 10 minutes. 
        </div>
      <div id="input" class="input-box">
        <input class="input" placeholder="password" type="password" v-model="password1" />
        <input class="input" placeholder="Re-enter password" type="password" v-model="password2" />
        <button class="btn" @click="generate">Generate Wallet</button>
      </div>
      <div class="input-info">
        Password 1 is {{password1}} and password 2  is {{password2}} <br>
        {{error}}
      </div>
    </div>
  </main>
</template>

<style scoped>
main{
  flex-direction: column;
}
.result {
  line-height: 20px;
  margin: 1.5rem auto;
  max-width: 500px;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
