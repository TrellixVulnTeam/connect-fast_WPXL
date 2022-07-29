<script>
export default{
 data(){
    return
        {
            
            open: false;
        }
 }
}
</script>
<template>
    <main>
        <div class="addButton" @click="open = !open">

        </div>
        
        <div v-if="open">
            add options here to choose from.
        </div>
    </main>
</template>
<style>

.addButton{
    position:fixed;
    left: 150px;
    bottom: 45px;
    overflow:hidden;
    background: green;
    border-radius: 50%;
    padding: 12px;
    margin: 0;
    display: flex;
    align-content: center;
    align-items: center;
    cursor:pointer;
}

</style>