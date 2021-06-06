import axios from 'axios'
axios.defaults.baseURL = 'http://localhost:10050';
axios.defaults.withCredentials = true;

export const login = (data, headers) => axios.post('/user/login', data, { headers: headers });
export const logoutReq = (data, headers) => axios.post('/user/logout', data, { headers: headers });
export const register = (data, headers) => axios.post('/user/register', data, { headers: headers });
export const Subscribe = (url, params, headers) => axios.post(url + '/subscribe', { params: params, headers: headers })
export const Unsubscribe = (url, params, headers) => axios.post(url + '/unsubscribe', { params: params, headers: headers })
export const getAllStrategies = (params, headers) => axios.get('/strategy/list', { params: params, headers: headers });
export const GetTradesPageByStrategyID = (url, params, headers) => axios.get(url + '/list', { params: params, headers: headers });
export const getPositionsByStrategyID = (url, params, headers) => axios.get(url + '/list', { params: params, headers: headers });
export const getPlanningTradesByStrategyID = (url, params, headers) => axios.get(url + '/list', { params: params, headers: headers });
