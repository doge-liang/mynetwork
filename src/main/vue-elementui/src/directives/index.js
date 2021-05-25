import elTableInfiniteScroll from "./table-infinite-scroll"

const directives = {
    elTableInfiniteScroll,
}

export default {
    install(Vue) {
        Object.keys(directives).forEach((key) => {
            Vue.directive(key, directives[key])
        })
    },
}