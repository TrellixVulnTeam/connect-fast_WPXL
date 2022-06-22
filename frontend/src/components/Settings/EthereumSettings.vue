<script>

import GenerateWallet from '../GenerateWallet.vue';
export default {
  name: "Ethereum Settings", 
  components: { GenerateWallet }
};
</script>

<template>
    <main>
        <div class="content">
              <h1>Ethereum  Settings</h1>
        <div>
            <h3>Wallets</h3>
            
            
            <div class="create-wallet"> 
                
                <GenerateWallet />
            </div>
            <div class="wallets">
                <div class="wallet">
                    Ethereum. 
                    <span>
                        50,000
                    </span>
                </div>
            </div>
            <div class="transactions">
                <div class="transaction">
                    <img src="" alt="transaction-icon"/>
                    <div class="details">
                        <h5>Health Pro Supplies Ltd.</h5>
                        <span>Saturday 20 October 2022</span>
                    </div>
                    <div class="details">
                        <span>- $150.00</span>
                        <span>- 4:45</span>
                    </div>
                    
                </div>
            </div>
        </div>
        </div>
      
    </main>
</template>

<style scoped>

.wallets{
    width: 100%;
}

.transactions{
    display: flex;

}
.transaction{
    display: flex;
    padding: 6px;
    flex-direction: row;
    
}
.transaction .details{
    display: flex;
    flex-direction: column;
    margin: 0px 12px; 
}
.transaction img{
    width: 32px;
    height: 32px;
    margin: 5px;
    margin-right: 12px;
}
.transaction h5{
    margin: 0;
    font-weight: 600;
    font-size: 14px;
    
    
}
.transaction span {
    font-weight: 300;
    font-size: 14px;
}

.transaction:hover{
    background: #d5ffe8;
    cursor: pointer;
}

</style>