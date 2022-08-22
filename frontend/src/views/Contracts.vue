<script>
//BUILD IT HERE INSTEAD OF IN THe OTHER.
import { ethers } from "ethers";
import contractInterface from "../assets/contracts/artifacts/Contracts.json";
import ContractElement from "../components/ContractElement.vue";
import CreateContractVue from "../components/Contracts/CreateContract.vue";
import Loading from "../components/DesignElements/Loading.vue";
import TitleVue from "../components/DesignElements/Title.vue";
import About from "./About.vue";
const publicEnvVar = import.meta.env.VITE_MAINCONTRACT;


export default {
  name: "Contracts",
  components: { ContractElement, TitleVue, CreateContractVue, About, Loading },
  data() {
    return {
      createOn: false,
      mainContract: "0xF5268BbBf2D33E99c63F9056dA853dB46C8018A8",
      loading: true, 
      contracts: [],
    };
  },
  mounted(){
    this.getContracts();
  },
  methods: {
    async getContracts() {
      const privatekey =
        "0xe6fb45e2b8128e4fd18e3a4e74ea506632bc687c0c7594ad1098dfbf2cefebaf";

      const provider = new ethers.providers.JsonRpcProvider(
        import.meta.env.VITE_INFURA
      );
      const wallet = new ethers.Wallet(privatekey, provider);

      const contract = new ethers.Contract(
        this.mainContract,
        contractInterface.abi,
        wallet
      );


      const totalContracts = await contract.countContracts();

      let total = totalContracts.toNumber();

      for (let i = 0; i < total; i++) {
        let address = await contract.pool(i);
        this.contracts.push(address)
      }
      
      
      this.loading = false;
    },
    showCreate() {
      console.log(publicEnvVar)
      this.createOn = !this.createOn;
    },
  },
};
</script>

<template>
  <main>
    <TitleVue title="Contracts" />
    <div class="actions">
     
      <button class="btn-success" @click="showCreate">Create Contract</button>
    </div>

    <div v-if="createOn">
      <CreateContractVue />
    </div>
  
    <Loading v-if="loading" />
    <div class="grid-container" v-else>
      <ContractElement v-for="(item, index) in contracts"  :data="item" :key="index"/>
    </div>
  </main>
</template>
<style scoped>
main {
  margin: 24px;
  padding: 0;
}
.grid-container {
  display: flex;
  flex-wrap: wrap;
}
.actions {
  display: flex;
  flex-direction: row;
}
</style>