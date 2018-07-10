<template>
    <footer id="app-footer" class="modal-card-foot">
        <div class="buttons has-addons">
            <span class="button" @click="$router.go(-1)">戻る</span>
        </div>
    </footer>
</template>

<script>
export default {
    name :"app-footer",
    props:[],
    data() {
        return {
        }
    }
}
</script>

<style>
    #app-footer{
        display: block;
        position: fixed;
        width: 100vw;
        height: 10vh;
        padding: 0vh;
        bottom: 8vh;
        left:0px;
    }
    #app-footer .buttons{
        display: block;
        position: relative;
        height: 10vh;
    }
    #app-footer .buttons .button{
        height: 6vh;
        margin: 2vh;
    }
</style>