<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <button class="button is-medium is-primary full-width" type="button" @click="registDevice">PINコード発行</button>
            <section class="modal-card-body">
                <b-input
                    type="text"
                    v-model="pin"
                    placeholder="ここにPINコードが表示されます。"
                    readonly>
                </b-input>
            </section>
            <div class="subtitle is-5">【デバイスリスト】：{{devices.nickname}}</div>
            <div class="list">
                <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
                <card v-for="(device, index) in devices.child_devices" 
                    :key="index"
                    :deviceId="device"
                    @remove="removeDevice">
                </card>
            </div>
        </div>
        <fab :icon="fabIcon" @click="getDevice"></fab>
        <app-footer></app-footer>
        <under-tab :index='0'></under-tab>
    </div>
</template>

<script>
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'
import Card from '../modules/deviceCard.vue'
import Fab from '../modules/fab.vue'

export default {
    name :"device",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
        Card,
        Fab
    },
    data() {
        return {
            title: "デバイス追加",
            child_id: "",
            pin: "",
            devices: {},
            fabIcon: "sync",
            isLoading: false,
        }
    },
    methods:{
        alertCustom() {
            this.$dialog.alert({
                title: 'PINコード',
                message: this.pin,
                confirmText: 'OK',
            })
        },
        getDevice(){
            this.isLoading = true
            http.getDevice()
                .then((response)=>{
                    console.log(response)
                    this.isLoading = false
                    var devices = response.data.devices
                    devices.forEach((device)=>{
                        if(device.child_id == this.child_id){
                            this.devices = device
                        }
                    })
                })
                .catch((err)=>{
                    this.isLoading = false
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
        registDevice(){
            http.registDevice(Number(this.child_id))
                .then((response)=>{
                    console.log(response.data)
                    this.pin = response.data.pin
                    this.alertCustom()
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
        removeDevice(id){
            this.$dialog.confirm({
                title: 'デバイス削除',
                message: '『'+ id +'』を削除しますか？',
                confirmText: '削除',
                type: 'is-danger',
                hasIcon: true,
                onConfirm: () => 
                    http.removeDevice(id)
                        .then((response)=>{
                            console.log('delete')
                            this.getDevice()
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
                        })
            })
        }
    },
    created(){
        this.child_id = localStorage.getItem('child_id')
        this.getDevice()
    }
}
</script>

<style>
    button{
        display: block;
    }
    .list{
        position: relative;
        overflow: auto;
        height: 43vh;
    }
</style>