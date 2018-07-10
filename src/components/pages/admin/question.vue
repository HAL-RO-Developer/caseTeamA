<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">

            <section class="modal-card-body" lebel="問題データ">
                <b-field>
                    <b-input type="number" placeholder="ドリルID" v-model="data.book_id" required></b-input>
                </b-field>
                <b-field>
                    <b-input type="number" placeholder="問題No" v-model="data.question_no" required></b-input>               
                </b-field>
                <b-field>
                    <b-select placeholder="分野を選択" v-model="data.genre" required>
                        <option v-for="option in options.genre"
                            :value="option.genre_id"
                            :key="option.genre_id">
                            {{ option.genre_name }}
                        </option>
                    </b-select>
                    <button class="button is-info" type="button" @click="isComponentModalActive = true"><b-icon icon="plus"></b-icon></button>
                </b-field>
                <b-field>
                    <b-input type="textarea" maxlength="140" placeholder="問題文" v-model="data.sentence.text" required></b-input>
                </b-field>
                <b-field>
                    <b-input type="text" maxlength="30" placeholder="回答1" v-model="data.answer[0].text" required></b-input>
                </b-field>
                <b-field>       
                    <b-input type="text" maxlength="30" placeholder="回答2" v-model="data.answer[1].text" required></b-input>
                </b-field>
                <b-field>
                    <b-input type="text" maxlength="30" placeholder="回答3" v-model="data.answer[2].text" required></b-input>
                </b-field>
                <b-field label="正解を選択">
                    <b-select placeholder="正解を選択" v-model="corr" required expanded>
                        <option value="0">{{data.answer[0].text}}</option>
                        <option value="1">{{data.answer[1].text}}</option>
                        <option value="2">{{data.answer[2].text}}</option>
                    </b-select>
                </b-field>
            </section>
            <button class="button is-medium is-primary full-width" type="button" @click="createQuestion">作成</button>
        </div>
        <app-footer></app-footer>
        <b-modal :active.sync="isComponentModalActive" has-modal-card>
            <b-field>
                <b-input type="text" maxlength="10" placeholder="分野名" v-model="new_genre" required></b-input>
            </b-field>
            <button class="button is-primary full-width" type="button" @click="addGenre">分野追加</button>
        </b-modal>
    </div>
</template>

<script>
import http from '../../../service/service';
import UnderTab from '../../modules/underTab.vue'
import AppHeader from '../../modules/header.vue'
import AppFooter from '../../modules/footer.vue'

export default {
    name :"question",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
    },
    data() {
        return {
            title: "問題作成",
            data:{
                book_id: null,          // ドリルID
                question_no: null,      // 問題No
                sentence:{              // 問題文
                    tag_id:"",          // タグID
                    text:""             // 本文
                },
                answer:[                // 回答
                    {tag_id: "", text: "", answer: false},     // 回答１(タグID, 本文)
                    {tag_id: "", text: "", answer: false},     // 回答２(タグID, 本文)
                    {tag_id: "", text: "", answer: false}      // 回答３(タグID, 本文)
                ],
                genre:null              // 分野(id)
            },
            corr:"",
            options:{
                genre:[
                    /*
                    {genre_id: 1, genre_name:"国語"},
                    {genre_id: 2, genre_name:"算数"},
                    {genre_id: 3, genre_:"雑学"},
                    */
                ]
            },
            isComponentModalActive: false,
            new_genre:""
        }
    },
    methods:{
        createQuestion(){
            this.data.sentence.tag_id = "question" + this.data.book_id + "_" + this.data.question_no
            for(let i in this.data.answer) {
                this.data.answer[i].tag_id = "answer" + this.data.book_id + "_" + this.data.question_no + "_" + i
                if(i==this.corr) this.data.answer[i].answer = true
            }            
            console.log(JSON.stringify(this.data))
            //this.clear()
            http.createQuestion(this.data)
                .then((response)=>{
                    this.$toast.open('作成しました')
                    this.clear()
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
                        this.clear()
                    }
                })
        },
        clear(){
            this.data = {
                book_id: null,          // ドリルID
                question_no: null,      // 問題No
                sentence:{              // 問題文
                    tag_id:"",          // タグID
                    text:""             // 本文
                },
                answer:[                // 回答
                    {tag_id: "", text: "", answer: false},     // 回答１(タグID, 本文)
                    {tag_id: "", text: "", answer: false},     // 回答２(タグID, 本文)
                    {tag_id: "", text: "", answer: false}      // 回答３(タグID, 本文)
                ],
                genre:null              // 分野(id)
            }
            this.corr = ""
        },
        addGenre(){
            this.isComponentModalActive = false 
            http.addGenre(this.new_genre)
                .then((response)=>{
                    this.$toast.open("【" + this.new_genre + "】を追加しました")
            
                    this.getGenre()
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
                    }
                })
        },
        getGenre(){
            http.getGenre(this.new_genre)
                .then((response)=>{
                    this.options.genre = response.data.genre
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
                    }
                })
        }
    },
    created(){
        this.getGenre()
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