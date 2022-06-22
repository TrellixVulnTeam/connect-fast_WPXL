<template>
  <div class="home">
    <h1>
      Setup
    </h1>

    <div>
        <GenerateWallet />
        
        <Transfer />
    </div>
  </div>
</template>

<script>
import HelloWorld from '../components/HelloWorld.vue';
import GenerateWallet from '../components/GenerateWallet.vue';
import Transfer from '../components/Transfer.vue';
export default {
    name: "Home",
    components: { HelloWorld, GenerateWallet, Transfer }
};
</script>