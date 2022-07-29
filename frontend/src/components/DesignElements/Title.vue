<template>
    <main>
        <div class="title-section">
            <h1>
                {{title}}
            </h1>
            <FileOptions v-if="title == 'Files'" class="menu-options" />
        </div>
    </main>
</template>
<script>
import FileOptions from "./Options/FileOptions.vue";
export default {
    props: {
        title: {
            type: String,
            default: "Default Title"
        },
        options: {
            type: Object
        }
    },
    components: { FileOptions }
}
</script>
<style scoped>
    .title-section{
        margin-top: 24px;
        padding-bottom: 12px;
        border-bottom: 1px solid #ccc;
        display: flex;
        align-content: center;
        justify-content: space-between;
    }
    .menu-options{
       margin-right: 12px;
    }
    .title-section h1{
        color:white;
        font-family:monospace;
        margin: 0px;
        font-size: 1.6em;
        font-weight: 400;
        text-transform: capitalize;
        margin-bottom: 0.6em;
    }
</style>