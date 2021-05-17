import axios from 'axios'
axios.defaults.baseURL = 'http://localhost:10050';
axios.defaults.withCredentials = true;

export const login = (params, headers) => axios.post('/user/login/', params, headers);
export const register = (params, headers) => axios.post('/user/register/', params, headers);
export const getAllStrategies = (params, headers) => axios.get('/strategy/list/', params, headers);
export const getTradesByStrategyID = (strategy_url, page, params, headers) => axios.get(url + '/trades/' + page, params, headers);
export const getPositionsByStrategyID = (strategy_url, params, headers) => axios.get(url, '/positions', params, headers);
export const getPlanningTradesByStrategyID = (strategy_url, params, headers) => axios.get(url, '/planningTrades', params, headers);
