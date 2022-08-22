
<template>
  <div class="content">
    <TitleVue title="Dashboard" />

    <div v-if="loginStep == 0" class="step0">
      <span>
        We need to create a password in order complete the first creation of a
        password.
      </span>

      <input v-model="password1" />
      <input v-model="password2" />

      <button @click="registerAccount">Create Account</button>
    </div>
    <div v-if="store.address.length < 1" class="step1">
      <span> Login in here to your accounts. </span>
      <input v-model="password1" type="password" />

      <button @click="loginAccount">Unlock</button>

    </div>

    <div v-if="store.address.length > 1" class="step2">
      <span> Account: {{store.address}}</span>

      <button style="btn-danger" @click="logoutAccount()">Log out</button>
    </div>
  </div>
</template>
<script>
import {
CreateAccount,
GetAccount,
GetAccountExist
} from "../../../wailsjs/go/main/App";
import { store } from "../../compositions/store";
import TitleVue from "../DesignElements/Title.vue";
import GenerateWallet from "../GenerateWallet.vue";

import { ethers } from "ethers";

export default {
  name: "Ethereum Settings",
  components: { GenerateWallet, TitleVue },
  data() {
    return {
      password1: "",
      password2: "",
      loginStep: 0,
      privateKey: "",
      store,
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
    loginAccount() {
      console.log("loggin to the account");
      GetAccount(this.password1).then((result) => {
        console.log(result);

        console.log(result[0].address);

        this.store.updateAddress(result[0].address);
        this.privateKey = result[0].private_key;
      });
    },
    registerAccount() {
      if (this.password1 === this.password2) {
        CreateAccount(this.password1).then((result) => {
          console.log(result);
          this.loginStep = 2;
        });
      } else {
        console.log("password don't match");
      }
    },
    async SendTransaction() {
      const provider = new ethers.providers.JsonRpcProvider(
       import.meta.env.VITE_INFURA
      );
      const wallet = new ethers.Wallet(this.privateKey, provider);

      const account2 = "0xfFf09621F09CAa2C939386b688e62e5BE19D2D56";

      const tx = await wallet.sendTransaction({
        to: account2,
        value: ethers.utils.parseEther("0.1"),
      });

      console.log(tx);
    },
    logoutAccount(){
      store.updateAddress("");
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
.content input {
  padding: 12px;
}

.content button {
  background: #21212a;
  color: #afafb2;
  border: none;
  padding: 8px;
  font-size: 18px;
  border-radius: 8px;
}

.content div {
  display: flex;
  flex-direction: column;
  width: 50%;
  align-self: center;
  gap: 24px;
  margin-left: 40px;
}
</style>