<template>
  <div class="contract_create">
    <!--  What information do we need to add here?  -->

    <div class="class-form">
      <FormKit type="form" @submit="onSubmit" v-model="data" >
        <FormKitSchema :schema="form" />
      </FormKit>
    </div>

    <div class="createForm">
      <div class="enterForm">
      <FormKit type="form" @submit="onAddSubmit" v-model="inputFormat">
        <FormKitSchema :schema="inputForm"  />
      </FormKit>
      </div>
      
      
      <div class="previewForm">
        <FormKit
          v-for="(value, key) in dataCode"
          v-bind:key="key"
          :type="value.$formkit"
          :name="value.name"
          :label="value.label"
          :classes="{
              outer: 'mb-5 color-white',
              label: 'block mb-1 font-bold text-sm',
              inner: 'max-w-md border border-gray-400 rounded-lg mb-1 overflow-hidden focus-within:border-blue-500',
              input: 'w-full h-10 px-3 border-none text-base text-white-700 placeholder-gray-400',
              help: 'text-xs text-gray-500'
            }"
        />
      </div>
    </div>

    <div class="result">
      {{ dataCode }}
    </div>
    
    <button @click="startContract" class="btn-succes">Submit</button>
  </div>
</template>


<script>
import { ethers } from "ethers";
import { WriteFile } from "../../../wailsjs/go/main/App";
import contractInterface from "../../assets/contracts/artifacts/Contracts.json";
import { store } from "../../compositions/store";
import createForm from "./forms/createContract.json";
import createInput from "./forms/createInput.json";

export default {
  data() {
    return {
      data: {},
      inputFormat: {}, 
      store,
      title: "",
      total_reward: 0.0,
      reward_per_account: 0.0,
      min_people: 1,
      data_format: "",
      end_data: "",
      country: "",
      description: "",
      form: "",
      inputForm: "",
      dataCode: [],
    };
  },
  methods: {
    async createContract(link) {
      const mainContract = "0xF5268BbBf2D33E99c63F9056dA853dB46C8018A8";
      const privatekey =
        "0xe6fb45e2b8128e4fd18e3a4e74ea506632bc687c0c7594ad1098dfbf2cefebaf";

      const provider = new ethers.providers.JsonRpcProvider(
        import.meta.env.VITE_INFURA
      );
      const wallet = new ethers.Wallet(privatekey, provider);

      const contract = new ethers.Contract(
        mainContract,
        contractInterface.abi,
        wallet
      );

      const contractWithSigner = contract.connect(wallet);

      console.log("====  CONTRACT WITH SIGNER.... ====");

      console.log(contractWithSigner);

      try {
        const tx = await contractWithSigner.CreateNewContract(
          link,
          this.data.title,
          ethers.utils.parseEther(this.data.fee),
          true
        );

        console.log(tx);
      } catch (err) {
        console.log("ERROR", err);
      }
    },

    startContract() {
      //create a jsonStringify object to write to IPFS, and retireve back a link.

      var data = {
        data: this.data, 
        form: this.dataCode
        
      };

      const stringedJSON = JSON.stringify(data);

      WriteFile(stringedJSON, "json").then((result) => {
        console.log(result);
        this.createContract(result);
      });
    },
    
    async onSubmit() {
      console.log(this.data);
    },
    onAddSubmit(){
      this.inputFormat.$formkit = this.inputFormat.type;
      this.dataCode.push(this.inputFormat);
      this.inputFormat = {};
    },
  },
  mounted() {
    this.form = createForm;
    this.inputForm = createInput;
  },
};
</script>
<style lang="scss" scoped>
.contract_create {
  border: 1px solid #c2c2c2;
  padding: 24px;
  margin: 24px;
}

.two_section_form {
  display: flex;
  flex-direction: row;
  gap: 24px;
  justify-content: center;
}
.form_input {
  display: flex;
  flex-direction: column;
  border: none;
  margin-bottom: 8px;

  input {
    margin: 12px;
    margin-left: 0;
    padding: 8px;
    font-size: 14px;
    border-radius: 6px;
    border: none;
    color: white !important;
  }
  .formkit-input {
    color: white !important;
  }

  textarea {
    border-radius: 6px;
    padding: 8px;
    font-size: 12px;
  }
  

  .createForm {
    display: flex;
    flex-direction: row;
    gap: 12px;
  }
}
</style>