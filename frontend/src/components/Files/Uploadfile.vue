<script>
import { AddFile, SelectFile } from '../../../wailsjs/go/main/App';
import { store } from '../../compositions/store';
import File from './File.vue';

export default{
    data() {
        return {
            selectfiles: [],
            store
        };
    },
    methods: {
        selectFile() {
            SelectFile().then((result) => {
                console.log(result);
                this.selectfiles.push(result);
            });
        }, 
        async encryptFile(path){
            //first we need to retireve the bytes from io.Reader. 
            var secretmessage = "secret message to be encrypted";
            const privateKey =  "sionethg"
            
            const signature  = EthCrypto.sign(
                privateKey, EthCrypto.hash.keccak256(secretmessage)
            ); 
            
            const payload = {
                message: secretmessage, 
                signature
            }
            
            const publicKey = "8284930234"
            
            const encrypt = await EthCrypto.encryptWithPublicKey(
                publicKey, 
                JSON.stringify(payload)
            );
            
            const encryptedString = EthCrypto.cipher.stringify(encrypt);
            console.log(encryptedString)
            
        }, 
        uploadFile(file){
            //[√]need a path             
            var newFilePath = store.filePath + "/" + file.name
            console.log("AddFile(" +newFilePath+ "," + file.path +")")
            AddFile(newFilePath, file.path).then((result) => {
                console.log(result)
            })
            
        }
    },
    components: { File }
}
</script>
<template>
	<main>
		<div class="content">
            Path: {{store.filePath}}
			<button @click="selectFile">Select file</button>
		</div>
		
		<File v-for="file in selectfiles" :key="file.path" :file="file" @click="uploadFile(file.path)" />
	</main>
</template>