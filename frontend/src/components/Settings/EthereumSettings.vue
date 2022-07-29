
<template>
  <div class="content">
    <TitleVue title="Ethereum" />

    <div v-if="loginStep == 0" class="step0">
      <span>
        We need to create a password in order complete the first creation of a
        password.
      </span>

      <input v-model="password1" />
      <input v-model="password2" />
      
      <button @click="registerAccount">
        Create Account
      </button>
    </div>
    <div v-if="loginStep == 1" class="step1">
        <span>
            Login in here to your accounts. 
        </span>
        <input v-model="password1" />
        
        <button @click="loginAccount">
            Unlock
        </button>
    </div>
    
    <div v-if="loginStep == 2" class="step2">
        <span>
            Account: 
        </span>
        
        <span>
            Balance: 
        </span>
    </div>
    </div>
</template>
<script>
import { CreateAccount, GetAccountExist } from "../../../wailsjs/go/main/App";
import TitleVue from "../DesignElements/Title.vue";
import GenerateWallet from "../GenerateWallet.vue";

export default {
  name: "Ethereum Settings",
  components: { GenerateWallet, TitleVue },
  data() {
    return {
        password1: "", 
        password2: "",
      loginStep: 0,
    };
  },
  methods: {
    checkAccount() {
      GetAccountExist().then((result) => {
        console.log(result);
        if (result) {
          this.loginStep = 1;
        }
      });
    },
    loginAccount(){
        console.log("loggin to the account")
    },
    registerAccount(){
        if(this.password1 === this.password2){
        CreateAccount(this.password1).then((result) =>
        {
            console.log(result)
            this.loginStep = 2;
        })
    }
    else{
            console.log("password don't match")
        }
    }
  },
  beforeMount() {
    this.checkAccount();
  },
  
};
</script>

<style scoped>

.step0 {
  display: flex;
  flex-direction: column;
  width: 50%;
  align-self: center;
  gap: 24px;
}
.step0 input{
    padding: 12px;
    
}
</style>