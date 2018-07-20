<template>
    <footer id="app-footer" class="modal-card-foot">
        <p>
            <img src="/image/logo_yukai.png">
            <b-icon icon="close"/>
            <img src="/image/logo_hal.png">
        </p>
        <div class="buttons has-addons">
            <!--<span class="button" @click="$router.go(-1)">戻る</span>-->
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
        height: 8vh;
        padding: 0vh;
        bottom: 0vh;
        border-top:0;
        left:0px;
        background-color: #88CC29;
        z-index: 10;
    }

    .modal-card-foot {
        border-radius: 0;
    }

    #app-footer p{
        display: flex;
        color: white;
        height:100%;
        justify-content: center;
        align-items: center;
    }

    #app-footer img{
        height: auto;
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

    #app-footer img{
        height: 75%;
    }
</style>