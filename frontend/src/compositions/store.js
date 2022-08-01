import { reactive } from 'vue'

export const store = reactive({
    address: "",
    updateAddress(addr) {
        this.address = addr
    },

    filePath: "/",
    updatePath(newPath) {
        if (this.filePath.charAt(this.filePath.lastIndexOf()) == "/") {
            this.filePath += newPath
        } else {
            this.filePath += "/" + newPath
        }
    }

})