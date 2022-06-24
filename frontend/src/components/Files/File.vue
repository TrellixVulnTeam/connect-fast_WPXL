<template>
   <div class="file">
      <div class="type" v-if="file.fileType == 'directory'">
         <img src="../../assets/icons/folder.svg" />
      </div>
      <div class="type" v-else>
         <img src="../../assets/icons/file-1.svg" />
      </div>
      <div class="name">
        {{file.name}}
        <span class="date">
        {{file.last_modified}}
        </span>
      </div>
      <div class="size">
        {{Math.round(file.size /10000, 2)/100}} MB
      </div>
      <div class="contract">
        Integrated contract.   
      </div>
      
   </div>  
</template>
<script>
export default {
    props: {
    file: {
        type: Object
    }}
    
}
</script>

<style>
img {
  height: 24px;
  width: 24px;
}
.file{
  display: flex;
  gap: 8px;
  width: 100%;
}
.file div{
  margin: 4px 8px;
}
span::before{
  content: "\A";
  white-space: pre;
}
.date{
  font-size: 10px;
  font-style: italic;
  flex: 2;
}
.name {
  flex:3 
}
.type{
  flex: 1
}
.size{
  flex:1
}
</style>