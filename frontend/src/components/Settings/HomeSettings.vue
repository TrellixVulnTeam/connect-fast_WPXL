<template>
        <div class="content">
            <div class="standard_info">
                <div class="info_photo">
                    <img src="../../assets/images/logo-universal.png" />
                    
                    <button>
                        Edit Photo
                    </button>
                </div>
                <div class="account_info">
                    <span class="main_account_info">
                        0x234235434534
                    </span>

                     <div class="info_row">
                        <div class="info_label">
                            Balance:
                        </div>
                        <div class="info_value">
                            200.029384   CTX
                        </div>
                        <div class="info_action">
                            Transactions
                        </div>
                    </div>
                    <div class="info_row">
                        <div class="info_label">
                            Email:
                        </div>
                        <div class="info_value">
                            connectfastv1@gmail.com
                        </div>
                        <div class="info_action">
                            Edit
                        </div>
                    </div>
                     <div class="info_row">
                        <div class="info_label">
                            Security check:
                        </div>
                        <div class="info_value">
                            not registered
                        </div>
                        <div class="info_action" >
                            Edit
                        </div>
                    </div>
                </div>
                
            </div>
        </div>

</template>
<style scoped>
.standard_info{
    display:flex;
    flex-direction: row;
}

.info_photo {
    display: flex;
    flex-direction: column;
}
.info_photo img{
  height: 128px;
  width: 128px;
  border-radius: 12px;  
}
.info_photo button{
    border: none;
    background: #21212a;
    color: white;
    font-weight: 600;
    margin-top: 12px;
    text-transform: uppercase;
    border-radius: 6px;
    padding: 12px;
}
.account_info{
    flex: 1;
    margin-left: 48px;
}
.main_account_info{
   font-size: 18px;
   font-weight: 600;
    color: white;
    padding-bottom: 24px;
    margin-bottom: 12px;
}

.info_row{
    display: flex;
    flex-direction: row;
    gap: 12px;
    padding: 12px;
    margin: 12px;
    border-bottom: 1px solid #c2c2c2;
}
.info_row .info_label{
    flex:1;
}
.info_row .info_value{
    flex: 3;
    color: #c2c2c2;
}

.info_action:hover{
    color: #c2c2c2;
    cursor: pointer;
}
hr {
    border-top: 1px solid #c2c2c2;
    margin-top: 2rem;
}

</style>