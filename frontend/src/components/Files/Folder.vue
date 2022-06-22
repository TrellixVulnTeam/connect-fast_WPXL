<template>
  <main>
    <div>
      <button @click="getFiles('')">Get Files</button>

      <!-- /* FILE MANAGMENT */ -->
      <table>
        <thead>
          <tr>
            <th></th>
            <th>Name</th>
            <th>Size</th>
            <th>Contract</th>
            <th>Last modified</th>
          </tr>
        </thead>
        <tbody>
          <FileVue
            v-for="file in files"
            :key="file.id"
            :name="file.name"
            :type="file.type"
            :file="file"
          />
        </tbody>
      </table>
    </div>
  </main>
</template>
<script>
import { ReadDirectory } from "../../../wailsjs/go/main/App";
import FileVue from "./File.vue";

export default {
  name: "Folders",
  components: { FileVue },
  data() {
    return {
      files: [
      
      ],
    };
  },
  methods: {
    getFiles(path) {
      ReadDirectory(path).then((result) => {
        console.log(result)
      });
    },
  },
  created() {
    ReadDirectory("").then((result) => {
      for (var i = 0; i < result.length; i++) {
        console.log(result[i])
        this.files.push({
          id: i,
          type: result[i].filetype,
          name: result[i].cid,
          size: result[i].size,
        });
      }
    });
  },
};
</script>

<style>
table {
  padding: 5px;
  align-content: center;
}
</style>