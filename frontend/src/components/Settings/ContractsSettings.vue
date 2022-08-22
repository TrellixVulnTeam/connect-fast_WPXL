<template>
    <div>
        <Title title="Contracts" />
        <div> Contracts. 
            <button @click="deployContract">Show ABI</button>
            <div class="class">
                {{abi}}
            </div>
            
            <div class="metadata">
                {{metadata}}
            </div>
        </div>
    </div>
    
</template>

<script>
import { ContractFactory, ethers } from "ethers";
import {
GetAccount
} from "../../../wailsjs/go/main/App";

import contracts from '../../assets/contracts/artifacts/Contracts.json';
import Title from "../DesignElements/Title.vue";

export default{
    data() {
        return {
            abi: "", 
            metadata: "",
            privatekey: ""
        };
    },
    methods: {
        loginAccount(password) {
            GetAccount(password).then((result) => {
                console.log(result);
                this.privatekey = result[0].private_key

            });
        },
        async deployContract() {
            this.loginAccount("password");
            this.abi = contracts.abi;
            this.metadata = contracts.data.bytecode;
            
            const provider = new ethers.providers.JsonRpcProvider(import.meta.env.VITE_INFURA)
            const wallet = new ethers.Wallet(this.privatekey, provider)
            
            
            const factory = new ContractFactory(this.abi, this.metadata, wallet)
            
            
            const tokenAddress = "0x53ED56F9AaDF3AADE8e8C7CB730AF4e6BDa90815"
            const title = "First Deployed Contract"
            const link = "no link avaible"
            
            const contract = await factory.deploy(tokenAddress, title, link)
            
            console.log(contract);
            console.log(contract.address)
        }
    },
    beforeMount(){
        this.loginAccount("password");
    },
    components: { Title }
}
</script>
