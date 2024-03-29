const express = require('express')
const { json, urlencoded } = require('body-parser')
const app = express()
const port =3000

app.use(json())
app.use(urlencoded({extended:true}))

app.get('/',(req,res)=>{
    res.status(200).send("Welcome to server")
})

app.get('/get',(req,res)=>{
    res.status(200).json("Hello from surbhi")
})

app.post('/post',(req,res)=>{
    let myJson=req.body
    res.status(200).send(myJson)
})

app.post('/postform',(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(port,()=>{
    console.log(`Example app listening at http://localhost:${port}`)
})