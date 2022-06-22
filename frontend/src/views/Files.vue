<template>
  <div class="files">
    <h1>
      Files System
    </h1>
  <button @click="message">
      Open Messaage
    </button>
    <div>
       <Uploadfile />
    </div>
     <h1>File Management</h1>
    
        <Folder />
   
   
  </div>
</template>

<script>
import { OpenMessage } from '../../wailsjs/go/main/App';
import Folder from '../components/Files/Folder.vue';
import Uploadfile from '../components/Uploadfile.vue';
undefined
export default {
    name: "Files",
    components: { Uploadfile, Folder }, 
    methods:{
      message(){
        OpenMessage();
      }
    }
};
</script>

<style>
.files{
  display: flex;
  flex-direction: column;
  align-content: center;
  align-items: center; 
}
</style>