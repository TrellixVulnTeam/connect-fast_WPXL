<script>
//TODO: Check if there is a connection with ethereum. 
import { store } from './compositions/store'

//check if there is localhost storage
export default {
  data(){
    return {
      store
    }
  }
}
</script>
<template>
  <div id="app">
    <div id="leftSide">
      <div id="nav">
        <div id="logo">
          <Dashboard  class="nav-item-icon"/>
            
            
        </div>
        <div class="nav-item" to="/" @click="$router.push('/')">
            <img
              class="nav-item-icon"
              src="./assets/icons/dashboard.svg"
              alt="Dashboard"
            />
          <router-link to="/">Dashboard</router-link>
        </div>
        <div class="nav-item"  @click="$router.push('/files')">
          <img
              class="nav-item-icon"
              src="./assets/icons/files.svg"
              alt="Dashboard"
            />
          <router-link to="/files">Files</router-link>
        </div>
         <div class="nav-item"  @click="$router.push('/settings')">
           <img
              class="nav-item-icon"
              src="./assets/icons/cog.svg"
              alt="Dashboard"
              fill="white"
            />
          <router-link to="/settings">Settings</router-link>
        </div>
         <div class="nav-item" >
           <img
              class="nav-item-icon"
              src="./assets/icons/file-1.svg"
              alt="Dashboard"
              fill="white"
            />
          <router-link to="/contracts">Contracts</router-link>
        </div>
      </div>
    </div>
    <div id="rightSide">
      <div id="top-nav">
        <div class="blockchain">
         {{store.address}}
        </div>
      
      </div>
      <div id="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<style>
#leftSide {
  width: 100px;
  height: 100%;
  background-color: #8aaae5;
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  overflow-y: auto;
}
#rightSide {
  width: calc(100% - 100px);
  height: 100%;
  left: 100;
  top: 0;
  bottom: 0;
  overflow-y: auto;
}
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
#nav {
  width: 100%;
  height: 50px;
  display: flex;
  flex-direction: column;
}
.nav-item {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1rem;
  color: #F4F9FF;
  font-weight: 300;
  padding-top: 20px;
  padding-bottom: 20px;
  display:flex;
  flex-direction: column;
}
.nav-item:hover {
  background-color: #F4F9FF;
  color: #8aaae5 !important;
  cursor: pointer
}
.nav-item:hover a {
  color: #8aaae5 !important;
}

.nav-item a {
  color: #F4F9FF;
  text-decoration: none;
}
.nav-item-icon {
  width: 42px;
  height: 42px;
  margin-right: 10px;
  padding: 10px;
  fill: #F4F9FF;
  color: #F4F9FF;
}

/* TOPBAR SETTINGS AND STYLING */
 #top-nav {
  background-color: #E6F4F1;
  display: flex;
  flex-direction:row;
  position: fixed;
  top: 0;
  width: calc(100% - 150px);
  align-content: center;
  padding: 24px;
  margin-left: 100px;
}


#content {
  margin-top: 100px;
}
</style>


