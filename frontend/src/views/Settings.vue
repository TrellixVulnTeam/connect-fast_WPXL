
<template>
  <main>
    <div class="navigation-items">
        <div @click="$router.push('/settings/')" :active="$route.matched.some(({ name }) => name === '')">
          Overview
        </div>
        <div @click="$router.push('/settings/ethereum')">
          Ethereum 
        </div>
        <div @click="$router.push('/settings/extra')">
          IPFS
        </div>
         <div @click="$router.push('/settings/contracts')">
          Contracts
        </div>
    </div>

    <div class="settings-content">
      <router-view></router-view>
    </div>
  </main>
</template>

<script>
import Address from "../components/Address.vue";
import Title from "../components/DesignElements/Title.vue";
import GenerateWallet from "../components/GenerateWallet.vue";

//need to generate the params that are here.

export default {
  name: "Settings",
  components: { GenerateWallet, Address, Title },
  data() {
    return {
      address: null,
    };
  },
  methods: {
    getAddress() {
      console.log("nothing here yet");
    },
  },
};
</script>

<style scoped>
.navigation-items {
  display: flex;
  flex-direction: row;
  padding:  12px;
  gap: 12px;
  flex: 1;
}
.navigation-items {
  border-radius: 6px;
  padding-left: 0;
  margin: 6px;
}
.navigation-items div{
  list-style-type: none;
  border-radius: 6px;
  padding: 12px;
  margin: 0px 0px;
  font-weight: 600;
}
.navigation-items div:hover {
  background: #d0f4ec;
  cursor: pointer;
}

.active {
  background: #b0f3e4;
  border-radius: 8px;
}

</style>