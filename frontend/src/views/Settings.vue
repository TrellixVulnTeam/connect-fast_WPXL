<script>
import Address from "../components/Address.vue";
import GenerateWallet from "../components/GenerateWallet.vue";

//need to generate the params that are here. 

export default {
  name: "Settings",
  components: { GenerateWallet, Address },
  data() {
    return {
      address: null,
    };
  },
  methods: {
    getAddress() {
        console.log("nothing here yet")
    },
  },
};
</script>
<template>
  <main>
    <div class="left-items">
      <ul>
        <li @click="$router.push('/settings/')" :active="$route.matched.some(({ name }) => name === '')"
>Overview</li>
        <li @click="$router.push('/settings/ethereum')">Ethereum Details</li>
        <li @click="$router.push('/settings/ipfs')">IPFS.</li>
      </ul>
    </div>

      <div class="settings-content">
          <router-view></router-view>
       
      </div>
    
  </main>
</template>

<style scoped>
main{
    display: flex;
    flex-direction: row;
    gap: 24px;
    align-content: center;
    
}
.left-items{
    padding: 32px 12px;
    flex: 1;
    
}
.left-items ul{
  
    border-radius: 6px;
    
}
.left-items ul li{
    list-style-type: none;
    padding: 12px;
    margin: 12px 0px;
    font-weight: 600;
    
}
.left-items ul li:hover{
    background: #d0f4ec;
    cursor: pointer;
}

.active{
    background: #b0f3e4;
    border-radius: 8px;
}

/* SETTINGS CONTENT */
.settings-content{
    padding: 24px;
    width: 100%;
    flex:4;
    display: flex;
    flex-direction:column;
}
</style>