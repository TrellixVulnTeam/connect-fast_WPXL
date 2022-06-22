<template>
   
        <tr v-if="file.type === 'folder'">
    <td>
      <img src="../../assets/icons/folder.svg" />
    </td>
    <td>{{file.name}}</td>
    <td>{{file.size}}</td>
    <td>Smart Contract 1</td>
    <td>January 24, 2022</td>
  </tr>
  <tr v-else>
    <td>
      <img src="../../assets/icons/file-1.svg" />
    </td>
    <td>{{file.name}}</td>
    <td>{{file.size}}</td>
    <td>Smart Contract 1</td>
    <td>January 24, 2022</td>
  </tr>
   

</template>
<script>
export default {
    props: {
    file: {
        type: Object
    }}
    
}
</script>

<style scoped>
img {
  height: 24px;
  width: 24px;
}
td {
  padding: 12px;
  font-weight: 100;
  margin-right: 12px;
  border:none;
}
tr:hover td {
  background-color: grey;
  border:none;
  cursor: pointer;
}
</style>