<template>
<div>
    <div class="options">
            <img src="../../../assets/icons/web-technology/add-files.svg" @click="changeSelect('add_directory')" alt='add_files' />
            <img src="../../../assets/icons/web-technology/search-alt.svg" alt="search_alt"/>
            <img src="../../../assets/icons/construction/ruler-pencil.svg" alt="ruler_pencil"/>
        </div>
        
        <NewDirectoryVue v-if="selected == 'add_directory'"  @click="changeSelect('')" />
</div>


</template>
<script>
import NewDirectoryVue from '../Modals/NewDirectory.vue';
export default {
    components: {NewDirectoryVue},
    data() {
        return{
            selected: ""
        }
    },
    methods:{
        changeSelect(value){
            this.selected = value;
        }
    }
}
</script>
<style>
.options{
    display: flex;
    gap: 1.5em;

}
img{
    filter: invert(1);
    opacity: 50%;
}
img:hover {
 opacity: 100%;
 cursor: pointer;
}
</style>