<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <span class="title is-6">メッセージ</span>
        </header>
        <section class="modal-card-body">
            <b-field label="条件:">
                <b-select placeholder="Select a condition" v-model="data.condition" @input="change" required expanded>
                    <option
                        v-for="option in options.condition"
                        :value="option.id"
                        :key="option.id">
                        {{ option.name }}
                    </option>
                </b-select>
            </b-field>
            <b-field label="回数:">
                <b-input type="number" v-model="data.message_call" min="1" :disabled='count_disabled'></b-input>
            </b-field>
            <b-field label="メッセージ:">
                <b-input maxlength="140" type="textarea" v-model="data.message" required></b-input>
            </b-field>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">キャンセル</button>
            <button class="button is-primary" @click="add">追加</button>
        </footer>
    </div>
</template>

<script>
    export default {
        data(){
            return{
                count_disabled: true,
                data:{
                    condition: null,
                    message_call: null,
                    message: ""
                },
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
            change(){
                switch(this.data.condition){
                    case 3:
                        this.count_disabled = false
                        break
                    default:
                        this.count_disabled = true
                        break
                }
            },
            add(){
                this.$emit('add', this.data)
            }
        },
        created() {
        },
    }
</script>