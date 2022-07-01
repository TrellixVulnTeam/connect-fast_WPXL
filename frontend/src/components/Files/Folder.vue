<template>
  <main>
    
    <div>
      <button class="success" @click="getFiles('')"> Go Back</button>
    
      <div class="create">
          <input type="text" v-model="new_dir"/>
          <button @click="createDirectory(name)">Create Directory</button>
      </div>
    
      
      <!-- /* FILE MANAGMENT */ -->
        <div v-if="files.length > 0">
             <FileVue
            v-for="file in files"
            :key="file.id"
     
            :file="file"
            
            @click="navigate(file.name, file.fileType)"
          />
        </div>
        <div v-else  @click="getFiles('')" class="empty">
            No files found in the directory, click to go back
        </div>
    </div>
  </main>
</template>
<script>
import { CreateDirectory, FileStat, ReadDirectory, ReadFile } from "../../../wailsjs/go/main/App";
import FileVue from "./File.vue";

export default {
  name: "Folders",
  components: { FileVue },
  data() {
    return {
      files: [
      
      ],
      path:"", 
      navigation: [], 
      new_dir: ""
    };
  },
  methods: {
    getFiles(path) {
      ReadDirectory(path).then((result) => {
        this.files =[]
        if(result == null){
            console.timeLog("no files in the directory")
        }else{
             for(var i = 0; i < result.length; i++){
            this.files.push(result[i])
        }
        }
       
      });
    },
    fileStat(){
        FileStat().then(result => console.log(result))
    }, 
    navigate(newPath, fileType){
        if(fileType == "directory"){
            console.log(this.path + "/"+ newPath)
    
            this.path +="/" + newPath
            this.getFiles(this.path)
        }
        else{
          //this is a file and we need to read this. 
            this.readFile(this.path + "/" +newPath)
            console.log("this is a file so we will be reading it")
        }
    }, 
    readFile(path){
      console.log(path);
      ReadFile(path).then((result) => {
        console.log("reading the file...")
        console.log(result);
      })
    },
    createDirectory(name){
        // replaice all the whitespaces
        this.new_dir = this.new_dir.replaceAll(/\s/g,'')
        CreateDirectory(this.path + "/" + this.new_dir).then((result) => {
            console.log(result)
        }).catch(err => console.log(err))
    }
    

    
  },
  created() {
    ReadDirectory("").then((result) => {
        if(result.length < 1){
            console.timeLog("no files in the directory")
        }
       for(var i = 0; i < result.length; i++){
            this.files.push(result[i])
        }
    })
  },
};

</script>

<style>
table {
  padding: 5px;
  align-content: center;
  width: 100%;
}
path{
    margin: 5px;
}

.empty{
  margin: 5px;
  margin-top: 12px;
  background: rgb(244, 164, 164);
  padding: 12px;
  border-radius: 8px;
}
.empty:hover{
  cursor: pointer;
  background: rgb(255, 140, 140);
}
</style>