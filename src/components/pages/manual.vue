<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
        </div>
        <app-footer></app-footer>
        <under-tab :index='2'></under-tab>
    </div>
</template>

<script>
    import http from '../../service/service';
    import UnderTab from '../modules/underTab.vue'
    import AppHeader from '../modules/header.vue'
    import AppFooter from '../modules/footer.vue'
    export default {
        name :"settings",
        components:{
            UnderTab,
            AppHeader,
            AppFooter
        },
        data() {
            return {
                title: "取扱説明書"
            }
        },
        methods:{
        },
        created(){
        }
    }
</script>