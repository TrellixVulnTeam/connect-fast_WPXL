<script>
import { ethers } from "ethers";
import { EncryptToSmart, GetAccount, GetIPFS } from "../../wailsjs/go/main/App";
import contractInterface from "../assets/contracts/artifacts/ShareContract.json";
import OwnerContractDetails from "../components/Contracts/OwnerContractDetails.vue";
import Password from "../components/DesignElements/Modals/Password.vue";
import ProgressBar from "../components/DesignElements/ProgressBar.vue";
import Title from '../components/DesignElements/Title.vue';


export default {
  name: "Contract Detail",
  components: { ProgressBar, Password, Title, Title, OwnerContractDetails },
  data() {
    return {
      registered: false,
      menuSelect: "standard",
      result: "nor esult",
      contract: null,
      show: false, 
      password: "",
      privateKey: null, 
      inputData: {},
      newLink: "",
      entryForm: false,
      details: {
        owner: "",
        parentContract: "",
        link: "",
        title: "DefaultTitle",
      },
      data: {"data": {
        "fee": "undefined", 
        "endDate": undefined, 
        "description": "undefined"
      }},
      formData: {},
    };
  },
  props: {
    contractLink: {
      type: String,
      default: "nothing found",
    },
  },

  mounted() {
    this.contract = this.$route.params.id;
    this.retrieveContract(this.contract);
    //now we want to call something that retireve the contrat.
  },

  methods: {
    async getContractCreatorKey() {
        const provider = new ethers.providers.JsonRpcProvider(
         import.meta.env.VITE_INFURA
        );

        const contractAddress = "0xf0bbE607A76089FD7A716b8E55e9E349ad8A956C";
        const contract = new ethers.Contract(
          contractAddress,
          contractInterface.abi,
          provider
        );

        console.log(contract);
        //const transactionDeploy = contract.deployTransaction()

        console.log(provider);

        const tx = await provider.getTransaction(
          "0xf20dc3ef4c826e205e2ff5ffa0c33f32b1abcae4fbe210a17249e704a3b15a5d"
        );
        const expandedSig = {
          r: tx.r,
          s: tx.s,
          v: tx.v,
        };
        
        console.log(expandedSig);
        const signature = ethers.utils.joinSignature(expandedSig);
        const txData = {
          gasPrice: tx.gasPrice,
          gasLimit: tx.gasLimit,
          value: tx.value,
          nonce: tx.nonce,
          data: tx.data,
          chainId: tx.chainId,
          to: tx.to, // you might need to include this if it's a regular tx and not simply a contract deployment
        };
        const rsTx = await ethers.utils.resolveProperties(txData);
        const raw = ethers.utils.serializeTransaction(rsTx); // returns RLP encoded tx
        const msgHash = ethers.utils.keccak256(raw); // as specified by ECDSA
        const msgBytes = ethers.utils.arrayify(msgHash); // create binary hash
        const recoveredPubKey = ethers.utils.recoverPublicKey(
          msgBytes,
          signature
        );
        const recoveredAddress = ethers.utils.recoverAddress(msgBytes, signature);

        console.log(recoveredPubKey);
        
        return recoveredPubKey;
    },

    async retrieveContract(contractAddress) {
      const provider = new ethers.providers.JsonRpcProvider(
        import.meta.env.VITE_INFURA
      );

      const contract = new ethers.Contract(
        contractAddress,
        contractInterface.abi,
        provider
      );

      const contractDetails = await contract.getContractDetails();
      if (contractDetails.length > 0) {
        this.details.owner = contractDetails[0];
        this.details.link = contractDetails[2];
        this.details.title = contractDetails[3];
        this.details.parentContract = contractDetails[1];
      }

      //check if
      if (this.details.link.length > 0) {
        this.retrieveIPFS();
      }
    },

    async retrieveIPFS() {
      GetIPFS(this.details.link).then((result) => {
        this.data = JSON.parse(result);
        if (this.data.form.length > 0) {
          this.formData = this.data.form;
          this.entryForm = true;
        }
      });
    },

    async submitDataToContract() {
      console.log(this.inputData);
      
        const stringedJSON  = JSON.stringify(this.inputData); 
        console.log(stringedJSON);
        
        //now we need to make sure that we get the publickey. 
        var pubKey = await this.getContractCreatorKey();
        console.log(pubKey)
        
        //nowe we want to create file to submit to there. 
        EncryptToSmart(pubKey, stringedJSON).then(result => {console.log(result)
           this.show = true;
          this.newLink = result;
        });
    },
    
    async submitToContract(password){
      console.log("submittingToContract now");
      console.log(this.password)
       const provider = new ethers.providers.JsonRpcProvider(
        import.meta.env.VITE_INFURA
      );
      

      if(password.length > 1){
        
         await GetAccount(password).then((result) => {
            console.log(result);

            console.log(result[0].address);

            this.privateKey = result[0].private_key;
            
          
          });
      }
      
      console.log(this.privateKey)
       const wallet = new ethers.Wallet(this.privateKey, provider);

      const submitContract = new ethers.Contract(
        this.contract,
        contractInterface.abi,
        wallet
      );
      
      console.log(submitContract)
      
      try{
        const tx = await submitContract.userSubmission(this.newLink, false)
        console.log(tx)
      }catch(err){
        console.log(err)
      }
    },
   
    },
    
    
      
};
</script>


