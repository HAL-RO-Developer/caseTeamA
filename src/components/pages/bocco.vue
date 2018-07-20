<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <section class="modal-card-body">
                    <!--
                    <b-field label="E-mail">
                        <b-input
                                type="email"
                                v-model="email"
                                placeholder="Your E-mail"
                                size="is-medium"
                                required>
                        </b-input>
                    </b-field>
                    <b-field label="Password">
                        <b-input
                                type="password"
                                v-model="password"
                                password-reveal
                                placeholder="Your password"
                                size="is-medium"
                                required>
                        </b-input>
                    </b-field>
                    -->
                    <div v-if="email">
                        <button class="button is-danger is-medium full-width" style="margin-bottom:1vh;" type="button" @click="disconnect">
                            <b-icon icon="power-plug-off"></b-icon>
                            <span style="padding-left:.5vh;">切断する</span>
                        </button>
                        <b-notification type="is-success" :closable="false">
                            <h3 class="title is-6"><b>BOCCOアカウント</b></h3>
                            <p class="subtitle is-6">{{email}}</p>
                        </b-notification>
                    </div>
                    <b-collapse v-else :open="false">
                        <button class="button is-success is-medium full-width" style="margin-bottom:1vh;" type="button" slot="trigger">
                            <b-icon icon="power-plug"></b-icon>
                            <span style="padding-left:.5vh;">接続する</span>
                        </button>
                        <b-notification type="is-success" :closable="false">
                            <section>
                                <h2 class="title is-6">BOCCOアカウントを入力</h2>
                                <b-field label="E-mail">
                                    <b-input
                                            type="email"
                                            v-model="form.email"
                                            placeholder="Your E-mail"
                                            size="is-medium"
                                            required>
                                    </b-input>
                                </b-field>
                                <b-field label="Password">
                                    <b-input
                                            type="password"
                                            v-model="form.password"
                                            password-reveal
                                            placeholder="Your password"
                                            size="is-medium"
                                            required>
                                    </b-input>
                                </b-field>
                                <button class="button is-info is-medium full-width" style="margin-bottom:1vh;" type="button" @click="connect">
                                    <b-icon icon="send"></b-icon>
                                    <span style="padding-left:.5vh;">送信</span>
                                </button>
                            </section>
                        </b-notification>
                    </b-collapse>
                </section>
        </div>
        <!--<img class="illust" src="/image/bocco.png">-->
        <under-tab :index='2'></under-tab>
        <app-footer></app-footer>
        
    </div>
</template>
<script>
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'

export default {
    name :"bocco",
    components:{
        UnderTab,
        AppHeader,
        AppFooter
    },
    data() {
        return {
            title:"BOCCOとつなげる",
            email:"",
            form:{
                email:"",
                password:"",
            }
        }
    },
    methods:{
       connect(){
           http.connectBocco(this.form.email,this.form.password)
                .then((response)=>{
                    console.log(response)
                    this.getBocco()
                    this.$toast.open('接続しました')
                })
                .catch((err)=>{
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                            icon: 'alert-circle',
                        })
                        switch(err.response.status){
                            case 401:
                                http.RemoveToken()
                                this.$router.push({path:'/'})
                                break;
                            default:
                                break;
                        }
                    }
                })
       },
       disconnect(){
           this.$dialog.confirm({
                title: 'BOCCO切断',
                message: 'BOCCOアカウントとの連携を解除します。',
                confirmText: '切断',
                type: 'is-danger',
                position: 'is-bottom',
                hasIcon: true,
                onConfirm: () => {
                    http.disconnectBocco()
                        .then((response)=>{
                            this.email = null
                            localStorage.removeItem('email')
                            this.$toast.open('切断しました')
                        })
                        .catch((err)=>{
                            if(err){
                                this.$dialog.alert({
                                    title: 'Error',
                                    message: err.response.data.error,
                                    type: 'is-danger',
                                    hasIcon: true,
                                    icon: 'alert-circle',
                                })
                                switch(err.response.status){
                                    case 401:
                                        http.RemoveToken()
                                        this.$router.push({path:'/'})
                                        break;
                                    default:
                                        break;
                                }
                            }
                        })
                    }
           })          
       },
       getBocco(){
            http.getBocco()
                .then((response)=>{
                    console.log(response.data)
                    this.email = response.data.email
                    localStorage.setItem("email", this.email);
                })
                .catch((err)=>{
                    this.isLoading = false
                    localStorage.removeItem('email')
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                            icon: 'alert-circle',
                        })
                        switch(err.response.status){
                            case 401:
                                http.RemoveToken()
                                this.$router.push({path:'/'})
                                break;
                            default:
                                break;
                        }
                    }
                });
       }
    },
    created(){
        this.getBocco();
    }
}
</script>

<style>
</style>