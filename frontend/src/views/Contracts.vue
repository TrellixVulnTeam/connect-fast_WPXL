<script>

import ContractElement from '../components/ContractElement.vue';
import TitleVue from '../components/DesignElements/Title.vue';
export default {
    name: "Contracts",
    components: { ContractElement, TitleVue }
};
</script>

<template>
<main>
   <TitleVue />
    
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