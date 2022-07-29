<script>
//TODO: Check if there is a connection with ethereum.
import { store } from "./compositions/store";

//check if there is localhost storage
export default {
  data() {
    return {
      store,
    };
  },
};
</script>
<template>
  <div id="app">
    <div id="leftSide">
      <div id="nav">
        <div id="logo">
          <img
            src="./assets/icons/web-technology/display.svg"
            class="nav-item-icon"
          />
        </div>
        <div class="nav-item" to="/" @click="$router.push('/')">
          <img
            src="./assets/icons/web-technology/display.svg"
            class="nav-item-icon"
          />
          <router-link to="/">Home</router-link>
        </div>
        <div class="nav-item" @click="$router.push('/files')">
          <img
            class="nav-item-icon"
            src="./assets/icons/files.svg"
            alt="Files"
          />
          <router-link to="/files">Files</router-link>
        </div>
        <div class="nav-item" @click="$router.push('/settings')">
          <img
            class="nav-item-icon"
            src="./assets/icons/cog.svg"
            alt="Dashboard"
          />
          <router-link to="/settings">Settings</router-link>
        </div>
        <div class="nav-item" @click="$router.push('/contracts')">
          <img
            class="nav-item-icon"
            src="./assets/icons/file-1.svg"
            alt="Dashboard"
          />
          <router-link to="/contracts">Contracts</router-link>
        </div>
      </div>
    </div>
    <div id="content">
      <router-view />
    </div>
  </div>
</template>

<style>
#leftSide {
  width: 100px;
  height: 100%;
  background-color: #21212a;
  position: fixed;
  left: 0;
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
  color: #cacacc;
  font-weight: 600;
  padding-top: 20px;
  padding-bottom: 20px;
  font-size: 12px;
  text-transform: uppercase;
  display: flex;
  flex-direction: column;
}
.nav-item:hover {
  filter: brightness(0) invert(1);
  cursor: pointer;
  border-right: solid 1px #fff;
}
.nav-item:hover a {
  color: #fff !important;
}

.nav-item a {
  color: #cacacc;
  text-decoration: none;
}
.nav-item-icon {
  width: 24px;
  height: 24px;
  margin-right: 10px;
  padding: 10px;
  fill: #f4f9ff;
  color: #f4f9ff;
  filter: brightness(3) invert(1);
}

#content {
  margin-left: 100px;
  padding: 12px;
}
</style>


