<template>
<main>
    <div class="container">
        <span>Progress:</span>
       <div class="progress-bar">   
        <div class="progress" width="20%"></div>
    </div>
    </div>
    
</main>
 
</template>
<style scoped>
.progress-bar{
    width: 100%;
    height: 20px;
    border-radius: 8px;
    border: 1px solid black;
}
.progress{
    background-color: 'green';
    height: 20px;
    width: 20%;
}
</style>