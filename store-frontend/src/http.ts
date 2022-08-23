import axios from "axios";

const http = axios.create({
    //baseURL: process.env.NEXT_PUBLIC_STORE_API_URL
    baseURL: 'http://app:3000/api'
})

export default http;