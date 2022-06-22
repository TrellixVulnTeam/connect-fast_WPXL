<script>

export default {
    data(){
        return {
            loggedIn: false, 
            address: "null"
        }
    }, 
    methods: {
        login(){
            this.address = "0x24"
            
        }, 
        logout(){
            this.loggedIn = false
            this.address = null
        }
    } 
    
    

}

</script>
<template>
<main>
     <div>
        <div v-if="loggedIn">
        <h1>Address: {{address}}</h1>
        Works    
        
        <button @click="logout">Logout</button>
    </div>
    <div v-else>
        <h1>Not logged in</h1>
        <p>
            {{address}}
        </p>
       
    </div>  
      <button @click="login">Log in</button>

   </div>
</main>
  
   
</template>

<style>
</style>