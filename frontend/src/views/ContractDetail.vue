<script>
import ProgressBar from "../components/DesignElements/ProgressBar.vue";
export default {
  name: "Contract Detail",
  components: { ProgressBar },
  data(){
      return {
          registered: false
          
      }
  },
  methods:{
     register(){
         this.registered = true
     }, 
     unregister(){
         this.registered = false
     }
  }
};
</script>
<template>
  <main class="content">
    <div class="top-content">
      <div class="back" @click="$router.push('/contracts')">
        <img src="../assets/icons/backward.svg"
            class="back-icon"
            alt="go back"
        />
        </div>
        <div class="top-title-box">
               <h1 class="title">
                Data Collection for improvement of product {contract.title}
                </h1>
                
                <span class="status" v-if="registered">Status: Register 
                <button class="btn-danger" @click="unregister()">Cancel</button>
                </span>
                <button class="btn-success" @click="register()" v-else>Register</button>
        </div>
   
    </div>

    <div>
      <p>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent auctor
        tempor libero, eget consectetur nisl viverra eu. Ut fringilla tempor
        metus, non imperdiet eros rhoncus eu. Cras euismod orci quam, eget
        blandit augue porttitor a. Aliquam quam erat, tincidunt sed ornare quis,
        convallis a mi. Etiam sit amet accumsan purus. Sed id justo a nisi
        blandit condimentum. Nam eu libero turpis. In varius quam ante, sed
        euismod neque condimentum at. Maecenas efficitur fermentum interdum.
        Suspendisse rhoncus pharetra mi, non vulputate urna suscipit at.
        Vestibulum in maximus ante, sed consectetur lectus.
      </p>
    </div>
    <div class="box-extra">
      
      <table class="additional">
        <tr>
          <th>Owner</th>
          <td>Facebook</td>
        </tr>
        <tr>
          <th>Reward</th>
          <td>10 dCFX</td>
        </tr>
        <tr>
          <th>Min. participants</th>
          <td>100.000</td>
        </tr>
        <tr>
          <th>Data Privacy</th>
          <td>High</td>
        </tr>
        <tr>
          <th>Resell rights</th>
          <td>No</td>
        </tr>
      </table>
      <table class="additional">
        <tr>
          <th>Start date</th>
          <td>10-10-2022 10:23</td>
        </tr>
        <tr>
          <th>End Date</th>
          <td>04-11-2022 21:34</td>
        </tr>
        <tr>
          <th>Current Participants</th>
          <td>1.000</td>
        </tr>
        <tr>
          <th>Max Participants</th>
          <td>unlimited</td>
        </tr>
        <tr>
          <th>Country</th>
          <td>United States of America</td>
        </tr>
      </table>
     
    </div>
     <div class="progress-box">
        <ProgressBar />
      </div>
  </main>
</template>
<style >
main {
  display: flex;
  padding: 32px;
  margin-left: 32px;
  flex-direction: column;
  text-align: left;
}
.back-icon{
    height: 16px;
    width: 16px;
}
.top-content {
  display: flex;
  align-items: center;
}


.back {
  padding: 24px;
  cursor: pointer;
}
.title {
  font-size: 32px;
  font-weight: 300;
}
.additional {
  font-size: 14px;
}
td, th{
    padding-right: 24px;
}
.box-extra{
    gap: 12px;
}
.btn-success{
    margin: 6px;
    padding: 8px 12px;
    background-color: #b0f3e4;
    font-size:14px;
    border: 1px solid #b0f3e4;
    border-radius: 6px;
    
}
.btn-danger{
    margin: 6px;
    padding: 8px 12px;
    background-color: #f3b0b0;
    font-size:14px;
    border: 1px solid #f3b0b0;
    border-radius: 6px;
    
}
</style>