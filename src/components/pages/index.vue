<template>
    <div>
        <under-tab :index='1'></under-tab>
    </div>
</template>

<script>
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'

export default {
    name :"dashboad",
    components:{
        UnderTab
    },
    data() {
        return {
        }
    },
    methods:{
        
    },
    created(){
        if(!http.GetToken()){
            this.$router.push({ path: '/login' })
        }else{
            this.$router.push({ path: '/device' })
        }
    }
}
</script>
