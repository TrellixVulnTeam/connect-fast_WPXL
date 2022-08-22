<template>
  <div>
    <button @click="getAllData">Get all data</button>
    <div>
      {{ profiles }}
    </div>

    <div v-if="profiles.length > 1">
      <button @click="readAllFiles">File reading...</button>
    </div>
  </div>
</template>

<script>
import { ethers } from "ethers";
import { DecryptData, GetAccount } from "../../../wailsjs/go/main/App";
import contractInterface from "../../assets/contracts/artifacts/ShareContract.json";
import { store } from "../../compositions/store";

export default {
  data() {
    return {
      store,
      privateKey: null,
      password: "password",
      profiles: [],
      data: "",
    };
  },
  setup() {},
  methods: {
    async getAllData() {
      console.log();
      const provider = new ethers.providers.JsonRpcProvider(
        import.meta.env.VITE_INFURA
      );

      await GetAccount(this.password).then((result) => {
        console.log(result);

        console.log(result[0].address);

        this.privateKey = result[0].private_key;
      });

      const wallet = new ethers.Wallet(this.privateKey, provider);

      const someContract = new ethers.Contract(
        this.$parent.contract,
        contractInterface.abi,
        wallet
      );

      this.profiles = await someContract.getProfiles();
    },

    async readAllFiles() {
      for (var i = 0; i < this.profiles.length; i++) {
        console.log(this.profiles[i][3]);

        await DecryptData(this.profiles[i][3]).then((result) => {
          this.data = result;
            console.log(result)
        });
        
       
        
      }
    },
  
  
    
  },
};
</script>