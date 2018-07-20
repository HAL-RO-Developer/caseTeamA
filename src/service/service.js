import axios from 'axios'

class Http{
    
    constructor(){
        this.Load()
    }

    Load(){
        this.api = axios.create({
            baseURL: localStorage.getItem('server'), 
            headers: {
              'ContentType': 'application/json',
              'Authorization': this.GetToken()
            },
            responseType: 'json'  
        })
    }

    SetToken(t){
        localStorage.setItem("token", t)
        http.Load();
    }
    GetToken(){
        return localStorage.getItem("token")
    }
    RemoveToken(){
        localStorage.removeItem("token")
    }
    
    signin(name, pass){
        return  this.api.post('/signin',{
            name,
            pass
        })
    }
    signup(name, pass){
        return  this.api.post('/signup',{
            name,
            pass
        })
    }
    
    registDevice(child_id){
        return  this.api.post('/device',{
            child_id
        })
    }
    getDevice(){
        return  this.api.get('/device')
    }
    removeDevice(device_id){
        return  this.api.delete('/device/'+device_id)
    }
    getChild(){
        return  this.api.get('/child')
    }
    addChild(nickname, birthday, sex){
        return  this.api.post('/child',{
            nickname,
            birthday,
            sex
        })
    }
    removeChild(child_id){
        return  this.api.delete('/child/'+child_id)
    }
    getRecords(child_id, filter){
        return  this.api.get('/work/graph/'+child_id,{
            params:{
                filter
            }
        })
    }
    getDetail(child_id, date, genre){
        return  this.api.get('/work/detail/'+child_id,{
            params:{
                date,
                genre
            }
        })
    }
    addMessage(child_id, condition, message_call, message ){
        return  this.api.post('/work/message',{
            child_id,
            condition,
            message_call,
            message
        })
    }
    getMessage(child_id){
        return  this.api.get('/work/message/'+child_id)
    }
    removeMessage(message_id){
        return  this.api.delete('/work/message/'+message_id)
    }
    createQuestion(data){
        let book_id = Number(data.book_id)
        let question_no = Number(data.question_no)
        let sentence = data.sentence
        let answer = data.answer
        let genre_id = Number(data.genre) 
        let correct = data.correct
        return  this.api.post('/question/create',{
            book_id,
            question_no,
            sentence,
            answer,
            genre_id,
            correct
        })
    }
    addGenre(genre_name){
        return  this.api.post('/question/genre',{
            genre_name
        })
    }
    getGenre(){
        return  this.api.get('/question/genre')
    }
    connectBocco(email, pass){
        return  this.api.post('/bocco',{
            email,
            pass
        })
    }
    disconnectBocco(){
        return  this.api.delete('/bocco')
    }
    getBocco(){
        return  this.api.get('/bocco')
    }
}
var http = new Http()
export default http;
