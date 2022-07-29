
<template>
  <main>
    <div v-if="openUI" class="modal">
      <div class="closeIcon" @click="openUI = !openUI">Close</div>
      <div class="content">
        <div class="drag-and-drop">
          <span> Drag and drop your files here. </span>
          <span class="seperateText"> ----- OR ----- </span>
          <button class="upload-button" @click="selectFile">Select file</button>
        </div>
      </div>
    </div>
    <div v-else class="upload-button" @click="openUI = !openUI">Open Modal</div>

    <File
      v-for="file in selectfiles"
      :key="file.path"
      :file="file"
      @click="uploadFile(file.path)"
    />
  </main>
</template>
<style scoped>
.open-button {
  cursor: pointer;
}
.modal {
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
  background: white;
  margin: 12px;
  width: 100%;
  border-radius: 12px;
}
.top-modal {
  border-bottom: 1px solid #acacac;
  padding: 4px 4;
  display: flex;
}
.closeIcon {
  position: relative;
  top: -10px;
  right: -270px;
  cursor: pointer;
}
.drag-and-drop {
  border: 2px dashed rgb(196, 196, 196);
  border-radius: 8px;
  background: #f4f7fd;
  width: 80%;
  height: 20%;
  padding: 24px;
  text-align: center;
  cursor: pointer;
  display: flex;
  flex-direction: column;
}
.content {
  padding: 12px;
}

.upload-button {
  border-radius: 4px;
  background-color: #5765ee;
  border: none;
  font-size: 1em;
  color: white;
  font-weight: 600;
  padding: 12px;
}

.seperateText {
  color: grey;
  font-weight: 200;
}
</style>
<script>
import { AddFile, SelectFile } from "../../../wailsjs/go/main/App";
import { store } from "../../compositions/store";
import File from "./File.vue";

export default {
  data() {
    return {
      selectfiles: [],
      store,
      openUI: false,
    };
  },
  methods: {
    selectFile() {
      SelectFile().then((result) => {
        console.log(result);
        this.selectfiles.push(result);
      });
    },
    async encryptFile(path) {
      //first we need to retireve the bytes from io.Reader.
      var secretmessage = "secret message to be encrypted";
      const privateKey = "sionethg";

      const signature = EthCrypto.sign(
        privateKey,
        EthCrypto.hash.keccak256(secretmessage)
      );

      const payload = {
        message: secretmessage,
        signature,
      };

      const publicKey = "8284930234";

      const encrypt = await EthCrypto.encryptWithPublicKey(
        publicKey,
        JSON.stringify(payload)
      );

      const encryptedString = EthCrypto.cipher.stringify(encrypt);
      console.log(encryptedString);
    },
    uploadFile(file) {
      //[√]need a path
      console.log(store.filePath);
      var newFilePath = store.filePath + "/" + file.name;
      console.log("AddFile(" + newFilePath + "," + file.path + ")");
      AddFile(newFilePath, file.path).then((result) => {
        console.log(result);
      });
    },
  },
  components: { File },
};
</script>