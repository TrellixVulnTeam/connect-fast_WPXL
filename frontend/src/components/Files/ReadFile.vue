<template>
    <div>
        <button @click="displayFile">Ceck if variable is used</button>
       Here should we display the file.
       <pre>
            {{fileData}}
        </pre>
    </div>
</template>

<script>
export default {
    name: "ReadFile", 
    components: {}, 
    props: {
        fileData:{
            type: String, 
            default: "nothing found"
        }
    },
    methods: {
        displayFile(){
            console.log(fileData)
        }
    }
}
</script>>