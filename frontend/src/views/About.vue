<template>
  <div class="home">
    <h1>
      About page Page
    </h1>

    <GetBalance />
  </div>
</template>

<script>
import GetBalance from '../components/GetBalance.vue';

export default {
    name: "Home",
    components: { GetBalance }
};
</script>