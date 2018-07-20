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
                    <button v-if="email" class="button is-danger is-medium full-width" style="margin-bottom:1vh;" type="button" @click="connect">
                        <b-icon icon="power-plug-off"></b-icon>
                        <span style="padding-left:.5vh;">切断する</span>
                    </button>
                    <button v-else class="button is-success is-medium full-width" style="margin-bottom:1vh;" type="button" @click="connect">
                        <b-icon icon="power-plug"></b-icon>
                        <span style="padding-left:.5vh;">接続する</span>
                    </button>
                    
                    <!-- <b-checkbox>Remember me</b-checkbox> -->
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
            password:""
        }
    },
    methods:{
       connect(){
           console.log(this.email)
           console.log(this.password)
       }
    },
    created(){
        this.email = localStorage.getItem("email");
    }
}
</script>

<style>
</style>