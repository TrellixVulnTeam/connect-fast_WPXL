import { createRouter, createWebHistory } from 'vue-router'
import ContractsSettings from '../components/Settings/ContractsSettings.vue'
import EthereumSettings from '../components/Settings/EthereumSettings.vue'
import ExtraSettings from '../components/Settings/ExtraSettings.vue'
import HomeSettings from '../components/Settings/HomeSettings.vue'
import IpfsSettings from '../components/Settings/IpfsSettings.vue'
import About from '../views/About.vue'
import ContractDetail from '../views/ContractDetail.vue'
import Contracts from '../views/Contracts.vue'
import Files from '../views/Files.vue'
import Settings from '../views/Settings.vue'


const routes = [
  {
    path: '/', 
    name: 'Home', 
    component: EthereumSettings
  }, 
  {
    path: '/about',
    name: 'About',
    component: About
  }, 
  {
    path: '/files',
    name: 'Files',
    component: Files
  }
  , 
  {
    path: '/settings',
    name: 'Settings',
    component: Settings, 
    children: [
      {path: '', component: HomeSettings  }, 
      {path: 'ethereum', component: EthereumSettings}, 
      {path: 'ipfs', component: IpfsSettings}, 
      {path: 'contracts', component: ContractsSettings},
      {path: 'extra', component: ExtraSettings}
    ]
  }, 
  {
    path: '/contracts',
    name: 'Contracts',
    component: Contracts
  },
  {
    path: '/contract/:id',
    name: 'ContractDetail',
    component: ContractDetail
  },
  
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router;