<template>
  <main class="content">
    <div class="top-content">
      <div class="back" @click="$router.push('/contracts')">
        <img
          src="../assets/icons/backward.svg"
          class="back-icon"
          alt="go back"
        />
      </div>
     
     <Title :title="details.title" />
    </div>


    <!-- We want to build a switch here for the view, that provides the owner to look atmmore details. -->
    
    <div class="submenu">
      <div class="menu-item" @click="menuSelect ='standard'">
        Standard details. 
      </div>
      <div class="menu-item" @click="menuSelect ='owner'">
        Onwer options. 
      </div>
    </div>
    
    <div v-if="menuSelect === 'standard'">
      <div class="basic-info">
        <div class="extra-details">
          <table class="additional">
            <tr>
              <th>Owner</th>
              <td>{{ details.owner }}</td>
            </tr>
            <tr>
              <th>Parent Contract</th>
              <td>{{ details.parentContract }}</td>
            </tr>
            <tr>
              <th>Contract Address</th>
              <td>{{ contract }}</td>
            </tr>
            <tr>
              <th>IPFS link</th>
              <td @click="retrieveIPFS">{{ details.link }}</td>
            </tr>
            <tr>
              <th>Fee</th>
              <td>{{ data.data.fee }}</td>
            </tr>
            <tr>
              <th>End date</th>
              <td>{{ data.data.endDate }}</td>
            </tr>
          </table>
        </div>
        <div class="description">
          {{ data.data.description }}
        </div>
      </div>
      <div class="box-extra"></div>

      <div class="registerForm">
        <h4>Data form for submission.</h4>
        <FormKit
          type="form"
          @submit="submitDataToContract"
          v-model="inputData"
          v-if="entryForm"
        >
          <FormKitSchema :schema="formData" />
        </FormKit>
      </div>
    </div>
    <div v-if="menuSelect === 'owner'">
      <OwnerContractDetails />
    </div>
    <Password v-if="show" />
  </main>
</template>
<style >
main {
  display: flex;
  padding: 32px;
  margin-left: 32px;
  flex-direction: column;
  text-align: left;
}
.back-icon {
  height: 16px;
  width: 16px;
}
.top-content {
  display: flex;
  align-items: center;
  align-self: center;
}

.back {
  padding: 24px;
  cursor: pointer;
}
.title {
  font-size: 32px;
  font-weight: 300;
}

.basic-info {
  display: flex;
  flex-direction: row;
}

.additional {
  font-size: 14px;
}
td,
th {
  padding-right: 24px;
}
.box-extra {
  gap: 12px;
}
.registerForm {
  margin-top: 24px;
  align-items: center;
}

.submenu{
  display: flex;
  gap: 20px;
  margin-bottom: 24px;
}

.menu-item{
  border: 1px solid green; 
  padding: 12px;
  border-radius: 8px;  
}

.menu-item:hover{
  cursor: pointer;
  background: green;
}

</style>