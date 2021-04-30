import { get, post } from 'axios'
export const login = (params, headers) => post('http://localhost:10050/user/login/', params, headers)
export const register = (params, headers) => post('http://localhost:10050/user/register/', params, headers)