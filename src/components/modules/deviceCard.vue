<template>
    <div class="card">
        <div class="card-content">
            <div class="media">
                <div class="media-content">
                    <p class="title is-4">{{deviceId}}</p>
                </div>
                <div class="media-right">
                    <button class="button is-danger" @click="remove">
                        <b-icon icon="delete"></b-icon>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
            }
        },
        methods:{
            remove(){
                this.$emit('remove', this.deviceId )
            }
        },
        props:['deviceId']
    }
</script>