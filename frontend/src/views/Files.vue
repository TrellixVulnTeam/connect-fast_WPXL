<template>
  <main>
    <Title title="Files" />
    <Folder />
   
    <AddContent />
  </main>
  
</template>

<script>
import { OpenMessage } from '../../wailsjs/go/main/App';
import Title from '../components/DesignElements/Title.vue';
import AddContent from '../components/Files/AddContent.vue';
import Folder from '../components/Files/Folder.vue';
import Uploadfile from '../components/Files/Uploadfile.vue';

export default {
    name: "Files",
    components: { Uploadfile, Folder, AddContent, Title }, 
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