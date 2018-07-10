<template>
    <form action="">
        <div class="modal-card" style="width: auto">
            <header class="modal-card-head">
                <p class="modal-card-title">ログイン</p>
            </header>
            <section class="modal-card-body">
                <b-field label="Name">
                    <b-input
                            type="text"
                            v-model="name"
                            placeholder="Your name"
                            required>
                    </b-input>
                </b-field>

                <b-field label="Password">
                    <b-input
                            type="password"
                            v-model="password"
                            password-reveal
                            placeholder="Your password"
                            required>
                    </b-input>
                </b-field>
                <!-- <b-checkbox>Remember me</b-checkbox> -->
            </section>
            <footer class="modal-card-foot">
                <button class="button" type="button" @click="signup">新規登録</button>
                <button class="button" type="button" @click="signin">ログイン</button>
            </footer>
        </div>
    </form>
</template>

<script>
import axios from 'axios'
import http from '../../service/service'
export default {
    data() {
        return {
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
