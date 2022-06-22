import { ref } from 'vue'

export default function useAddress(){
    const address = ref(null)
    const loggedIn = ref(false)
    
    function Login(addr){
        address.value = addr
        loggedIn.value = true
    }
    
    function Logout(){
        loggedIn.value = false
        address.value = null
    }
}