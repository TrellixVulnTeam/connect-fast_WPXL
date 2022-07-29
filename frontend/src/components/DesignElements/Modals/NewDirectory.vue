<template>

        <div class="modal-background">
                <div class="modal" @click.stop="nothing()">
                    <div class="modal_title">
                        <div class="modal_title_icon">
                            <img src="../../../assets/icons/web-technology/share.svg" />
                        </div>
                        <div class="modal_title_name">
                            Create folder 
                        </div>
                        <div class="modal_title_close" @click="close">
                            <img src="../../../assets/icons/interface-sign/close.svg" />
                        </div>
                    </div>
                    <div class="modal_content">
                        <input placeholder="Create new Directory" v-model="new_dir"/>
                    </div>
                    <div class="modal_footer">
                        <button @click="createDirectory">
                            Create Directory
                        </button>
                    </div>
                </div>
        </div>
</template>
<script>
import { CreateDirectory } from "../../../../wailsjs/go/main/App";
export default {
    data(){
        return{
            new_dir: '',    
            path:'',
        }
    }, 
    methods:{
        nothing(){
            console.log("nothing")
        },
        close(){
            this.$parent.changeSelect('');
        },
         createDirectory(name){
        // replaice all the whitespaces
            this.new_dir = this.new_dir.replaceAll(/\s/g,'')
            CreateDirectory(this.path + "/" + this.new_dir).then((result) => {
                console.log(result)
                this.$parent.changeSelect('');
            }).catch(err => console.log(err))
    }
    }
}
</script>>

<style scoped>
    .modal-background{
        background: RGBA(128,128,128, 0.4);
        position: fixed;
        top:0;
        left: 0;
        width: 100%;
        height: 100%;
        z-index:1;
        color: white;
        overflow: auto;
        padding-top: 25%;
    }
    .modal{
        background-color: #17141c;
        color: #c2c2c2;
        margin: 5% auto 15% auto; /* 5% from the top, 15% from the bottom and centered */
        width: 40%;
        padding: 32px;
        border-radius: 6px;
    }
    .modal_title{
        display:flex;
        align-content: center;
        align-items: center;
        gap: 12px;
        padding-bottom: 12px;
        border-bottom: 1px solid #c2c2c2;
    }
    .modal_title_name{
        flex: 1;
        font-size: 16px;
    }
    .modal_title_close{
        background: #20202a;
        padding: 6px;
        border-radius: 6px;
    }
.modal_title_close img{
        height: 16px;
        width: 16px;
    }
    .modal_content{
        padding: 24px;
        text-align:center;
    }
    .modal_content input{
        padding: 12px;
        font-size: 14px;
        width: 300px;
        border-radius: 5px;
        border: none;
    }
    .modal_footer{
        background:#20202a;
        border-radius: 6px;
        margin: -32px;
        padding: 24px;
        margin-top: 24px;
        text-align: right;
    }
    button{
        background: #3b3a41;
        color: white;
        border: none;
        font-size: 14px;
        text-transform: uppercase;
        font-weight: 700;
        padding: 16px;
    }
</style>