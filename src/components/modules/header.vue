<template>
    <header id="app-header" class="modal-card-head">
        <p class="modal-card-title"><img src="/image/bocco_w.png">{{title}}</p>
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
        margin-bottom: 5%;
    }
    #app-header img{
       padding: 0;
    }
</style>