import {useState,useEffect} from "react";
import messageService from "../../services/message";


const BackendCall = () =>{
    const [message,setMessage] = useState<string>('')

    const fetchData = async () =>{
        try{
            const response = await messageService.getMessage()
            setMessage(response.Message)
        }catch(error){
            setMessage("Woops, could not communicate with server")
        }
    }


    useEffect(()=>{
        fetchData()
    },[])

    return <section>
        {message.slice(0,5) == 'Woops'?
        <h1>{message}</h1>
        :<>
        <h1>You received response from server~</h1>
        <p>{message}</p></>}
        
    </section>
}




export default BackendCall