<template>
    <b-collapse class="card" :open="false">
            <div slot="trigger" slot-scope="props" class="card-header">
                <p class="card-header-title">
                    <span v-if="condition==3">{{messageCall}}回</span>
                    {{target.name}}
                </p>
                <a class="card-header-icon">
                    <b-icon
                        :icon="props.open ? 'menu-down' : 'menu-up'">
                    </b-icon>
                </a>
            </div>
            <div  v-for="(message,index) in messages" :key="index" class="card-content">
                <div class="media">
                    <div class="media-content">
                        <p class="subtitle is-6">{{message.message}}</p>
                    </div>
                    <div class="media-right">
                        <!--
                        <button class="button is-info" @click="edit">
                            <b-icon icon="pencil"></b-icon>
                        </button>
                        -->
                        <button class="button is-danger" @click="remove(message.message_id)">
                            <b-icon icon="delete"></b-icon>
                        </button>
                    </div>
                </div>
            </div>
        </b-collapse>
</template>

<script>
    export default {
        data() {
            return {
                target:{},
                options: {
                    condition:[
                        {id: 1, name: '正解のとき'},
                        {id: 2, name: '不正解のとき'},
                        {id: 3, name: '連続正解のとき'},
                    ]
                }
            }
        },
        methods:{
            edit(){},
            remove(id){
                this.$emit('remove', id )
            }
        },
        created() {
            this.target = this.options.condition.find((item, index)=>{
                if (item.id == this.condition) return true;
            });
        },
        props:['condition', 'messageCall','messages']
    }
</script>