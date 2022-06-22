<script>

import ContractElement from '../components/ContractElement.vue';
export default {
    name: "Contracts",
    components: { ContractElement }
};
</script>

<template>
<main>
   <h2>
       Contracts
       </h2>
    
    <div class="grid-container">
           <ContractElement />
            <ContractElement />
            <ContractElement />
            <ContractElement />
    </div>

</main>
    
</template>
<style scoped>
.grid-container{
    display: flex;
    flex-wrap: wrap;
}
</style>