import { reactive } from 'vue'

export const store = reactive({
    address: localStorage.getItem('address'), 
    updateAddress(addr){
        this.address = addr
    }
})