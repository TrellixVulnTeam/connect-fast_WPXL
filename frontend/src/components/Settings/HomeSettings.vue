<template>
    <main>
        <div class="content">
            <h1>Home Settings</h1>
            
            <div class="overview">
                <div class="box">
                    <h5>IPFS</h5>
                    <table>
                        <tr>
                            <td>Connected</td>
                            <td>True</td>
                        </tr>
                        <tr>
                            <td>Local Path</td>
                            <td>~/root/ipfs/</td>
                        </tr>                      
                         <tr>
                            <td>API</td>
                            <td>/ipfs/127.0.0.1/tcp/5001</td>
                        </tr>
                    </table>
                </div>
                 <div class="box">
                    <h5>Ethereum</h5>
                     <table>
                        <tr>
                            <td>Network</td>
                            <td>Ethereum</td>
                        </tr>
                         <tr>
                            <td>Default Account</td>
                            <td>3000</td>
                        </tr>
                        <tr>
                            <td>Wallet Balance</td>
                            <td>1000.000</td>
                        </tr>
                        <tr>
                            <td></td>
                            <td>
                                More details..
                            </td>
                        </tr>
                        </table>
                </div>
            </div>
        </div>
    </main>
</template>
<style scoped>
.overview{
    display: flex;
    flex-direction: row; 
    gap: 24px;
}
.overview .box{
    flex: 1;
    padding: 24px;
    border-radius: 6px;
    box-shadow: rgba(50, 50, 93, 0.25) 0px 2px 5px -1px,
    rgba(0, 0, 0, 0.3) 0px 1px 3px -1px;
}
.box h5{
    font-size: 20px;
    margin: 12px;
    text-transform: uppercase;
    font-weight:300;
}
</style>