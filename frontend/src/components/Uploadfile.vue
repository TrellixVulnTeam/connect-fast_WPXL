<script setup>
// Components
// File Management
import useFileList from '../compositions/file-list';
// Uploader
import createUploader from '../compositions/file-uploader';
import DropZone from './DropZone.vue';
import FilePreview from './FilePreview.vue';


const { files, addFiles, removeFile } = useFileList()
function onInputChange(e) {
	addFiles(e.target.files)
	e.target.value = null // reset so that selecting the same file again will still cause it to fire this change
}

//needs to be fixed
const { uploadFiles } = createUploader('/tmp/files/s')

</script>

<template>
    <main>
     <div>
        <h1>File upload</h1>
        
       
       <DropZone class="drop-area" @files-dropped="addFiles" #default="{ dropZoneActive }">
			<label for="file-input">
				<span v-if="dropZoneActive">
					<span>Drop Them Here</span>
					<span class="smaller">to add them</span>
				</span>
				<span v-else>
					<span>Drag Your Files Here</span>
					<span class="smaller">
						or <strong><em>click here</em></strong> to select files
					</span>
				</span>

				<input type="file" id="file-input" multiple @change="onInputChange" />
			</label>
			<ul class="image-list" v-show="files.length">
				<FilePreview v-for="file of files" :key="file.id" :file="file" tag="li" @remove="removeFile" />
			</ul>
		</DropZone>
		<button @click.prevent="uploadFiles(files)" class="upload-button">Upload</button>
     </div>
    </main>
</template>

<style scoped>
.result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
}
</style>
