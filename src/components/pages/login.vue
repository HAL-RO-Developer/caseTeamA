<template>
    <form action="">
        <div class="modal-card" style="width: auto">
            <app-header :title="title"></app-header>
            <div class="contents">
                <section class="modal-card-body">
                    <b-field label="Name">
                        <b-input
                                type="text"
                                v-model="name"
                                placeholder="Your name"
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
                    <!-- <b-checkbox>Remember me</b-checkbox> -->
                </section>
                <button class="button is-info is-medium full-width" style="margin-bottom:1vh;" type="button" @click="signup">
                    <b-icon icon="account-plus"></b-icon>
                    <span style="padding-left:.5vh;"> 新規作成</span>
                </button>
                <button class="button is-success is-medium full-width" style="margin-bottom:1vh;" type="button" @click="signin">
                    <b-icon icon="login-variant"></b-icon>
                    <span style="padding-left:.5vh;"> ログイン</span>
                </button>
            </div>
            <app-footer></app-footer>
        </div>
    </form>
</template>

<script>
import axios from 'axios'
import http from '../../service/service'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'
export default {
    components:{
        AppHeader,
        AppFooter,
    },
    data() {
        return {
            title: "ログイン",
            name: "",
            password: ""
        }
    },
    methods:{
        signup(){
            http.signup(this.name, this.password)
            .then((response)=>{
                console.log(response)
                this.$dialog.alert({
                        title: 'ユーザー作成',
                        message: response.data.success,
                        type: 'is-info',
                        hasIcon: true,
                        icon: 'times-circle',
                        iconPack: 'fa'
                })
                this.signin()
            })
            .catch((err)=>{
                if(err){
                    this.$dialog.alert({
                        title: 'Error',
                        message: err.response.data.error,
                        type: 'is-danger',
                        hasIcon: true,
                        icon: 'times-circle',
                        iconPack: 'fa'
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
        },
        signin(){
            http.signin(this.name, this.password)
            .then((response)=> { 
                http.SetToken(response.data.token);
                this.$router.push({ path: '/' })
            })
            .catch((err)=> {
                if(err){
                    this.$dialog.alert({
                        title: 'Error',
                        message: err.response.data.error,
                        type: 'is-danger',
                        hasIcon: true,
                        icon: 'times-circle',
                        iconPack: 'fa'
                    })
                    switch(err.response.status){
                        case 401:
                            this.$router.push({path:'/'})
                            break;
                        default:
                            break;
                    }
                }
            });
        }
    }
}
</script>
