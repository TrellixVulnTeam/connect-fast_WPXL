<template>
  <main>
    <Title title="Setup" />

    <div>
        <GenerateWallet />
        
        <Transfer />
    </div>
  </main>
</template>

<script>
import Title from '../components/DesignElements/Title.vue';
import GenerateWallet from '../components/GenerateWallet.vue';
import HelloWorld from '../components/HelloWorld.vue';
import Transfer from '../components/Transfer.vue';
export default {
    name: "Home",
    components: { HelloWorld, GenerateWallet, Transfer, Title }
};
</script>