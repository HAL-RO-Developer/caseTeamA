<template>
    <header id="app-header" class="modal-card-head">
        <p class="modal-card-title">
            <span class="title"><img src="/image/bocco_w.png">{{title}}</span>
        </p>
    </header>
</template>

<script>
export default {
    name :"app-header",
    props:['title'],
    data() {
        return {
        }
    }
}
</script>

<style>
    #app-header{
        display: block;
        position: fixed;
        width: 100vw;
        height: 10vh;
        top: 0px;
        left:0px;
        margin: 0px;        
        padding: 0vh;
        background-color: #88CC29;
    }

    .modal-card-head {
        border-radius: 0;
    }

    #app-header .modal-card-title{
        height:100%;
    }
    
    #app-header .modal-card-title .title{
        display: flex;
        color: white;
        height:100%;
        justify-content: center;
        align-items: center;
        color: white;
    }

    #app-header .modal-card-title img{
        height: 100%;
    }
    
</style>