<template>
  <main>
    <div>
        <ReadFileVue :fileData="filePreviewData"  v-if="filePreview" />
    </div>
    <div>
   
    <div class="folder_header">
      <div class="map_icon">
        <img src="../../assets/icons/web-technology/folder.svg" />
      </div>
      <div class="folder_title">
          Folder Directory name
      </div>
      <div class="folder_options">
        <img src="../../assets/icons/web-technology/menu.svg" @click="row"/>
        <img src="../../assets/icons/design/grid-alt.svg" @click="bigGrid"/>
        <img src="../../assets/icons/design/grid.svg" @click="smallGrid"/>
      </div>
    </div>
      <!-- /* FILE MANAGMENT */ -->
    <Loading v-if="loading" />
    <div v-else>
         <!-- <button class="success" @click="getFiles('')"> Go Back</button>
    
      <div class="create">
          <input type="text" v-model="new_dir"/>
          <button @click="createDirectory(name)">Create Directory</button>
      </div> -->
      
      
       <div  v-if="files.length > 0">
          <FileVue
            v-for="file in files"
            :key="file.id"
            :file="file"
            @dblclick="navigate(file.name, file.fileType)"
          />
        </div>
        <div v-else  @click="getFiles('')" class="empty">
            No files found in the directory, click to go back
        </div>
    </div>
       
    </div>
  </main>
</template>
<script>
import { CreateDirectory, FileStat, ReadDirectory, ReadFile } from "../../../wailsjs/go/main/App";
import Loading from "../DesignElements/Loading.vue";
import FileVue from "./File.vue";
import ReadFileVue from "./ReadFile.vue";


export default {
  name: "Folders",
  components: { FileVue, ReadFileVue, Loading },
  data() {
    return {
      files: [
      
      ],
      loading: false,
      path:"", 
      filePreview: false,
      filePreviewData: "",
      navigation: [], 
      new_dir: "", 
      viewOption: "row",
    };
  },
  methods: {
    getFiles(path) {
      this.loading = true;
      ReadDirectory(path).then((result) => {
        this.files =[]
        if(result == null){
            console.timeLog("no files in the directory")
        }else{
        for(var i = 0; i < result.length; i++){
            this.files.push(result[i])
        }
        }
       this.loading = false;
      });
    },
    selectView(view){
      this.viewOption = view;
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
        this.filePreviewData = result;
        this.filePreview = true;
        
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
   
    
    if(this.files.length <= 0){
       this.loading = true;
       ReadDirectory("").then((result) => {
        if(result.length < 1){
            console.timeLog("no files in the directory")
        }
       for(var i = 0; i < result.length; i++){
            this.files.push(result[i])
            
            this.loading = false;
        }
    })
    }
   
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


/* FLDER HEADER SETTINGS */

.folder_header{
  display: flex;
  align-content: center;
  gap: 24px;
  margin-bottom: 24px;
}
.folder_title{
  flex: 1;
  color: white;
  font-size: 16px;
  margin: 0;
  padding: 0;
}
.folder_options{
  background: #21212a;
  padding: 2px;
  align-self: flex-end;
}
.folder_options img{
  margin: 6px;
  height: 12px;
  width: 12px;
  padding: 3px;
}
.folder_options img:hover{
  background: #303036;
}
</style>