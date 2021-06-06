import { reactive } from "@vue/reactivity"

export const StateSymbol = Symbol("state");
export const CreateState = () => reactive({ isLogin: sessionStorage.getItem("isLogin") });