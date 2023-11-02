import axios from 'axios'
const baseUrl = 'http://localhost:8080'


const getMessage =  () =>{
    try{
        const request = axios.get(`${baseUrl}/hello`)
        return request.then(response =>{
            return response.data
        })

    }catch(error){
        throw new Error('Error fetching from backend')
    }
}

const messageService = {
    getMessage
}

export default messageService