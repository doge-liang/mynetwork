import axios from 'axios'
axios.defaults.baseURL = 'http://localhost:10050';
axios.defaults.withCredentials = true;

export const login = (params, headers) => axios.post('/user/login', { params: params, headers: headers });
export const register = (params, headers) => axios.post('/user/register', { params: params, headers: headers });
export const getAllStrategies = (params, headers) => axios.get('/strategy/list', { params: params, headers: headers });
export const GetTradesPageByStrategyID = (url, params, headers) => axios.get(url + '/list', { params: params, headers: headers });
export const getPositionsByStrategyID = (url, params, headers) => axios.get(url + '/list', { params: params, headers: headers });
export const getPlanningTradesByStrategyID = (url, params, headers) => axios.get(url + '/list', { params: params, headers: headers });
